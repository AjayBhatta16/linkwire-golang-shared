package utilities

import (
	"context"
	"encoding/base64"
    "encoding/json"
    "fmt"
	"io"
    "net/http"

    "cloud.google.com/go/pubsub/v2"
)

type Publisher struct {
    client *pubsub.Client
}

func NewPublisher(ctx context.Context) (*Publisher, error) {
    client, err := pubsub.NewClient(ctx, pubsub.DetectProjectID)

    if err != nil {
        return nil, fmt.Errorf("Error creating pub/sub client: %w", err)
    }

    return &Publisher{client: client}, nil
}

func (p *Publisher) Close() error {
    return p.client.Close()
}

func Publish[T any](p *Publisher, topicID string, payload T, attrs ...map[string]string) error {
    data, err := json.Marshal(payload)

    if err != nil {
        return fmt.Errorf("Error marshaling payload: %w", err)
    }

    msg := &pubsub.Message{Data: data}

    if len(attrs) > 0 {
        msg.Attributes = attrs[0]
    }

    publisher := p.client.Publisher(topicID)

    defer publisher.Stop()

	result := publisher.Publish(context.Background(), msg)

    if _, err := result.Get(context.Background()); err != nil {
        return fmt.Errorf("Error publishing message: %w", err)
    }

    return nil
}

type Subscriber struct {
    client *pubsub.Client
}

func NewSubscriber(ctx context.Context) (*Subscriber, error) {
    client, err := pubsub.NewClient(ctx, pubsub.DetectProjectID)
    if err != nil {
        return nil, fmt.Errorf("create pubsub client: %w", err)
    }
    return &Subscriber{client: client}, nil
}

func (s *Subscriber) Close() error {
    return s.client.Close()
}

type PubSubEnvelope struct {
    Message struct {
        Data       string            `json:"data"`
        Attributes map[string]string `json:"attributes"`
        MessageID  string            `json:"messageId"`
    } `json:"message"`
    Subscription string `json:"subscription"`
}

func Receive[T any](s *Subscriber, r *http.Request) (T, map[string]string, error) {
    var zero T

    body, err := io.ReadAll(r.Body)
    if err != nil {
        return zero, nil, fmt.Errorf("read request body: %w", err)
    }
    defer r.Body.Close()

    var envelope PubSubEnvelope
    if err := json.Unmarshal(body, &envelope); err != nil {
        return zero, nil, fmt.Errorf("unmarshal envelope: %w", err)
    }

    decoded, err := base64.StdEncoding.DecodeString(envelope.Message.Data)
    if err != nil {
        return zero, nil, fmt.Errorf("decode message data: %w", err)
    }

    var payload T
    if err := json.Unmarshal(decoded, &payload); err != nil {
        return zero, nil, fmt.Errorf("unmarshal payload: %w", err)
    }

    return payload, envelope.Message.Attributes, nil
}