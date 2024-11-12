package user

type User struct {
	ID          string
	Email       Email
	Name        Name
	Preferences map[string]string
	Role        Role
}
