package models

import "time"

type Subscriber struct {
	ID                int       `json:"id,omitempty"`
	UserID            int       `json:"userID"`
	SubscriptionID    int       `json:"subscriptionID"`
	SubscriptionStart time.Time `json:"subscriptionStart"`
	SubscriptionEnd   time.Time `json:"subscriptionEnd"`
}

type SubscribersStorage interface {
	AddSubscriber(sub *Subscriber) error
	SubscriberByID(id int) (*Subscriber, error)
	SubscriptionsByUserID(usrID int) ([]*Subscriber, error)
	DeleteSubscriberByID(id int) error
	UpdateSubscriber(sub *Subscriber) error
}
