package configs

var Conf *Config

type Config struct {
	App        *app        `yaml:"app"`
	DataSource *datasource `yaml:"datasource"`
}

type app struct {
	Name                  string `yaml:"name"`
	Env                   string `yaml:"env"`
	Port                  int    `yaml:"port"`
	PaginationDefaultSize int    `yaml:"pagination_default_size"`
	PaginationMaxSize     int    `yaml:"pagination_max_size"`
}

type datasource struct {
	MySQL struct {
		DsnWithDefault string `yaml:"dsn_with_default"`
		MaxOpen        int    `yaml:"max_open"`
		MaxIdle        int    `yaml:"max_idle"`
		MaxLifeTime    int    `yaml:"max_life_time"`
	} `yaml:"mysql"`
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
		PoolSize int    `yaml:"pool_size"`
	} `yaml:"redis"`
}
