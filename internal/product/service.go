package product

type Service interface {
	CreateProduct(Product *Product) (*Product, error)
	GetProduct(id string) (*Product, error)
	ListProducts() ([]*Product, error)
	GetProductByCategory(category string) ([]*Product, error)
}

type ProductService struct {
	repo      Repository
}

func NewProductService(repo Repository) Service {
	return &ProductService{
		repo:      repo,
	}
}

func (s *ProductService) CreateProduct(Product *Product) (*Product, error) {
	return Product, s.repo.Create(Product)
}

func (s *ProductService) GetProduct(id string) (*Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) ListProducts() ([]*Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) GetProductByCategory(category string) ([]*Product, error) {
	return s.repo.GetByCategory(category)
}