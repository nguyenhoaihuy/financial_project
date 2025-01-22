package main
import (
	"os"
	"financial_project/api"
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
	function := "EARNINGS_CALENDAR"
	horizon := "2month"
	reader, cleanup, err := api.FetchCSVReader(function,horizon,cfg.APIKey)
	if err != nil {
		logmanager.Errorf("Error fetching CSV: %v\n", err)
		return 1
	}
	defer cleanup()

	// Read and process the CSV file
	if err := services.EarningCalendarProcessor(reader, dbManager); err != nil {
		logmanager.Errorf("Error processing earning calendar: %v", err)
		return 1
	}
	
	return 0
}
