package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.params) == 0 {
		return fmt.Errorf("no username given")
	}

	err := s.conf.SetUser(cmd.params[0])
	if err != nil {
		return err
	}
	fmt.Println("Set user to:", s.conf.CurrentUserName)
	return nil
}
