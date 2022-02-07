package kafka

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReplConKafka struct {
	pulumi.ResourceState

	Service *corev1.Service
	ReplicationController *corev1.ReplicationController
}

type ReplConKafkaArgs struct {
	Image           pulumi.StringInput
	Replicas        pulumi.IntPtrInput
}


func NewReplConKafkaBroker(ctx *pulumi.Context, name string, args *ReplConKafkaArgs, opts ...pulumi.ResourceOption,)(*ReplConKafka, error){

	replConKafka := &ReplConKafka{}

	err := ctx.RegisterComponentResource("kafka-app", name, replConKafka, opts...)

		if err != nil {
			return nil, err
		}


	var last = name[len(name)-1:]


	replConKafka.ReplicationController, err = corev1.NewReplicationController(ctx , name, &corev1.ReplicationControllerArgs {

		Metadata: &metav1.ObjectMetaArgs {
			Name: pulumi.String(name),
			Labels: pulumi.StringMap {
				"name" : pulumi.String(name),
				"broker_id" : pulumi.String(last),
			},
		},

		Spec: corev1.ReplicationControllerSpecArgs{
			Replicas: args.Replicas,
			Selector: pulumi.StringMap{
				"name": pulumi.String(name),
				"broker_id": pulumi.String(last),
			},
			Template: corev1.PodTemplateSpecArgs{
				 Metadata: &metav1.ObjectMetaArgs{
					 Labels: pulumi.StringMap{
						 "name": pulumi.String(name),
						 "broker_id": pulumi.String(last),
						 "service": pulumi.String("kafka"),
					 },
				 },

				 Spec: corev1.PodSpecArgs{
					 Containers: corev1.ContainerArray{
							 corev1.ContainerArgs{
							  Name: pulumi.String("kafka"),
							  ImagePullPolicy: pulumi.String("Always"),
							  Image: args.Image.ToStringOutput(),
							  Resources: &corev1.ResourceRequirementsArgs{
								  Limits: pulumi.StringMap{
									  "memory": pulumi.String("750Mi"),
										  },
									   },
							   Env: corev1.EnvVarArray{
								   corev1.EnvVarArgs{
									   Name: pulumi.String("KAFKA_HEAP_OPTS"),
									   Value: pulumi.String("-Xmx512M -Xms512M"),
									   },
								   corev1.EnvVarArgs{
									   Name: pulumi.String("BROKER_ID"),
									   Value: pulumi.String(last),
									   },
								   corev1.EnvVarArgs{
									   Name: pulumi.String("NUM_PARTITIONS"),
									    Value: pulumi.String("3"),
									   },
								   corev1.EnvVarArgs{
									   Name: pulumi.String("ADVERTISED_HOSTNAME"),
									   Value: pulumi.String(name),
									   },
								   corev1.EnvVarArgs{
									   Name: pulumi.String("ZOOKEEPER_CONNECT"),
									   Value: pulumi.String("zoo1:2181,zoo2:2181,zoo3:2181"),
									   },
								   corev1.EnvVarArgs{
									   Name: pulumi.String("RETENTION_HOURS"),
									   Value: pulumi.String("72"),
									   },
								   },
							   Ports: &corev1.ContainerPortArray{
								   corev1.ContainerPortArgs{
									   ContainerPort: pulumi.Int(9092),
									   Name: pulumi.String("broker"),
									   Protocol: pulumi.String("TCP"),
									   },
								   },
							   VolumeMounts: &corev1.VolumeMountArray{
								   &corev1.VolumeMountArgs{
									   MountPath: pulumi.String("/data"),
									   Name: pulumi.String("data"),
									   },
								   },

							 },
						 },
						 Volumes: &corev1.VolumeArray{
							 &corev1.VolumeArgs{
								 Name: pulumi.String("data"),
							 },
						 },
					 },

				 },
			 },

		 }, pulumi.Parent(replConKafka))

		if err != nil {
			return nil, err
		}

		replConKafka.Service, err = corev1.NewService(ctx, name, &corev1.ServiceArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String(name),
				Labels: pulumi.StringMap{
					"service": pulumi.String("kafka"),
					"name": pulumi.String(name),
					"broker_id": pulumi.String(last),
				},
			},
			Spec: &corev1.ServiceSpecArgs{
				Ports: corev1.ServicePortArray{
					&corev1.ServicePortArgs{
						Name: pulumi.String("broker"),
						Port: pulumi.Int(9092),
					},
				},
				Selector: pulumi.StringMap{
					"service": pulumi.String("kafka"),
					"broker_id": pulumi.String(last),
				},
			},

		}, pulumi.Parent(replConKafka))

		if err != nil {
			return nil, err
		}



		return replConKafka, nil
}

