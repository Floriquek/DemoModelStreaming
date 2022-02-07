package main

import (
	"zookeeper"
	"kafka"
	"flask"
	"fmt"
        "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {

		fmt.Println(zookeeper.Zookeeper(ctx))
		fmt.Println(kafka.Kafka(ctx))
		fmt.Println(flask.Flask(ctx))

		return nil

    })
  }
