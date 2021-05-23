package sleep

import (
	"testing"
	"time"
)

func TestTmpExcutable(t *testing.T) {
	t.Log("TestTmpExcutable")
	time.Sleep(2 * time.Minute)
}
