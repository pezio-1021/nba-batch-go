package usecase

type UseCase struct {
	playerRepo PlayerRepository
}

func NewUsecase(p PlayerRepository) *UseCase {
	return &UseCase{
		playerRepo: p,
	}
}
