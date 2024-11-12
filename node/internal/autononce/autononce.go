package autononce

import (
	"crypto/rand"

	"google.golang.org/protobuf/proto"
	"source.quilibrium.com/quilibrium/monorepo/node/protobufs"
)

func clone[T proto.Message](message T) T {
	return proto.Clone(message).(T)
}

func AddTokenRequest(request *protobufs.TokenRequest) *protobufs.TokenRequest {
	request.Nonce = make([]byte, 32)
	if _, err := rand.Read(request.Nonce); err != nil {
		panic(err)
	}
	return request
}

func Add(message proto.Message) proto.Message {
	if tokenRequest, ok := message.(*protobufs.TokenRequest); ok {
		return AddTokenRequest(clone(tokenRequest))
	}
	return message
}
