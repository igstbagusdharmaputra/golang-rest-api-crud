package product

type Service interface {
	CreateProduct(input CreateProductInput) (Product, error)
	GetAllProduct() ([]Product, error)
	GetByIDProduct(ID int) (Product, error)
	UpdateProduct(ID int, input CreateProductInput) (Product, error)
	DeleteProduct(ID int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateProduct(input CreateProductInput) (Product, error) {
	newProduct := Product{}
	newProduct.Name = input.Name
	newProduct.Price = input.Price
	data, err := s.repository.Create(newProduct)
	if err != nil {
		return data, err
	}
	return data, nil
}
func (s *service) GetAllProduct() ([]Product, error) {
	data, err := s.repository.FindAll()
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *service) GetByIDProduct(id int) (Product, error) {
	data, err := s.repository.FindByID(id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *service) UpdateProduct(id int, input CreateProductInput) (Product, error) {
	product, err := s.repository.FindByID(id)
	if err != nil {
		return product, err
	}
	product.Name = input.Name
	product.Price = input.Price
	data, err := s.repository.Update(product)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *service) DeleteProduct(id int) (Product, error) {
	product, err := s.repository.FindByID(id)
	if err != nil {
		return product, err
	}

	data, err := s.repository.Delete(product)
	if err != nil {
		return data, err
	}
	return data, nil
}
