package services

import (
	"encoding/csv"
	// "financial_project/models"
	"financial_project/db"
	"financial_project/logmanager"
)

func EarningCalendarProcessor(reader *csv.Reader, dbManager *db.DBManager) error {
	// Read and process the CSV file
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			logmanager.Errorf("Error reading CSV: %v\n", err)
			return err
		}
		// Extract and print Symbol and ReportDate
		// logmanager.Infof("Symbol: %s, ReportDate: %s\n", record[0], record[2])
		exists, err := dbManager.CompanyExists(record[0])
		if err != nil {
			return err
		}

		if exists {
			if err := dbManager.UpdateCompanyEarningDate(record[0], record[2]); err != nil {
				return err
			}
			logmanager.Infof("Updated %s with new earning date %s\n", record[0], record[2])
		}
	}

	return nil
}