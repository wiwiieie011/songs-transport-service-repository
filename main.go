package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wiwiieie011/songs/config"
	"github.com/wiwiieie011/songs/models"
	"github.com/wiwiieie011/songs/repository"
	"github.com/wiwiieie011/songs/services"
	"github.com/wiwiieie011/songs/transport"
)

func main() {
	db := config.SetUpDatabaseConnection()
	server := gin.Default()

	if err := db.AutoMigrate(&models.Category{}, &models.Song{}, &models.User{}, &models.PlayList{}, &models.PlayListItems{}); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	logs:= logrus.New()
	logs.SetLevel(logrus.InfoLevel)

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: false,
		FullTimestamp:    true,
		PadLevelText:     true,
	})

	songsRepo := repository.NewSongRepository(db,logs)
	categoryRepo := repository.NewCategoryRepository(db,logs)
	playlistRepo := repository.NewPlayListRepository(db,logs)
	playlistItemRepo := repository.NewPlayListItemsRepository(db,logs)
	userRepo := repository.NewUserRepository(db,logs)

	songsServise := services.NewSongService(songsRepo, logs)
	categoryServise := services.NewCategoryService(categoryRepo, logs)
	userService := services.NewUserService(userRepo, logs)
	playlistService := services.NewPlayListServices(playlistRepo, userService,logs)
	playlistItemService := services.NewPlayListItemsServices(playlistItemRepo, playlistService, songsServise, logs)

	if tableList, err := db.Migrator().GetTables(); err == nil {
		fmt.Println("tables:", tableList)
	}

	transport.RegisterRoutes(
		server,
		songsServise,
		categoryServise,
		userService,
		playlistService,
		playlistItemService,
		logs,
	)

	logrus.Infof("server started addr=:%v env=%v", os.Getenv("PORT"), os.Getenv("DB_HOST"))
	if err := server.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}