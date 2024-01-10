package response

import (
	"context"
	"distributed-kv-db/common/rslt"
)

func channelWithContextAndNodes[Response, Node any](ctx context.Context, nodes []Node) <-chan Response {
	return Channel(ctx, requestDummy[Response, Node], nodes)
}

func channelWithNodes[Response, Node any](nodes []Node) <-chan Response {
	return channelWithContextAndNodes[Response, Node](context.Background(), nodes)
}

func channelWithRequestAndNodes[Response, Node any](request func(Node) rslt.Of[Response], nodes []Node) <-chan Response {
	return Channel[Response, Node](context.Background(), request, nodes)
}
