/*
Package logger contains a simple logger. The logger is thread-safe.
*/

package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Logger struct {
	mu         sync.Mutex     // ensures atomic writes; protects the following fields
	prefix     string         // prefix to write at beginning of each line
	flag       int            // properties
	out        io.WriteCloser // destination for output
	buf        []byte         // for accumulating text to write
	level      int            // One of DEBUG, ERROR, INFO
	consoleLog bool           // log in console if true
	slackUrl   string         // url for Slack webhook
}

const (
	DEBUG = 1 << iota
	INFO
	ERROR
)

var DefaultLogger *Logger

func init() {
	f, _ := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	DefaultLogger = &Logger{
		out:        f,
		consoleLog: true,
		slackUrl:   "https://hooks.slack.com/services/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",
	}
}

// Error logs messages of level ERROR
func Error(v interface{}) {
	DefaultLogger.mu.Lock()
	defer DefaultLogger.mu.Unlock()
	if DefaultLogger.level <= ERROR {
		s := fmt.Sprintf(time.Now().Local().Format("2006/01/02 15:04:05")+" ERROR: %v\n", v)
		logInConsole(s)
		DefaultLogger.out.Write([]byte(s))
		if v != nil {
			err := postSlack(s)
			if err != nil {
				logInConsole(err.Error())
				DefaultLogger.out.Write([]byte(err.Error() + "\n"))
			}
		}
	}
}

// LogError is a convenience function to log errors of type error
func LogError(err error) {
	if err != nil {
		Error(err)
	}
}

// Info log messages of level INFO
func Info(v interface{}) {
	DefaultLogger.mu.Lock()
	defer DefaultLogger.mu.Unlock()
	if DefaultLogger.level <= INFO {
		s := fmt.Sprintf(time.Now().Local().Format("2006/01/02 15:04:05")+" INFO: %s\n", PrettyPrint(v))
		logInConsole(s)
		DefaultLogger.out.Write([]byte(s))
	}
}

// Debug logs messages of level DEBUG
func Debug(v interface{}) {
	DefaultLogger.mu.Lock()
	defer DefaultLogger.mu.Unlock()
	if DefaultLogger.level <= DEBUG {
		s := fmt.Sprintf(time.Now().Local().Format("2006/01/02 15:04:05")+" DEBUG: %s\n", PrettyPrint(v))
		logInConsole(s)
		DefaultLogger.out.Write([]byte(s))
	}
}

// Slack logs messages to Slack
func Slack(v interface{}) {
	DefaultLogger.mu.Lock()
	defer DefaultLogger.mu.Unlock()
	s := fmt.Sprintf(time.Now().Local().Format("2006/01/02 15:04:05 %s\n"), PrettyPrint(v))
	postSlack(s)
}

// SetLogLevel sets a log level for thg logger
func SetLogLevel(lvl int) {
	DefaultLogger.mu.Lock()
	defer DefaultLogger.mu.Unlock()
	DefaultLogger.level = lvl
}

// NewFile creates a new file to log to.
func NewFile(filename string) error {
	DefaultLogger.mu.Lock()
	defer DefaultLogger.mu.Unlock()
	Close()
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	DefaultLogger.out = f
	return nil
}

// Close closes the output file
func Close() {
	DefaultLogger.out.Close()
}

// PrettyPrint prints structs nicely
func PrettyPrint(v interface{}) string {
	pretty, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return "Error formatting string"
	}
	return string(pretty)
}

func logInConsole(v interface{}) {
	if DefaultLogger.consoleLog {
		log.Println(v)
	}
}

func postSlack(s string) error {
	text, err := json.Marshal(struct {
		Text string `json:"text"`
	}{s})
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", DefaultLogger.slackUrl, bytes.NewBuffer(text))
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		resp.Body.Close()
		return fmt.Errorf("Post failed: %s", resp.Status)
	}
	return nil
}