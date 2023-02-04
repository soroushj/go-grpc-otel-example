package client

import (
	"context"
	"errors"

	"github.com/soroushj/go-grpc-otel-example/notes"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

var (
	trc = otel.Tracer("go-grpc-otel-example/client")
)

type Client struct {
	client notes.NotesClient
}

func New(client notes.NotesClient) *Client {
	return &Client{client}
}

func (c *Client) GetNoteText(ctx context.Context, id string) (_ string, _err error) {
	ctx, span := trc.Start(ctx, "GetNoteText")
	span.SetAttributes(attribute.String("id", id))
	defer func() {
		if _err != nil {
			span.SetStatus(codes.Error, _err.Error())
		}
		span.End()
	}()

	resp, err := c.client.GetNote(ctx, &notes.GetNoteRequest{Id: id})
	if err != nil {
		return "", err
	}
	if resp.Note == nil {
		return "", errors.New("client: bad response")
	}
	return resp.Note.Text, nil
}
