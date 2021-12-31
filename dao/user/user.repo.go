package user

import "gorm.io/gorm"

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Create(user *User) error {
	return u.db.Create(&user).Error
}

func (u *UserRepo) Update(user *User) error {
	return u.db.Updates(&user).Error
}

func (u *UserRepo) GetByID(id uint) (user *User, err error) {
	err = u.db.Where("id", id).First(&user).Error
	return
}

func (u *UserRepo) GetByAccount(account string) (user *User, err error) {
	err = u.db.Where("account", account).First(&user).Error
	return
}
