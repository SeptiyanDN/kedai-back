package categories

type Services interface {
	SaveCategory(input AddCategoryInput) (Category, error)
	FindAll() ([]Category, error)
}

type services struct {
	repository Repository
}

func NewServices(repository Repository) *services {
	return &services{repository}
}

func (s *services) SaveCategory(input AddCategoryInput) (Category, error) {
	category := Category{}
	category.Category_name = input.Category_name
	category.Business_id = input.Business_id
	newCategory, err := s.repository.Save(category)
	if err != nil {
		return category, err
	}
	return newCategory, nil
}

func (s *services) FindAll() ([]Category, error) {
	Category, err := s.repository.FindAll()
	if err != nil {
		return Category, err
	}
	return Category, nil
}
