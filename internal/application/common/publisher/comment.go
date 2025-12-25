package publisher

import (
	domain "github.com/literally_user/gozon/internal/domain/comment"
)

type CommentCreatedEvent struct {
	Comment domain.Comment
}

func (CommentCreatedEvent) Name() string {
	return "create.comment"
}

type CommentRemovedEvent struct {
	Comment domain.Comment
}

func (CommentRemovedEvent) Name() string {
	return "remove.comment"
}
