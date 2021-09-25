package article

type Service interface {
	FindAll() ([]Posts, error)
	NewArticle(articleRequest ArticleRequest) (Posts, error)
	FindByID(ID int) (Posts, error)
	UpdateArticle(ID int, articleRequest ArticleRequest) (Posts, error)
	ArticleDelete(ID int) (Posts, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Posts, error) {
	post, err := s.repository.FindAll()
	if err != nil {
		return post, err
	}

	return post, nil

}

func (s *service) FindByID(ID int) (Posts, error) {
	post, err := s.repository.FindByID(ID)
	if err != nil {
		return post, err
	}

	return post, nil

}

// Menambahkan Article
func (s *service) NewArticle(articleRequest ArticleRequest) (Posts, error) {
	posts := Posts{
		Title:    articleRequest.Title,
		Content:  articleRequest.Content,
		Category: articleRequest.Category,
		Status:   articleRequest.Status,
	}

	newArticles, err := s.repository.NewArticle(posts)
	return newArticles, err
}

func (s *service) UpdateArticle(ID int, articleRequest ArticleRequest) (Posts, error) {
	articels, err := s.repository.FindByID(ID)

	articels.Title = articleRequest.Title
	articels.Content = articleRequest.Content
	articels.Category = articleRequest.Category
	articels.Status = articleRequest.Status

	newArticles, err := s.repository.UpdateArticle(articels)
	return newArticles, err
}

func (s *service) ArticleDelete(ID int) (Posts, error) {
	articels, err := s.repository.FindByID(ID)
	newArticles, err := s.repository.ArticleDelete(articels)
	return newArticles, err
}
