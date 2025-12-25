package manageComment

import (
	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	domain "github.com/literally_user/gozon/internal/domain/comment"
)

type GetAllCommentsInteractor struct {
	Repository repositories.CommentRepository
}

func (i *GetAllCommentsInteractor) Execute(productUUID uuid.UUID) []domain.Comment {
	comments := i.Repository.GetAllByProductUUID(productUUID)
	return comments
}
