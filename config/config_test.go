package config

import (
	"fmt"
	"os"
	"testing"
)

func TestNew_config(t *testing.T) {
	wantPort := 3333
	os.Setenv("PORT", fmt.Sprintf("%d", wantPort))

	// test port
	got, err := New()
	if err != nil {
		t.Errorf("failed to create config")
	}

	// めっちゃtableテストにした方がいい気がしている
	if got.Port != wantPort {
		t.Errorf("invalid port address want %v, but got %v", wantPort, got.Port)
	}
	// test default env
	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Errorf("invalid port address want %v, but got %v", wantPort, got.Port)
	}
}
