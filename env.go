package zap_log

import (
	"bytes"
	"errors"
	"fmt"
)

var errUnmarshalNilEnv = errors.New("can't unmarshal a nil *Env")

type Env int8

const (
	Env__Dev Env = iota - 1
	Env__Test
	Env__Release
	Env__Product
)

func (e Env) String() string {
	switch e {
	case Env__Dev:
		return "dev"
	case Env__Test:
		return "test"
	case Env__Release:
		return "release"
	case Env__Product:
		return "product"
	default:
		return fmt.Sprintf("Env(%d)", e)
	}
}
func (e Env) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}
func (e *Env) UnmarshalText(text []byte) error {
	if e == nil {
		return errUnmarshalNilEnv
	}
	if !e.unmarshalText(text) && !e.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized env: %q", text)
	}
	return nil
}
func (e *Env) unmarshalText(text []byte) bool {
	switch string(text) {
	case "dev", "Dev", "DEV", "": // 默认为dev
		*e = Env__Dev
	case "test", "Test", "TEST":
		*e = Env__Test
	case "release", "Release", "RELEASE":
		*e = Env__Release
	case "product", "Product", "PRODUCT":
		*e = Env__Product
	default:
		return false
	}
	return true
}
func (e *Env) Set(s string) error {
	return e.UnmarshalText([]byte(s))
}
func (e *Env) Get() interface{} {
	return *e
}

// Enabled returns true if the given env is at or above this env.
func (e Env) Enabled(env Env) bool {
	return env >= e
}

type envEnabler interface {
	Enabled(Env) bool
}
