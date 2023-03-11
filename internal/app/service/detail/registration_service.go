package detail

import (
	"context"
	"new-filterization/internal/app/model/domain"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegistrationRepository interface {
	GetByFirstChoiceLevel(ctx context.Context, optType string, schoolByOpt []*domain.PpdbOption, firstChoice primitive.ObjectID) []domain.PpdbRegistration
}

type RegistrationService struct {
	registrationRepo RegistrationRepository
}

func NewRegistrationService(registrationRepo RegistrationRepository) *RegistrationService {
	return &RegistrationService{registrationRepo}
}

func (service RegistrationService) InitStudent(group *sync.WaitGroup, mutex sync.Mutex,
	messages chan *domain.PpdbOption, ctx context.Context,
	schoolByOpt []*domain.PpdbOption, level string, opt *domain.PpdbOption) {
	//TODO implement me

	defer group.Done()
	group.Add(1)

	mutex.Lock()

	var studentRegistrations []domain.PpdbRegistration
	studentRegistrations = service.registrationRepo.GetByFirstChoiceLevel(ctx, "abk", schoolByOpt, opt.Id)

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
		//AddQuota:            0,
	}
	mutex.Unlock()
	//fmt.Println("RunAsyncInitStudentSenior done ", opt.Type, "-", opt.Name)
	//logger.Info("aft getFiltered", optType)
	messages <- tmpOpt
}

/*
func (controller InitialControllerImpl) InitSchoolStep1Senior(ctx context.Context, level string) map[string][]*domain.PpdbOption {
	//TODO implement me
	var optionTypes map[string][]*domain.PpdbOption
	var schoolOptions map[string][]*domain.PpdbOption
	optionTypes = map[string][]*domain.PpdbOption{}
	schoolOptions = map[string][]*domain.PpdbOption{}

	fmt.Println("start InitAll InitSchoolSenior")
	schoolOptions = controller.GetSchoolOptionStep1SeniorByType(ctx, schoolOptions, level)

	fmt.Println("finish InitAll InitSchoolSenior:", len(schoolOptions))

	group := &sync.WaitGroup{}
	var mutex sync.Mutex

	var channelAbk = make(chan []*domain.PpdbOption)
	var channelKetm = make(chan []*domain.PpdbOption)
	var channelKondisiTertentu = make(chan []*domain.PpdbOption)
	var channelPerpindahan = make(chan []*domain.PpdbOption)
	var channelPrestasi = make(chan []*domain.PpdbOption)
	var channelPrestasiRapor = make(chan []*domain.PpdbOption)

	defer close(channelAbk)
	defer close(channelKetm)
	defer close(channelKondisiTertentu)
	defer close(channelPerpindahan)
	defer close(channelPrestasi)
	defer close(channelPrestasiRapor)

	go controller.RunAsyncSenior(group, mutex, channelAbk, ctx, optionTypes["abk"], level, "abk", schoolOptions["abk"])
	go controller.RunAsyncSenior(group, mutex, channelKetm, ctx, optionTypes["ketm"], level, "ketm", schoolOptions["ketm"])
	go controller.RunAsyncSenior(group, mutex, channelKondisiTertentu, ctx, optionTypes["kondisi-tertentu"], level, "kondisi-tertentu", schoolOptions["kondisi-tertentu"])
	go controller.RunAsyncSenior(group, mutex, channelPerpindahan, ctx, optionTypes["perpindahan"], level, "perpindahan", schoolOptions["perpindahan"])
	go controller.RunAsyncSenior(group, mutex, channelPrestasi, ctx, optionTypes["prestasi"], level, "prestasi", schoolOptions["prestasi"])
	go controller.RunAsyncSenior(group, mutex, channelPrestasiRapor, ctx, optionTypes["prestasi-rapor"], level, "prestasi-rapor", schoolOptions["prestasi-rapor"])

	dataAbk := <-channelAbk
	dataKetm := <-channelKetm
	dataKondisiTertentu := <-channelKondisiTertentu
	dataPerpindahan := <-channelPerpindahan
	dataPrestasi := <-channelPrestasi
	dataPrestasiRapor := <-channelPrestasiRapor
	//dataZonasi := <-channelZonasi
	group.Wait()

	optionTypes["abk"] = dataAbk
	optionTypes["ketm"] = dataKetm
	optionTypes["kondisi-tertentu"] = dataKondisiTertentu
	optionTypes["perpindahan"] = dataPerpindahan
	optionTypes["prestasi"] = dataPrestasi
	optionTypes["prestasi-rapor"] = dataPrestasiRapor
	//optionTypes["zonasi"] = dataZonasi

	return optionTypes
}*/
