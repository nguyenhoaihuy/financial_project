package models

type MissingKPIRecord struct {
	Symbol     string
	FiscalDate string
}

type KPI struct {
	Symbol          string
	FiscalDate      string
	RevenueGrowth   float64
	NetProfitMargin float64
	EPS             float64
	ROE             float64
	ROA             float64
	GrossMargin     float64
	CurrentRatio    float64
	CashRatio	   float64
	QuickRatio      float64
	DebtToEquity    float64
	InterestCoverage float64
	AssetTurnover   float64
	InventoryTurnover float64
	ReceivablesTurnover float64
	AccountsPayableTurnover float64
	DSO             float64
	DIO             float64
	DPO             float64
	PERatio         float64
	PBRatio         float64
	DividendYield   float64
	EVToEBITDA      float64
}