BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE product_type AS ENUM (
    'TypeA',
    'TypeB',
    'TypeC',
    'TypeD'
    );


CREATE TABLE IF NOT EXISTS "products"(
"id" UUID NOT NULL UNIQUE DEFAULT uuid_generate_v4(),
"name" VARCHAR(100) PRIMARY KEY NOT NULL CHECK(LENGTH(name)>5),
"quantity" INT,
"type" product_type NOT NULL ,
"price" NUMERIC(10, 2) NOT NULL CHECK(price>=0) DEFAULT 0.00,
"description" varchar(500)
);

COMMIT;
