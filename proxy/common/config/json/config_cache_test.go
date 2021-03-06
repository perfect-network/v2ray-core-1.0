package json

import (
	"testing"

	"github.com/v2ray/v2ray-core/proxy/common/config"
	"github.com/v2ray/v2ray-core/testing/unit"
)

func TestRegisterInboundConfig(t *testing.T) {
	assert := unit.Assert(t)
	initializeConfigCache()

	protocol := "test_protocol"
	creator := func() interface{} {
		return true
	}

	err := RegisterInboundConnectionConfig(protocol, creator)
	assert.Error(err).IsNil()

	configObj := CreateConfig(protocol, config.TypeInbound)
	assert.Bool(configObj.(bool)).IsTrue()

	configObj = CreateConfig(protocol, config.TypeOutbound)
	assert.Pointer(configObj).IsNil()
}

func TestRegisterOutboundConfig(t *testing.T) {
	assert := unit.Assert(t)
	initializeConfigCache()

	protocol := "test_protocol"
	creator := func() interface{} {
		return true
	}

	err := RegisterOutboundConnectionConfig(protocol, creator)
	assert.Error(err).IsNil()

	configObj := CreateConfig(protocol, config.TypeOutbound)
	assert.Bool(configObj.(bool)).IsTrue()

	configObj = CreateConfig(protocol, config.TypeInbound)
	assert.Pointer(configObj).IsNil()
}
