package config

// Environment defines SharePoint environment settings binding
type Environment struct {
	Name   string
	Code   string
	Config string
}

// Environments configuration
var Environments = []*Environment{
	&Environment{
		Name:   "SharePoint Online (Standard Release)",
		Code:   "spo",
		Config: "./config/private.spo.json",
	},
	&Environment{
		Name:   "SharePoint Online (Target Release)",
		Code:   "spo.target",
		Config: "./config/private.spo.target.json",
	},
	&Environment{
		Name:   "SharePoint 2019",
		Code:   "2019",
		Config: "./config/private.2019.json",
	},
	&Environment{
		Name:   "SharePoint 2016",
		Code:   "2016",
		Config: "./config/private.2016.json",
	},
	&Environment{
		Name:   "SharePoint 2013",
		Code:   "2013",
		Config: "./config/private.2013.json",
	},
}
