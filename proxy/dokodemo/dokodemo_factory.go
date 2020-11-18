package dokodemo

import (
	"github.com/v2ray/v2ray-core/app"
	"github.com/v2ray/v2ray-core/proxy/common/connhandler"
	"github.com/v2ray/v2ray-core/proxy/dokodemo/config/json"
)

type DokodemoDoorFactory struct {
}

func (this DokodemoDoorFactory) Create(dispatcher app.PacketDispatcher, rawConfig interface{}) (connhandler.InboundConnectionHandler, error) {
	config := rawConfig.(*json.DokodemoConfig)
	return NewDokodemoDoor(dispatcher, config), nil
}

func init() {
	connhandler.RegisterInboundConnectionHandlerFactory("dokodemo-door", DokodemoDoorFactory{})
}
