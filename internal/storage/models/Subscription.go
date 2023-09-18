package models

import "time"

type Subscription struct {
	ID           int       `json:"id,omitempty"`
	Name         string    `json:"subscription_name"`
	Description  string    `json:"description"`
	AdminID      int       `json:"admin_id"`
	MaxMembers   int       `json:"max_members"`
	Price        int       `json:"price"`
	Currency     string    `json:"currency"`
	Commission   int       `json:"commission"`
	ChargePeriod int       `json:"charge_period"`
	Creation     time.Time `json:"creation"`
	Start        time.Time `json:"start"`
	Ending       time.Time `json:"ending"`
}

type SubscriptionsStorage interface {
	AddSubscription(sub *Subscription) error
	SubscriptionByID(subID int) (*Subscription, error)
	SubscriptionByAdminID(subAdminID int) (*Subscription, error)
	SubscriptionByName(name string) (*Subscription, error)
	UpdateSubscription(sub *Subscription) error
	DeleteSubscription(sub *Subscription) error
}
