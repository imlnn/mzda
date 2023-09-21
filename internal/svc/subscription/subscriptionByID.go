package subscription

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (svc *Svc) SubscriptionByID(id int) (res []byte, statusCode int, err error) {
	const fn = "internal/svc/subscription/subscriptionByID/SubscriptionByID"
	sub, err := svc.subStorage.SubscriptionByID(id)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, http.StatusNotFound, fmt.Errorf("%s %v", fn, err)
	}
	res, err = json.Marshal(&sub)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to generate response")
	}
	return res, http.StatusOK, nil
}
