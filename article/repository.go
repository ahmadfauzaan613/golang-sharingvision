package article

import "gorm.io/gorm"

type Repository interface {
	NewArticle(article Posts) (Posts, error)
	FindAll() ([]Posts, error)
	FindByID(ID int) (Posts, error)
	UpdateArticle(article Posts) (Posts, error)
	ArticleDelete(article Posts) (Posts, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Menambahkan Article
func (r *repository) NewArticle(article Posts) (Posts, error) {
	err := r.db.Create(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

// Menampilkan seluruh article
func (r *repository) FindAll() ([]Posts, error) {
	var post []Posts

	err := r.db.Find(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil
}

// Menampilkan Berdasarkan ID article
func (r *repository) FindByID(ID int) (Posts, error) {
	var post Posts

	err := r.db.Where("id = ?", ID).Find(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

// Update Article
func (r *repository) UpdateArticle(article Posts) (Posts, error) {
	err := r.db.Save(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

// Delete Article
func (r *repository) ArticleDelete(article Posts) (Posts, error) {
	err := r.db.Delete(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}
