package main

import (
	"context"

	firestoreRepository "github.com/yanuar-nc/migration-database-microservice/src/repository/firestore"
	gormRepository "github.com/yanuar-nc/migration-database-microservice/src/repository/gorm"
	usecasePackage "github.com/yanuar-nc/migration-database-microservice/src/usecase"

	"github.com/yanuar-nc/migration-database-microservice/config"
)

func NewScheduler(cfg config.Config) error {
	firestoreRepository := firestoreRepository.NewRepository(cfg.Firestore.Client)
	mysqlRepository := gormRepository.NewRepository(cfg.WriteDB.DB)

	usecase := usecasePackage.NewUsecaseImplementation(cfg).
		PutRepository(firestoreRepository).
		PutRepositoryMigration(mysqlRepository)
	if err := usecase.Migration(context.Background()); err != nil {
		return err
	}
	return nil
}
