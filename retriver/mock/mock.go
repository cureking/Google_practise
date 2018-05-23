package mock

type Restriever struct {
	Constents string
}

func (r Restriever) Post(url string, string map[string]string) string {
	panic("implement me")
}

func (r Restriever) Get(url string) string {
	return r.Constents
}
