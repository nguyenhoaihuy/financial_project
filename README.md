# Introduction
This project collects financial data of public company

# Prerequisite
[Docker](https://docs.docker.com/engine/install/)

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
### Run company_info_collector from docker image
```
./bin/company_info_collector
```

### Run company_info_collector from docker image

```
docker run --rm financial_homelab:latest ./company_info_collector
```