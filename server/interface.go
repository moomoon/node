package server

import "github.com/mysterium/node/server/dto"

type Client interface {
	NodeRegister(nodeKey, connectionConfig string) (err error)
	NodeSendStats(nodeKey string, sessionStats []dto.SessionStats) (err error)
	SessionCreate(nodeKey string) (session dto.Session, err error)
}
