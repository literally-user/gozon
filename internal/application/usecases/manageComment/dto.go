package manageComment

import "github.com/google/uuid"

type DTO struct {
	Rate    float32
	Content string

	ProductUUID uuid.UUID
	UserUUID    uuid.UUID
}
