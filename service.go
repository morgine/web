package web

type Service struct {
	Title   string
	Short   string
	Comment string
	In      interface{}
	Out     interface{}
	Errs    []*ServErr
	H       HandlerFunc
}

type ServErr struct {
	Code int
	Msg  string
}
