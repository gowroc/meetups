package main

type Service struct{}

func (s *Service) Close() error {
	return s.doSomething()
}

func (s *Service) User(id int) string {
	return "name"
}

func (s *Service) Users() []string {
	return []string{"name1", "name2"}
}

func (s *Service) SetUser(id int, name string) error {
	return nil
}
