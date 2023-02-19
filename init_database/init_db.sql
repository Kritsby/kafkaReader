CREATE TABLE IF NOT EXISTS "public"."exgauster" (
    guid uuid unique not null primary key ,
    name text,
    work double precision
);

CREATE TABLE IF NOT EXISTS "public"."bearing" (
    guid uuid unique not null primary key ,
    number int,
    exgauster_guid uuid
);

CREATE TABLE IF NOT EXISTS "public"."temperature" (
    guid uuid unique  not null primary key,
    temperature double precision,
    bearing_guid uuid
);

CREATE TABLE IF NOT EXISTS "public"."vibration" (
    guid uuid unique not null primary key,
    type text,
    bearing_guid uuid
);

CREATE TABLE IF NOT EXISTS "public"."alarm_and_warning" (
    guid uuid unique not null primary key,
    alarm_max double precision,
    alarm_min double precision,
    warning_max double precision,
    warning_min double precision,
    main_guid uuid
);

CREATE TABLE IF NOT EXISTS "public"."cooler" (
    guid uuid unique not null primary key,
    type text,
    exgauster_guid uuid
);

CREATE TABLE IF NOT EXISTS "public"."cooler_temperature" (
    guid uuid unique not null primary key,
    temperature_before double precision,
    temperature_after double precision,
    cooler_guid uuid
);

CREATE TABLE IF NOT EXISTS "public"."gas_manifold" (
    guid uuid unique not null primary key,
    temperature_before double precision,
    underpresure_before double precision,
    exgauster_guid uuid
);

CREATE TABLE IF NOT EXISTS "public"."gate_valve" (
    guid uuid unique not null primary key,
    closed double precision,
    open double precision,
    position double precision,
    exgauster_guid uuid
);

CREATE TABLE IF NOT EXISTS "public"."main_drive" (
    guid uuid unique not null primary key,
    rotor_current double precision,
    rotor_voltage double precision,
    stator_current double precision,
    stator_voltage double precision,
    exgauster_guid uuid
);

CREATE TABLE IF NOT EXISTS "public"."oil_system" (
    guid uuid unique not null primary key,
    oil_level double precision,
    oil_pressure double precision,
    exgauster_guid uuid
);








