package configs

type Configs struct {
	PostgreSQL PostgreSQL
	App        Fiber
}

type Fiber struct {
	Host string
	Post string
}

type PostgreSQL struct {
	Host     string
	Post     string
	Protocol string
	Username string
	Password string
	Database string
	SSLMode  string
}
