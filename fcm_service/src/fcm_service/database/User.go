package database

type User struct {
	Phone string	`json:"phone"`
	Tokens []string	`json:"tokens"`
}
