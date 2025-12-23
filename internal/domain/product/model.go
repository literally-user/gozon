package product

import "github.com/google/uuid"

type Product struct {
	UUID uuid.UUID

	title       string
	description string
	productType string

	count int
	price float64

	rating       float32
	shadowRating float32
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
	if len(title) >= 30 {
		return ErrTitleWrongFormat
	}
	if title == p.title {
		return ErrTitleDoesntChanged
	}

	p.title = title
	return nil
}

func (p *Product) ChangeDescription(description string) error {
	if len(description) >= 500 {
		return ErrDescriptionWrongFormat
	}
	if description == p.description {
		return ErrDescriptionDoesntChanged
	}

	p.description = description
	return nil
}

func (p *Product) ChangeType(productType string) error {
	if productType == p.productType {
		return ErrTypeDoesntChanged
	}

	p.productType = productType
	return nil
}

func (p *Product) ChangePrice(price float64) error {
	if price == p.price {
		return ErrPriceDoesntChanged
	}

	p.price = price
	return nil
}

func (p *Product) ChangeRating(rating float32) error {
	if rating == p.rating {
		return ErrRatingDoesntChanged
	}

	p.rating = rating
	return nil
}

func (p *Product) ChangeShadowRating(rating float32) error {
	if rating == p.shadowRating {
		return ErrShadowRatingDoesntChanged
	}

	p.shadowRating = rating
	return nil
}

func (p *Product) ChangeCount(count int) error {
	if count == p.count {
		return ErrCountDoesntChanged
	}

	p.count = count
	return nil
}

func (p *Product) ProductTitle() string { return p.title }

func (p *Product) ProductDescription() string { return p.description }

func (p *Product) ProductPrice() float64 { return p.price }

func (p *Product) ProductType() string { return p.productType }

func (p *Product) ProductCount() int {
	return p.count
}

func (p *Product) ProductShadowRating() float32 {
	return p.shadowRating
}

func (p *Product) ProductRating() float32 {
	return p.rating
}
