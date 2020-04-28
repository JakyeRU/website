package config

type (
	Config struct {
		Server   Server
		Commands []Command
	}

	Server struct {
		Address string
	}

	Command struct {
		Name        string
		Description string
		Example     string
		Permission  string
	}
)
