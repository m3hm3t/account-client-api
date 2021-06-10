package test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Setup()
	code := m.Run()
	Shutdown()
	os.Exit(code)
}
