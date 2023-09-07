package storage

type UserStorage interface {
	AddUser() error
	GetUser() error
	DeleteUser() error
	UpdateUser() error
}

type SubscriptionsStorage interface {
	AddSubscription()
	GetSubscription()
	DeleteSubscription()
	UpdateSubscription()
}

type SubscribersStorage interface {
	AddSubscriber()
	GetSubscriber()
	DeleteSubscriber()
	UpdateSubscriber()
}

type InvoiceStorage interface {
	AddInvoice()
	GetInvoice()
	DeleteInvoice()
	UpdateInvoice()
}
