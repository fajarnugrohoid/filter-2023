package service

import (
	"context"
	"fmt"
	"new-filterization/internal/app/model/domain"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SchoolOptionService interface {
	//GetSchoolOptionByLevelAndOpt(ctx context.Context, level string) map[string][]*domain.PpdbOption
	GetSchoolOptionByLevelAndOpt(ctx context.Context, level string, optType string) []*domain.PpdbOption
	GetSchoolOptionTemporaryByLevel(ctx context.Context, level string, optType string) []*domain.PpdbOption
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

func (fs FilterService) MainInit(ctx context.Context) {

	var optionTypes = map[string][]*domain.PpdbOption{}
	var schoolOptions = map[string][]*domain.PpdbOption{}

	defer group.Done()
	group.Add(1)

	//logger.Info("bef getFiltered:", optType)
	fmt.Println("RunAsynInitStudentByOptType start ", optionTypes)
	//mutex.Lock()

	optionTypes = fs.InitSchoolOption(ctx, optionTypes, schoolOption, level, optType)

	//mutex.Unlock()
	fmt.Println("RunAsynInitStudentByOptType done ", optionTypes)
	//logger.Info("aft getFiltered", optType)

}

func (fs FilterService) InitSchoolOption(ctx context.Context, optionTypes []*domain.PpdbOption, schoolOption []*domain.PpdbOption, level string, optType string) []*domain.PpdbOption {

	var schoolOptionByType map[string][]*domain.PpdbOption
	schoolOptionByType[optType] = fs.schoolOptionService.GetSchoolOptionByLevelAndOpt(ctx, level, optType)

	group := &sync.WaitGroup{}
	var mutex sync.Mutex

	for _, opt := range schoolOptionByType[optType] {

		//fmt.Println("opt.id:", opt.Id, "-", opt.Name)

		var chanTmp = make(chan *domain.PpdbOption)
		defer close(chanTmp)

		go fs.registrationService.InitStudent(group, mutex, chanTmp, ctx, schoolOption, level, opt)
		tmpOpt := <-chanTmp

		optionTypes = append(optionTypes, tmpOpt)

	}

	var schoolOptionTemporary map[string][]*domain.PpdbOption
	schoolOptionTemporary[optType] = fs.schoolOptionService.GetSchoolOptionTemporaryByLevel(ctx, level, optType)
	for _, opt := range schoolOptionTemporary[optType] {
		TmpIdSMAAbk, _ := primitive.ObjectIDFromHex("000000000000000000000011")
		tmpSMAAbk := &domain.PpdbOption{
			Id:                  TmpIdSMAAbk,
			Name:                "TemporarySMAAbk",
			Type:                "abk",
			QuotaOld:            0,
			Quota:               0,
			QuotaFrom:           map[string]int{},
			Filtered:            1,
			UpdateQuota:         false,
			PpdbRegistration:    nil,
			RegistrationHistory: nil,
			HistoryShifting:     nil,
		}
		switch optType {
		case "abk":
			optionTypes = append(optionTypes, tmpSMAAbk)
			break
		}
	}

	/*

		TmpIdSMAKetm, _ := primitive.ObjectIDFromHex("000000000000000000000012")
		TmpIdSMAKondisiTertentu, _ := primitive.ObjectIDFromHex("000000000000000000000013")
		TmpIdSMAPerpindahan, _ := primitive.ObjectIDFromHex("000000000000000000000014")
		TmpIdSMAPrestasiNilaiRapor, _ := primitive.ObjectIDFromHex("000000000000000000000016")
		TmpIdSMAPrestasi, _ := primitive.ObjectIDFromHex("000000000000000000000017")
		TmpIdSMAZonasi, _ := primitive.ObjectIDFromHex("000000000000000000000018")


		tmpSMAKetm := &domain.PpdbOption{
			Id:                  TmpIdSMAKetm,
			Name:                "TemporarySMAKetm",
			Type:                "ketm",
			QuotaOld:            0,
			Quota:               0,
			QuotaFrom:           map[string]int{},
			Filtered:            1,
			UpdateQuota:         false,
			PpdbRegistration:    nil,
			RegistrationHistory: nil,
			HistoryShifting:     nil,
		}
		tmpSMAKondisiTertentu := &domain.PpdbOption{
			Id:                  TmpIdSMAKondisiTertentu,
			Name:                "TemporarySMAKondisiTertentu",
			Type:                "kondisi-tertentu",
			QuotaOld:            0,
			Quota:               0,
			QuotaFrom:           map[string]int{},
			Filtered:            1,
			UpdateQuota:         false,
			PpdbRegistration:    nil,
			RegistrationHistory: nil,
			HistoryShifting:     nil,
		}
		tmpSMAPerpindahan := &domain.PpdbOption{
			Id:                  TmpIdSMAPerpindahan,
			Name:                "TemporarySMAPerpindahan",
			Type:                "perpindahan",
			QuotaOld:            0,
			Quota:               0,
			QuotaFrom:           map[string]int{},
			Filtered:            1,
			UpdateQuota:         false,
			PpdbRegistration:    nil,
			RegistrationHistory: nil,
			HistoryShifting:     nil,
		}

		tmpSMAPrestasiRapor := &domain.PpdbOption{
			Id:                  TmpIdSMAPrestasiNilaiRapor,
			Name:                "TemporarySMAPrestasiRapor",
			Type:                "prestasi-rapor",
			QuotaOld:            0,
			Quota:               0,
			QuotaFrom:           map[string]int{},
			Filtered:            1,
			UpdateQuota:         false,
			PpdbRegistration:    nil,
			RegistrationHistory: nil,
			HistoryShifting:     nil,
		}
		tmpSMAPrestasi := &domain.PpdbOption{
			Id:                  TmpIdSMAPrestasi,
			Name:                "TemporarySMAPrestasi",
			Type:                "prestasi",
			QuotaOld:            0,
			Quota:               0,
			QuotaFrom:           map[string]int{},
			Filtered:            1,
			UpdateQuota:         false,
			PpdbRegistration:    nil,
			RegistrationHistory: nil,
			HistoryShifting:     nil,
		}
		tmpSMAZonasi := &domain.PpdbOption{
			Id:                  TmpIdSMAZonasi,
			Name:                "TemporarySMAZonasi",
			Type:                "zonasi",
			QuotaOld:            0,
			Quota:               0,
			QuotaFrom:           map[string]int{},
			Filtered:            1,
			UpdateQuota:         false,
			PpdbRegistration:    nil,
			RegistrationHistory: nil,
			HistoryShifting:     nil,
		}

		switch optType {
		case "abk":
			optionTypes = append(optionTypes, tmpSMAAbk)
			break

		case "ketm":
			optionTypes = append(optionTypes, tmpSMAKetm)
			break

		case "kondisi-tertentu":
			optionTypes = append(optionTypes, tmpSMAKondisiTertentu)
			break

		case "perpindahan":
			optionTypes = append(optionTypes, tmpSMAPerpindahan)
			break

		case "prestasi-rapor":
			optionTypes = append(optionTypes, tmpSMAPrestasi)
			break

		case "prestasi":
			optionTypes = append(optionTypes, tmpSMAPrestasiRapor)
			break

		case "zonasi":
			optionTypes = append(optionTypes, tmpSMAZonasi)
			break

		}

	*/

	fmt.Println("InitSeniorStudentByOptType with optType:", optType, " end")

	return schoolOption
}
