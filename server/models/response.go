package models

// JSON response
type Quote struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Time     string `json:"time"`
}

type User struct {
	Name      string `json:"name"`
	BdayMonth string `json:"bdaymonth"`
}

type ReturnUser struct {
	Name         string `json:"name"`
	BdayMonth    string `json:"bdaymonth"`
	NumericMonth int    `json:"numericmonth"`
}

type IncomingData map[string][]User

type ReturnData map[string][]ReturnUser
