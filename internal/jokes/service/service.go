package service

import (
	"fmt"

	"github.com/icelain/jokeapi"
)

type IService interface {
	GetRandomJoke() (*jokeapi.JokesResp, error)
}

type jokesService struct {
	client *jokeapi.JokeAPI
}

func New() *jokesService {
	fmt.Println("api")
	api := jokeapi.New()
	api.SetJokeType("twopart")
	service := &jokesService{
		client: api,
	}
	return service

}

func (s *jokesService) GetRandomJoke() (*jokeapi.JokesResp, error) {
	response, err := s.client.Fetch()
	if err != nil {
		return nil, fmt.Errorf("failed to get random joke from API: " + err.Error())
	}
	return &response, nil
}
