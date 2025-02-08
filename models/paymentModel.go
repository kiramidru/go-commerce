package models

type Payment struct {
	Amount          string `json:"amount"`
	Currency        string `json: "currency"`
	First_name      string `json: "first_name"`
	Last_name       string `json: "last_name"`
	Email           string `json: "email"`
	Callback_URL    string `json: "callback_url"`
	Transaction_ref string `json: "transaction_ref"`
	Title           string `json: "title"`
	Description     string `json: "description"`
	Logo            string `json: "logo"`
}
