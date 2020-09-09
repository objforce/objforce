package providers

import (
	"github.com/micro/go-micro/v2"
	"lucfish.com/uim/uim-srv/proto/uim/gen-go"
	"lucfish.com/uim/uim-srv/app/handlers"
)

func RegisterHandlers(service micro.Service,
	echoHandler *handlers.EchoHandler,
	channelHandler *handlers.ChannelHandler,
	chatHandler *handlers.ChatHandler,
) {
	uim.RegisterEchoServiceHandler(service.Server(), echoHandler)
	uim.RegisterChannelServiceHandler(service.Server(), channelHandler)
	uim.RegisterChatServiceHandler(service.Server(), chatHandler)
}