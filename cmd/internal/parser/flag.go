package parser

import (
	"flag"
	"fmt"
	"time"
)

type ConsoleFlag struct {
	Url     string
	Depth   int
	Workers int
	Timeout time.Duration
}

func ParseFlag(consoleFlag *ConsoleFlag) error {
	flag.StringVar(&consoleFlag.Url, "url", "", "flag with start url")
	flag.IntVar(&consoleFlag.Depth, "depth", 3, "max crawl depth")
	flag.IntVar(&consoleFlag.Workers, "workers", 10, "max workers")
	flag.DurationVar(&consoleFlag.Timeout, "timeout", 5*time.Second, "request timeout")

	flag.Parse()

	return consoleFlag.Validate()
}

func (c *ConsoleFlag) Validate() error {
	switch {
	case c.Url == "":
		return fmt.Errorf("url is required")
	case c.Depth < 0:
		return fmt.Errorf("depth cannot be negative")
	case c.Workers <= 0:
		return fmt.Errorf("workers cannot be zero or negative")
	case c.Timeout <= 0:
		return fmt.Errorf("timeout cannot be zero or negative")
	}
	return nil
}
