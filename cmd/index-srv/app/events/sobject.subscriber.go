package events

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"
	"github.com/objforce/objforce/idl/data/gen-go"
)

type IndexSubscriber struct {
}

func RegisterSObjectSubscriber(service micro.Service) error {
	s := &IndexSubscriber{}
	err := micro.RegisterSubscriber("sobject.created", service.Server(), s, server.SubscriberQueue("indexes"))

	return err
}


func (e *IndexSubscriber) Handle(c context.Context, msg *data.SObject) error {
	logger.Debug("Handler Received message: ", msg)
	return nil
}