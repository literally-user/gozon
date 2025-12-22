package publisher

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/domain/product"
)

type ProductCreatedEvent struct {
	Product product.Product
}

func (ProductCreatedEvent) Name() string {
	return "create.product"
}

type ProductRemovedEvent struct {
	Product product.Product
}

func (ProductRemovedEvent) Name() string {
	return "remove.product"
}

type ProductChangedTitleEvent struct {
	UUID     uuid.UUID
	OldTitle string
	NewTitle string
}

func (ProductChangedTitleEvent) Name() string {
	return "change.product.title"
}

type ProductChangedDescriptionEvent struct {
	UUID           uuid.UUID
	OldDescription string
	NewDescription string
}

func (ProductChangedDescriptionEvent) Name() string {
	return "change.product.description"
}

type ProductChangedPriceEvent struct {
	UUID     uuid.UUID
	OldPrice float64
	NewPrice float64
}

func (ProductChangedPriceEvent) Name() string {
	return "change.product.price"
}

type ProductChangedCountEvent struct {
	UUID     uuid.UUID
	OldCount int
	NewCount int
}

func (ProductChangedCountEvent) Name() string {
	return "change.product.count"
}

type ProductChangedRatingEvent struct {
	UUID      uuid.UUID
	OldRating float32
	NewRating float32
}

func (ProductChangedRatingEvent) Name() string {
	return "change.product.rating"
}

type ProductChangedShadowRatingEvent struct {
	UUID            uuid.UUID
	OldShadowRating float32
	NewShadowRating float32
}

func (ProductChangedShadowRatingEvent) Name() string {
	return "change.product.shadow_rating"
}

type ProductChangedTypeEvent struct {
	UUID    uuid.UUID
	OldType string
	NewType string
}

func (ProductChangedTypeEvent) Name() string {
	return "change.product.type"
}
