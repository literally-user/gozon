package events

type EventBus interface {
	Notify(event Event) error
}

type Event interface {
	Name() string
}
