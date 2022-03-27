package repository

type MessageBroker interface {
	Publish(req interface{}) error
}
