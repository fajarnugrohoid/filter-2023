package handler

import (
	"context"
	"new-filterization/internal/app/model/domain"
)

type RegistrationService interface {
	//GetSchoolOptionByLevelAndOpt(ctx context.Context, level string) map[string][]*domain.PpdbOption
}

type RegistrationHandler struct {
	registrationService RegistrationService
}

func NewRegistrationHandler(registrationService RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{registrationService}
}

func (uh RegistrationHandler) GetBySchoolIdOption(ctx context.Context, level string) map[string][]*domain.PpdbOption {
	var schoolOption map[string][]*domain.PpdbOption
	//var schoolOption ["abk"]domain.PpdbOption

	//schoolOption = uh.registrationService.GetSchoolOptionByLevelAndOpt(ctx, level)
	/*
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"token": token,
		})*/
	return schoolOption
}
