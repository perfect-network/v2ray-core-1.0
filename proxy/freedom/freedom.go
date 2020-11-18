package freedom

import (
	"net"
	"sync"

	"github.com/v2ray/v2ray-core/common/log"
	v2net "github.com/v2ray/v2ray-core/common/net"
	"github.com/v2ray/v2ray-core/transport/ray"
)

type FreedomConnection struct {
}

func NewFreedomConnection() *FreedomConnection {
	return &FreedomConnection{}
}

func (this *FreedomConnection) Dispatch(firstPacket v2net.Packet, ray ray.OutboundRay) error {
	conn, err := net.Dial(firstPacket.Destination().Network(), firstPacket.Destination().Address().String())
	log.Info("Freedom: Opening connection to %s", firstPacket.Destination().String())
	if err != nil {
		close(ray.OutboundOutput())
		log.Error("Freedom: Failed to open connection: %s : %v", firstPacket.Destination().String(), err)
		return err
	}

	input := ray.OutboundInput()
	output := ray.OutboundOutput()
	var readMutex, writeMutex sync.Mutex
	readMutex.Lock()
	writeMutex.Lock()

	if chunk := firstPacket.Chunk(); chunk != nil {
		conn.Write(chunk.Value)
		chunk.Release()
	}

	if !firstPacket.MoreChunks() {
		writeMutex.Unlock()
	} else {
		go func() {
			v2net.ChanToWriter(conn, input)
			writeMutex.Unlock()
		}()
	}

	go func() {
		defer readMutex.Unlock()
		defer close(output)

		response, err := v2net.ReadFrom(conn, nil)
		log.Info("Freedom receives %d bytes from %s", response.Len(), conn.RemoteAddr().String())
		if response.Len() > 0 {
			output <- response
		} else {
			response.Release()
		}
		if err != nil {
			return
		}
		if firstPacket.Destination().IsUDP() {
			return
		}

		v2net.ReaderToChan(output, conn)
	}()

	writeMutex.Lock()
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		tcpConn.CloseWrite()
	}
	readMutex.Lock()
	conn.Close()

	return nil
}
