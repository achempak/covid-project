-- name: GetCases :many
SELECT * FROM covid_usa.cases_by_date;

-- name: GetCasesByUID :many
SELECT * FROM covid_usa.cases_by_date
WHERE uid = $1;

-- name: GetCasesByState :many
SELECT c.* FROM covid_usa.cases_by_date c
LEFT JOIN covid_usa.locations l on l.uid = c.uid
WHERE l."Province_State" = $1;

-- name: GetCasesByDate :many
SELECT * FROM covid_usa.cases_by_date
WHERE "Created_At" = $1;

-- name: GetCasesSinceDate :many
SELECT * FROM covid_usa.cases_by_date
WHERE "Created_At" >= $1;

-- name: GetCaseByUIDOnDate :one
SELECT * FROM covid_usa.cases_by_date
WHERE uid = $1 AND "Created_At" = $2;

-- name: GetCasesByUIDSinceDate :many
SELECT * FROM covid_usa.cases_by_date
WHERE uid = $1 AND "Created_At" >= $2;

-- name: GetLocationByUID :one
SELECT * FROM covid_usa.locations
WHERE uid = $1;

-- name: GetLocationByName :many
SELECT * FROM covid_usa.locations
WHERE "Province_State" = $1;

-- name: GetLocations :many
SELECT * FROM covid_usa.locations;