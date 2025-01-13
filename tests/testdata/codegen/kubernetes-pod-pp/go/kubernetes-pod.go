package main

import (
	corev1 "github.com/khulnasoft/codeinfra-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/khulnasoft/codeinfra-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		bar, err := corev1.NewPod(ctx, "bar", &corev1.PodArgs{
			ApiVersion: codeinfra.String("v1"),
			Metadata: &metav1.ObjectMetaArgs{
				Namespace: codeinfra.String("foo"),
				Name:      codeinfra.String("bar"),
				Labels: codeinfra.StringMap{
					"app.kubernetes.io/name":    codeinfra.String("cilium-agent"),
					"app.kubernetes.io/part-of": codeinfra.String("cilium"),
					"k8s-app":                   codeinfra.String("cilium"),
				},
			},
			Spec: &corev1.PodSpecArgs{
				Containers: corev1.ContainerArray{
					&corev1.ContainerArgs{
						Name:  codeinfra.String("nginx"),
						Image: codeinfra.String("nginx:1.14-alpine"),
						Ports: corev1.ContainerPortArray{
							&corev1.ContainerPortArgs{
								ContainerPort: codeinfra.Int(80),
							},
						},
						Resources: &corev1.ResourceRequirementsArgs{
							Limits: codeinfra.StringMap{
								"memory": codeinfra.String("20Mi"),
								"cpu":    codeinfra.String("0.2"),
							},
						},
					},
					&corev1.ContainerArgs{
						Name:  codeinfra.String("nginx2"),
						Image: codeinfra.String("nginx:1.14-alpine"),
						Ports: corev1.ContainerPortArray{
							&corev1.ContainerPortArgs{
								ContainerPort: codeinfra.Int(80),
							},
						},
						Resources: &corev1.ResourceRequirementsArgs{
							Limits: codeinfra.StringMap{
								"memory": codeinfra.String("20Mi"),
								"cpu":    codeinfra.String("0.2"),
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		// Test that we can assign from a constant without type errors
		_ := bar.Kind
		return nil
	})
}
