package strategies

import (
	"fmt"
	"testing"
)

type Server struct {
	config int
}

func New() *Server {
	return &Server{
		config: 1,
	}
}

func TestDefineStrategy(t *testing.T) {
	server := Server{}
	fmt.Println("fmt.Println(\"\", server)\n", server)

}
