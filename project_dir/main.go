package main

import (
	"regexp"
	"runtime"
)

func GetProjectDir() string {
	_, file, _, _ := runtime.Caller(0)
	return regexp.MustCompile("/project_dir/main.go").ReplaceAllString(file, "")
}
