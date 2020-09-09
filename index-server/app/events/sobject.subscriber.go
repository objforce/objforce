package events

import (
	"github.com/micro/go-micro/v2/broker"
)

type IndexSubscriber struct {
	broker broker.Broker
}

func NewSObjectSubscriber(b broker.Broker) (*IndexSubscriber, error) {
	s := &IndexSubscriber{
		b,
	}

	_, err := b.Subscribe("sobject.created", func(evt broker.Event) error {
		// m := evt.Message()

		// message 的 SObject 中携带了 CustomObject 的 objId
		// 根据 sObejct.ObjId 查询 customObj信息

		// 根据CustomObject元信息对需要索引的字段建立索引
		return nil
	}, broker.Queue("indexes"))

	return s, err
}

