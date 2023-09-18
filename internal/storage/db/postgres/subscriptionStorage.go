package postgres

import (
	"fmt"
	"log"
	"mzda/internal/storage/models"
)

func (c *Connection) AddSubscription(sub *models.Subscription) error {
	const fn = "internal/storage/db/postgres/subscriptionStorage/AddSubscription"
	stmt, err := c.db.Prepare("INSERT INTO subscription (name, admin_id, description, max_members, price, currency, commission, charge_period, creation, start, ending) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	_, err = stmt.Exec(
		sub.Name, sub.AdminID, sub.Description, sub.MaxMembers, sub.Price, sub.Currency, sub.Commission,
		sub.ChargePeriod, sub.Creation, sub.Start, sub.Ending)

	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	return nil
}

func (c *Connection) SubscriptionByID(subID int) (*models.Subscription, error) {
	const fn = "internal/storage/db/postgres/subscriptionStorage/SubscriptionByID"
	stmt, err := c.db.Prepare("SELECT * FROM subscription WHERE id = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	var sub models.Subscription

	err = stmt.QueryRow(subID).Scan(
		&sub.Name, &sub.AdminID, &sub.Description, &sub.MaxMembers, &sub.Price, &sub.Currency, &sub.Commission,
		&sub.ChargePeriod, &sub.Creation, &sub.Start, &sub.Ending)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	return &sub, nil
}

func (c *Connection) SubscriptionByAdminID(subAdminID int) (*models.Subscription, error) {
	const fn = "internal/storage/db/postgres/subscriptionStorage/SubscriptionByAdminID"
	stmt, err := c.db.Prepare("SELECT * FROM subscription WHERE admin_id = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	var sub models.Subscription

	err = stmt.QueryRow(subAdminID).Scan(
		&sub.Name, &sub.AdminID, &sub.Description, &sub.MaxMembers, &sub.Price, &sub.Currency, &sub.Commission,
		&sub.ChargePeriod, &sub.Creation, &sub.Start, &sub.Ending)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	if sub.Name == "" {
		return nil, fmt.Errorf("subscription not found")
	}

	return &sub, nil
}

func (c *Connection) SubscriptionByName(name string) (*models.Subscription, error) {
	const fn = "internal/storage/db/postgres/subscriptionStorage/SubscriptionByAdminID"
	stmt, err := c.db.Prepare("SELECT * FROM subscription WHERE name = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	var sub models.Subscription

	err = stmt.QueryRow(name).Scan(
		&sub.Name, &sub.AdminID, &sub.Description, &sub.MaxMembers, &sub.Price, &sub.Currency, &sub.Commission,
		&sub.ChargePeriod, &sub.Creation, &sub.Start, &sub.Ending)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	return &sub, nil
}

func (c *Connection) UpdateSubscription(sub *models.Subscription) error {
	const fn = "internal/storage/db/postgres/subscriptionStorage/UpdateSubscription"
	stmt, err := c.db.Prepare("UPDATE subscription SET name = $1, description = $2, admin_id = $3, max_members = $4, price = $5, currency = $6, commission = $7, charge_period = $8, creation = $9, start = $10, ending = $11 WHERE id = $12")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	_, err = stmt.Exec(sub.Name, sub.AdminID, sub.Description, sub.MaxMembers, sub.Price, sub.Currency, sub.Commission,
		sub.ChargePeriod, sub.Creation, sub.Start, sub.Ending, sub.ID)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	return nil
}

func (c *Connection) DeleteSubscription(sub *models.Subscription) error {
	const fn = "internal/storage/db/postgres/subscriptionStorage/DeleteSubscription"
	stmt, err := c.db.Prepare("DELETE FROM subscription WHERE id = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	_, err = stmt.Exec(sub.ID)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	return nil
}
