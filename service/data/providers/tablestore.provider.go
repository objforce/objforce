package providers

import (
	"context"
	"fmt"
	"strings"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/duolacloud/microbase/datasource"
	"github.com/duolacloud/microbase/multitenancy"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/config"
)

func NewTableStoreTenancy(config config.Config, entityMap datasource.EntityMap) (multitenancy.Tenancy, error) {
	endpoint := config.Get("tablestore", "endpoint").String("")
	instanceName := config.Get("tablestore", "instance_name").String("")
	accessKeyId := config.Get("tablestore", "access_key_id").String("")
	accessKeySecret := config.Get("tablestore", "access_key_secret").String("")

	clientCreateFn := func(ctx context.Context, tenantId string) (multitenancy.Resource, error) {
		// TODO config := configService.Retrieve("objforce", tenantId, "dataSource")
		client := tablestore.NewClient(endpoint, instanceName, accessKeyId, accessKeySecret)

		autoMigrate(tenantId, entityMap, client)
		return client, nil
	}

	clientCloseFunc := func(resource multitenancy.Resource) {}

	tenancy := multitenancy.NewCachedTenancy(clientCreateFn, clientCloseFunc)

	return tenancy, nil
}

// DBName returns the prefixed database name in order to avoid collision with MySQL internal databases.
func DBName(prefix string, tenantId string) string {
	if len(tenantId) == 0 {
		return prefix
	}
	return fmt.Sprintf("%s_%s", prefix, tenantId)
}

// FromDBName returns the source name of the tenant.
func FromDBName(serviceName string, name string) string {
	return strings.TrimPrefix(name, fmt.Sprintf("%s_", serviceName))
}

func DBFromContext(tenancy multitenancy.Tenancy, ctx context.Context) (*gorm.DB, error) {
	tenantName, _ := multitenancy.FromContext(ctx)

	db, err := tenancy.ResourceFor(ctx, tenantName)
	if err != nil {
		return nil, err
	}
	return db.(*gorm.DB), nil
}

func TableName(tableName string, tenantId string) string {
	if len(tenantId) == 0 {
		return tableName
	}

	return fmt.Sprintf("%s_%s", tableName, tenantId)
}

func autoMigrate(tenantId string, entityMap datasource.EntityMap, client *tablestore.TableStoreClient) error {
	// entities := entityMap.GetEntities()

	return nil
}
