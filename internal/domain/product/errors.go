package product

import "errors"

var ErrTitleDoesntChanged = errors.New("product title doesn't changed")
var ErrTitleWrongFormat = errors.New("product title wrong format")

var ErrDescriptionDoesntChanged = errors.New("product description doesn't changed")
var ErrDescriptionWrongFormat = errors.New("product description wrong format")

var ErrTypeDoesntChanged = errors.New("product type doesn't changed")
var ErrPriceDoesntChanged = errors.New("product price doesn't changed")
var ErrRatingDoesntChanged = errors.New("product rating doesn't changed")
var ErrShadowRatingDoesntChanged = errors.New("product shadow rating doesn't changed")
var ErrCountDoesntChanged = errors.New("product count doesn't changed")
