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

	// symbols := []string{"AAPL", "MSFT", "GOOGL", "META", "AMZN", "NFLX", "TSLA", "NVDA", "INTC", "AMD", "KO","PEP", "MCD", "SBUX", "NKE", "LULU", "GPS", "TGT", "WMT", "COST"}
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
	
	return 0
}
