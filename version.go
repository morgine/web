package web

var Versions []*Version

type Version struct {
	Number   string
	UpdateAt int64
	Changes  []string
}
