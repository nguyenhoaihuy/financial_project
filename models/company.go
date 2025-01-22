package models

type Company struct {
    Symbol                     string  `json:"Symbol"`
    AssetType                  string  `json:"AssetType"`
    Name                       string  `json:"Name"`
    Description                string  `json:"Description"`
    CIK                        string  `json:"CIK"`
    Exchange                   string  `json:"Exchange"`
    Currency                   string  `json:"Currency"`
    Country                    string  `json:"Country"`
    Sector                     string  `json:"Sector"`
    Industry                   string  `json:"Industry"`
    Address                    string  `json:"Address"`
    OfficialSite               string  `json:"OfficialSite"`
    FiscalYearEnd              string  `json:"FiscalYearEnd"`
}