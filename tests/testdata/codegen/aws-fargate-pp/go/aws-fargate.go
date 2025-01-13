package main

import (
	"encoding/json"

	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/ec2"
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/ecs"
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/elasticloadbalancingv2"
	"github.com/khulnasoft/codeinfra-aws/sdk/v5/go/aws/iam"
	"github.com/khulnasoft/codeinfra/sdk/v3/go/codeinfra"
)

func main() {
	codeinfra.Run(func(ctx *codeinfra.Context) error {
		// Read the default VPC and public subnets, which we will use.
		vpc, err := ec2.LookupVpc(ctx, &ec2.LookupVpcArgs{
			Default: codeinfra.BoolRef(true),
		}, nil)
		if err != nil {
			return err
		}
		subnets, err := ec2.GetSubnetIds(ctx, &ec2.GetSubnetIdsArgs{
			VpcId: vpc.Id,
		}, nil)
		if err != nil {
			return err
		}
		// Create a security group that permits HTTP ingress and unrestricted egress.
		webSecurityGroup, err := ec2.NewSecurityGroup(ctx, "webSecurityGroup", &ec2.SecurityGroupArgs{
			VpcId: codeinfra.String(vpc.Id),
			Egress: ec2.SecurityGroupEgressArray{
				&ec2.SecurityGroupEgressArgs{
					Protocol: codeinfra.String("-1"),
					FromPort: codeinfra.Int(0),
					ToPort:   codeinfra.Int(0),
					CidrBlocks: codeinfra.StringArray{
						codeinfra.String("0.0.0.0/0"),
					},
				},
			},
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					Protocol: codeinfra.String("tcp"),
					FromPort: codeinfra.Int(80),
					ToPort:   codeinfra.Int(80),
					CidrBlocks: codeinfra.StringArray{
						codeinfra.String("0.0.0.0/0"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		// Create an ECS cluster to run a container-based service.
		cluster, err := ecs.NewCluster(ctx, "cluster", nil)
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2008-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Sid":    "",
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"Service": "ecs-tasks.amazonaws.com",
					},
					"Action": "sts:AssumeRole",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		// Create an IAM role that can be used by our service's task.
		taskExecRole, err := iam.NewRole(ctx, "taskExecRole", &iam.RoleArgs{
			AssumeRolePolicy: codeinfra.String(json0),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "taskExecRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      taskExecRole.Name,
			PolicyArn: codeinfra.String("arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"),
		})
		if err != nil {
			return err
		}
		// Create a load balancer to listen for HTTP traffic on port 80.
		webLoadBalancer, err := elasticloadbalancingv2.NewLoadBalancer(ctx, "webLoadBalancer", &elasticloadbalancingv2.LoadBalancerArgs{
			Subnets: toCodeinfraStringArray(subnets.Ids),
			SecurityGroups: codeinfra.StringArray{
				webSecurityGroup.ID(),
			},
		})
		if err != nil {
			return err
		}
		webTargetGroup, err := elasticloadbalancingv2.NewTargetGroup(ctx, "webTargetGroup", &elasticloadbalancingv2.TargetGroupArgs{
			Port:       codeinfra.Int(80),
			Protocol:   codeinfra.String("HTTP"),
			TargetType: codeinfra.String("ip"),
			VpcId:      codeinfra.String(vpc.Id),
		})
		if err != nil {
			return err
		}
		webListener, err := elasticloadbalancingv2.NewListener(ctx, "webListener", &elasticloadbalancingv2.ListenerArgs{
			LoadBalancerArn: webLoadBalancer.Arn,
			Port:            codeinfra.Int(80),
			DefaultActions: elasticloadbalancingv2.ListenerDefaultActionArray{
				&elasticloadbalancingv2.ListenerDefaultActionArgs{
					Type:           codeinfra.String("forward"),
					TargetGroupArn: webTargetGroup.Arn,
				},
			},
		})
		if err != nil {
			return err
		}
		tmpJSON1, err := json.Marshal([]map[string]interface{}{
			map[string]interface{}{
				"name":  "my-app",
				"image": "nginx",
				"portMappings": []map[string]interface{}{
					map[string]interface{}{
						"containerPort": 80,
						"hostPort":      80,
						"protocol":      "tcp",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json1 := string(tmpJSON1)
		// Spin up a load balanced service running NGINX
		appTask, err := ecs.NewTaskDefinition(ctx, "appTask", &ecs.TaskDefinitionArgs{
			Family:      codeinfra.String("fargate-task-definition"),
			Cpu:         codeinfra.String("256"),
			Memory:      codeinfra.String("512"),
			NetworkMode: codeinfra.String("awsvpc"),
			RequiresCompatibilities: codeinfra.StringArray{
				codeinfra.String("FARGATE"),
			},
			ExecutionRoleArn:     taskExecRole.Arn,
			ContainerDefinitions: codeinfra.String(json1),
		})
		if err != nil {
			return err
		}
		_, err = ecs.NewService(ctx, "appService", &ecs.ServiceArgs{
			Cluster:        cluster.Arn,
			DesiredCount:   codeinfra.Int(5),
			LaunchType:     codeinfra.String("FARGATE"),
			TaskDefinition: appTask.Arn,
			NetworkConfiguration: &ecs.ServiceNetworkConfigurationArgs{
				AssignPublicIp: codeinfra.Bool(true),
				Subnets:        toCodeinfraStringArray(subnets.Ids),
				SecurityGroups: codeinfra.StringArray{
					webSecurityGroup.ID(),
				},
			},
			LoadBalancers: ecs.ServiceLoadBalancerArray{
				&ecs.ServiceLoadBalancerArgs{
					TargetGroupArn: webTargetGroup.Arn,
					ContainerName:  codeinfra.String("my-app"),
					ContainerPort:  codeinfra.Int(80),
				},
			},
		}, codeinfra.DependsOn([]codeinfra.Resource{
			webListener,
		}))
		if err != nil {
			return err
		}
		ctx.Export("url", webLoadBalancer.DnsName)
		return nil
	})
}
func toCodeinfraStringArray(arr []string) codeinfra.StringArray {
	var codeinfraArr codeinfra.StringArray
	for _, v := range arr {
		codeinfraArr = append(codeinfraArr, codeinfra.String(v))
	}
	return codeinfraArr
}
