package main

import (
	appsv1 "github.com/khulnasoft/codeinfra-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/khulnasoft/codeinfra-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/khulnasoft/codeinfra-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := appsv1.NewDeployment(ctx, "argocd_serverDeployment", &appsv1.DeploymentArgs{
			ApiVersion: codeinfra.String("apps/v1"),
			Kind:       codeinfra.String("Deployment"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: codeinfra.String("argocd-server"),
			},
			Spec: &appsv1.DeploymentSpecArgs{
				Selector: &metav1.LabelSelectorArgs{
					MatchLabels: codeinfra.StringMap{
						"app": codeinfra.String("server"),
					},
				},
				Replicas: codeinfra.Int(1),
				Template: &corev1.PodTemplateSpecArgs{
					Metadata: &metav1.ObjectMetaArgs{
						Labels: codeinfra.StringMap{
							"app": codeinfra.String("server"),
						},
					},
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							&corev1.ContainerArgs{
								Name:  codeinfra.String("nginx"),
								Image: codeinfra.String("nginx"),
								ReadinessProbe: &corev1.ProbeArgs{
									HttpGet: &corev1.HTTPGetActionArgs{
										Port: codeinfra.Any(8080),
									},
								},
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
