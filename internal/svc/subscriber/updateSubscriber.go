package subscriber

import (
	"fmt"
	"log"
	"mzda/internal/utils"
	"net/http"
)

func (svc *Svc) UpdateSubscriber(req *http.Request) (statusCode int, err error) {
	const fn = "internal/svc/subscription/updateSubscription/UpdateSubscription"
	sub, err := parseSubscriber(req.Body)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusBadRequest, fmt.Errorf("failed to parse request")
	}

	oldSubscriber, err := svc.subscriberStorage.SubscriberByID(sub.ID)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusBadRequest, err
	}

	subscription, err := svc.subscriptionStorage.SubscriptionByID(oldSubscriber.SubscriptionID)

	jwt := req.Context().Value("jwt").(*utils.JWT)
	if jwt.UserID != subscription.AdminID || !jwt.Admin {
		err = fmt.Errorf("token given for another user")
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusUnauthorized, err
	}

	err = svc.subscriberStorage.UpdateSubscriber(sub)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusInternalServerError, fmt.Errorf("failed to update")
	}

	return http.StatusOK, nil
}
