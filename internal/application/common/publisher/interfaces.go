package publisher

type Publisher interface {
	Publish(Event) error
}

type Event interface {
	Name() string
}
