package service

import(
	moviesapi "github.com/LiaNestazi/moviesApi"
	"github.com/LiaNestazi/moviesApi/pkg/repository"
)

type ActorsListService struct{
	repo repository.ActorsList
}

func NewActorsListService(repo repository.ActorsList) *ActorsListService {
	return &ActorsListService{repo: repo}
}

func (s *ActorsListService) CreateActor(actor moviesapi.Actor) (int, error) {
	return s.repo.CreateActor(actor)
}

func (s *ActorsListService) GetActors() ([]moviesapi.Actor, error) {
	return s.repo.GetActors()
}

func (s *ActorsListService) DeleteActor(actorId int) error {
	return s.repo.DeleteActor(actorId)
}

func (s *ActorsListService) UpdateActor(actorId int, input moviesapi.Actor) error {
	return s.repo.UpdateActor(actorId, input)
}