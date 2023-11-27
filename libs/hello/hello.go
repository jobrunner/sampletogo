package hello

import (
	"errors"
	"fmt"
	"strings"
)

type Greeting struct {
	audience string
}

func NewAudience(audience string) (Greeting, error) {
	a := strings.TrimSpace(audience)
	if a == "" {
		return Greeting{}, errors.New("Audience must not be empty")
	}
	return Greeting{audience: a}, nil
}

func (g Greeting) Greet() string {
	return fmt.Sprintf("Hello, %s!", g.audience)
}

func (g Greeting) Audience() string {
	return g.audience
}
