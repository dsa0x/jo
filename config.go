package docman

type Config struct {
	PORT       string
	DbHost     string
	DbPort     string
	DbUser     string
	DbName     string
	DbPassword string
}

var Cfg *Config

func NewConfig() {
	cfg := &Config{}
	Cfg = cfg
}
