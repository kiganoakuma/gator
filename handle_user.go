package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("username required")
	}
	s.cfg.SetUser(cmd.Args[0])
	fmt.Printf("%s set as user", s.cfg.CurrentUserName)
	return nil

}
