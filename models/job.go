package models

type Job struct {
	Name      string
	Namespace string
	Image     string
	Command   []string
	TTL       int32
}
