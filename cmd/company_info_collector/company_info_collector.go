package main

import (
	"os"
	"financial_project/api"
	"financial_project/config"
	"financial_project/db"
	"financial_project/services"
	"financial_project/logmanager"
)
// 
// logmanager.Info("Task started")
// 	logmanager.Errorf("Error initializing database: %v", err)
// 	logmanager.Info("Task completed")
func main() {
	// Call the main application logic
	code := run()
	
	// Exit with the appropriate code
	os.Exit(code)
}

func run() int {
	cfg := config.LoadConfig()
	dbManager, err := db.NewDBManager(cfg.DBDSN)
	if err != nil {
		logmanager.Errorf("Error initializing database: %v", err)
		return 1
	}
	defer dbManager.Close()

	symbols := cfg.SYMBOLS
	logmanager.Infof("Processing company %s", symbols)
	for _, symbol := range symbols {
		
		result, err := dbManager.CompanyExists(symbol)
		if err != nil {
			logmanager.Errorf("Error checking company exists: %v", err)
			return 1
		}
		if result {
			logmanager.Infof("Company %s already exists", symbol)
			continue
		}
		function := "OVERVIEW"
		data, err := api.FetchAPIData(function, symbol, cfg.APIKey)
		if err != nil {
			logmanager.Errorf("Error fetching data for %s: %v", function, err)
			return 1
		}
		if err := services.ProcessCompanyInfo(data, dbManager); err != nil {
			logmanager.Errorf("Error processing company info: %v", err)
			return 1
		}
	}

	for _, symbol := range symbols {
		functions := []string{"INCOME_STATEMENT", "BALANCE_SHEET", "CASH_FLOW"}
		for _, function := range functions {
			// add initial data to financial statement tables
			switch function {
			case "INCOME_STATEMENT":
				// check if the data is missing
				missingData, err := dbManager.IsMissingIncomeStatement(symbol)
				if err != nil {
					return 1
				}
				if missingData {
					data, err := api.FetchAPIData(function, symbol, cfg.APIKey)
					if err != nil {
						return 1
					}
					if err := services.ProcessIncomeStatement(data, dbManager); err != nil {
						return 1
					}
				}
				
			case "BALANCE_SHEET":
				missingData, err := dbManager.IsMissingBalanceSheet(symbol)
				if err != nil {
					return 1
				}
				if missingData {
					data, err := api.FetchAPIData(function, symbol, cfg.APIKey)
					if err != nil {
						return 1
					}
					// Call ProcessBalanceSheet
					if err := services.ProcessBalanceSheet(data, dbManager); err != nil {
						return 1
					}
				}
			case "CASH_FLOW":
				missingData, err := dbManager.IsMissingCashFlow(symbol)
				if err != nil {
					return 1
				}
				if missingData {
					data, err := api.FetchAPIData(function, symbol, cfg.APIKey)
					if err != nil {
						return 1
					}
					// Call ProcessCashFlow
					if err := services.ProcessCashFlow(data, dbManager); err != nil {
						return 1
					}
				}
			}
		}
	}
	
	return 0
}
