package zookeeper

import (
        "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"fmt"
)


  func	Zookeeper(ctx *pulumi.Context) error {


		x, err := NewDeploymentZookeeper(ctx, "zoo1", &DeploymentZookeeperArgs{
			Image: pulumi.String("navicore/zookeeper:3.4.9"),
			Replicas: pulumi.Int(1),


		})

		//
		if err != nil {
			return err
		}

                y, err := NewDeploymentZookeeper(ctx, "zoo2", &DeploymentZookeeperArgs{
                        Image: pulumi.String("navicore/zookeeper:3.4.9"),

                        Replicas: pulumi.Int(1),
                })

                if err != nil {
                        return err
                }

                z, err := NewDeploymentZookeeper(ctx, "zoo3", &DeploymentZookeeperArgs{
                        Image: pulumi.String("navicore/zookeeper:3.4.9"),

                        Replicas: pulumi.Int(1),
                })

                if err != nil {
                        return err
                }

		fmt.Println("x", x)
		fmt.Println("y", y)
		fmt.Println("z", z)

		return nil
  }

