package models

type SubscribersStorage interface {
	AddSubscriber()
	GetSubscriber()
	DeleteSubscriber()
	UpdateSubscriber()
}
