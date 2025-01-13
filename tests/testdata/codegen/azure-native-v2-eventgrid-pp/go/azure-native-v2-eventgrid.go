package main

import (
	eventgrid "github.com/khulnasoft/codeinfra-azure-native-sdk/eventgrid/v2"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := eventgrid.NewEventSubscription(ctx, "example", &eventgrid.EventSubscriptionArgs{
			Destination: &eventgrid.EventHubEventSubscriptionDestinationArgs{
				EndpointType: codeinfra.String("EventHub"),
				ResourceId:   codeinfra.String("example"),
			},
			ExpirationTimeUtc: codeinfra.String("example"),
			Scope:             codeinfra.String("example"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
