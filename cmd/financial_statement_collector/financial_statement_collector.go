package main

import (
	"os"
	"time"
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

	functions := []string{"INCOME_STATEMENT", "BALANCE_SHEET", "CASH_FLOW"}
	// symbols := []string{"AAPL", "MSFT", "GOOGL", "META", "AMZN", "NFLX", "TSLA", "NVDA", "INTC", "AMD", "KO","PEP", "MCD", "SBUX", "NKE", "LULU", "GPS", "TGT", "WMT", "COST"}
	symbols := cfg.SYMBOLS
	for _, symbol := range symbols {
		earningDateToday, err := dbManager.IsEarningDateToday(symbol)
		if err != nil {
			logmanager.Errorf("Error checking earning date: %v", err)
			return 1
		}
		if earningDateToday {
			today := time.Now()
			formattedDate := today.Format("2006-01-02")
			// Print today's date
			logmanager.Infof("%s earning report today: %s", symbol, formattedDate)
			for _, function := range functions {
				data, err := api.FetchAPIData(function, symbol, cfg.APIKey)
				if err != nil {
					logmanager.Errorf("Error fetching data for %s: %v", function, err)
					return 1
				}
				switch function {
				case "INCOME_STATEMENT":
					if err := services.ProcessIncomeStatement(data, dbManager); err != nil {
						logmanager.Errorf("Error processing income statement: %v", err)
						return 1
					}
				case "BALANCE_SHEET":
					// Call ProcessBalanceSheet
					if err := services.ProcessBalanceSheet(data, dbManager); err != nil {
						logmanager.Errorf("Error processing balance sheet: %v", err)
						return 1
					}
				case "CASH_FLOW":
					// Call ProcessCashFlow
					if err := services.ProcessCashFlow(data, dbManager); err != nil {
						logmanager.Errorf("Error processing cash flow statement: %v", err)
						return 1
					}
				}
				
			}
		} else {
			logmanager.Infof("No earning report today for %s", symbol)
		}
	}
	
	return 0
}
