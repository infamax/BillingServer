package models

type UserBalance struct {
	ID      int     `json:"ID"`
	Balance float64 `json:"Balance"`
}

type Payment struct {
	ID      int     `json:"ID"`
	UserID  int     `json:"UserID"`
	Payment float64 `json:"Payment"`
	Date    string  `json:"Date"`
	Reason  string  `json:"Reason"`
}

type Accrual struct {
	ID      int     `json:"ID"`
	UserID  int     `json:"UserID"`
	Accrual float64 `json:"Accrual"`
	Date    string  `json:"Date"`
	Reason  string  `json:"Reason"`
}
