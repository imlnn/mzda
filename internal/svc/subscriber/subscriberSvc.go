package subscriber

import (
	"mzda/internal/storage/models"
	"net/http"
)

type Service interface {
	AddSubscriber(req *http.Request) (statusCode int, err error)
	SubscriberByID(id int) (sub []byte, statusCode int, err error)
	ListSubscribersByUserID(usrID int) (subs []byte, statusCode int, err error)
	UpdateSubscriber(req *http.Request) (statusCode int, err error)
	DeleteSubscriberByID(id int, req *http.Request) (statusCode int, err error)
}

type Svc struct {
	subscriberStorage models.SubscribersStorage
}

func NewSubscriberSvc(storage models.SubscribersStorage) *Svc {
	return &Svc{subscriberStorage: storage}
}
