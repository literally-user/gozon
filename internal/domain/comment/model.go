package comment

import (
	"fmt"

	"github.com/google/uuid"
)

type Comment struct {
	UUID uuid.UUID

	content string
	rate    float32

	ProductUUID uuid.UUID
	UserUUID    uuid.UUID
}

func NewComment(content string, rate float32, productUUID, userUUID uuid.UUID) (Comment, error) {
	var err error
	newComment := Comment{
		UUID: uuid.New(),

		ProductUUID: productUUID,
		UserUUID:    userUUID,
	}

	err = newComment.ChangeContent(content)
	if err != nil {
		return Comment{}, fmt.Errorf("failed to set content: %w", err)
	}

	err = newComment.ChangeRate(rate)
	if err != nil {
		return Comment{}, fmt.Errorf("failed to set rate: %w", err)
	}

	return newComment, nil
}

func (c *Comment) Rate() float32 {
	return c.rate
}

func (c *Comment) Content() string {
	return c.content
}

func (c *Comment) ChangeRate(rate float32) error {
	if c.rate == rate {
		return ErrRateDoesntChanged
	}

	c.rate = rate
	return nil
}

func (c *Comment) ChangeContent(content string) error {
	if c.content == content {
		return ErrContentDoesntChanged
	}
	if len(content) < 5 {
		return ErrWrongContentFormat
	}

	c.content = content
	return nil
}
