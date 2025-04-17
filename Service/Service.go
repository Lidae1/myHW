package Service

import "fmt"

type Producer interface {
	Produce(path string) string
}

type Presenter interface {
	Present(string) string
}

type Service struct {
	prod Producer
	pres Presenter
}

func NewService(prod Producer, pres Presenter) *Service {
	return &Service{prod: prod, pres: pres}
}

func (s *Service) Run(path string) {
	data := s.prod.Produce(path)
	fmt.Println(s.pres.Present(data))

}
