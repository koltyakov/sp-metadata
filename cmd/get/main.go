package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/go-xmlfmt/xmlfmt"
	"github.com/koltyakov/gosip"
	"github.com/koltyakov/gosip-sandbox/samples/dynauth"
	"github.com/koltyakov/gosip/api"

	"github.com/koltyakov/sp-metadata/config"
	"github.com/koltyakov/sp-metadata/pkg/edmx"
)

func main() {

	for _, env := range config.Environments {
		if err := envFlow(env); err != nil {
			log.Printf("Error while processing \"%s\": %v.\n", env.Name, err)
		}
	}

}

func envFlow(env *config.Environment) error {
	if !fileExists(env.Config) {
		log.Printf("Warning: Config \"%s\" doesn't exist, skipping.\n", env.Config)
		return nil
	}

	log.Printf("Processing \"%s\" environment.\n", env.Name)

	data, err := ioutil.ReadFile(env.Config)
	if err != nil {
		return err
	}

	conf := &struct {
		Strategy string `json:"strategy"`
	}{}

	if err := json.Unmarshal(data, &conf); err != nil {
		return err
	}

	if conf.Strategy == "" {
		return fmt.Errorf("no auth strategy provided")
	}

	authCnfg, err := dynauth.NewAuthCnfg(conf.Strategy, env.Config)
	if err != nil {
		return err
	}

	client := &gosip.SPClient{AuthCnfg: authCnfg}
	sp := api.NewSP(client)

	rsp, err := sp.Metadata()
	if err != nil {
		return err
	}

	x := strings.TrimSpace(
		xmlfmt.FormatXML(
			edmx.Clean(fmt.Sprintf("%s\n", rsp)),
			"",
			"  ",
		),
	)

	if err := ioutil.WriteFile(fmt.Sprintf("./edmx/%s.xml", env.Code), []byte(x), 0644); err != nil {
		return err
	}

	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
