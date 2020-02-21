package config

// Environment defines SharePoint environment settings binding
type Environment struct {
	Name   string
	Code   string
	Config string
}

// Environments configuration
var Environments = []*Environment{
	{
		Name:   "SPO",
		Code:   "spo",
		Config: "./config/private.spo.json",
	},
	{
		Name:   "SPO (Target)",
		Code:   "spo.target",
		Config: "./config/private.spo.target.json",
	},
	{
		Name:   "SP 2019",
		Code:   "2019",
		Config: "./config/private.2019.json",
	},
	{
		Name:   "SP 2016",
		Code:   "2016",
		Config: "./config/private.2016.json",
	},
	{
		Name:   "SP 2013",
		Code:   "2013",
		Config: "./config/private.2013.json",
	},
}
