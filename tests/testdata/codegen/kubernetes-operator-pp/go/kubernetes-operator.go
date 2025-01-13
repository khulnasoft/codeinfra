package main

import (
	appsv1 "github.com/khulnasoft/codeinfra-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/khulnasoft/codeinfra-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/khulnasoft/codeinfra-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	rbacv1 "github.com/khulnasoft/codeinfra-kubernetes/sdk/v3/go/kubernetes/rbac/v1"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		_, err := appsv1.NewDeployment(ctx, "codeinfra_kubernetes_operatorDeployment", &appsv1.DeploymentArgs{
			ApiVersion: codeinfra.String("apps/v1"),
			Kind:       codeinfra.String("Deployment"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: codeinfra.String("codeinfra-kubernetes-operator"),
			},
			Spec: &appsv1.DeploymentSpecArgs{
				Replicas: codeinfra.Int(1),
				Selector: &metav1.LabelSelectorArgs{
					MatchLabels: codeinfra.StringMap{
						"name": codeinfra.String("codeinfra-kubernetes-operator"),
					},
				},
				Template: &corev1.PodTemplateSpecArgs{
					Metadata: &metav1.ObjectMetaArgs{
						Labels: codeinfra.StringMap{
							"name": codeinfra.String("codeinfra-kubernetes-operator"),
						},
					},
					Spec: &corev1.PodSpecArgs{
						ServiceAccountName: codeinfra.String("codeinfra-kubernetes-operator"),
						ImagePullSecrets: corev1.LocalObjectReferenceArray{
							&corev1.LocalObjectReferenceArgs{
								Name: codeinfra.String("codeinfra-kubernetes-operator"),
							},
						},
						Containers: corev1.ContainerArray{
							&corev1.ContainerArgs{
								Name:  codeinfra.String("codeinfra-kubernetes-operator"),
								Image: codeinfra.String("codeinfra/codeinfra-kubernetes-operator:v0.0.2"),
								Command: codeinfra.StringArray{
									codeinfra.String("codeinfra-kubernetes-operator"),
								},
								Args: codeinfra.StringArray{
									codeinfra.String("--zap-level=debug"),
								},
								ImagePullPolicy: codeinfra.String("Always"),
								Env: corev1.EnvVarArray{
									&corev1.EnvVarArgs{
										Name: codeinfra.String("WATCH_NAMESPACE"),
										ValueFrom: &corev1.EnvVarSourceArgs{
											FieldRef: &corev1.ObjectFieldSelectorArgs{
												FieldPath: codeinfra.String("metadata.namespace"),
											},
										},
									},
									&corev1.EnvVarArgs{
										Name: codeinfra.String("POD_NAME"),
										ValueFrom: &corev1.EnvVarSourceArgs{
											FieldRef: &corev1.ObjectFieldSelectorArgs{
												FieldPath: codeinfra.String("metadata.name"),
											},
										},
									},
									&corev1.EnvVarArgs{
										Name:  codeinfra.String("OPERATOR_NAME"),
										Value: codeinfra.String("codeinfra-kubernetes-operator"),
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
		_, err = rbacv1.NewRole(ctx, "codeinfra_kubernetes_operatorRole", &rbacv1.RoleArgs{
			ApiVersion: codeinfra.String("rbac.authorization.k8s.io/v1"),
			Kind:       codeinfra.String("Role"),
			Metadata: &metav1.ObjectMetaArgs{
				CreationTimestamp: nil,
				Name:              codeinfra.String("codeinfra-kubernetes-operator"),
			},
			Rules: rbacv1.PolicyRuleArray{
				&rbacv1.PolicyRuleArgs{
					ApiGroups: codeinfra.StringArray{
						codeinfra.String(""),
					},
					Resources: codeinfra.StringArray{
						codeinfra.String("pods"),
						codeinfra.String("services"),
						codeinfra.String("services/finalizers"),
						codeinfra.String("endpoints"),
						codeinfra.String("persistentvolumeclaims"),
						codeinfra.String("events"),
						codeinfra.String("configmaps"),
						codeinfra.String("secrets"),
					},
					Verbs: codeinfra.StringArray{
						codeinfra.String("create"),
						codeinfra.String("delete"),
						codeinfra.String("get"),
						codeinfra.String("list"),
						codeinfra.String("patch"),
						codeinfra.String("update"),
						codeinfra.String("watch"),
					},
				},
				&rbacv1.PolicyRuleArgs{
					ApiGroups: codeinfra.StringArray{
						codeinfra.String("apps"),
					},
					Resources: codeinfra.StringArray{
						codeinfra.String("deployments"),
						codeinfra.String("daemonsets"),
						codeinfra.String("replicasets"),
						codeinfra.String("statefulsets"),
					},
					Verbs: codeinfra.StringArray{
						codeinfra.String("create"),
						codeinfra.String("delete"),
						codeinfra.String("get"),
						codeinfra.String("list"),
						codeinfra.String("patch"),
						codeinfra.String("update"),
						codeinfra.String("watch"),
					},
				},
				&rbacv1.PolicyRuleArgs{
					ApiGroups: codeinfra.StringArray{
						codeinfra.String("monitoring.coreos.com"),
					},
					Resources: codeinfra.StringArray{
						codeinfra.String("servicemonitors"),
					},
					Verbs: codeinfra.StringArray{
						codeinfra.String("get"),
						codeinfra.String("create"),
					},
				},
				&rbacv1.PolicyRuleArgs{
					ApiGroups: codeinfra.StringArray{
						codeinfra.String("apps"),
					},
					ResourceNames: codeinfra.StringArray{
						codeinfra.String("codeinfra-kubernetes-operator"),
					},
					Resources: codeinfra.StringArray{
						codeinfra.String("deployments/finalizers"),
					},
					Verbs: codeinfra.StringArray{
						codeinfra.String("update"),
					},
				},
				&rbacv1.PolicyRuleArgs{
					ApiGroups: codeinfra.StringArray{
						codeinfra.String(""),
					},
					Resources: codeinfra.StringArray{
						codeinfra.String("pods"),
					},
					Verbs: codeinfra.StringArray{
						codeinfra.String("get"),
					},
				},
				&rbacv1.PolicyRuleArgs{
					ApiGroups: codeinfra.StringArray{
						codeinfra.String("apps"),
					},
					Resources: codeinfra.StringArray{
						codeinfra.String("replicasets"),
						codeinfra.String("deployments"),
					},
					Verbs: codeinfra.StringArray{
						codeinfra.String("get"),
					},
				},
				&rbacv1.PolicyRuleArgs{
					ApiGroups: codeinfra.StringArray{
						codeinfra.String("codeinfra.com"),
					},
					Resources: codeinfra.StringArray{
						codeinfra.String("*"),
					},
					Verbs: codeinfra.StringArray{
						codeinfra.String("create"),
						codeinfra.String("delete"),
						codeinfra.String("get"),
						codeinfra.String("list"),
						codeinfra.String("patch"),
						codeinfra.String("update"),
						codeinfra.String("watch"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = rbacv1.NewRoleBinding(ctx, "codeinfra_kubernetes_operatorRoleBinding", &rbacv1.RoleBindingArgs{
			Kind:       codeinfra.String("RoleBinding"),
			ApiVersion: codeinfra.String("rbac.authorization.k8s.io/v1"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: codeinfra.String("codeinfra-kubernetes-operator"),
			},
			Subjects: rbacv1.SubjectArray{
				&rbacv1.SubjectArgs{
					Kind: codeinfra.String("ServiceAccount"),
					Name: codeinfra.String("codeinfra-kubernetes-operator"),
				},
			},
			RoleRef: &rbacv1.RoleRefArgs{
				Kind:     codeinfra.String("Role"),
				Name:     codeinfra.String("codeinfra-kubernetes-operator"),
				ApiGroup: codeinfra.String("rbac.authorization.k8s.io"),
			},
		})
		if err != nil {
			return err
		}
		_, err = corev1.NewServiceAccount(ctx, "codeinfra_kubernetes_operatorServiceAccount", &corev1.ServiceAccountArgs{
			ApiVersion: codeinfra.String("v1"),
			Kind:       codeinfra.String("ServiceAccount"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: codeinfra.String("codeinfra-kubernetes-operator"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
