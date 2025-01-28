package main

import (
	"os"
	// "time"
	// "financial_project/api"
	"financial_project/config"
	"financial_project/db"
	"financial_project/services"
	"financial_project/logmanager"
)

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

	// calculate annual KPI
	income_statement_table := "income_statement_annual"
	balance_sheet_table := "balance_sheet_annual"
	kpi_table := "kpi_annual"
	err = services.ProcessKPI(income_statement_table, balance_sheet_table, kpi_table, dbManager)
	if err != nil {
		logmanager.Errorf("Error processing annual KPI: %v", err)
		return 1
	}

	// calculate quarterly KPI
	income_statement_table = "income_statement_quarter"
	balance_sheet_table = "balance_sheet_quarter"
	kpi_table = "kpi_quarter"
	err = services.ProcessKPI(income_statement_table, balance_sheet_table, kpi_table, dbManager)
	if err != nil {
		logmanager.Errorf("Error processing quarterly KPI: %v", err)
		return 1
	}
	return 0
}