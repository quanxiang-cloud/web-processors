package config

// Config config.
type Config struct {
	Port         string
	UploadDir    string
	CommandDir   string
	UploadPrefix string

	PersonaEndpoint  string
	PersonaVersion   string
	PersonaKeySuffix string
}
