package config

type (
	Config struct {
		Commands []Command
	}

	Command struct {
		Name        string
		Description string
		Example     string
		Permission  string
	}
)
