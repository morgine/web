package web

type Service struct {
	Title   string
	Short   string
	Comment string
	In      interface{}
	Out     interface{}
	Errs    []*ServErr
	Version string
	H       HandlerFunc
}

func (s *Service) Handle(ctx *Context) error {
	panic("implement me")
}

type ServErr struct {
	Code int
	Msg  string
}
