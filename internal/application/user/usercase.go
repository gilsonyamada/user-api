package user

import (
    "context"
    "errors"
    "internal/domain/user"
)

type UseCase struct {
    Repo user.UserRepository
}

// RegisterUser registers a new user with the given details.
func (uc *UseCase) RegisterUser(ctx context.Context, email, firstName, lastName string, role user.Role) (*user.User, error) {
    // Business logic for user registration
    newUser := &user.User{
        Email:     email,
        Name: {
			FirstName: firstName,
			LastName:  lastName
		},
        Role:      role,
    }
    if err := uc.Repo.Save(newUser); err != nil {
        return nil, err
    }
    return newUser, nil
}

// GetUserProfile retrieves a user's profile by ID.
func (uc *UseCase) GetUserProfile(ctx context.Context, id string) (*user.User, error) {
    return uc.Repo.FindByID(id)
}