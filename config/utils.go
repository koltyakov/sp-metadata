package config

// GetEnvironmentsMap ...
func GetEnvironmentsMap() map[string]*Environment {
	e := map[string]*Environment{}
	for _, env := range Environments {
		e[env.Code] = env
	}
	return e
}

// GetEnvironmentsCodes ...
func GetEnvironmentsCodes() []string {
	var codes []string
	for _, env := range Environments {
		codes = append(codes, env.Code)
	}
	return codes
}
