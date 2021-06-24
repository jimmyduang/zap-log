package zap_log

import (
	"testing"
)

func TestLogInit(t *testing.T) {
	logger, err := LogInit("test.log", false, "dev")
	if err != nil {
		t.Fatal("init log fail")
	}
	if logger == nil {
		t.Fatal("init log fail")
	}
	if logger.logger == nil {
		t.Fatal("init log fail")
	}

}

func TestLog_Info(t *testing.T) {
	logger, err := LogInit("test.log", true, "dev")
	if err != nil {
		t.Fatal("init log fail")
	}
	if logger == nil {
		t.Fatal("init log fail")
	}
	if logger.logger == nil {
		t.Fatal("init log fail")
	}
	logger.Info(map[string]interface{}{"foo": "bar"})
}
