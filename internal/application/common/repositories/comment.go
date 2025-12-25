package repositories

import (
	"github.com/google/uuid"
	domain "github.com/literally_user/gozon/internal/domain/comment"
)

type CommentRepository interface {
	GetAllByUserUUID(uuid uuid.UUID) []domain.Comment
	GetAllByProductUUID(uuid uuid.UUID) []domain.Comment
	GetByUUID(uuid uuid.UUID) (domain.Comment, error)
	Create(comment domain.Comment) error
	Update(comment domain.Comment) error
	Remove(comment domain.Comment) error
}
