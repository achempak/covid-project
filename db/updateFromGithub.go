package db

import (
	"covidProject/logger"
	"fmt"
	"io/ioutil"
	"os"
	//"os/exec"

	git "github.com/go-git/go-git/v5"
)

func PullData(dataDir string) error {
	//dataDir := os.Getenv("DATA_DIR")
	r, err := git.PlainOpen(dataDir)
	if err != nil {
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	logger.Info("git pull origin")
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		return err
	}
	ref, err := r.Head()
	if err != nil {
		return err
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return err
	}
	logger.Info(fmt.Sprintf("commit hash %s", commit))
	return nil
}

func GetAllDates() ([]string, error) {
	dataDir := os.Getenv("DATA_DIR") + "/csse_covid_19_data/csse_covid_19_daily_reports_us/"
	files, err := ioutil.ReadDir(dataDir)
	var dates []string
	if err != nil {
		return nil, fmt.Errorf("Problem getting dates from filenames\n%s", err)
	}
	for _, file := range files {
		if file.Name() == "README.md" {
			continue
		}
		dates = append(dates, dataDir + file.Name())
	}
	return dates, nil
}

//func getExistingDates() (map[string]struct{}, error) {
//	dataDir := os.Getenv("DATA_DIR")
//	file, err := os.OpenFile(dataDir + "/../existing_dates.txt", os.O_RDONLY, 0644)
//	if err != nil {
//		return nil, fmt.Errorf("Problem opening/creating existing_dates.txt\n%s", err)
//	}
//	defer file.Close()
//
//	fileSet := make(map[string]struct{}, 2)
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//		if _, ok := fileSet[scanner.Text()]; !ok {
//			fileSet[scanner.Text()] = struct{}{}
//		}
//	}
//	if err := scanner.Err(); err != nil {
//		return nil, fmt.Errorf("Problem scanning existing dates file\n%s", err)
//	}
//	return fileSet, nil
//}
//
//func updateExistingDates(processedDates []string) error {
//	dataDir := os.Getenv("DATA_DIR")
//	file, err := os.OpenFile(dataDir + "/../existing_dates.txt", os.O_WRONLY, 0644)
//	if err != nil {
//		return fmt.Errorf("Problem opening existing_dates.txt\n%s", err)
//	}
//	defer file.Close()
//	for _, d := range processedDates {
//		_, err := file.WriteString(d + "\n")
//		if err != nil {
//			return fmt.Errorf("Failed to write %s to existing_dates.txt\n%s", d, err)
//		}
//	}
//	return nil
//}