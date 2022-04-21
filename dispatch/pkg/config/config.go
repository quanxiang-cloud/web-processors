package config

// Config config.
type Config struct {
	Port            string
	UploadDir       string
	CommandDir      string
	StorePathPrefix string

	PersonaEndpoint  string
	PersonaVersion   string
	PersonaKeySuffix string
}
