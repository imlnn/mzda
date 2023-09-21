package subscription

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mzda/internal/storage/models"
	"net/http"
)

func parseSubscription(b io.ReadCloser) (*models.Subscription, error) {
	const fn = "internal/svc/subscription/addSubscription/parseSubscription"
	var req models.Subscription

	err := json.NewDecoder(b).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("%s %v", fn, err)
	}

	return &req, nil
}

func (svc *Svc) AddSubscription(req *http.Request) (statusCode int, err error) {
	const fn = "internal/svc/subscription/addSubscription/AddSubscription"
	sub, err := parseSubscription(req.Body)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusBadRequest, fmt.Errorf("failed to parse request")
	}

	err = svc.subStorage.AddSubscription(sub)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusInternalServerError, fmt.Errorf("failed to add subscription")
	}

	return http.StatusOK, nil
}
