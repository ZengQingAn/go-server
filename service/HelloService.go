package service

type HelloService struct{}

func (s *HelloService) SayHello(name string) (string, error) {
	return "Hello, " + name, nil
}
