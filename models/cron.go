package models

type Cron struct {
	Name      string
	Namespace string
	Image     string
	Command   []string
	TTL       int32
	Schedule  string
}
