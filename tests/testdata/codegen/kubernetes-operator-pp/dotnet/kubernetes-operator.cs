using System.Collections.Generic;
using System.Linq;
using Codeinfra;
using Kubernetes = Codeinfra.Kubernetes;

return await Deployment.RunAsync(() => 
{
    var codeinfra_kubernetes_operatorDeployment = new Kubernetes.Apps.V1.Deployment("codeinfra_kubernetes_operatorDeployment", new()
    {
        ApiVersion = "apps/v1",
        Kind = "Deployment",
        Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
        {
            Name = "codeinfra-kubernetes-operator",
        },
        Spec = new Kubernetes.Types.Inputs.Apps.V1.DeploymentSpecArgs
        {
            Replicas = 1,
            Selector = new Kubernetes.Types.Inputs.Meta.V1.LabelSelectorArgs
            {
                MatchLabels = 
                {
                    { "name", "codeinfra-kubernetes-operator" },
                },
            },
            Template = new Kubernetes.Types.Inputs.Core.V1.PodTemplateSpecArgs
            {
                Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
                {
                    Labels = 
                    {
                        { "name", "codeinfra-kubernetes-operator" },
                    },
                },
                Spec = new Kubernetes.Types.Inputs.Core.V1.PodSpecArgs
                {
                    ServiceAccountName = "codeinfra-kubernetes-operator",
                    ImagePullSecrets = new[]
                    {
                        new Kubernetes.Types.Inputs.Core.V1.LocalObjectReferenceArgs
                        {
                            Name = "codeinfra-kubernetes-operator",
                        },
                    },
                    Containers = new[]
                    {
                        new Kubernetes.Types.Inputs.Core.V1.ContainerArgs
                        {
                            Name = "codeinfra-kubernetes-operator",
                            Image = "codeinfra/codeinfra-kubernetes-operator:v0.0.2",
                            Command = new[]
                            {
                                "codeinfra-kubernetes-operator",
                            },
                            Args = new[]
                            {
                                "--zap-level=debug",
                            },
                            ImagePullPolicy = "Always",
                            Env = new[]
                            {
                                new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                {
                                    Name = "WATCH_NAMESPACE",
                                    ValueFrom = new Kubernetes.Types.Inputs.Core.V1.EnvVarSourceArgs
                                    {
                                        FieldRef = new Kubernetes.Types.Inputs.Core.V1.ObjectFieldSelectorArgs
                                        {
                                            FieldPath = "metadata.namespace",
                                        },
                                    },
                                },
                                new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                {
                                    Name = "POD_NAME",
                                    ValueFrom = new Kubernetes.Types.Inputs.Core.V1.EnvVarSourceArgs
                                    {
                                        FieldRef = new Kubernetes.Types.Inputs.Core.V1.ObjectFieldSelectorArgs
                                        {
                                            FieldPath = "metadata.name",
                                        },
                                    },
                                },
                                new Kubernetes.Types.Inputs.Core.V1.EnvVarArgs
                                {
                                    Name = "OPERATOR_NAME",
                                    Value = "codeinfra-kubernetes-operator",
                                },
                            },
                        },
                    },
                },
            },
        },
    });

    var codeinfra_kubernetes_operatorRole = new Kubernetes.Rbac.V1.Role("codeinfra_kubernetes_operatorRole", new()
    {
        ApiVersion = "rbac.authorization.k8s.io/v1",
        Kind = "Role",
        Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
        {
            CreationTimestamp = null,
            Name = "codeinfra-kubernetes-operator",
        },
        Rules = new[]
        {
            new Kubernetes.Types.Inputs.Rbac.V1.PolicyRuleArgs
            {
                ApiGroups = new[]
                {
                    "",
                },
                Resources = new[]
                {
                    "pods",
                    "services",
                    "services/finalizers",
                    "endpoints",
                    "persistentvolumeclaims",
                    "events",
                    "configmaps",
                    "secrets",
                },
                Verbs = new[]
                {
                    "create",
                    "delete",
                    "get",
                    "list",
                    "patch",
                    "update",
                    "watch",
                },
            },
            new Kubernetes.Types.Inputs.Rbac.V1.PolicyRuleArgs
            {
                ApiGroups = new[]
                {
                    "apps",
                },
                Resources = new[]
                {
                    "deployments",
                    "daemonsets",
                    "replicasets",
                    "statefulsets",
                },
                Verbs = new[]
                {
                    "create",
                    "delete",
                    "get",
                    "list",
                    "patch",
                    "update",
                    "watch",
                },
            },
            new Kubernetes.Types.Inputs.Rbac.V1.PolicyRuleArgs
            {
                ApiGroups = new[]
                {
                    "monitoring.coreos.com",
                },
                Resources = new[]
                {
                    "servicemonitors",
                },
                Verbs = new[]
                {
                    "get",
                    "create",
                },
            },
            new Kubernetes.Types.Inputs.Rbac.V1.PolicyRuleArgs
            {
                ApiGroups = new[]
                {
                    "apps",
                },
                ResourceNames = new[]
                {
                    "codeinfra-kubernetes-operator",
                },
                Resources = new[]
                {
                    "deployments/finalizers",
                },
                Verbs = new[]
                {
                    "update",
                },
            },
            new Kubernetes.Types.Inputs.Rbac.V1.PolicyRuleArgs
            {
                ApiGroups = new[]
                {
                    "",
                },
                Resources = new[]
                {
                    "pods",
                },
                Verbs = new[]
                {
                    "get",
                },
            },
            new Kubernetes.Types.Inputs.Rbac.V1.PolicyRuleArgs
            {
                ApiGroups = new[]
                {
                    "apps",
                },
                Resources = new[]
                {
                    "replicasets",
                    "deployments",
                },
                Verbs = new[]
                {
                    "get",
                },
            },
            new Kubernetes.Types.Inputs.Rbac.V1.PolicyRuleArgs
            {
                ApiGroups = new[]
                {
                    "codeinfra.com",
                },
                Resources = new[]
                {
                    "*",
                },
                Verbs = new[]
                {
                    "create",
                    "delete",
                    "get",
                    "list",
                    "patch",
                    "update",
                    "watch",
                },
            },
        },
    });

    var codeinfra_kubernetes_operatorRoleBinding = new Kubernetes.Rbac.V1.RoleBinding("codeinfra_kubernetes_operatorRoleBinding", new()
    {
        Kind = "RoleBinding",
        ApiVersion = "rbac.authorization.k8s.io/v1",
        Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
        {
            Name = "codeinfra-kubernetes-operator",
        },
        Subjects = new[]
        {
            new Kubernetes.Types.Inputs.Rbac.V1.SubjectArgs
            {
                Kind = "ServiceAccount",
                Name = "codeinfra-kubernetes-operator",
            },
        },
        RoleRef = new Kubernetes.Types.Inputs.Rbac.V1.RoleRefArgs
        {
            Kind = "Role",
            Name = "codeinfra-kubernetes-operator",
            ApiGroup = "rbac.authorization.k8s.io",
        },
    });

    var codeinfra_kubernetes_operatorServiceAccount = new Kubernetes.Core.V1.ServiceAccount("codeinfra_kubernetes_operatorServiceAccount", new()
    {
        ApiVersion = "v1",
        Kind = "ServiceAccount",
        Metadata = new Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs
        {
            Name = "codeinfra-kubernetes-operator",
        },
    });

});

