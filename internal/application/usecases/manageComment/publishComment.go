package manageComment

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
	"github.com/literally_user/gozon/internal/domain/comment"
)

type PublishCommentInteractor struct {
	ProductRepository repositories.ProductRepository
	CommentRepository repositories.CommentRepository
	Publisher         publisher.Publisher
}

func (i *PublishCommentInteractor) Execute(userUUID, productUUID uuid.UUID, content string, rate float32) error {
	product, err := i.ProductRepository.GetByUUID(productUUID)
	if err != nil {
		return fmt.Errorf("publish comment: failed to get product by uuid: %w", err)
	}

	newComment, err := comment.NewComment(content, rate, productUUID, userUUID)
	if err != nil {
		return fmt.Errorf("publish comment: failed to create domain comment: %w", err)
	}

	err = i.CommentRepository.Create(newComment)
	if err != nil {
		return fmt.Errorf("publish comment: failed to create comment: %w", err)
	}

	comments := i.CommentRepository.GetAllByProductUUID(productUUID)
	comments = append(comments, newComment)

	var sumRate float32
	var lenComments float32

	for _, val := range comments {
		sumRate += val.Rate()
		lenComments++
	}

	err = product.ChangeRating(sumRate / lenComments)
	if err != nil {
		return fmt.Errorf("publish comment: failed to change product rating: %w", err)
	}

	err = i.ProductRepository.Update(product)
	if err != nil {
		return fmt.Errorf("publish comment: failed to update product: %w", err)
	}

	err = i.Publisher.Publish(publisher.CommentCreatedEvent{
		Comment: newComment,
	})
	if err != nil {
		return fmt.Errorf("publish comment: failed to publish comment: %w", err)
	}

	return nil
}
