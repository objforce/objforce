package commands

import (
	"context"

	"github.com/jinzhu/gorm"
	proto "github.com/micro/go-micro/v2/agent/proto"
	"github.com/objforce/objforce/service/meta/app/domain/models"
)

type MigrateCommand struct {
	db *gorm.DB
}

func NewMigrateCommand(db *gorm.DB) *MigrateCommand {
	return &MigrateCommand{db: db}
}

// Help returns the command usage
func (c *MigrateCommand) Help(ctx context.Context, req *proto.HelpRequest, rsp *proto.HelpResponse) error {
	// Usage should include the name of the command
	rsp.Usage = "migrate"
	rsp.Description = "This is an example bot command as a micro service which echos the message"
	return nil
}

// Exec executes the command
func (c *MigrateCommand) Exec(ctx context.Context, req *proto.ExecRequest, rsp *proto.ExecResponse) error {
	// rsp.Error could be set to return an error instead
	// the function error would only be used for service level issues

	c.db.AutoMigrate(
		models.CustomObject{},
		models.CustomField{},
	)

	rsp.Result = []byte("数据库模式构建完毕")

	return nil
}
