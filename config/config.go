package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

var config = make(map[string]any)

func init() {
	var path string

	if file, ok := os.LookupEnv("CONFIG_FILE"); ok {
		path = file
	} else {
		path = "config.toml"
	}

	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(fmt.Errorf("config io: %w", err))
		return
	}

	if err = toml.Unmarshal(data, &config); err != nil {
		fmt.Println(fmt.Errorf("config parsing: %w", err))
		return
	}

	// allows us to configure other features like Postgres via environment variables
	if env, ok := config["env"]; ok {
		for key, value := range env.(map[string]any) {
			err = os.Setenv(key, fmt.Sprint(value))
		}
	}

	abs, _ := filepath.Abs(path)
	log.Printf("Configuration file: %q\n", abs)
}

func value(path string) (value any) {
	root := config
	keys := strings.Split(path, ".")
	for i := 0; i < len(keys); i++ {
		value = root[keys[i]]

		if value == nil {
			return
		} else if i == len(keys)-1 {
			return
		} else if next, ok := value.(map[string]any); ok {
			root = next
		} else {
			return
		}
	}
	return
}

func IntOrDefault(path string, default_value int) int {
	if value := value(path); value != nil {
		n, _ := strconv.ParseInt(fmt.Sprint(value), 0, 0)
		return int(n)
	}
	return default_value
}

func StringOrDefault(path, default_value string) string {
	if value := value(path); value == nil {
		return default_value
	} else {
		return fmt.Sprint(value)
	}
}

func DurationOrDefault(path string, default_value time.Duration) time.Duration {
	if value := String(path); value != "" {
		duration, _ := time.ParseDuration(value)
		return duration
	}
	return default_value
}

func BoolOrDefault(path string, default_value bool) bool {
	if value := value(path); value == nil {
		return default_value
	} else {
		return value.(bool)
	}
}

// Bool returns true or false depending on the value of `path`. If `path` is not defined then panic.
func Bool(path string) bool {
	if value := value(path); value != nil {
		switch value := value.(type) {
		case bool:
			return value
		case string:
			b, _ := strconv.ParseBool(value)
			return b
		}
	}
	panic(fmt.Sprint("bool not defined:", path))
}

// Stringf is shorthand for String(fmt.Sprintf(format, args...))
func Stringf(format string, args ...any) string {
	return String(fmt.Sprintf(format, args...))
}

// String returns the string value of `path`. If `path` is not defined then panic.
func String(path string) (value string) {
	value = StringOrDefault(path, "")
	if value == "" {
		panic(fmt.Sprint("string not defined:", path))
	}
	return
}
