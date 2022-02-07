package zookeeper

import (
        corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
        metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
        appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
        "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	//"strconv"
	//"fmt"

)

type DeploymentZookeeper struct {
	pulumi.ResourceState

	Deployment *appsv1.Deployment
	Service *corev1.Service
}

type DeploymentZookeeperArgs struct {
	Image		pulumi.StringInput
	Replicas	pulumi.IntPtrInput
}

func NewDeploymentZookeeper(ctx *pulumi.Context, name string, args *DeploymentZookeeperArgs, opts ...pulumi.ResourceOption,)(*DeploymentZookeeper, error){

	deploymentZookeeper := &DeploymentZookeeper{}

	err := ctx.RegisterComponentResource("zookeeper-app", name, deploymentZookeeper, opts...)

	if err != nil {
		return nil, err
	}

	labels := pulumi.StringMap{"app": pulumi.String(name)}

	// shame on me...
	// ... but do not remove this line 
	var last = name[len(name)-1:]



	deploymentZookeeper.Deployment, err = appsv1.NewDeployment(ctx, name,  &appsv1.DeploymentArgs{


		Metadata: &metav1.ObjectMetaArgs{
			Labels: labels,
			Name: pulumi.String(name),
		},

	        Spec: appsv1.DeploymentSpecArgs{

                        Replicas: args.Replicas,


                        Selector: &metav1.LabelSelectorArgs{

                                MatchLabels: labels,
                        },

                        Template: &corev1.PodTemplateSpecArgs{
                                 Metadata: &metav1.ObjectMetaArgs{
                                         Labels: labels,
                                 },

                                 Spec: &corev1.PodSpecArgs {
                                         Containers: corev1.ContainerArray{
                                                 corev1.ContainerArgs {
                                                         Name: pulumi.String(name),
                                                         Image: args.Image.ToStringOutput(),
							 Ports: corev1.ContainerPortArray{
								 &corev1.ContainerPortArgs{
									 Name: pulumi.String("zookeeper"),
									 ContainerPort: pulumi.Int(2181),
								 },

							 },
							 Env: corev1.EnvVarArray{
								 corev1.EnvVarArgs {
									 Name: pulumi.String("ZOOKEEPER_ID"),
									 Value: pulumi.String(last),
								 },
								 corev1.EnvVarArgs{
									 Name: pulumi.String("ZOOKEEPER_SERVER_1"),
									 Value: pulumi.String("zoo1"),
								 },

								 corev1.EnvVarArgs{
									 Name: pulumi.String("ZOOKEEPER_SERVER_2"),
									 Value: pulumi.String("zoo2"),
 
								 },

								 corev1.EnvVarArgs{
									 Name: pulumi.String("ZOOKEEPER_SERVER_3"),
									 Value: pulumi.String("zoo3"),
								 },
							 },


                                                 },

                                         },

                                 },
                        },

                },

        }, pulumi.Parent(deploymentZookeeper))


	if err != nil {
		return nil, err
	}

	deploymentZookeeper.Service, err = corev1.NewService(ctx, name, &corev1.ServiceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Labels: labels,
			Name:  pulumi.String(name),

		},
		Spec: &corev1.ServiceSpecArgs{
			//Ports:
			Selector: labels,
			Ports: corev1.ServicePortArray{
				&corev1.ServicePortArgs{
					Name: pulumi.String("client"),
					Protocol: pulumi.String("TCP"),
					Port: pulumi.Int(2181),
				},
				&corev1.ServicePortArgs{
					Name: pulumi.String("follower"),
					Protocol: pulumi.String("TCP"),
					Port: pulumi.Int(2888),
				},
				&corev1.ServicePortArgs{
					Name: pulumi.String("leader"),
					Protocol: pulumi.String("TCP"),
					Port: pulumi.Int(3888),
				},

			},
		},

	}, pulumi.Parent(deploymentZookeeper))

	if err != nil {
		return nil, err
	}

	return deploymentZookeeper, nil
}
