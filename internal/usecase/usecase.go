package usecase

type giftService interface {
	GetGifts()
}

type rateService interface {
	UpdateRates()
}

type healthService interface {
	DBHealth()
}

type UseCase struct {
	healthSvc healthService
	rateSvc   rateService
	giftsvc   giftService
}
