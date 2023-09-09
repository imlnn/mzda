package mzda

type SubscribersStorage interface {
	AddSubscriber()
	GetSubscriber()
	DeleteSubscriber()
	UpdateSubscriber()
}
