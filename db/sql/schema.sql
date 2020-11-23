create schema if not exists covid_usa;

create table if not exists covid_usa.cases_by_date
(
    last_update          timestamp,
    confirmed            integer,
    deaths               integer,
    recovered            double precision,
    active               double precision,
    incident_rate        numeric,
    people_tested        double precision,
    people_hospitalized  double precision,
    mortality_rate       numeric,
    uid                  bigint    not null
        constraint cases_by_date_locations_uid_fk
            references locations,
    testing_rate         numeric,
    hospitalization_rate numeric,
    created_at           timestamp not null,
    constraint cases_by_date_pk
        primary key (uid, created_at)
);

create table if not exists covid_usa.locations
(
    uid            bigint not null
        constraint locations_pk
            primary key,
    iso2           char(2),
    iso3           char(3),
    code3          integer,
    fips           real,
    admin2         varchar,
    province_state varchar,
    country_region varchar,
    lat            double precision,
    long_          double precision,
    combined_key   varchar,
    population     bigint
);