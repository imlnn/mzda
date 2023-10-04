package subscriber

import (
	"mzda/internal/storage/models"
	"net/http"
)

type Service interface {
	AddSubscriber(req *http.Request) (statusCode int, err error)
	SubscriberByID(id int) (res []byte, statusCode int, err error)
	ListSubscribersByUserID(id int) (res []byte, statusCode int, err error)
	UpdateSubscriber(req *http.Request) (statusCode int, err error)
	DeleteSubscriberByID(id int, req *http.Request) (statusCode int, err error)
}

type Svc struct {
	subscriberStorage   models.SubscribersStorage
	subscriptionStorage models.SubscriptionsStorage
}

func NewSubscriberSvc(storage models.SubscribersStorage) *Svc {
	return &Svc{subscriberStorage: storage}
}
