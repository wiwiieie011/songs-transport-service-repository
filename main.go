package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/wiwiieie011/songs/config"
	"github.com/wiwiieie011/songs/models"
	"github.com/wiwiieie011/songs/repository"
	"github.com/wiwiieie011/songs/services"
	"github.com/wiwiieie011/songs/transport"
)

func main() {
	db := config.SetUpDatabaseConnection()
	server := gin.Default()

	if err := db.AutoMigrate(&models.Category{},&models.Song{}, &models.User{}, &models.PlayList{}, &models.PlayListItems{}); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}


	songsRepo := repository.NewSongRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	playlistRepo := repository.NewPlayListRepository(db)
	playlistItemRepo := repository.NewPlayListItemsRepository(db)
	userRepo := repository.NewUserRepository(db)

	songsServise := services.NewSongService(songsRepo)
	categoryServise := services.NewCategoryService(categoryRepo)
	userService := services.NewUserService(userRepo)
	playlistService := services.NewPlayListServices(playlistRepo, userService)
	playlistItemService := services.NewPlayListItemsServices(playlistItemRepo,playlistService,songsServise)

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
	)

	if err := server.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}