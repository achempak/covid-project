-- name: GetCasesByUID :many
SELECT c.*, l."Province_State" FROM covid_usa.cases_by_date c
LEFT JOIN covid_usa.locations l on c.uid = l.uid
WHERE c.uid = $1;

-- name: GetCasesByState :many
SELECT c.*, l."Province_State" FROM covid_usa.cases_by_date c
LEFT JOIN covid_usa.locations l on l.uid = c.uid
WHERE l."Province_State" = $1;

-- name: GetCasesByDate :many
SELECT c.*, l."Province_State" FROM covid_usa.cases_by_date c
LEFT JOIN covid_usa.locations l on l.uid = c.uid
WHERE c."Created_At" = $1;

-- name: GetCasesSinceDate :many
SELECT c.*, l."Province_State" FROM covid_usa.cases_by_date c
LEFT JOIN covid_usa.locations l on l.uid = c.uid
WHERE c."Created_At" >= $1;

-- name: GetCasesByStateOnDate :many
SELECT c.*, l."Province_State" FROM covid_usa.cases_by_date c
LEFT JOIN covid_usa.locations l on l.uid = c.uid
WHERE l."Province_State" = $1 AND c."Created_At" = $2;

-- name: GetCasesByStateSinceDate :many
SELECT c.*, l."Province_State" FROM covid_usa.cases_by_date c
LEFT JOIN covid_usa.locations l on l.uid = c.uid
WHERE l."Province_State" = $1 AND c."Created_At" >= $2;