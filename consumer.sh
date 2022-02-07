#!/bin/sh

if [ $# -lt 1 ]
then
        echo "provide arguments: ./consumer.sh KAFKA_ARG"
fi

kubectl exec $1 -ti -- bash /kafka/bin/kafka-console-consumer.sh --bootstrap-server kafka-1:9092 --topic test1
