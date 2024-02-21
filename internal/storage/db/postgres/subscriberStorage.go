package postgres

import (
	"fmt"
	"log"
	"mzda/internal/storage/models"
)

func (c *Connection) AddSubscriber(sub *models.Subscriber) error {
	const fn = "internal/storage/db/postgres/subscriberStorage/AddSubscriber"
	stmt, err := c.db.Prepare("INSERT INTO subscribers (userid, subscriptionid, subscription_start, subscription_ending) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	_, err = stmt.Exec(sub.UserID, sub.SubscriptionID, sub.SubscriptionStart, sub.SubscriptionEnd)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	return nil
}

func (c *Connection) SubscriberByID(id int) (*models.Subscriber, error) {
	const fn = "internal/storage/db/postgres/subscriberStorage/SubscriberByID"
	stmt, err := c.db.Prepare("SELECT * FROM subscribers WHERE id = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	var sub models.Subscriber

	err = stmt.QueryRow(id).Scan(
		&sub.ID, &sub.UserID, &sub.SubscriptionID, &sub.SubscriptionStart, &sub.SubscriptionEnd)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	return &sub, nil
}

func (c *Connection) SubscriptionsByUserID(usrID int) ([]*models.Subscriber, error) {
	const fn = "internal/storage/db/postgres/subscriberStorage/SubscriptionsByUserID"
	stmt, err := c.db.Prepare("SELECT * FROM subscribers WHERE userid = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	rows, err := stmt.Query(usrID)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	res := make([]*models.Subscriber, 0)

	for rows.Next() {
		sub := new(models.Subscriber)
		err = rows.Scan(
			&sub.ID, &sub.UserID, &sub.SubscriptionID, &sub.SubscriptionStart, &sub.SubscriptionEnd)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			return nil, err
		}
		res = append(res, sub)
	}

	return res, nil
}

func (c *Connection) DeleteSubscriberByID(id int) error {
	const fn = "internal/storage/db/postgres/subscriberStorage/DeleteSubscriberByID"
	stmt, err := c.db.Prepare("DELETE FROM subscribers WHERE id = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	return nil
}

func (c *Connection) UpdateSubscriber(sub *models.Subscriber) error {
	const fn = "internal/storage/db/postgres/subscriberStorage/UpdateSubscriber"
	stmt, err := c.db.Prepare("UPDATE subscribers SET userid = $1, subscriptionid = $2, subscription_start = $3, subscription_ending = $4 WHERE id = $5")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	_, err = stmt.Exec(sub.UserID, sub.SubscriptionID, sub.SubscriptionStart, sub.SubscriptionEnd, sub.ID)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	return nil
}
