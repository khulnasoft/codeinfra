import codeinfra
import codeinfra_kubernetes as kubernetes

codeinfra_kubernetes_operator_deployment = kubernetes.apps.v1.Deployment("codeinfra_kubernetes_operatorDeployment",
    api_version="apps/v1",
    kind="Deployment",
    metadata={
        "name": "codeinfra-kubernetes-operator",
    },
    spec={
        "replicas": 1,
        "selector": {
            "match_labels": {
                "name": "codeinfra-kubernetes-operator",
            },
        },
        "template": {
            "metadata": {
                "labels": {
                    "name": "codeinfra-kubernetes-operator",
                },
            },
            "spec": {
                "service_account_name": "codeinfra-kubernetes-operator",
                "image_pull_secrets": [{
                    "name": "codeinfra-kubernetes-operator",
                }],
                "containers": [{
                    "name": "codeinfra-kubernetes-operator",
                    "image": "codeinfra/codeinfra-kubernetes-operator:v0.0.2",
                    "command": ["codeinfra-kubernetes-operator"],
                    "args": ["--zap-level=debug"],
                    "image_pull_policy": "Always",
                    "env": [
                        {
                            "name": "WATCH_NAMESPACE",
                            "value_from": {
                                "field_ref": {
                                    "field_path": "metadata.namespace",
                                },
                            },
                        },
                        {
                            "name": "POD_NAME",
                            "value_from": {
                                "field_ref": {
                                    "field_path": "metadata.name",
                                },
                            },
                        },
                        {
                            "name": "OPERATOR_NAME",
                            "value": "codeinfra-kubernetes-operator",
                        },
                    ],
                }],
            },
        },
    })
codeinfra_kubernetes_operator_role = kubernetes.rbac.v1.Role("codeinfra_kubernetes_operatorRole",
    api_version="rbac.authorization.k8s.io/v1",
    kind="Role",
    metadata={
        "creation_timestamp": None,
        "name": "codeinfra-kubernetes-operator",
    },
    rules=[
        {
            "api_groups": [""],
            "resources": [
                "pods",
                "services",
                "services/finalizers",
                "endpoints",
                "persistentvolumeclaims",
                "events",
                "configmaps",
                "secrets",
            ],
            "verbs": [
                "create",
                "delete",
                "get",
                "list",
                "patch",
                "update",
                "watch",
            ],
        },
        {
            "api_groups": ["apps"],
            "resources": [
                "deployments",
                "daemonsets",
                "replicasets",
                "statefulsets",
            ],
            "verbs": [
                "create",
                "delete",
                "get",
                "list",
                "patch",
                "update",
                "watch",
            ],
        },
        {
            "api_groups": ["monitoring.coreos.com"],
            "resources": ["servicemonitors"],
            "verbs": [
                "get",
                "create",
            ],
        },
        {
            "api_groups": ["apps"],
            "resource_names": ["codeinfra-kubernetes-operator"],
            "resources": ["deployments/finalizers"],
            "verbs": ["update"],
        },
        {
            "api_groups": [""],
            "resources": ["pods"],
            "verbs": ["get"],
        },
        {
            "api_groups": ["apps"],
            "resources": [
                "replicasets",
                "deployments",
            ],
            "verbs": ["get"],
        },
        {
            "api_groups": ["codeinfra.com"],
            "resources": ["*"],
            "verbs": [
                "create",
                "delete",
                "get",
                "list",
                "patch",
                "update",
                "watch",
            ],
        },
    ])
codeinfra_kubernetes_operator_role_binding = kubernetes.rbac.v1.RoleBinding("codeinfra_kubernetes_operatorRoleBinding",
    kind="RoleBinding",
    api_version="rbac.authorization.k8s.io/v1",
    metadata={
        "name": "codeinfra-kubernetes-operator",
    },
    subjects=[{
        "kind": "ServiceAccount",
        "name": "codeinfra-kubernetes-operator",
    }],
    role_ref={
        "kind": "Role",
        "name": "codeinfra-kubernetes-operator",
        "api_group": "rbac.authorization.k8s.io",
    })
codeinfra_kubernetes_operator_service_account = kubernetes.core.v1.ServiceAccount("codeinfra_kubernetes_operatorServiceAccount",
    api_version="v1",
    kind="ServiceAccount",
    metadata={
        "name": "codeinfra-kubernetes-operator",
    })
