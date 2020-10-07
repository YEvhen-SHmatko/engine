package mq

import (
	"context"
	"github.com/webitel/engine/model"
)

type LayeredMQLayer interface {
	MQ
}

type LayeredMQ struct {
	context context.Context
	MQLayer LayeredMQLayer
}

func NewMQ(mq LayeredMQLayer) MQ {
	return &LayeredMQ{
		context: context.TODO(),
		MQLayer: mq,
	}
}

func (l *LayeredMQ) SendJSON(name string, data []byte) *model.AppError {
	return l.MQLayer.SendJSON(name, data)
}

func (l *LayeredMQ) Start() {
	l.MQLayer.Start()
}

func (l *LayeredMQ) Close() {
	l.MQLayer.Close()
}

func (l *LayeredMQ) BindCallEvents(domainId, userId int64) error {
	return l.MQLayer.BindCallEvents(domainId, userId)
}

func (l *LayeredMQ) UnBindCallEvents(domainId, userId int64) error {
	return l.MQLayer.UnBindCallEvents(domainId, userId)
}

func (l *LayeredMQ) NewDomainQueue(domainId int64, bindings model.GetAllBindings) (DomainQueue, *model.AppError) {
	return l.MQLayer.NewDomainQueue(domainId, bindings)
}

func (l *LayeredMQ) RegisterWebsocket(domainId int64, event *model.RegisterToWebsocketEvent) *model.AppError {
	return l.MQLayer.RegisterWebsocket(domainId, event)
}

func (l *LayeredMQ) UnRegisterWebsocket(domainId int64, event *model.RegisterToWebsocketEvent) *model.AppError {
	return l.MQLayer.UnRegisterWebsocket(domainId, event)
}