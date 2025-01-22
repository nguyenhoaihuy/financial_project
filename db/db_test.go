package db

import (
	"financial_project/config"
	"testing"
	"financial_project/logmanager"
	"financial_project/models"
)

func TestCheckCompanyExists(t *testing.T) {
	cfg := config.LoadConfig()
	dbManager, err := NewDBManager(cfg.DBDSN)
	if err != nil {
		logmanager.Errorf("Error initializing database: %v", err)
	}
	defer dbManager.Close()

	result, err := dbManager.CompanyExists("HUYTEST")
	expect := false
	if result != expect {
		t.Errorf("Expected %v but got %v", expect, result)
	}
	company := models.Company{
        Symbol: "HUYTEST", 
		Name: "Huy",
    }
	err = dbManager.InsertCompany(company)
	if err != nil {
		logmanager.Errorf("Error adding a company: %v", err)
	}

	result, err = dbManager.CompanyExists("HUYTEST")
	expect = true
	if result != expect {
		t.Errorf("Expected %v but got %v", expect, result)
	}

	err = dbManager.DeleteCompany("HUYTEST")
	if err != nil {
		logmanager.Errorf("Error deleting a company: %v", err)
	}

	result, err = dbManager.CompanyExists("HUYTEST")
	expect = false
	if result != expect {
		t.Errorf("Expected %v but got %v", expect, result)
	}
}