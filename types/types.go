package types

type User struct {
	ID int `json:"user_id"`
	User string `json:"user_name"`
	First string `json:"first_name"`
	Last string `json:"last_name"`
}

type BankAcct struct {
	ID int `json:"user_id"`
	Balance string `json:"balance"`
	Deposits []string
	Withdrawls []string
}

type CreditCard struct {
	ID int `json:"user_id"`
	Balance string `json:"balance"`
	CreditAvalible string `json:"credit_avalible"`
	Apr string `json:"apr"`
}