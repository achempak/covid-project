-- name: GetCases :many
SELECT * FROM covid_usa.cases_by_date;

-- name: GetCasesByUID :many
SELECT * FROM covid_usa.cases_by_date
WHERE uid = $1;

-- name: GetCasesByState :many
SELECT c.* FROM covid_usa.cases_by_date c
INNER JOIN covid_usa.locations l on l.uid = c.uid
WHERE l."province_state" = $1;

-- name: GetCasesByDate :many
SELECT * FROM covid_usa.cases_by_date
WHERE "created_at" = $1;

-- name: GetCasesSinceDate :many
SELECT * FROM covid_usa.cases_by_date
WHERE "created_at" >= $1;

-- name: GetCaseByUIDOnDate :one
SELECT * FROM covid_usa.cases_by_date
WHERE uid = $1 AND "created_at" = $2;

-- name: GetCasesByUIDSinceDate :many
SELECT * FROM covid_usa.cases_by_date
WHERE uid = $1 AND "created_at" >= $2;

-- name: GetLocationByUID :one
SELECT * FROM covid_usa.locations
WHERE uid = $1;

-- name: GetLocationByName :many
SELECT * FROM covid_usa.locations
WHERE "province_state" = $1;

-- name: GetLocations :many
SELECT * FROM covid_usa.locations;

-- name: ListLocationsByCaseIDs :many
SELECT l.*, c.uid AS case_id FROM covid_usa.cases_by_date c
INNER JOIN covid_usa.locations l on l.uid = c.uid AND c.uid = ANY($1::bigint[]);