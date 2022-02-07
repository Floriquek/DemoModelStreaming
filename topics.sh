#!/bin/bash

if [ $# -lt 1 ]
then
        echo "provide arguments: ./script.sh KAFKA_ARG"
fi

KAFKA=$1

listTopics(){
        echo "topics, if any"
	topic=$(kubectl exec $KAFKA -- /kafka/bin/kafka-topics.sh --list --zookeeper zoo1,zoo2,zoo3)
        echo "here are the topics: $topic"
}

checkTopics(){
	echo "checking and/or creating topic test1"
	topic=$(kubectl exec $KAFKA -- /kafka/bin/kafka-topics.sh --list --zookeeper zoo1,zoo2,zoo3)

	if [[ $topic = "test1"  ]] 
	then
		echo "found topic $topic"
		exit 1
	else
		echo "creating topic"
		new_topic=`kubectl exec $KAFKA -- /kafka/bin/kafka-topics.sh --create --zookeeper zoo1 --topic test1 --partitions 1 --replication-factor 1`
		echo $new_topic
	fi
}	



listTopicsAgain(){
	echo "list again topics"
	topic=$(kubectl exec $KAFKA -- /kafka/bin/kafka-topics.sh --list --zookeeper zoo1,zoo2,zoo3)
	echo $topic
}

main(){
        listTopics 
	checkTopics
	listTopicsAgain
}

main
