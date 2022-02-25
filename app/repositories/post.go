package repositories

import (
	"gorm.io/gorm"

	"clean-arch-2/app/models"
)

type PostRepository struct {
	Conn *gorm.DB
}

func NewPostRepository(Conn *gorm.DB) PostRepository {
	return PostRepository{Conn}
}

func (r *PostRepository) Save(post *models.Post) (error) {
	err := r.Conn.Create(&post).Error

	return err 
}

func (r *PostRepository) FetchAll() ([]models.Post, error) {
	var posts []models.Post
	err := r.Conn.Find(&posts).Error

	return posts, err
}


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