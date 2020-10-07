package db

import (
	"context"
	"covidProject/logger"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var columnMap = map[string]int {
	"Last_Update": 2,
	"Confirmed": 5,
	"Deaths": 6,
	"Recovered": 7,
	"Active": 8,
	"Incident_Rate": 10,
	"People_Tested": 11,
	"People_Hospitalized": 12,
	"Mortality_Rate": 13,
	"UID": 14,
	"Testing_Rate": 16,
	"Hospitalization_Rate": 17,
}

func fixData(f string) *string {
	if f == "" || f == " " {
		return nil
	}
	return &f
}

func ReadFile(filename string) ([][]*string, error) {
	pathParts := strings.Split(filename, "/")
	dateUnformatted := pathParts[len(pathParts)-1][:10]
	dateParts := strings.Split(dateUnformatted, "-")
	date := dateParts[2] + "-" + dateParts[0] + "-" + dateParts[1]
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Problem opening %s\n%s", filename, err)
	}
	defer file.Close()

	var records [][]*string
	r := csv.NewReader(file)

	// Skip csv header
	_, err = r.Read()
	if err == io.EOF {
		return records, nil
	}

	for {
		recordAllCols, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Problem reading %s\n%s", filename, err)
		}
		record := make([]string, 0, 13) // Extract for 13 columns in target table
		record = append(record, recordAllCols[2])
		record = append(record, recordAllCols[5:9]...)
		record = append(record, recordAllCols[10:15]...)
		record = append(record, recordAllCols[16:]...)

		recordPointer := make([]*string, 0, 13)
		// Fix empty columns for numeric values.
		for _, val := range record {
			recordPointer = append(recordPointer, fixData(val))
		}
		records = append(records, append(recordPointer, &date))
	}
	return records, nil
}

func upsertStatement(rows [][]*string) string {
	query := "INSERT INTO covid_usa.cases_by_date(" +
		"\"Last_Update\"," +
		"\"Confirmed\"," +
		"\"Deaths\"," +
		"\"Recovered\"," +
		"\"Active\"," +
		"\"Incident_Rate\"," +
		"\"People_Tested\"," +
		"\"People_Hospitalized\"," +
		"\"Mortality_Rate\"," +
		"\"uid\"," +
		"\"Testing_Rate\"," +
		"\"Hospitalization_Rate\"," +
		"\"Created_At\") " +
		"VALUES "
	for rowNum, row := range rows {
		for i := 1; i <= len(row); i++ {
			if i == 1 {
				query += "("
			}
			if i < len(row) {
				query += "$" + strconv.Itoa(i+rowNum*13) + ","
			} else {
				query += "$" + strconv.Itoa(i+rowNum*13) + "),"
			}
		}
	}
	query = query[:len(query)-1] + " " // get rid of trailing comma
	query +=
		"ON CONFLICT(uid, \"Created_At\")" +
		"DO UPDATE SET " +
			"\"Last_Update\" = EXCLUDED.\"Last_Update\"," +
			"\"Confirmed\" = EXCLUDED.\"Confirmed\"," +
			"\"Deaths\" = EXCLUDED.\"Deaths\"," +
			"\"Recovered\" = EXCLUDED.\"Recovered\"," +
			"\"Active\" = EXCLUDED.\"Active\"," +
			"\"Incident_Rate\" = EXCLUDED.\"Incident_Rate\"," +
			"\"People_Tested\" = EXCLUDED.\"People_Tested\"," +
			"\"People_Hospitalized\" = EXCLUDED.\"People_Hospitalized\"," +
			"\"Mortality_Rate\" = EXCLUDED.\"Mortality_Rate\"," +
			"\"Testing_Rate\" = EXCLUDED.\"Testing_Rate\"," +
			"\"Hospitalization_Rate\" = EXCLUDED.\"Hospitalization_Rate\" " +
		"WHERE covid_usa.cases_by_date.\"Last_Update\" != EXCLUDED.\"Last_Update\";"
	return query
}

// Upsert data for one date
func (p *Pool) Upsert(ctx context.Context, rows [][]*string) error {
	query := upsertStatement(rows)
	rowsFlat := make([]interface{}, 0, 1100)
	for _, row := range rows {
		for _, v := range row {
			rowsFlat = append(rowsFlat, v)
		}
	}

	// Execute query
	_, err := p.Exec(ctx, query, rowsFlat...)
	if err != nil {
		return err
	}
	return nil
}

// Upsert all data
func (p *Pool) UpsertAll(ctx context.Context) error {
	dates, err := GetAllDates()
	if err != nil {
		return err
	}
	var errs []string
	for _, date := range dates {
		data, err := ReadFile(date)
		if err != nil {
			logger.Error(err)
			errs = append(errs, date)
			continue
		}
		err = p.Upsert(ctx, data)
		if err != nil {
			logger.Error(err)
			errs = append(errs, date)
			continue
		}
		logger.Info("Upserted data for " + date)
	}
	if len(errs) != 0 {
		errString := "failed to upsert data for "
		for _, e := range errs {
			errString += e + ", "
		}
		errString = errString[:len(errString)-2]
		return fmt.Errorf(errString)
	}
	return nil
}