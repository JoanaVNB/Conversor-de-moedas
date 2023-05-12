package main

import (
	"exchange/service"
	"exchange/service/RateAPI"
	"exchange/domain"
	"exchange/repository"
	"exchange/handlers"
	"exchange/cron"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
)

func main(){

	user := "adm"
	password := "Pass123!"
	database := "store"
	dns := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, database)
	DB, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
			log.Panic("failed to connect to database")
	}
	DB.AutoMigrate(&domain.Exchange{})

	mySQLrepository := repository.NewRepository(DB)

	createUC := service.NewCreateUseCase(mySQLrepository)
	createHandler := handlers.NewCreateHandler(createUC)

	createWithRateUC := rate.NewCreateWithRateUseCase(mySQLrepository)
	createRHandler := handlers.NewCreateRHandler(createWithRateUC)

	getUC := service.NewGetUseCase(mySQLrepository)
	getHandler := handlers.NewGetHandler(getUC)

	getAllUC := service.NewGetAllUseCase(mySQLrepository)
	getAllHandler := handlers.NewGetAllHandler(getAllUC)

	getByFromUC := service.NewGetByFromUseCase(mySQLrepository)
	getByFromHandler := handlers.NewGetByFromHandler(getByFromUC)

	deleteUC := service.NewDeleteUseCase(mySQLrepository)
	deleteHandler := handlers.NewDeleteHandler(deleteUC)

	r := gin.Default()
	r.GET("/exchange/:amount/:from/:to/:rate", createHandler.Create)
	r.GET("/exchange/:amount/:from/:to", createRHandler.Create)
	r.GET("/exchange/id/:id", getHandler.Get)
	r.GET("/exchange", getAllHandler.GetAll)
	r.GET("/exchange/from/:from", getByFromHandler.GetByFrom)
	r.DELETE("/exchange/id/:id", deleteHandler.Delete)



	go cron.AutomatedRoutine(DB)

	r.Run(":8000")
}	
