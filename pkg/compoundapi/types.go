package compoundapi

type CToken struct {
	BorrowRate struct {
		Value string `json:"value,omitempty"`
	} `json:"borrow_rate,omitempty"`
	Cash struct {
		Value string `json:"value,omitempty"`
	} `json:"cash,omitempty"`
	CollateralFactor struct {
		Value string `json:"value,omitempty"`
	} `json:"collateral_factor,omitempty"`
	ExchangeRate struct {
		Value string `json:"value,omitempty"`
	} `json:"exchange_rate,omitempty"`
	InterestRateModelAddress string `json:"interest_rate_model_address,omitempty"`
	Name                     string `json:"name,omitempty"`
	NumberOfBorrowers        int    `json:"number_of_borrowers,omitempty"`
	NumberOfSuppliers        int    `json:"number_of_suppliers,omitempty"`
	Reserves                 struct {
		Value string `json:"value,omitempty"`
	} `json:"reserves,omitempty"`
	SupplyRate struct {
		Value string `json:"value,omitempty"`
	} `json:"supply_rate,omitempty"`
	Symbol       string `json:"symbol,omitempty"`
	TokenAddress string `json:"token_address,omitempty"`
	TotalBorrows struct {
		Value string `json:"value,omitempty"`
	} `json:"total_borrows,omitempty"`
	TotalSupply struct {
		Value string `json:"value,omitempty"`
	} `json:"total_supply,omitempty"`
	UnderlyingAddress string `json:"underlying_address,omitempty"`
	UnderlyingName    string `json:"underlying_name,omitempty"`
	UnderlyingPrice   struct {
		Value string `json:"value,omitempty"`
	} `json:"underlying_price,omitempty"`
	UnderlyingSymbol string `json:"underlying_symbol,omitempty"`
}

type CTokens []struct {
	BorrowRate struct {
		Value string `json:"value,omitempty"`
	} `json:"borrow_rate,omitempty"`
	Cash struct {
		Value string `json:"value,omitempty"`
	} `json:"cash,omitempty"`
	CollateralFactor struct {
		Value string `json:"value,omitempty"`
	} `json:"collateral_factor,omitempty"`
	ExchangeRate struct {
		Value string `json:"value,omitempty"`
	} `json:"exchange_rate,omitempty"`
	InterestRateModelAddress string `json:"interest_rate_model_address,omitempty"`
	Name                     string `json:"name,omitempty"`
	NumberOfBorrowers        int    `json:"number_of_borrowers,omitempty"`
	NumberOfSuppliers        int    `json:"number_of_suppliers,omitempty"`
	Reserves                 struct {
		Value string `json:"value,omitempty"`
	} `json:"reserves,omitempty"`
	SupplyRate struct {
		Value string `json:"value,omitempty"`
	} `json:"supply_rate,omitempty"`
	Symbol       string `json:"symbol,omitempty"`
	TokenAddress string `json:"token_address,omitempty"`
	TotalBorrows struct {
		Value string `json:"value,omitempty"`
	} `json:"total_borrows,omitempty"`
	TotalSupply struct {
		Value string `json:"value,omitempty"`
	} `json:"total_supply,omitempty"`
	UnderlyingAddress string `json:"underlying_address,omitempty"`
	UnderlyingName    string `json:"underlying_name,omitempty"`
	UnderlyingPrice   struct {
		Value string `json:"value,omitempty"`
	} `json:"underlying_price,omitempty"`
	UnderlyingSymbol string `json:"underlying_symbol,omitempty"`
}
