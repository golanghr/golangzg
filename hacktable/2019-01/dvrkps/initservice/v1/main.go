package main

// https://play.golang.org/p/eRaZXLE0E-C

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func main() {
	s, err := NewService(&Config{
		ID:   1,
		Name: "One",
	})

	if err != nil {
		log.Printf("one: %v\n", err)
		return
	}
	err = s.Start(context.Background())
	if err != nil {
		log.Printf("one: %v\n", err)
		return
	}
}

type Config struct {
	ID   int
	Name string
}

func (c *Config) Verify() error {
	if c.ID < 1 {
		return fmt.Errorf("ID %v < 0", c.ID)
	}
	if c.Name == "" {
		return errors.New("empty Name")
	}
	return nil
}

type Service struct {
	id   int
	name string
}

func NewService(c *Config) (*Service, error) {
	err := c.Verify()
	if err != nil {
		return nil, fmt.Errorf("new service: %v", err)
	}
	s := Service{
		id:   c.ID,
		name: c.Name,
	}
	return &s, nil
}

func (s *Service) Start(ctx context.Context) error {
	if ctx == nil {
		return errors.New("nil context")
	}
	fmt.Printf("%s: started\n", s.name)
	return nil
}
