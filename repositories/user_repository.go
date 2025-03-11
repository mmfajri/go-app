package repositories

import (
	"go-app/models"
	"log"

	"gorm.io/gorm"
)

//Construct -> Build the "Class"
type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	AddUser(models.User) (models.User, error)
	GetUser(int) (models.User, error)
	GetByEmail(string) (models.User, error)
	GetAllUser() ([]models.User, error)
	UpdateUser(models.User) (models.User, error) 
	DeleteUser(models.User) (models.User, error)
	SoftDeleteUser(uint) (int, error)
	Migrate() error
}

//Construct -> return new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) Migrate() error {
	log.Print("[UserRepository]...Migrate")
	return u.DB.AutoMigrate(&models.User{})
}

func (u userRepository) GetUser(id int) (user models.User, err error) {
	return user, u.DB.First(&user, id).Error
}

func (u userRepository) GetByEmail(email string) (user models.User, err error) {
	return user, u.DB.First(&user, "email=?", email).Error
}

func (u userRepository) GetAllUser() (users []models.User, err error) {
	return users, u.DB.Find(&users).Error 
}

func (u userRepository) AddUser(user models.User) (models.User,  error) {
	return user, u.DB.Create(&user).Error
}

func (u userRepository) UpdateUser(user models.User) (models.User,error) {
	if err := u.DB.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, u.DB.Model(&user).Updates(&user).Error
}

func (u userRepository) DeleteUser(user models.User) (models.User,error) {
	if err := u.DB.First(&user, user.ID).Error; err != nil {
		return user, err
	}
	return user, u.DB.Delete(&user).Error
}

func (u userRepository) SoftDeleteUser(id uint) (int, error) {
	result := u.DB.Model(&models.User{}).Where("id = ?", id).Update("is_deleted = ? ", true)
	return int(result.RowsAffected), result.Error 
}

