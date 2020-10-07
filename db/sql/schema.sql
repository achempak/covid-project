create schema if not exists covid_usa;

create table if not exists covid_usa.cases_by_date
(
    "Last_Update"          timestamp,
    "Confirmed"            integer,
    "Deaths"               integer,
    "Recovered"            double precision,
    "Active"               double precision,
    "Incident_Rate"        numeric,
    "People_Tested"        double precision,
    "People_Hospitalized"  double precision,
    "Mortality_Rate"       numeric,
    uid                    bigint    not null
        constraint cases_by_date_locations_uid_fk
            references locations,
    "Testing_Rate"         numeric,
    "Hospitalization_Rate" numeric,
    "Created_At"           timestamp not null,
    constraint cases_by_date_pk
        primary key (uid, "Created_At")
);

create table if not exists covid_usa.locations
(
    uid              bigint not null
        constraint locations_pk
            primary key,
    iso2             char(2),
    iso3             char(3),
    code3            integer,
    fips             real,
    "Admin2"         varchar,
    "Province_State" varchar,
    "Country_Region" varchar,
    "Lat"            double precision,
    "Long_"          double precision,
    "Combined_Key"   varchar,
    "Population"     bigint
);