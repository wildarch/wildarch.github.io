package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/yuin/goldmark"
)

type ConfigDescription struct {
	OneLine string
	Full    string
}

type NewsItem struct {
	Year    int
	Month   int
	Day     int
	Date    string
	Content template.HTML
}

type Config struct {
	Title       string
	Name        string
	Description ConfigDescription
	AboutMe     template.HTML
	News        []NewsItem
}

func renderMarkdownFile(path string) (string, error) {
	mdBytes, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", path, err)
	}

	var htmlBytes bytes.Buffer
	err = goldmark.Convert(mdBytes, &htmlBytes)
	if err != nil {
		return "", fmt.Errorf("failed to render markdown from %s: %w", path, err)
	}

	return htmlBytes.String(), nil
}

func parseConfig() Config {
	configText, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatalf("failed to read config.json: %v", err)
	}

	var config Config
	err = json.Unmarshal(configText, &config)
	if err != nil {
		log.Fatalf("failed to parse config.json: %v", err)
	}

	return config
}

var monthNames = [...]string{
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

func renderNewsItemFromFile(path string) (*NewsItem, error) {
	newsFilenameRegex := regexp.MustCompile(`(?P<Year>\d{4})-(?P<Month>\d{2})-(?P<Day>\d{2})-.+\.md$`)
	match := newsFilenameRegex.FindStringSubmatch(path)
	if match == nil {
		return nil, fmt.Errorf("invalid filename: %s", path)
	}
	year, err := strconv.Atoi(match[1])
	if err != nil {
		return nil, fmt.Errorf("not a valid year: %s", match[1])
	}
	month, err := strconv.Atoi(match[2])
	if err != nil {
		return nil, fmt.Errorf("not a valid month: %s", match[2])
	}
	day, err := strconv.Atoi(match[3])
	if err != nil {
		return nil, fmt.Errorf("not a valid day: %s", match[3])
	}

	html, err := renderMarkdownFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to render news item: %w", err)
	}

	return &NewsItem{
		Year:    year,
		Month:   month,
		Day:     day,
		Date:    fmt.Sprintf("%s'%d", monthNames[month], year-2000),
		Content: template.HTML(html),
	}, nil
}

func sortNewToOld(news []NewsItem) {
	sort.Slice(news, func(i, j int) bool {
		a := news[i]
		b := news[j]
		if a.Year > b.Year {
			return true
		}

		if a.Year == b.Year && a.Month > b.Month {
			return true
		}

		if a.Year == b.Year && a.Month == b.Month && a.Day > b.Day {
			return true
		}

		return false
	})
}

func main() {
	tmpl, err := template.ParseFiles("index.template.html")
	if err != nil {
		log.Fatalf("failed to open index.template.html: %v", err)
	}

	config := parseConfig()

	aboutMe, err := renderMarkdownFile("about_me.md")
	if err != nil {
		log.Fatalf("failed to render about me: %v", err)
	}
	config.AboutMe = template.HTML(aboutMe)

	newsDir, err := os.ReadDir("news/")
	if err != nil {
		log.Fatalf("failed to read news/: %v", err)
	}

	for _, newsDirEntry := range newsDir {
		path := "news/" + newsDirEntry.Name()
		item, err := renderNewsItemFromFile(path)
		if err != nil {
			log.Fatalf("failed to render news item from %s: %v", path, err)
		}
		config.News = append(config.News, *item)
	}
	sortNewToOld(config.News)

	out, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("failed to create out.html: %v", err)
	}

	tmpl.Execute(out, config)

	out.Close()
}
