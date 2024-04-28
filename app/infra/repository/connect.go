package repository

import (
	"github.com/ganganbiz1/go-echo-gorm-template/app/config"
	"gorm.io/gorm"
)

type PrimaryDB struct {
	*gorm.DB
}

type ReplicaDB struct {
	*gorm.DB
}

func NewPrimaryDBConnect(c *config.PrimaryDBConfig) (PrimaryDB, func(), error) {
	// TODO インフラ構築したらコメントアウトはずす
	return PrimaryDB{nil}, nil, nil
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=%s", c.User, c.Password, c.Host, c.Port, c.Database, c.TimeZone)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	Logger: gormLogger.Default.LogMode(c.LogLevel),
	// })

	// if err != nil {
	// 	return PrimaryDB{}, nil, infra.HandleError(err)
	// }

	// d, err := db.DB()

	// if err != nil {
	// 	return PrimaryDB{}, nil, infra.HandleError(err)
	// }

	// err = d.Ping()

	// if err != nil {
	// 	return PrimaryDB{}, nil, infra.HandleError(err)
	// }
	// cl := func() {
	// 	d.Close()
	// }
	// logger.Info("PrimaryDB Connect Success")
	// return PrimaryDB{db}, cl, nil
}

func NewReplicaDBConnect(c *config.ReplicaDBConfig) (ReplicaDB, func(), error) {
	// TODO インフラ構築したらコメントアウトはずす
	return ReplicaDB{nil}, nil, nil
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=%s", c.User, c.Password, c.Host, c.Port, c.Database, c.TimeZone)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	Logger: gormLogger.Default.LogMode(c.LogLevel),
	// })

	// if err != nil {
	// 	return ReplicaDB{}, nil, infra.HandleError(err)
	// }

	// d, err := db.DB()

	// if err != nil {
	// 	return ReplicaDB{}, nil, infra.HandleError(err)
	// }

	// err = d.Ping()

	// if err != nil {
	// 	return ReplicaDB{}, nil, infra.HandleError(err)
	// }
	// cl := func() {
	// 	d.Close()
	// }
	// logger.Info("ReplicaDB Connect Success")
	// return ReplicaDB{db}, cl, nil
}
