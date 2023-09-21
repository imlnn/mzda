package subscription

import (
	"fmt"
	"log"
	"mzda/internal/utils"
	"net/http"
)

func (svc *Svc) DeleteSubscription(id int, req *http.Request) (statusCode int, err error) {
	const fn = "internal/svc/subscription/deleteSubscription/DeleteSubscription"
	sub, err := svc.subStorage.SubscriptionByID(id)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusBadRequest, err
	}

	jwt := req.Context().Value("jwt").(*utils.JWT)
	if jwt.UserID != sub.AdminID {
		err = fmt.Errorf("token given for another user")
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusUnauthorized, err
	}

	err = svc.subStorage.DeleteSubscription(sub)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusInternalServerError, fmt.Errorf("failed to update")
	}

	return http.StatusOK, nil
}
