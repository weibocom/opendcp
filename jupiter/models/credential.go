package models

type Credential struct {
	name   string `required:"true" description:"key name"`
	key    string `required:"true"`
	secret string `required:"true"`
}
