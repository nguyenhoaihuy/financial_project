package models

type IncomeStatement struct {
	// Symbol                         string `json:"symbol"`
	FiscalDateEnding               string `json:"fiscalDateEnding"`
	ReportedCurrency               string `json:"reportedCurrency"`
	TotalRevenue                   string    `json:"totalRevenue"`
	GrossProfit                    string    `json:"grossProfit"`
	CostOfRevenue                  string    `json:"costOfRevenue"`
	CostOfGoodsAndServicesSold     string   `json:"costofGoodsAndServicesSold"`
	OperatingIncome                string    `json:"operatingIncome"`
	SellingGeneralAndAdministrative string   `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment         string   `json:"researchAndDevelopment"`
	OperatingExpenses              string    `json:"operatingExpenses"`
	InvestmentIncomeNet            string   `json:"investmentIncomeNet"`
	NetInterestIncome              string   `json:"netInterestIncome"`
	InterestIncome                 string   `json:"interestIncome"`
	InterestExpense                string   `json:"interestExpense"`
	NonInterestIncome              string   `json:"nonInterestIncome"`
	OtherNonOperatingIncome        string   `json:"otherNonOperatingIncome"`
	Depreciation                   string   `json:"depreciation"`
	DepreciationAndAmortization    string    `json:"depreciationAndAmortization"`
	IncomeBeforeTax                string    `json:"incomeBeforeTax"`
	IncomeTaxExpense               string    `json:"incomeTaxExpense"`
	InterestAndDebtExpense         string   `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations string `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax    string   `json:"comprehensiveIncomeNetOfTax"`
	Ebit                           string    `json:"ebit"`
	Ebitda                         string    `json:"ebitda"`
	NetIncome                      string    `json:"netIncome"`
}

type BalanceSheet struct {
	// Symbol                                string `json:"symbol"`
	FiscalDateEnding                      string `json:"fiscalDateEnding"`
	ReportedCurrency                      string `json:"reportedCurrency"`
	TotalAssets                           string    `json:"totalAssets"`
	TotalCurrentAssets                    string    `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue string    `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments           string    `json:"cashAndShortTermInvestments"`
	Inventory                             string    `json:"inventory"`
	CurrentNetReceivables                 string    `json:"currentNetReceivables"`
	TotalNonCurrentAssets                 string    `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                string    `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE string   `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                      string   `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill     string   `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                              string   `json:"goodwill"`
	Investments                           string    `json:"investments"`
	LongTermInvestments                   string    `json:"longTermInvestments"`
	ShortTermInvestments                  string    `json:"shortTermInvestments"`
	OtherCurrentAssets                    string    `json:"otherCurrentAssets"`
	OtherNonCurrentAssets                 string    `json:"otherNonCurrentAssets"`
	TotalLiabilities                      string    `json:"totalLiabilities"`
	TotalCurrentLiabilities               string    `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                string    `json:"currentAccountsPayable"`
	DeferredRevenue                       string   `json:"deferredRevenue"`
	CurrentDebt                           string   `json:"currentDebt"`
	ShortTermDebt                         string    `json:"shortTermDebt"`
	TotalNonCurrentLiabilities            string    `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations               string   `json:"capitalLeaseObligations"`
	LongTermDebt                          string    `json:"longTermDebt"`
	CurrentLongTermDebt                   string    `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                string    `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                string    `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities               string    `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities            string    `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                string    `json:"totalShareholderEquity"`
	TreasuryStock                         string   `json:"treasuryStock"`
	RetainedEarnings                      string    `json:"retainedEarnings"`
	CommonStock                           string    `json:"commonStock"`
	CommonStockSharesOutstanding          string    `json:"commonStockSharesOutstanding"`
}

type CashFlow struct {
	// Symbol                                string `json:"symbol"`
	FiscalDateEnding                      string `json:"fiscalDateEnding"`
	ReportedCurrency                      string `json:"reportedCurrency"`
	OperatingCashflow                     string    `json:"operatingCashflow"`
	PaymentsForOperatingActivities        string   `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities       string   `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities          string   `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets               string   `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization  string    `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                   string    `json:"capitalExpenditures"`
	ChangeInReceivables                   string    `json:"changeInReceivables"`
	ChangeInInventory                     string    `json:"changeInInventory"`
	ProfitLoss                            string    `json:"profitLoss"`
	CashflowFromInvestment                string    `json:"cashflowFromInvestment"`
	CashflowFromFinancing                 string    `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt string   `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock    string    `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity         string    `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock string   `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                        string    `json:"dividendPayout"`
	ProceedsFromIssuanceOfCommonStock     string   `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebt    string   `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ChangeInCashAndCashEquivalents        string   `json:"changeInCashAndCashEquivalents"`
	NetIncome                             string    `json:"netIncome"`
}


