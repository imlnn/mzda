package subscription

import (
	"net/http"
)

type SubscriptionService interface {
	AddSubscription(req *http.Request) (statusCode int, err error)
	SubscriptionByID(req *http.Request) (payload []byte, statusCode int, err error)
	SubscriptionByAdminID(req *http.Request) (payload []byte, statusCode int, err error)
	SubscriptionByName(req *http.Request) (payload []byte, statusCode int, err error)
	UpdateSubscription(req *http.Request) (statusCode int, err error)
	DeleteSubscription(req *http.Request) (statusCode int, err error)
}
