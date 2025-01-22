package services

import (
	"encoding/json"
	"financial_project/models"
	"financial_project/db"
	"financial_project/logmanager"
)

func ProcessIncomeStatement(data []byte, dbManager *db.DBManager) error {
	var response struct {
		Symbol        string                      `json:"symbol"`
		AnnualReports []models.IncomeStatement `json:"annualReports"`
		QuarterReports []models.IncomeStatement `json:"quarterlyReports"`
	}
	if err := json.Unmarshal(data, &response); err != nil {
		logmanager.Errorf("Error parsing data from json: %v", err)
		return err
	}
	for _, report := range response.AnnualReports {
		exists, err := dbManager.RecordExists("income_statement_annual", response.Symbol, report.FiscalDateEnding)
		if err != nil {
			return err
		}

		if !exists {
			if err := dbManager.InsertIncomeStatement(report,"income_statement_annual",response.Symbol); err != nil {
				return err
			}
		} else {
			logmanager.Infof("Record (%s, %s) exists in table %s", response.Symbol, report.FiscalDateEnding, "income_statement_annual")
		}
	}

	for _, report := range response.QuarterReports {
		exists, err := dbManager.RecordExists("income_statement_quarter", response.Symbol, report.FiscalDateEnding)
		if err != nil {
			return err
		}

		if !exists {
			if err := dbManager.InsertIncomeStatement(report,"income_statement_quarter",response.Symbol); err != nil {
				return err
			}
		} else {
			logmanager.Infof("Record (%s, %s) exists in table %s", response.Symbol, report.FiscalDateEnding, "income_statement_quarter")
		}
	}

	return nil
}

func ProcessBalanceSheet(data []byte, dbManager *db.DBManager) error {
	var response struct {
		Symbol        string                      `json:"symbol"`
		AnnualReports []models.BalanceSheet `json:"annualReports"`
		QuarterReports []models.BalanceSheet `json:"quarterlyReports"`
	}

	if err := json.Unmarshal(data, &response); err != nil {
		logmanager.Errorf("Error parsing data from json: %v", err)
		return err
	}

	for _, report := range response.AnnualReports {
		exists, err := dbManager.RecordExists("balance_sheet_annual", response.Symbol, report.FiscalDateEnding)
		if err != nil {
			return err
		}

		if !exists {
			if err := dbManager.InsertBalanceSheet(report,"balance_sheet_annual",response.Symbol); err != nil {
				return err
			}
		} else {
			logmanager.Infof("Record (%s, %s) exists in table %s", response.Symbol, report.FiscalDateEnding, "balance_sheet_annual")
		}
	}

	for _, report := range response.QuarterReports {
		exists, err := dbManager.RecordExists("balance_sheet_quarter", response.Symbol, report.FiscalDateEnding)
		if err != nil {
			return err
		}

		if !exists {
			if err := dbManager.InsertBalanceSheet(report,"balance_sheet_quarter",response.Symbol); err != nil {
				return err
			}
		} else {
			logmanager.Infof("Record (%s, %s) exists in table %s", response.Symbol, report.FiscalDateEnding, "balance_sheet_quarter")
		}
	}

	return nil
}

func ProcessCashFlow(data []byte, dbManager *db.DBManager) error {
	var response struct {
		Symbol        string                      `json:"symbol"`
		AnnualReports []models.CashFlow `json:"annualReports"`
		QuarterReports []models.CashFlow `json:"quarterlyReports"`
	}

	if err := json.Unmarshal(data, &response); err != nil {
		logmanager.Errorf("Error parsing data from json: %v", err)
		return err
	}

	for _, report := range response.AnnualReports {
		exists, err := dbManager.RecordExists("cash_flow_annual", response.Symbol, report.FiscalDateEnding)
		if err != nil {
			return err
		}

		if !exists {
			if err := dbManager.InsertCashFlow(report,"cash_flow_annual",response.Symbol); err != nil {
				return err
			}
		} else {
			logmanager.Infof("Record (%s, %s) exists in table %s", response.Symbol, report.FiscalDateEnding, "cash_flow_annual")
		}
	}

	for _, report := range response.QuarterReports {
		exists, err := dbManager.RecordExists("cash_flow_quarter", response.Symbol, report.FiscalDateEnding)
		if err != nil {
			return err
		}

		if !exists {
			if err := dbManager.InsertCashFlow(report,"cash_flow_quarter",response.Symbol); err != nil {
				return err
			}
		} else {
			logmanager.Infof("Record (%s, %s) exists in table %s", response.Symbol, report.FiscalDateEnding, "cash_flow_quarter")
		}
	}

	return nil
}