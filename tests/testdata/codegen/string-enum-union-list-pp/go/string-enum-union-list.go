package main

import (
	"github.com/khulnasoft/codeinfra-azure-native/sdk/go/azure/servicebus"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := servicebus.NewNamespaceAuthorizationRule(ctx, "namespaceAuthorizationRule", &servicebus.NamespaceAuthorizationRuleArgs{
			AuthorizationRuleName: codeinfra.String("sdk-AuthRules-1788"),
			NamespaceName:         codeinfra.String("sdk-Namespace-6914"),
			ResourceGroupName:     codeinfra.String("ArunMonocle"),
			Rights: codeinfra.StringArray{
				codeinfra.String(servicebus.AccessRightsListen),
				codeinfra.String(servicebus.AccessRightsSend),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
