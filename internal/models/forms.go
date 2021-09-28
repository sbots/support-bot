package models

type SignUpForm struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
	Company  string `json:"company"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type SignInForm struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type NewCompanyForm struct {
	Name string `json:"name"`
}
