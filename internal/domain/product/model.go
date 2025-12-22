package product

import "github.com/google/uuid"

type Product struct {
	UUID uuid.UUID

	Title       string
	Description string
	Type        string

	Count int
	Price float64

	Rating       float32
	ShadowRating float32
}

func NewProduct(title, description, productType string, price float64) (Product, error) {
	var err error
	var product Product

	err = product.ChangeTitle(title)
	if err != nil {
		return Product{}, err
	}

	err = product.ChangeDescription(description)
	if err != nil {
		return Product{}, err
	}

	err = product.ChangeType(productType)
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

func (p *Product) ChangeType(productType string) error {
	if productType == p.Type {
		return ErrTypeDoesntChanged
	}

	p.Type = productType
	return nil
}

func (p *Product) ChangePrice(price float64) error {
	if price == p.Price {
		return ErrPriceDoesntChanged
	}

	p.Price = price
	return nil
}

func (p *Product) ChangeRating(rating float32) error {
	if rating == p.Rating {
		return ErrRatingDoesntChanged
	}

	p.Rating = rating
	return nil
}

func (p *Product) ChangeShadowRating(rating float32) error {
	if rating == p.ShadowRating {
		return ErrShadowRatingDoesntChanged
	}

	p.ShadowRating = rating
	return nil
}

func (p *Product) ChangeCount(count int) error {
	if count == p.Count {
		return ErrCountDoesntChanged
	}

	p.Count = count
	return nil
}
