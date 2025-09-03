package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/yanplayer/YanplayerBackend/cmd/seed/seeders"
	"github.com/yanplayer/YanplayerBackend/internal/config"
	"github.com/yanplayer/YanplayerBackend/internal/constants"
	"github.com/yanplayer/YanplayerBackend/internal/utils"
	"github.com/yanplayer/YanplayerBackend/pkg/logger"
)

func init() {
	if err := config.InitializeAppConfig(); err != nil {
		logger.Fatal(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
	}
	logger.Info("configuration loaded", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryConfig})
}

func main() {
	db, err := utils.SetupPostgresConnection()
	if err != nil {
		logger.Panic(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})
	}
	defer db.Close()

	logger.Info("seeding...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})

	seeder := seeders.NewSeeder(db)
	err = seeder.UserSeeder(seeders.UserData)
	if err != nil {
		logger.Panic(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})
	}

	logger.Info("seeding success!", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySeeder})
}
