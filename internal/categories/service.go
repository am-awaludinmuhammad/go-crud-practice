package categories

type CategoryService interface {
	CreateCategory(category *Category) error
	GetCategoryByID(id uint) (*Category, error)
	GetAllCategories() ([]Category, error)
	UpdateCategory(id uint, updatedCategory *Category) error
	DestroyCategory(id uint) error
}

type categoryService struct {
	repository CategoryRepository
}

func NewCategoryService(repo CategoryRepository) CategoryService {
	return &categoryService{repository: repo}
}

func (s *categoryService) GetAllCategories() ([]Category, error) {
	return s.repository.GetAllCategories()
}

func (s *categoryService) CreateCategory(category *Category) error {
	return s.repository.CreateCategory(category)
}

func (s *categoryService) GetCategoryByID(id uint) (*Category, error) {
	return s.repository.GetCategoryByID(id)
}

func (s *categoryService) DestroyCategory(id uint) error {
	return s.repository.DestroyCategory(id)
}

func (s *categoryService) UpdateCategory(id uint, updatedCategory *Category) error {
	return s.repository.UpdateCategory(id, updatedCategory)
}
