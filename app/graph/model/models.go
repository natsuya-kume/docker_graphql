package model

type Service struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (Service) IsNode() {}

type ServiceAccount struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	ServiceID string `json:"service"`
}

func (ServiceAccount) IsNode() {}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ServiceID string `json:"service"`
}

func (User) IsNode() {}

type PersonalTag struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"review"`
}

func (PersonalTag) IsNode() {}

type ReviewTag struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"review"`
}

func (ReviewTag) IsNode() {}
