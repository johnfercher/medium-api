package mysqldriver

import (
	"context"
	"fmt"

	"github.com/johnfercher/medium-api/pkg/observability/log"
	"github.com/johnfercher/medium-api/pkg/observability/log/field"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Start(ctx context.Context, url string, dbName string, adminUser string, adminPassword string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", adminUser, adminPassword, url, dbName)

	log.Info(ctx, "connecting to mysql", field.String("url", url),
		field.String("dbname", dbName),
		field.String("username", adminUser),
		field.String("password", adminPassword))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(ctx, "could not connect to mysql", field.String("url", url),
			field.String("dbname", dbName),
			field.String("username", adminUser),
			field.String("password", adminPassword))
		return nil, err
	}

	log.Info(ctx, "connected to mysql", field.String("url", url),
		field.String("dbname", dbName),
		field.String("username", adminUser),
		field.String("password", adminPassword))

	return db, nil
}
