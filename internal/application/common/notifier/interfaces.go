package notifier

type Adapter interface {
	Send(to string) error
}
