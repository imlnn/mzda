package subscriber

import (
	"fmt"
	"log"
	"mzda/internal/utils"
	"net/http"
)

func (svc *Svc) DeleteSubscriberByID(id int, req *http.Request) (statusCode int, err error) {
	const fn = "internal/svc/subscriber/deleteSubscriberByID/DeleteSubscriberByID"

	subscriber, err := svc.subscriberStorage.SubscriberByID(id)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusInternalServerError, err
	}

	subscription, err := svc.subscriptionStorage.SubscriptionByID(subscriber.SubscriptionID)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusInternalServerError, err
	}

	jwt := req.Context().Value("jwt").(*utils.JWT)
	if jwt.UserID != subscription.AdminID || !jwt.Admin {
		err = fmt.Errorf("token given for another user")
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusUnauthorized, err
	}

	err = svc.subscriberStorage.DeleteSubscriberByID(id)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusInternalServerError, fmt.Errorf("failed to update")
	}

	return http.StatusOK, nil
}
