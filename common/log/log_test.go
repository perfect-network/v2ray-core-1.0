package log

import (
	"bytes"
	"log"
	"testing"

	"github.com/v2ray/v2ray-core/testing/unit"
)

func TestLogLevelSetting(t *testing.T) {
	assert := unit.Assert(t)

	assert.Pointer(debugLogger).Equals(noOpLoggerInstance)
	SetLogLevel(DebugLevel)
	assert.Pointer(debugLogger).Equals(streamLoggerInstance)

	SetLogLevel(InfoLevel)
	assert.Pointer(debugLogger).Equals(noOpLoggerInstance)
	assert.Pointer(infoLogger).Equals(streamLoggerInstance)
}

func TestStreamLogger(t *testing.T) {
	assert := unit.Assert(t)

	buffer := bytes.NewBuffer(make([]byte, 0, 1024))
	logger := &streamLogger{
		logger: log.New(buffer, "", 0),
	}
	logger.WriteLog("TestPrefix: ", "Test %s Format", "Stream Logger")
	assert.Bytes(buffer.Bytes()).Equals([]byte("TestPrefix: Test Stream Logger Format\n"))

	buffer.Reset()
	logger.WriteLog("TestPrefix: ", "Test No Format")
	assert.Bytes(buffer.Bytes()).Equals([]byte("TestPrefix: Test No Format\n"))
}
