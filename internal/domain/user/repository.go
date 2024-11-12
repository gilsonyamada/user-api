package user

type UserRepository interface {
    Save(user *User) error
    FindByID(id string) (*User, error)
    Delete(id string) error
}