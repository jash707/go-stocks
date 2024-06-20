package models

type Stock struct {
	StockID int64  `json:"stocksid"`
	Name    string `json:"name"`
	Price   int64  `json:"price"`
	Company string `json:"company"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
