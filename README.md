# Introduction
This project collects financial data of public company

# Prerequisite
1. [Docker](https://docs.docker.com/engine/install/)
2. Data source [alphavantage.co](https://www.alphavantage.co/). Make sure you get an API key first. Note that free API key limits 25 calls per day

# Components

# To initialize database
Use ```initial.sql``` to initialize database

# Generate .env file

in /bin directory
```
cp .env_template .env
```
then change the info as you want

# How to build project 
### Build bin 
```
// home dir
go build -o ./bin/company_info_collector ./cmd/company_info_collector/company_info_collector.go
go build -o ./bin/financial_statement_collector ./cmd/financial_statement_collector/financial_statement_collector.go
```

### Build bin and add it to a docker image
```
docker build -t finanical_homelab:latest .
```

# Run components
## Option 1
### Run company_info_collector
```
./bin/company_info_collector
```
I run this service once a day. Recommand to setup a cronjob

### Run financial_statement_collector
```
./bin/company_info_collector
```
I run this service once a day

## Option 2:
### Run company_info_collector from docker image
```
docker run --rm financial_homelab:latest ./company_info_collector
```

### Run financial_statement_collector from docker image
```
docker run --rm financial_homelab:latest ./company_info_collector
```