package publisher

import (
	"fmt"

	"github.com/literally_user/gozon/internal/application/common/publisher"
)

type Publisher struct{}

func NewMockPublisher() Publisher {
	return Publisher{}
}

func (Publisher) Publish(event publisher.Event) error {
	fmt.Printf("Published: %s | %v\n", event.Name(), event)
	return nil
}
