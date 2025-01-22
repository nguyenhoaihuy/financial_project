package db

import (
	"strconv"
	"fmt"
	"financial_project/models"
	"financial_project/logmanager"
)

// Convert values to thousands
func convert(value string) int64 {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0 // Default to 0 if conversion fails
	}
	return v / 1000
}

func (m *DBManager) RecordExists(table, symbol, fiscalDateEnding string) (bool, error) {
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE symbol = ? AND fiscal_date_ending = ?)", table)
	var exists bool
	err := m.DB.QueryRow(query, symbol, fiscalDateEnding).Scan(&exists)
	if err != nil {
		logmanager.Errorf("Error checking record exists: %v", err)
	}
	return exists, err
}

func (m *DBManager) InsertIncomeStatement(report models.IncomeStatement, tableName string, symbol string) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			symbol,
			fiscal_date_ending,
			reported_currency,
			total_revenue,
			gross_profit,
			cost_of_revenue,
			cost_of_goods_and_services_sold,
			operating_income,
			selling_general_and_administrative,
			research_and_development,
			operating_expenses,
			investment_income_net,
			net_interest_income,
			interest_income,
			interest_expense,
			non_interest_income,
			other_non_operating_income,
			depreciation,
			depreciation_and_amortization,
			income_before_tax,
			income_tax_expense,
			interest_and_debt_expense,
			net_income_from_continuing_operations,
			comprehensive_income_net_of_tax,
			ebit,
			ebitda,
			net_income
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, tableName)
	
	_, err := m.DB.Exec(
		query, 
		symbol, 
		report.FiscalDateEnding,
		report.ReportedCurrency,
		convert(report.TotalRevenue),
		convert(report.GrossProfit),
		convert(report.CostOfRevenue),
		convert(report.CostOfGoodsAndServicesSold),
		convert(report.OperatingIncome),
		convert(report.SellingGeneralAndAdministrative),
		convert(report.ResearchAndDevelopment),
		convert(report.OperatingExpenses),
		convert(report.InvestmentIncomeNet),
		convert(report.NetInterestIncome),
		convert(report.InterestIncome),
		convert(report.InterestExpense),
		convert(report.NonInterestIncome),
		convert(report.OtherNonOperatingIncome),
		convert(report.Depreciation),
		convert(report.DepreciationAndAmortization),
		convert(report.IncomeBeforeTax),
		convert(report.IncomeTaxExpense),
		convert(report.InterestAndDebtExpense),
		convert(report.NetIncomeFromContinuingOperations),
		convert(report.ComprehensiveIncomeNetOfTax),
		convert(report.Ebit),
		convert(report.Ebitda),
		convert(report.NetIncome))
	return err
}

// Similarly, create functions for BalanceSheet and CashFlow
func (m *DBManager) InsertBalanceSheet(report models.BalanceSheet, tableName string, symbol string) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			symbol,
			fiscal_date_ending,
			reported_currency,
			total_assets,
			total_current_assets,
			cash_and_cash_equivalents,
			cash_and_short_term_investments,
			inventory,
			current_net_receivables,
			total_non_current_assets,
			property_plant_equipment,
			accumulated_depreciation_amortization_ppe,
			intangible_assets,
			intangible_assets_excluding_goodwill,
			goodwill,
			investments,
			long_term_investments,
			short_term_investments,
			other_current_assets,
			other_non_current_assets,
			total_liabilities,
			total_current_liabilities,
			current_accounts_payable,
			deferred_revenue,
			current_debt,
			short_term_debt,
			total_non_current_liabilities,
			capital_lease_obligations,
			long_term_debt,
			current_long_term_debt,
			long_term_debt_noncurrent,
			short_long_term_debt_total,
			other_current_liabilities,
			other_non_current_liabilities,
			total_shareholder_equity,
			treasury_stock,
			retained_earnings,
			common_stock,
			common_stock_shares_outstanding
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, tableName)
	
	_, err := m.DB.Exec(
		query, 
		symbol, 
		report.FiscalDateEnding,
		report.ReportedCurrency,
		convert(report.TotalAssets),
		convert(report.TotalCurrentAssets),
		convert(report.CashAndCashEquivalentsAtCarryingValue),
		convert(report.CashAndShortTermInvestments),
		convert(report.Inventory),
		convert(report.CurrentNetReceivables),
		convert(report.TotalNonCurrentAssets),
		convert(report.PropertyPlantEquipment),
		convert(report.AccumulatedDepreciationAmortizationPPE),
		convert(report.IntangibleAssets),
		convert(report.IntangibleAssetsExcludingGoodwill),
		convert(report.Goodwill),
		convert(report.Investments),
		convert(report.LongTermInvestments),
		convert(report.ShortTermInvestments),
		convert(report.OtherCurrentAssets),
		convert(report.OtherNonCurrentAssets),
		convert(report.TotalLiabilities),
		convert(report.TotalCurrentLiabilities),
		convert(report.CurrentAccountsPayable),
		convert(report.DeferredRevenue),
		convert(report.CurrentDebt),
		convert(report.ShortTermDebt),
		convert(report.TotalNonCurrentLiabilities),
		convert(report.CapitalLeaseObligations),
		convert(report.LongTermDebt),
		convert(report.CurrentLongTermDebt),
		convert(report.LongTermDebtNoncurrent),
		convert(report.ShortLongTermDebtTotal),
		convert(report.OtherCurrentLiabilities),
		convert(report.OtherNonCurrentLiabilities),
		convert(report.TotalShareholderEquity),
		convert(report.TreasuryStock),
		convert(report.RetainedEarnings),
		convert(report.CommonStock),
		convert(report.CommonStockSharesOutstanding))
	return err
}

func (m *DBManager) InsertCashFlow(report models.CashFlow, tableName string, symbol string) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			symbol,
			fiscal_date_ending,
			reported_currency,
			operating_cashflow,
			payments_for_operating_activities,
			proceeds_from_operating_activities,
			change_in_operating_liabilities,
			change_in_operating_assets,
			depreciation_depletion_and_amortization,
			capital_expenditures,
			change_in_receivables,
			change_in_inventory,
			profit_loss,
			cashflow_from_investment,
			cashflow_from_financing,
			proceeds_from_repayments_of_short_term_debt,
			payments_for_repurchase_of_common_stock,
			payments_for_repurchase_of_equity,
			payments_for_repurchase_of_preferred_stock,
			dividend_payout,
			proceeds_from_issuance_of_common_stock,
			proceeds_from_issuance_of_long_term_debt,
			change_in_cash_and_cash_equivalents,
			net_income
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, tableName)
	
	_, err := m.DB.Exec(
		query, 
		symbol, 
		report.FiscalDateEnding,
		report.ReportedCurrency,
		convert(report.OperatingCashflow),
		convert(report.PaymentsForOperatingActivities),
		convert(report.ProceedsFromOperatingActivities),
		convert(report.ChangeInOperatingLiabilities),
		convert(report.ChangeInOperatingAssets),
		convert(report.DepreciationDepletionAndAmortization),
		convert(report.CapitalExpenditures),
		convert(report.ChangeInReceivables),
		convert(report.ChangeInInventory),
		convert(report.ProfitLoss),
		convert(report.CashflowFromInvestment),
		convert(report.CashflowFromFinancing),
		convert(report.ProceedsFromRepaymentsOfShortTermDebt),
		convert(report.PaymentsForRepurchaseOfCommonStock),
		convert(report.PaymentsForRepurchaseOfEquity),
		convert(report.PaymentsForRepurchaseOfPreferredStock),
		convert(report.DividendPayout),
		convert(report.ProceedsFromIssuanceOfCommonStock),
		convert(report.ProceedsFromIssuanceOfLongTermDebt),
		convert(report.ChangeInCashAndCashEquivalents),
		convert(report.NetIncome))
	return err
}

func (m *DBManager) CompanyExists(symbol string) (bool, error) {
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM company WHERE symbol = ?)")
	var exists bool
	err := m.DB.QueryRow(query, symbol).Scan(&exists)
	if err != nil {
		logmanager.Errorf("Error checking company exists: %v", err)
	}
	return exists, err
}

func (m *DBManager) InsertCompany(report models.Company) error {
	query := fmt.Sprintf(`
		INSERT INTO company (
			symbol,
			asset_type,
			name,
			description,
			cik,
			exchange,
			currency,
			country,
			sector,
			industry,
			address,
			official_site,
			fiscal_year_end
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	_, err := m.DB.Exec(
		query, 
		report.Symbol,
		report.AssetType,
		report.Name,
		report.Description,
		report.CIK,
		report.Exchange,
		report.Currency,
		report.Country,
		report.Sector,
		report.Industry,
		report.Address,
		report.OfficialSite,
		report.FiscalYearEnd)
	return err
}

func (m *DBManager) DeleteCompany(symbol string) error {
	query := "DELETE FROM company WHERE symbol = ?"
    result, err := m.DB.Exec(query, symbol)
    if err != nil {
        return fmt.Errorf("failed to delete record: %v", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to fetch affected rows: %v", err)
    }

    fmt.Printf("Deleted %d record(s)\n", rowsAffected)
    return nil
}

