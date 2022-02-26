package repositories

import (
	"gorm.io/gorm"

	"clean-arch-2/app/models"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) UserRepository {
	return UserRepository{Conn}
}

func (r *UserRepository) Register(user *models.Users) error {
	err := r.Conn.Create(&user).Error

	return err
}

func (r *UserRepository) Login(
	userLogin *models.UserLoginInput,
) (*models.Users, error) {
	user := &models.Users{}

	result := r.Conn.
		Where("email = ?", userLogin.UsernameOrEmail).
		Or("username = ?", userLogin.UsernameOrEmail).
		First(&user)
	return user, result.Error
}

// func (r *UserRepository) FetchAll() ([]models.Users, error) {
// 	var users []models.Users
// 	err := r.Conn.Find(&users).Error

// 	return users, err
// }

// func NewPostRepository(Conn *gorm.DB) domains.PostRepository {
// 	return &PostRepository{Conn}
// }

// func Save(post *domains.Post) (*domains.Post, error) {
// 	ctx := context.Background()
// }

// func NewBookRepository(Conn *gorm.DB) domain.BookRepository {
// 	return &BookRepository{Conn}
// }

// func (m *BookRepository) Fetch(ctx context.Context) (res []domain.Book, err error) {
// 	var books []domain.Book
// 	m.Conn.Find(&books)

// 	return books, nil
// }
// func (m *BookRepository) GetByID(ctx context.Context, id string) (res domain.Book, err error) {
// 	var book domain.Book
// 	m.Conn.Where("id = ?", id).First(&book)

// 	return book, nil
// }
