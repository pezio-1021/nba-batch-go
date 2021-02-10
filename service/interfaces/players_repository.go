package interfaces

type PlayerRepository struct {
	dbHandler SQLHandler
}

func NewPlayerRepository(d SQLHandler) *PlayerRepository {
	return &PlayerRepository{
		dbHandler: d,
	}
}
