package postgres

import (
    "internal/domain/user"
    "gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Save(user *user.User) error {
    return r.db.Create(user).Error
}

func (r *UserRepository) FindByID(id string) (*user.User, error) {
    var u user.User
    if err := r.db.First(&u, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &u, nil
}

func (r *UserRepository) Delete(id string) error {
    return r.db.Delete(&user.User{}, "id = ?", id).Error
}