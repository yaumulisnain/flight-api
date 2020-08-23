# FLIGHT API

Simple CRUD API for flights data with unit test. This project use Go (Golang) as language with [Gin](https://github.com/gin-gonic/gin), [Gorm](github.com/jinzhu/gorm) and another library. The database only support for PostgreSQL. 

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/54fe5c2c9d5448320564)

## Contents
- [Prerequisites](#prerequisites)
- [Setup](#setup)
- [Quick Start](#quick-start)

## Prerequisites
* Go v1.14
* PostgreSQL

## Setup
1. Clone this project
```
git clone https://github.com/yaumulisnain/flight-api
```
2. Setup your database.
    * You can copy and run DDL script from ```app/migration/00_ddl.sql``` for create database, or create your database manually.
    * You also need to copy and run SQL script from ```app/migration/01_table_flight.sql``` for create Flight table.
    * You can copy and run SQL script from ```app/migration/02_dummy_flight.sql``` for create dummy flight data, or insert manually.

3. You must create ```.env``` file, you can copy from ```.env.dev``` and then customize the value.
```
APP_ENV=dev
TZ=UTC
PORT=:3000

DB_HOST=localhost
DB_DATABASE=db_flights
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_PORT=5432
```

## Quick Start
This project using Makefile for simplify your commands.

1. Run project
```
make run
```

2. Build project
```
make build
```

3. Run binary file
```
make start
```

4. Run Unit Test
```
make test
```