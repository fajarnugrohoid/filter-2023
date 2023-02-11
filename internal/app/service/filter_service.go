package service

import (
	"context"
	"new-filterization/internal/app/model/domain"
	"sync"
)

type SchoolOptionService interface {
	//GetSchoolOptionByLevelAndOpt(ctx context.Context, level string) map[string][]*domain.PpdbOption
	GetSchoolOptionByLevelAndOpt(ctx context.Context, level string, optType string) []*domain.PpdbOption
}
type RegistrationService interface {
	InitStudent(group *sync.WaitGroup, mutex sync.Mutex, messages chan *domain.PpdbOption, ctx context.Context, schoolByOpt []*domain.PpdbOption, level string, opt *domain.PpdbOption)
}

type FilterService struct {
	schoolOptionService SchoolOptionService
	registrationService RegistrationService
}

func NewFilterService(schoolOptionService SchoolOptionService, registrationService RegistrationService) *FilterService {
	return &FilterService{schoolOptionService, registrationService}
}

func (fs FilterService) InitSchoolOption(ctx context.Context, optionTypes []*domain.PpdbOption, schoolOption []*domain.PpdbOption, level string, optType string) []*domain.PpdbOption {

	var schoolOptionByType map[string][]*domain.PpdbOption
	schoolOptionByType[optType] = fs.schoolOptionService.GetSchoolOptionByLevelAndOpt(ctx, level, optType)

	group := &sync.WaitGroup{}
	var mutex sync.Mutex

	for _, opt := range schoolOptionByType[optType] {

		//fmt.Println("opt.id:", opt.Id, "-", opt.Name)
		/*
			var studentRegistrations []domain.PpdbRegistration
			studentRegistrations = controller.PpdbRegistrationService.FindByFirstChoiceLevel(ctx, level, opt.Type, opt.Id)
			studentHistories := make([]domain.PpdbRegistration, len(studentRegistrations), cap(studentRegistrations))
			copy(studentHistories, studentRegistrations)

			tmpOpt := &domain.PpdbOption{
				Id:                  opt.Id,
				Name:                opt.Name,
				Quota:               opt.QuotaOld,
				QuotaOld:            opt.QuotaOld,
				QuotaFrom:           map[string]int{"ketm": 0, "kondisi-tertentu": 0, "perpindahan": 0, "prestasi-rapor": 0, "prestasi": 0},
				Type:                opt.Type,
				Filtered:            0,
				UpdateQuota:         true,
				NeedQuota:           0,
				PpdbRegistration:    studentRegistrations,
				RegistrationHistory: studentHistories,
				HistoryShifting:     make([]domain.PpdbRegistration, 0),
			}*/

		var chanTmp = make(chan *domain.PpdbOption)
		defer close(chanTmp)

		go fs.registrationService.InitStudent(group, mutex, chanTmp, ctx, schoolOption, level, opt)
		tmpOpt := <-chanTmp

		optionTypes = append(optionTypes, tmpOpt)

	}
	//fs.registrationService.InitStudent(ctx, level)
	return schoolOption
}
