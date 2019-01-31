package main

// https://play.golang.org/p/jwxy18pWmK3

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func main() {
	s := &Service{
		ID:   1,
		Name: "One",
	}
	err := s.Start(context.Background())
	if err != nil {
		log.Printf("one: %v\n", err)
		return
	}
}

type Service struct {
	ID   int
	Name string
}

func (s *Service) Start(ctx context.Context) error {
	if ctx == nil {
		return errors.New("nil context")
	}
	err := s.verify()
	if err != nil {
		return fmt.Errorf("run: %v", err)
	}
	fmt.Printf("%s: started\n", s.Name)
	return nil
}

func (s *Service) verify() error {
	if s.ID < 1 {
		return fmt.Errorf("ID %v < 1", s.ID)
	}
	if s.Name == "" {
		return errors.New("empty Name")
	}
	return nil
}
