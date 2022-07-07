package product

type CreateProductInput struct {
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
}
