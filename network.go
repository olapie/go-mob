package mobile

import "google.golang.org/grpc"

const (
	NoNetwork = 0
	Cellular  = 1
	WIFI      = 2
)

const (
	Idle       = 0
	Connecting = 1
	Connected  = 2
)

type GRPCClientConn struct {
	Conn grpc.ClientConnInterface
}
