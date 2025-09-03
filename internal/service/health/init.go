package service_health

type healthRepository interface {
	CheckStatus() bool
}

type Service struct {
	healthRepo healthRepository
}

func NewService(hr healthRepository) *Service {
	return &Service{
		healthRepo: hr,
	}
}
