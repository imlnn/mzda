package models

type SubscriptionsStorage interface {
	AddSubscription()
	GetSubscription()
	DeleteSubscription()
	UpdateSubscription()
}
