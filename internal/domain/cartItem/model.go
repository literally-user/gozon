package cartItem

import "github.com/google/uuid"

type CartItem struct {
	UUID uuid.UUID

	UserUUID    uuid.UUID
	ProductUUID uuid.UUID
}

func NewCartItem(userUUID, productUUID uuid.UUID) (CartItem, error) {
	return CartItem{
		UUID:        uuid.New(),
		UserUUID:    userUUID,
		ProductUUID: productUUID,
	}, nil
}
