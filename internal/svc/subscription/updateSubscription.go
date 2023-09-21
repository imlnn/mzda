package subscription

import (
	"fmt"
	"log"
	"mzda/internal/utils"
	"net/http"
)

func (svc *Svc) UpdateSubscription(req *http.Request) (statusCode int, err error) {
	const fn = "internal/svc/subscription/updateSubscription/UpdateSubscription"
	sub, err := parseSubscription(req.Body)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusBadRequest, fmt.Errorf("failed to parse request")
	}

	oldSub, err := svc.subStorage.SubscriptionByID(sub.ID)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusBadRequest, err
	}

	jwt := req.Context().Value("jwt").(*utils.JWT)
	if jwt.UserID != oldSub.AdminID {
		err = fmt.Errorf("token given for another user")
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusUnauthorized, err
	}

	err = svc.subStorage.UpdateSubscription(sub)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusInternalServerError, fmt.Errorf("failed to update")
	}

	return http.StatusOK, nil
}
