package flask 

import (
	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


  func  Flask(ctx *pulumi.Context) error {

		appName := "flaskappservice"

		appLabels := pulumi.StringMap{
			"app": pulumi.String("flaskdep"),
		}


		deployment, err := appsv1.NewDeployment(ctx, "flaskdepdep", &appsv1.DeploymentArgs{
			Spec: appsv1.DeploymentSpecArgs{
				Selector: &metav1.LabelSelectorArgs{
					MatchLabels: appLabels,
				},
				Replicas: pulumi.Int(1),
				Template: &corev1.PodTemplateSpecArgs{
					Metadata: &metav1.ObjectMetaArgs{
						Labels: appLabels,
					},
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							corev1.ContainerArgs{
								Name:  pulumi.String("flaskdep"),
								Image: pulumi.String("pandas"),
								ImagePullPolicy: pulumi.String("Never"),
								Command: pulumi.StringArray{
									pulumi.String("/bin/bash"),


								},
								Args: pulumi.StringArray{
									pulumi.String("-c"),
									pulumi.String("while true; do echo Done ; sleep 3600;done "),
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

		var lbservice string
	        lbservice = "NodePort"

		servexpo, err := corev1.NewService(ctx, appName, &corev1.ServiceArgs {
			Metadata: &metav1.ObjectMetaArgs {
				Labels: appLabels,
			},

			Spec: &corev1.ServiceSpecArgs {
				Type: pulumi.String(lbservice),
				Ports: &corev1.ServicePortArray{
					corev1.ServicePortArgs {
						Port:  pulumi.Int(5000),
						TargetPort: pulumi.Int(5000),
						Protocol: pulumi.String("TCP"),
					},



				},


				Selector: appLabels,

			},
		})

		if err != nil {
			return err
		}


		ctx.Export("name flask pod", deployment.Metadata.Elem().Name())
		ctx.Export("Node Port", servexpo.Metadata.Elem().Name())


		return nil
}
