package config

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Config struct {
	Commands map[string]CliCommand
	Next     *string
	Previous *string
}
