package subscription

import (
	"mzda/internal/storage/models"
	"net/http"
)

type Svc struct {
	subStorage models.SubscriptionsStorage
}

type Service interface {
	AddSubscription(req *http.Request) (statusCode int, err error)
	SubscriptionByID(id int) (res []byte, statusCode int, err error)
	SubscriptionByAdminID(id int) (res []byte, statusCode int, err error)
	UpdateSubscription(req *http.Request) (statusCode int, err error)
	DeleteSubscription(id int, req *http.Request) (statusCode int, err error)
}

func NewSubscriptionSvc(storage models.SubscriptionsStorage) *Svc {
	const fn = "internal/svc/subscription/subscriptionSvc/NewSubscriptionSvc"
	return &Svc{
		subStorage: storage,
	}
}
