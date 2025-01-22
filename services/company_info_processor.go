package services

import (
	"encoding/json"
	"financial_project/models"
	"financial_project/db"
	"financial_project/logmanager"
)

func ProcessCompanyInfo(data []byte, dbManager *db.DBManager) error {
	var response models.Company
	if err := json.Unmarshal(data, &response); err != nil {
		logmanager.Errorf("Error parsing data from json: %v", err)
		return err
	}
	if response.Symbol == "" {
		logmanager.Errorf("Symbol is empty. Skipping record")
		return nil
	}
	exists, err := dbManager.CompanyExists(response.Symbol)
	if err != nil {
		return err
	}

	if !exists {
		if err := dbManager.InsertCompany(response); err != nil {
			return err
		}
		logmanager.Infof("Company %s added to database", response.Symbol)
	} else {
		logmanager.Infof("Record %s exists in table company_info", response.Symbol)
	}

	return nil
}