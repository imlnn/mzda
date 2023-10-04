package subscriber

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mzda/internal/storage/models"
	"net/http"
)

func parseSubscriber(b io.ReadCloser) (*models.Subscriber, error) {
	const fn = "internal/svc/subscriber/addSubscriber/parseSubscriber"
	var req models.Subscriber

	err := json.NewDecoder(b).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("%s %v", fn, err)
	}

	return &req, nil
}

func (svc *Svc) AddSubscriber(req *http.Request) (statusCode int, err error) {
	const fn = "internal/svc/subscriber/addSubscriber/AddSubscriber"
	sub, err := parseSubscriber(req.Body)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusBadRequest, fmt.Errorf("failed to parse request")
	}

	err = svc.subscriberStorage.AddSubscriber(sub)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusInternalServerError, fmt.Errorf("failed to add subscription")
	}

	return http.StatusOK, nil
}
