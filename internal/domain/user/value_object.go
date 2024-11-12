package user

type Email string

func (e Email) IsValid() bool {
	return true
}

type Role string

const (
	RoleAdmin    Role = "SystemAdmin"
	RoleReviewer Role = "Reviewer"
	RoleLeaner   Role = "Leaner"
)

func (r Role) IsValid() bool {
	return true
}

type Name struct {
	FirstName string
	LastName  string
}

func (n Name) IsValid() bool {
	return true
}
