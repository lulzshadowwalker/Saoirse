package bot

import "fmt"

type BotConfig struct {
	Token  string
	Prefix string `json:"command-prefix"`
}

func (c BotConfig) String() string {
	return fmt.Sprintf(`
token: %q
command-prefix: %q
`, c.Token, c.Prefix)
}
