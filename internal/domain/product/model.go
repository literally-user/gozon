package product

import "github.com/google/uuid"

type Product struct {
	UUID        uuid.UUID
	Title       string
	Description string
	ProductType string

	Price  float32
	Rating float32
}

func NewProduct(title, description, productType string, price float32) (Product, error) {
	var err error

	product := Product{
		UUID:   uuid.New(),
		Rating: 0.0,
	}

	err = product.ChangeTitle(title)
	if err != nil {
		return Product{}, err
	}

	err = product.ChangeDescription(description)
	if err != nil {
		return Product{}, err
	}

	err = product.ChangeProductType(productType)
	if err != nil {
		return Product{}, err
	}

	err = product.ChangePrice(price)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (p *Product) ChangeTitle(title string) error {
	if title == p.Title {
		return ErrTitleDoesntChanged
	}

	p.Title = title
	return nil
}

func (p *Product) ChangeDescription(description string) error {
	if description == p.Description {
		return ErrDescriptionDoesntChanged
	}

	p.Description = description
	return nil
}

func (p *Product) ChangeProductType(productType string) error {
	if productType == p.ProductType {
		return ErrProductTypeDoesntChanged
	}

	p.ProductType = productType
	return nil
}

func (p *Product) ChangePrice(price float32) error {
	if price == p.Price {
		return ErrPriceDoesntChanged
	}

	p.Price = price
	return nil
}

func (p *Product) ChangeProductRating(rating float32) error {
	if rating == p.Rating {
		return ErrProductRatingDoesntChanged
	}

	p.Rating = rating
	return nil
}
