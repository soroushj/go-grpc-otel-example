package server

import (
	"context"

	"github.com/soroushj/go-grpc-otel-example/notes"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otelcodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	trc = otel.Tracer("go-grpc-otel-example/server")
)

type Server struct {
	notes.UnimplementedNotesServer
	db map[string]string
}

func New() *Server {
	return &Server{
		db: map[string]string{
			"1": "buy coffee",
			"2": "buy bread",
			"3": "buy cheese",
		},
	}
}

func (s *Server) GetNote(ctx context.Context, req *notes.GetNoteRequest) (_ *notes.GetNoteResponse, _err error) {
	ctx, span := trc.Start(ctx, "GetNote")
	span.SetAttributes(attribute.Stringer("req", req))
	defer func() {
		if _err != nil {
			span.SetStatus(otelcodes.Error, _err.Error())
		}
		span.End()
	}()

	t := s.db[req.Id]
	if t == "" {
		return nil, status.Error(codes.NotFound, "note not found")
	}
	return &notes.GetNoteResponse{
		Note: &notes.Note{
			Id:   req.Id,
			Text: t,
		},
	}, nil
}
