package main

type globalConfig struct {
	TodoDir string `toml:"todo_dir"`
}

func getDefaultConfig() globalConfig {
	return globalConfig{
		TodoDir: ".",
	}
}
