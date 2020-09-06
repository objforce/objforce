package providers

import (
	"github.com/micro/go-micro/v2"
	proto "github.com/micro/go-micro/v2/agent/proto"
	"github.com/objforce/bot/app/commands"
)

func RegisterCommands(
	service micro.Service,
	migrateCommand *commands.MigrateCommand,
) {
	proto.RegisterCommandHandler(service.Server(), migrateCommand)
}
