package cartItem

import "github.com/google/uuid"

type CartItem struct {
	uuid uuid.UUID

	userUUID    uuid.UUID
	productUUID uuid.UUID
}

func NewCartItem(userUUID, productUUID uuid.UUID) (CartItem, error) {
	return CartItem{
		uuid:        uuid.New(),
		userUUID:    userUUID,
		productUUID: productUUID,
	}, nil
}
