package lg

import (
	"fmt"
	"log"
	"strconv"
)

// Logging levels
const (
	ErrorLevel   int = 2
	WarningLevel int = 1
	InfoLevel    int = 0
	DebugLevel   int = -1
)
const noTag int = -1

var levelMap = map[int]string{
	2:  "Error",
	1:  "Warning",
	0:  "Info",
	-1: "Debug",
}

type aLogTag struct {
	name    string
	group   string
	enabled bool
	level   int
}

var tags = make(map[int]*aLogTag)
var enabled = false
var logLevel = 0
var tagIDMax = 0
var tagLenMax = 0
var groupLenMax = 0

var (
	debugTag   int
	infoTag    int
	warningTag int
	errorTag   int
)

func init() {
	debugTag, _ = CreateTag("", "", DebugLevel)
	infoTag, _ = CreateTag("", "", InfoLevel)
	warningTag, _ = CreateTag("", "", WarningLevel)
	errorTag, _ = CreateTag("", "", ErrorLevel)
}

// Print prints a log item against a tag
func Print(tag int, s string, args ...interface{}) error {
	t, ok := tags[tag]
	if !ok {
		return fmt.Errorf("Unknown tag")
	}
	grpString := "%+" + strconv.Itoa(groupLenMax) + "v"
	tagString := grpString + "-%-" + strconv.Itoa(tagLenMax) + "v"
	if ok && enabled && t.enabled && t.level >= logLevel {
		s = fmt.Sprintf(s, args...)
		log.Printf("%+7v["+tagString+"]: %v", levelMap[t.level], t.group, t.name, s)
	}
	return nil
}

// Debug logs a line as debug
func Debug(s string, args ...interface{}) {
	Print(debugTag, s, args...)
}

// Info logs a line as info
func Info(s string, args ...interface{}) {
	Print(infoTag, s, args...)
}

// Warning logs a line as warning
func Warning(s string, args ...interface{}) {
	Print(warningTag, s, args...)
}

// Error logs a line as error
func Error(s string, args ...interface{}) {
	Print(errorTag, s, args...)
}

// CreateTag registers a new tag with set values
func CreateTag(name string, group string, level int) (id int, err error) {
	id = tagIDMax
	tags[id] = &aLogTag{name: name, group: group, enabled: true, level: level}
	tagIDMax++
	if len(name) > tagLenMax {
		tagLenMax = len(name)
	}
	if len(group) > groupLenMax {
		groupLenMax = len(group)
	}
	return id, nil
}

// EnableTag turns logginf on for this class
func EnableTag(tag int) {
	_, ok := tags[tag]
	if !ok {
		tags[tag] = &aLogTag{enabled: true, group: "default"}
	} else {
		tags[tag].enabled = true
	}
}

// DisableTag turns logginf off for this class
func DisableTag(tag int) {
	_, ok := tags[tag]
	if !ok {
		tags[tag] = &aLogTag{enabled: false, group: "default"}
	} else {
		tags[tag].enabled = false
	}
}

// SetTagLevel overrides the default (info) level of a tag
func SetTagLevel(tag int, level int) {
	_, ok := tags[tag]
	if !ok {
		tags[tag] = &aLogTag{enabled: false, group: "default", level: level}
	} else {
		tags[tag].level = level
	}
}

// DisableGroup Disables all tags in a certain group
func DisableGroup(group string) {
	for _, tag := range tags {
		if tag.group == group {
			tag.enabled = false
		}
	}
}

// EnableGroup Enables all tags in a certain group
func EnableGroup(group string) {
	for _, tag := range tags {
		if tag.group == group {
			tag.enabled = true
		}
	}
}

// Disable all logging
func Disable() {
	enabled = false
}

// Enable all logging
func Enable() {
	enabled = true
}

// SetLevel sets the overall logging level
func SetLevel(level int) {
	logLevel = level
}
