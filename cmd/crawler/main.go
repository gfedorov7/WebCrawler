package main

import (
	"flag"
	"fmt"
	"strconv"
)

type ConsoleFlag struct {
	Url     string
	Depth   int
	Workers int
	Timeout int
	Limit   int
}

var (
	urlFlag     = flag.String("url", "", "flag with start url")
	depthFlag   = flag.Int("depth", 3, "max crawl depth")
	workersFlag = flag.Int("workers", 10, "max workers")
	timeoutFlag = flag.Int("timeout", 5, "request timeout")
	limitFlag   = flag.Int("limit", 10, "max page count")
)
var requiredFlag = []string{"url"}

func CrawlerRecover() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from", r)
	}
}

func ParseFlag(consoleFlag *ConsoleFlag, requiredFlags *[]string) error {
	flag.Parse()

	providedFlags := make(map[string]bool)
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() != "" {
			providedFlags[f.Name] = true
			consoleFlag.AddValue(f.Name, f.Value)
		}
	})

	for _, v := range *requiredFlags {
		if _, ok := providedFlags[v]; !ok {
			return fmt.Errorf("required flag `%s` not defined", v)
		}
	}

	return nil
}

func (c *ConsoleFlag) AddValue(name string, value flag.Value) {
	valueStringFormat := value.String()

	switch name {
	case "url":
		c.Url = valueStringFormat
	case "depth":
		c.Depth, _ = strconv.Atoi(valueStringFormat)
	case "workers":
		c.Workers, _ = strconv.Atoi(valueStringFormat)
	case "timeout":
		c.Timeout, _ = strconv.Atoi(valueStringFormat)
	case "limit":
		c.Limit, _ = strconv.Atoi(valueStringFormat)
	}
}

func main() {
	defer CrawlerRecover()

	consoleFlag := ConsoleFlag{}
	err := ParseFlag(&consoleFlag, &requiredFlag)

	if err != nil {
		panic(err)
	}
}
