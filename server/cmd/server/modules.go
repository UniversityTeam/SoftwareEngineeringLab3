//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/roman-mazur/chat-channels-example/server/channels"
)

// ComposeApiServer will create an instance of CharApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*BalancerApiServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from channels package.
		balancers.Providers,
		// Provide BalancerApiServer instantiating the structure and injecting channels handler and port number.
		wire.Struct(new(BalancerApiServer), "Port", "ChannelsHandler"),
	)
	return nil, nil
}
