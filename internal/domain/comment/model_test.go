package comment

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewComment(t *testing.T) {
	productUUID := uuid.New()
	userUUID := uuid.New()

	tests := []struct {
		name        string
		content     string
		rate        float32
		productUUID uuid.UUID
		userUUID    uuid.UUID
		wantErr     bool
	}{
		{
			name:        "successful comment creation",
			content:     "Great product!",
			rate:        4.5,
			productUUID: productUUID,
			userUUID:    userUUID,
			wantErr:     false,
		},
		{
			name:        "comment with empty content",
			content:     "",
			rate:        3.0,
			productUUID: productUUID,
			userUUID:    userUUID,
			wantErr:     true,
		},
		{
			name:        "comment with maximum rate",
			content:     "Excellent!",
			rate:        5.0,
			productUUID: productUUID,
			userUUID:    userUUID,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment, err := NewComment(tt.content, tt.rate, tt.productUUID, tt.userUUID)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if comment.UUID == uuid.Nil {
					t.Error("NewComment() UUID should not be nil")
				}

				if comment.Content != tt.content {
					t.Errorf("NewComment() Content = %v, want %v", comment.Content, tt.content)
				}

				if comment.Rate != tt.rate {
					t.Errorf("NewComment() Rate = %v, want %v", comment.Rate, tt.rate)
				}

				if comment.ProductUUID != tt.productUUID {
					t.Errorf("NewComment() ProductUUID = %v, want %v", comment.ProductUUID, tt.productUUID)
				}

				if comment.UserUUID != tt.userUUID {
					t.Errorf("NewComment() UserUUID = %v, want %v", comment.UserUUID, tt.userUUID)
				}
			}
		})
	}
}

func TestComment_ChangeRate(t *testing.T) {
	tests := []struct {
		name     string
		initial  float32
		newRate  float32
		wantErr  error
		wantRate float32
	}{
		{
			name:     "successful rate change",
			initial:  3.0,
			newRate:  4.5,
			wantErr:  nil,
			wantRate: 4.5,
		},
		{
			name:     "rate doesn't change - same value",
			initial:  4.0,
			newRate:  4.0,
			wantErr:  ErrRateDoesntChanged,
			wantRate: 4.0,
		},
		{
			name:     "change from zero to non-zero",
			initial:  0.0,
			newRate:  2.5,
			wantErr:  nil,
			wantRate: 2.5,
		},
		{
			name:     "change to zero rate",
			initial:  3.5,
			newRate:  0.0,
			wantErr:  nil,
			wantRate: 0.0,
		},
		{
			name:     "change to negative rate",
			initial:  2.0,
			newRate:  -1.0,
			wantErr:  nil,
			wantRate: -1.0,
		},
		{
			name:     "change to maximum rate",
			initial:  3.0,
			newRate:  5.0,
			wantErr:  nil,
			wantRate: 5.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Comment{Rate: tt.initial}

			err := c.ChangeRate(tt.newRate)

			if err != tt.wantErr {
				t.Errorf("ChangeRate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if c.Rate != tt.wantRate {
				t.Errorf("ChangeRate() Rate = %v, want %v", c.Rate, tt.wantRate)
			}
		})
	}
}

func TestComment_ChangeContent(t *testing.T) {
	tests := []struct {
		name        string
		initial     string
		newContent  string
		wantErr     error
		wantContent string
	}{
		{
			name:        "successful content change",
			initial:     "Good product",
			newContent:  "Great product!",
			wantErr:     nil,
			wantContent: "Great product!",
		},
		{
			name:        "content doesn't change - same value",
			initial:     "Same content",
			newContent:  "Same content",
			wantErr:     ErrContentDoesntChanged,
			wantContent: "Same content",
		},
		{
			name:        "change from empty to non-empty",
			initial:     "",
			newContent:  "New content",
			wantErr:     nil,
			wantContent: "New content",
		},
		{
			name:        "change to empty content",
			initial:     "Old content",
			newContent:  "",
			wantErr:     nil,
			wantContent: "",
		},
		{
			name:        "change with special characters",
			initial:     "Simple text",
			newContent:  "Text with émojis 🎉 and symbols!",
			wantErr:     nil,
			wantContent: "Text with émojis 🎉 and symbols!",
		},
		{
			name:        "change with multiline content",
			initial:     "Single line",
			newContent:  "Line 1\nLine 2\nLine 3",
			wantErr:     nil,
			wantContent: "Line 1\nLine 2\nLine 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Comment{Content: tt.initial}

			err := c.ChangeContent(tt.newContent)

			if err != tt.wantErr {
				t.Errorf("ChangeContent() error = %v, wantErr %v", err, tt.wantErr)
			}

			if c.Content != tt.wantContent {
				t.Errorf("ChangeContent() Content = %v, want %v", c.Content, tt.wantContent)
			}
		})
	}
}
