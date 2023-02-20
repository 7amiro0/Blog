package storage

type Blog struct {
	Author string `yaml:"Author"`
	Title  string `yaml:"Title"`
	Body   string `yaml:"Body"`
}
