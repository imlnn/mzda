package models

type Subscriber struct {
}

type SubscribersStorage interface {
	AddSubscriber()
	GetSubscriber()
	DeleteSubscriber()
	UpdateSubscriber()
}
