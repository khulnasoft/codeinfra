import * as codeinfra from "@codeinfra/codeinfra";
import * as kubernetes from "@codeinfra/kubernetes";

const codeinfra_kubernetes_operatorDeployment = new kubernetes.apps.v1.Deployment("codeinfra_kubernetes_operatorDeployment", {
    apiVersion: "apps/v1",
    kind: "Deployment",
    metadata: {
        name: "codeinfra-kubernetes-operator",
    },
    spec: {
        replicas: 1,
        selector: {
            matchLabels: {
                name: "codeinfra-kubernetes-operator",
            },
        },
        template: {
            metadata: {
                labels: {
                    name: "codeinfra-kubernetes-operator",
                },
            },
            spec: {
                serviceAccountName: "codeinfra-kubernetes-operator",
                imagePullSecrets: [{
                    name: "codeinfra-kubernetes-operator",
                }],
                containers: [{
                    name: "codeinfra-kubernetes-operator",
                    image: "codeinfra/codeinfra-kubernetes-operator:v0.0.2",
                    command: ["codeinfra-kubernetes-operator"],
                    args: ["--zap-level=debug"],
                    imagePullPolicy: "Always",
                    env: [
                        {
                            name: "WATCH_NAMESPACE",
                            valueFrom: {
                                fieldRef: {
                                    fieldPath: "metadata.namespace",
                                },
                            },
                        },
                        {
                            name: "POD_NAME",
                            valueFrom: {
                                fieldRef: {
                                    fieldPath: "metadata.name",
                                },
                            },
                        },
                        {
                            name: "OPERATOR_NAME",
                            value: "codeinfra-kubernetes-operator",
                        },
                    ],
                }],
            },
        },
    },
});
const codeinfra_kubernetes_operatorRole = new kubernetes.rbac.v1.Role("codeinfra_kubernetes_operatorRole", {
    apiVersion: "rbac.authorization.k8s.io/v1",
    kind: "Role",
    metadata: {
        creationTimestamp: undefined,
        name: "codeinfra-kubernetes-operator",
    },
    rules: [
        {
            apiGroups: [""],
            resources: [
                "pods",
                "services",
                "services/finalizers",
                "endpoints",
                "persistentvolumeclaims",
                "events",
                "configmaps",
                "secrets",
            ],
            verbs: [
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
            apiGroups: ["apps"],
            resources: [
                "deployments",
                "daemonsets",
                "replicasets",
                "statefulsets",
            ],
            verbs: [
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
            apiGroups: ["monitoring.coreos.com"],
            resources: ["servicemonitors"],
            verbs: [
                "get",
                "create",
            ],
        },
        {
            apiGroups: ["apps"],
            resourceNames: ["codeinfra-kubernetes-operator"],
            resources: ["deployments/finalizers"],
            verbs: ["update"],
        },
        {
            apiGroups: [""],
            resources: ["pods"],
            verbs: ["get"],
        },
        {
            apiGroups: ["apps"],
            resources: [
                "replicasets",
                "deployments",
            ],
            verbs: ["get"],
        },
        {
            apiGroups: ["codeinfra.com"],
            resources: ["*"],
            verbs: [
                "create",
                "delete",
                "get",
                "list",
                "patch",
                "update",
                "watch",
            ],
        },
    ],
});
const codeinfra_kubernetes_operatorRoleBinding = new kubernetes.rbac.v1.RoleBinding("codeinfra_kubernetes_operatorRoleBinding", {
    kind: "RoleBinding",
    apiVersion: "rbac.authorization.k8s.io/v1",
    metadata: {
        name: "codeinfra-kubernetes-operator",
    },
    subjects: [{
        kind: "ServiceAccount",
        name: "codeinfra-kubernetes-operator",
    }],
    roleRef: {
        kind: "Role",
        name: "codeinfra-kubernetes-operator",
        apiGroup: "rbac.authorization.k8s.io",
    },
});
const codeinfra_kubernetes_operatorServiceAccount = new kubernetes.core.v1.ServiceAccount("codeinfra_kubernetes_operatorServiceAccount", {
    apiVersion: "v1",
    kind: "ServiceAccount",
    metadata: {
        name: "codeinfra-kubernetes-operator",
    },
});
