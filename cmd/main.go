package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"new-filterization/internal/app/database"
	"new-filterization/internal/app/handler"
	"new-filterization/internal/app/helper"
	"new-filterization/internal/app/repository"
	"new-filterization/internal/app/service/detail"
	"os"
	"time"
)

func main() {

	start := time.Now()

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	arg := os.Args[1]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	logger := helper.NewLogger(arg)

	database := database.NewDatabaseConn()
	seniorSchoolOptionRepo := repository.NewSeniorSchoolOptionRepository(database)
	vocationalSchoolOptionRepo := repository.NewVocationalSchoolOptionRepository(database)
	schoolOptionService := detail.NewSchoolOptionService(seniorSchoolOptionRepo, vocationalSchoolOptionRepo)

	registrationRepo := repository.NewRegistrationRepository(database)
	registrationService := detail.NewRegistrationService(registrationRepo)

	registrationHandler := handler.NewRegistrationHandler(schoolOptionService)

	//route.POST("/login", userHandler.Login)

	/*
		//ppdbRegistrationRepository := repository.NewPpdbRegistrationRepositoy()
		ppdbFilteredRepository := repository.NewPpdbFilteredRepository()
		ppdbStatisticRepository := repository.NewPpdbStatisticRepository()
		ppdbSchoolOptionRepository := repository.NewPpdbSchoolOptionRepositoy()

		ppdbRegistrationService := service.NewPpdbRegistrationService(ppdbRegistrationRepository, database)
		ppdbFilteredService := service.NewPpdbFilteredService(ppdbFilteredRepository, database)
		ppdbStatisticService := service.NewPpdbStatisticService(ppdbStatisticRepository, database)
		ppdbSchoolOptionService := service.NewPpdbSchoolOptionService(ppdbSchoolOptionRepository, database)

		initialController := controller.NewInitialController(ppdbSchoolOptionService, ppdbRegistrationService)
		finishController := controller.NewFinishController(ppdbFilteredService, ppdbStatisticService, ppdbRegistrationService, ppdbSchoolOptionService)

		var optionTypes map[string][]*domain.PpdbOption
		optionTypes = map[string][]*domain.PpdbOption{}
		var level string

		switch arg {
		case "sma-tahap1":
			level = "sma"
			optionTypes = initialController.InitStep1(ctx, level)
			optionTypes = logic.DoFilterSeniorStep1(optionTypes, logger)
			finishController.UpdatePpdbFilteredPpdbStatisticPpdbRegistrationPpdbOption(ctx, optionTypes, level, arg, logger)
			break
		case "sma-tahap2":
			level = "sma"
			optionTypes = initialController.InitStep2(ctx, level)
			optionTypes = logic.DoFilterSeniorStep2(optionTypes, logger)
			finishController.UpdatePpdbFilteredPpdbStatisticPpdbRegistrationPpdbOption(ctx, optionTypes, level, arg, logger)
			break
		case "smk-tahap1":
			level = "smk"
			fmt.Println("initialController.InitAll")
			optionTypes = initialController.InitStep1(ctx, level)
			optionTypes = logic.DoFilterVocationalStep1(optionTypes, logger)
			finishController.UpdatePpdbFilteredPpdbStatisticPpdbRegistrationPpdbOption(ctx, optionTypes, level, arg, logger)
			break
		case "smk-tahap2":
			level = "smk"
			optionTypes = initialController.InitStep2(ctx, level)
			optionTypes = logic.DoFilterVocationalStep2(optionTypes, logger)
			finishController.UpdatePpdbFilteredPpdbStatisticPpdbRegistrationPpdbOption(ctx, optionTypes, level, arg, logger)
			break
		case "sma-sendquotatahap1totahap2":
			level = "sma"
			logic.SendAllQuotaSeniorTahap1ToTahap2(optionTypes, logger)
			break
		case "smk-sendquotatahap1totahap2":
			level = "smk"
			logic.SendAllQuotaVocationalTahap1ToTahap2(optionTypes, logger)
			break
		}
	*/

	timeElapsed := time.Since(start)
	logger.Info("The `for` loop took %s", timeElapsed)
}
