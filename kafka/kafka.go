package kafka

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
        corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
        metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"

	"fmt"
)

func  Kafka(ctx *pulumi.Context) error {

	x, err := NewReplConKafkaBroker(ctx, "kafka-1", &ReplConKafkaArgs{
			Image: pulumi.String("navicore/kafka:0.10.1.1"),
			Replicas: pulumi.Int(1),

		})
	if err != nil {
		return err
	}

	y, err := NewReplConKafkaBroker(ctx, "kafka-2", &ReplConKafkaArgs{
			Image: pulumi.String("navicore/kafka:0.10.1.1"),
			Replicas: pulumi.Int(1),
		})
	
	z, err := NewReplConKafkaBroker(ctx, "kafka-3",  &ReplConKafkaArgs{
			Image: pulumi.String("navicore/kafka:0.10.1.1"),
			Replicas: pulumi.Int(1),
		})


	t, err := corev1.NewService(ctx, "kafka",  &corev1.ServiceArgs{
                        Metadata: &metav1.ObjectMetaArgs{
                                Name: pulumi.String("kafka"),
                                Labels: pulumi.StringMap{
                                        "service": pulumi.String("kafka"),
                                        "name": pulumi.String("kafka"),
                                },
                        },
                        Spec: &corev1.ServiceSpecArgs{
                                Ports: corev1.ServicePortArray{
                                        &corev1.ServicePortArgs{
                                                Name: pulumi.String("broker"),
                                                TargetPort: pulumi.Int(9092),
                                                Port: pulumi.Int(9092),
                                        },
                                },
                                Selector: pulumi.StringMap{
                                        "service": pulumi.String("kafka"),
                                },
                        },
                })

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(t)

	return nil

}

