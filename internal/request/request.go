package request

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Message struct {
	Body interface{} `json:"body"`
}

func New(body interface{}) *Message {
	return &Message{Body: body}
}

func (m *Message) Error(err error) error {
	logrus.WithField("id", uuid.NewString()).Info(m.Body)
	return errors.Wrap(err, "failed to return")
}

func (m *Message) Infof() {
	logrus.WithField("id", uuid.NewString()).Info()
}

func (m *Message) Reply() (*events.APIGatewayProxyResponse, error) {
	// Log the request
	m.Infof()
	//
	// // Handle the request body
	body, err := json.Marshal(m.Body)
	if err != nil {
		return nil, m.Error(err)
	}
	event := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}

	return &event, nil
}
