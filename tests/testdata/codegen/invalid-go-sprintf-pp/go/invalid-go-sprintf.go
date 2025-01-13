package main

import (
	appsv1 "github.com/khulnasoft/codeinfra-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	metav1 "github.com/khulnasoft/codeinfra-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		// example
		_, err := appsv1.NewDeployment(ctx, "argocd_serverDeployment", &appsv1.DeploymentArgs{
			ApiVersion: codeinfra.String("apps/v1"),
			Kind:       codeinfra.String("Deployment"),
			Metadata: &metav1.ObjectMetaArgs{
				Labels: codeinfra.StringMap{
					"app.kubernetes.io/component": codeinfra.String("server"),
					"aws:region":                  codeinfra.String("us-west-2"),
					"key%percent":                 codeinfra.String("percent"),
					"key...ellipse":               codeinfra.String("ellipse"),
					"key{bracket":                 codeinfra.String("bracket"),
					"key}bracket":                 codeinfra.String("bracket"),
					"key*asterix":                 codeinfra.String("asterix"),
					"key?question":                codeinfra.String("question"),
					"key,comma":                   codeinfra.String("comma"),
					"key&&and":                    codeinfra.String("and"),
					"key||or":                     codeinfra.String("or"),
					"key!not":                     codeinfra.String("not"),
					"key=>geq":                    codeinfra.String("geq"),
					"key==eq":                     codeinfra.String("equal"),
				},
				Name: codeinfra.String("argocd-server"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
