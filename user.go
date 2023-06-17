package gorf

type User interface {
	GetFirstName() string
	GetLastName() string
	GetEmail() string
	ID() string
}

type BaseUser struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u *BaseUser) GetFirstName() string {
	return u.FirstName
}

func (u *BaseUser) GetLastName() string {
	return u.LastName
}

func (u *BaseUser) GetEmail() string {
	return u.Email
}

func (u *BaseUser) ID() string {
	return u.Id
}

func NewUser(id, firstName, lastName, email string) *BaseUser {
	return &BaseUser{id, email, firstName, lastName}
}
