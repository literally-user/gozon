package manageComment

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/literally_user/gozon/internal/application/common/publisher"
	"github.com/literally_user/gozon/internal/application/common/repositories"
)

type RemoveCommentInteractor struct {
	ProductRepository repositories.ProductRepository
	CommentRepository repositories.CommentRepository
	Publisher         publisher.Publisher
}

func (i *RemoveCommentInteractor) Execute(commentUUID, productUUID uuid.UUID) error {
	product, err := i.ProductRepository.GetByUUID(productUUID)
	if err != nil {
		return fmt.Errorf("publish comment: failed to get product by uuid: %w", err)
	}
	comment, err := i.CommentRepository.GetByUUID(commentUUID)
	err = i.CommentRepository.Remove(comment)
	if err != nil {
		return fmt.Errorf("publish comment: failed to remove comment: %w", err)
	}

	comments := i.CommentRepository.GetAllByProductUUID(productUUID)

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

	err = i.Publisher.Publish(publisher.CommentRemovedEvent{
		Comment: comment,
	})
	if err != nil {
		return fmt.Errorf("publish comment: failed to publish comment: %w", err)
	}

	return nil
}
