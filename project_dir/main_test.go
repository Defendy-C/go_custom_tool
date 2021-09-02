package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCase(t *testing.T) {
	path01, _ := os.Getwd()
	path02 := GetProjectDir() + "/project_dir"
	
	assert.Equal(t, path01, path02)
}
