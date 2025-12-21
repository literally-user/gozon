package product

import "errors"

var ErrTitleDoesntChanged = errors.New("title doesn't changed")
var ErrDescriptionDoesntChanged error = errors.New("description doesn't changed")
var ErrProductTypeDoesntChanged error = errors.New("product type doesn't changed")
var ErrPriceDoesntChanged error = errors.New("price doesn't changed")
var ErrProductRatingDoesntChanged error = errors.New("product rating doesn't changed")
