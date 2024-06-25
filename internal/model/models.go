package model

type Person struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

type Employer struct {
	ID      int    `json:"id" db:"id"`
	Company string `json:"company" db:"company"`
	Person  Person `json:"person" db:"person"`
}
