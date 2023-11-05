package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"os"

	"github.com/yuin/goldmark"
)

type ConfigDescription struct {
	OneLine string
	Full    string
}

type Config struct {
	Title       string
	Name        string
	Description ConfigDescription
	AboutMe     template.HTML
}

func main() {
	tmpl, err := template.ParseFiles("index.template.html")
	if err != nil {
		log.Fatalf("failed to open index.template.html: %v", err)
	}

	configText, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("failed to read config.json: %v", err)
	}

	var config Config
	err = json.Unmarshal(configText, &config)
	if err != nil {
		log.Fatalf("failed to parse config.json: %v", err)
	}

	aboutMeMd, err := os.ReadFile("about_me.md")
	if err != nil {
		log.Fatalf("failed to reading about_me.md: %w", err)
	}

	var aboutMeHTML bytes.Buffer
	err = goldmark.Convert(aboutMeMd, &aboutMeHTML)
	config.AboutMe = template.HTML(aboutMeHTML.String())

	out, err := os.Create("out.html")
	if err != nil {
		log.Fatalf("failed to create out.html: %v", err)
	}

	tmpl.Execute(out, config)

	out.Close()
}
