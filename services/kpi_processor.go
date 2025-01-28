package services

import (
	// "encoding/json"
	"strconv"
	"financial_project/models"
	"financial_project/db"
	"financial_project/logmanager"
)

func ProcessKPI(income_statement_table string, balance_sheet_table string, kpi_table string, dbManager *db.DBManager) error {
	// 1. Get the list of missing records
	missingRecords, err := dbManager.GetMissingKPI(income_statement_table,kpi_table)
	for _, record := range missingRecords {

		// logmanager.Infof("Missing record (%s, %s) in table %s", record.Symbol, record.FiscalDate, kpi_table)
		incomeStatement, err := dbManager.GetIncomeStatement(income_statement_table, record.Symbol, record.FiscalDate)
		if err != nil {
			logmanager.Errorf("Error fetching income statement for symbol %s and date %s: %v", record.Symbol, record.FiscalDate, err)
			continue
		}
		balanceSheet, err := dbManager.GetBalanceSheet(balance_sheet_table, record.Symbol, record.FiscalDate)
		if err != nil {	
			logmanager.Errorf("Error fetching balance sheet for symbol %s and date %s: %v", record.Symbol, record.FiscalDate, err)
			continue
		}

		netIncome, _ := strconv.Atoi(incomeStatement.NetIncome)
		totalRevenue, _ := strconv.Atoi(incomeStatement.TotalRevenue)
		totalAssets, _ := strconv.Atoi(balanceSheet.TotalAssets)
		commonStockSharesOutstanding, _ := strconv.Atoi(balanceSheet.CommonStockSharesOutstanding)
		costOfGoodsAndServicesSold, _ := strconv.Atoi(incomeStatement.CostOfGoodsAndServicesSold)
		cashAndCashEquivalentsAtCarryingValue, _ := strconv.Atoi(balanceSheet.CashAndCashEquivalentsAtCarryingValue)
		grossProfit, _ := strconv.Atoi(incomeStatement.GrossProfit)
		totalCurrentAssets, _ := strconv.Atoi(balanceSheet.TotalCurrentAssets)
		totalCurrentLiabilities, _ := strconv.Atoi(balanceSheet.TotalCurrentLiabilities)
		inventory, _ := strconv.Atoi(balanceSheet.Inventory)
		interestExpense, _ := strconv.Atoi(incomeStatement.InterestExpense)
		totalLiabilities, _ := strconv.Atoi(balanceSheet.TotalLiabilities)
		totalShareholderEquity, _ := strconv.Atoi(balanceSheet.TotalShareholderEquity)
		currentAccountsPayable, _ := strconv.Atoi(balanceSheet.CurrentAccountsPayable)
		currentNetReceivables, _ := strconv.Atoi(balanceSheet.CurrentNetReceivables)
		ebit, _ := strconv.Atoi(incomeStatement.Ebit)
		// ebitda, _ := strconv.Atoi(incomeStatement.Ebitda)

		var KPIRecord models.KPI
		KPIRecord.NetProfitMargin = float64(netIncome) / float64(totalRevenue)
		KPIRecord.ROE = float64(netIncome) / float64(totalShareholderEquity)
		KPIRecord.EPS = float64(netIncome) / float64(commonStockSharesOutstanding)
		KPIRecord.ROA = float64(netIncome) / float64(totalAssets)
		KPIRecord.GrossMargin = float64(grossProfit) / float64(totalRevenue)
		KPIRecord.CurrentRatio = float64(totalCurrentAssets) / float64(totalCurrentLiabilities)
		KPIRecord.QuickRatio = float64(totalCurrentAssets - inventory) / float64(totalCurrentLiabilities)
		KPIRecord.CashRatio = float64(cashAndCashEquivalentsAtCarryingValue) / float64(totalCurrentLiabilities)
		KPIRecord.DebtToEquity = float64(totalLiabilities) / float64(totalShareholderEquity)
		KPIRecord.InterestCoverage = float64(ebit) / float64(interestExpense)
		KPIRecord.AssetTurnover = float64(totalRevenue) / float64(totalCurrentAssets)
		KPIRecord.InventoryTurnover = float64(costOfGoodsAndServicesSold) / float64(inventory)
		KPIRecord.ReceivablesTurnover = float64(totalRevenue) / float64(currentNetReceivables)
		KPIRecord.AccountsPayableTurnover = float64(costOfGoodsAndServicesSold) / float64(currentAccountsPayable)
		KPIRecord.DSO = float64(currentNetReceivables) / (float64(totalRevenue) / 365)
		KPIRecord.DIO = float64(inventory) / (float64(costOfGoodsAndServicesSold) / 365)
		KPIRecord.DPO = float64(currentAccountsPayable) / (float64(costOfGoodsAndServicesSold) / 365)
		
		dbManager.InsertKPI(KPIRecord, kpi_table, record.Symbol, record.FiscalDate)
		logmanager.Infof("Inserted KPI record for symbol %s and date %s", record.Symbol, record.FiscalDate)
		logmanager.Infof("KPI Record: %+v", KPIRecord)
	}
	return err
}	