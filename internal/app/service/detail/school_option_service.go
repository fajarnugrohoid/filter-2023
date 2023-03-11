package detail

import (
	"context"
	"fmt"
	"new-filterization/internal/app/model/domain"
)

type SeniorSchoolOptionRepository interface {
	GetSeniorSchoolOptionByOptType(ctx context.Context, level string, optType string) []*domain.PpdbOption
	GetSeniorSchoolOptionTemporary(ctx context.Context, level string, optType string) []*domain.PpdbOption
}
type VocationalSchoolOptionRepository interface {
	GetVocationalSchoolOptionByOptType(ctx context.Context, level string, optType string) []*domain.PpdbOption
	GetVocationalSchoolOptionTemporary(ctx context.Context, level string, optType string) []*domain.PpdbOption
}

type SchoolOptionService struct {
	seniorSchoolOptionRepo     SeniorSchoolOptionRepository
	vocationalSchoolOptionRepo VocationalSchoolOptionRepository
}

func NewSchoolOptionService(seniorSchoolOptionRepo SeniorSchoolOptionRepository, vocationalSchoolOptionRepo VocationalSchoolOptionRepository) *SchoolOptionService {
	return &SchoolOptionService{seniorSchoolOptionRepo, vocationalSchoolOptionRepo}
}

/*
func (service SchoolOptionService) GetSchoolOptionByLevelAndOpt(ctx context.Context, level string) map[string][]*domain.PpdbOption {
	//TODO implement me
	var schoolOption map[string][]*domain.PpdbOption
	schoolOption["abk"] = service.getSchoolOptionByLevelAndOptType(ctx, level, "abk")
	schoolOption["ketm"] = service.getSchoolOptionByLevelAndOptType(ctx, level, "ketm")

	return schoolOption
} */

func (service SchoolOptionService) GetSchoolOptionByLevelAndOpt(ctx context.Context, level string, optionType string) []*domain.PpdbOption {
	//TODO implement me
	var schoolOption []*domain.PpdbOption
	schoolOption = service.getSchoolOptionByLevelAndOptType(ctx, level, optionType)

	return schoolOption
}

func (service SchoolOptionService) GetSchoolOptionTemporaryByLevel(ctx context.Context, level string, optionType string) []*domain.PpdbOption {
	//TODO implement me
	var schoolOption []*domain.PpdbOption
	schoolOption = service.getSchoolOptionTemporaryCheckLevel(ctx, level, optionType)

	return schoolOption
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

func (service SchoolOptionService) getSchoolOptionByLevelAndOptType(ctx context.Context, level string, optType string) []*domain.PpdbOption {
	//TODO implement me

	var ppdbSchoolOptions []*domain.PpdbOption
	if level == "senior" {
		ppdbSchoolOptions = service.seniorSchoolOptionRepo.GetSeniorSchoolOptionByOptType(ctx, level, optType)
	} else {
		ppdbSchoolOptions = service.vocationalSchoolOptionRepo.GetVocationalSchoolOptionByOptType(ctx, level, optType)
	}

	fmt.Println("ppdbSchoolOptions ", optType, ":", len(ppdbSchoolOptions))
	return ppdbSchoolOptions
}

func (service SchoolOptionService) getSchoolOptionTemporaryCheckLevel(ctx context.Context, level string, optType string) []*domain.PpdbOption {
	//TODO implement me

	var ppdbSchoolOptions []*domain.PpdbOption
	if level == "senior" {
		ppdbSchoolOptions = service.seniorSchoolOptionRepo.GetSeniorSchoolOptionTemporary(ctx, level, optType)
	} else {
		ppdbSchoolOptions = service.vocationalSchoolOptionRepo.GetVocationalSchoolOptionTemporary(ctx, level, optType)
	}

	fmt.Println("ppdbSchoolOptions ", optType, ":", len(ppdbSchoolOptions))
	return ppdbSchoolOptions
}

/*
func (us UserService) Login(ctx context.Context, param domain.LoginParam) (string, error) {
	user, err := us.userRepo.GetByEmail(ctx, param.Email)
	if err != nil {
		return "", errors.New("wrong email/password")
	}

	if err := user.ComparePassword(param.Password); err != nil {
		return "", errors.New("wrong email/password")
	}
	return generateJWT(user.ID)
}

func (us UserService) Register(ctx context.Context, param domain.RegisterParam) (string, error) {
	user := domain.NewUser(param.Name, param.Email, param.Password)
	if err := us.userRepo.Create(ctx, user); err != nil {
		return "", errors.New("failed create user")
	}
	return generateJWT(user.ID)
}
*/
