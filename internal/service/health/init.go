package service_health

type HealthChecker interface {
	CheckStatus() bool
}

type Service struct {
	postgresRepo HealthChecker
	redisRepo    HealthChecker
}

func NewService(pgR HealthChecker, rdsR HealthChecker) *Service {
	return &Service{
		postgresRepo: pgR,
		redisRepo:    rdsR,
	}
}
