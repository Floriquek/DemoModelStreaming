#  Model Streaming with Flask and Kafka

<br>
<br>
<b> Scope: </b> Preparing and deploying with Pulumi-Golang a very simple Model Streaming with Flask and Kafka.
<br>
<br>

<b> <i> Local Environemnt specs: </i> </b>

```
root@tron-VirtualBox:~# more /etc/lsb-release
DISTRIB_ID=Ubuntu
DISTRIB_RELEASE=20.04
DISTRIB_CODENAME=focal
DISTRIB_DESCRIPTION="Ubuntu 20.04.3 LTS"
root@tron-VirtualBox:~# pulumi version
v3.22.1
root@tron-VirtualBox:~#
root@tron-VirtualBox:~#
root@tron-VirtualBox:~# kubeadm version
kubeadm version: &version.Info{Major:"1", Minor:"23", GitVersion:"v1.23.1", GitCommit:"86ec240af8cbd1b60bcc4c03c20da9b98005b92e", 
GitTreeState:"clean", BuildDate:"2021-12-16T11:39:51Z", GoVersion:"go1.17.5", Compiler:"gc", Platform:"linux/amd64"}
root@tron-VirtualBox:~#
root@tron-VirtualBox:~# kubectl version
Client Version: version.Info{Major:"1", Minor:"23", GitVersion:"v1.23.1", GitCommit:"86ec240af8cbd1b60bcc4c03c20da9b98005b92e", GitTreeState:"clean", BuildDate:"2021-12-16T11:41:01Z", GoVersion:"go1.17.5", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"23", GitVersion:"v1.23.1", GitCommit:"86ec240af8cbd1b60bcc4c03c20da9b98005b92e", GitTreeState:"clean", BuildDate:"2021-12-16T11:34:54Z", GoVersion:"go1.17.5", Compiler:"gc", Platform:"linux/amd64"}

```

Prepare environment

1. Under DemoModelStreaming/flask/buildimage, prepare local "pandas" image:'

```
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming/flask/buildimage# docker build -t pandas .
```

2. Create zookeeper and kafka modules for necessary dependencies and requirements:

```
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming/zookeeper# go mod init zookeeper
go: creating new go.mod: module zookeeper
go: to add module requirements and sums:
        go mod tidy
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming/zookeeper# cd ../kafka
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming/kafka# go mod init kafka
go: creating new go.mod: module kafka
go: to add module requirements and sums:
        go mod tidy
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming/kafka# 
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming/kafka# cd ../flask/
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming/flask# go mod init flask
go: creating new go.mod: module flask
go: to add module requirements and sums:
        go mod tidy

root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming/flask# cd ..
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming# more Pulumi.yaml | awk  'NR==1 {print $2}'
ZooKaBo

root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming# go mod init ZooKaBo
go: creating new go.mod: module ZooKaBo
go: to add module requirements and sums:
        go mod tidy
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming# go mod edit --replace zookeeper=./zookeeper
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming# go mod edit --replace kafka=./kafka
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming# go mod edit --replace flask=./flask
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming#
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming# go mod tidy
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming# go mod tidy
go: finding module for package github.com/pulumi/pulumi/sdk/v3/go/pulumi
go: found flask in flask v0.0.0-00010101000000-000000000000
go: found kafka in kafka v0.0.0-00010101000000-000000000000
go: found zookeeper in zookeeper v0.0.0-00010101000000-000000000000
[ ... snip ... ]
```

3. Initiate your Stack
```
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming# pulumi stack init
stack name: (dev) ZooKaBo
Created stack 'ZooKaBo'
Enter your passphrase to protect config/secrets:
Re-enter your passphrase to confirm:
```
And now create the resources:

```
pulumi up 

[ ... snip ... ]
Do you want to perform this update? yes

[ ... snip ... ] 
Updating (exampleModel):
     Type                                            Name                         Status      Info
 +   pulumi:pulumi:Stack                             anotherExample-exampleModel  created     11 messages
 +   ├─ zookeeper-app                                zoo1                         created
 +   │  ├─ kubernetes:core/v1:Service                zoo1                         created
 +   │  └─ kubernetes:apps/v1:Deployment             zoo1                         created
 +   ├─ zookeeper-app                                zoo2                         created
 +   │  ├─ kubernetes:core/v1:Service                zoo2                         created
 +   │  └─ kubernetes:apps/v1:Deployment             zoo2                         created
 +   ├─ kafka-app                                    kafka-1                      created
 +   │  ├─ kubernetes:core/v1:Service                kafka-1                      created
 +   │  └─ kubernetes:core/v1:ReplicationController  kafka-1                      created
 +   ├─ zookeeper-app                                zoo3                         created
 +   │  ├─ kubernetes:core/v1:Service                zoo3                         created
 +   │  └─ kubernetes:apps/v1:Deployment             zoo3                         created
 +   ├─ kafka-app                                    kafka-2                      created
 +   │  ├─ kubernetes:core/v1:ReplicationController  kafka-2                      created
 +   │  └─ kubernetes:core/v1:Service                kafka-2                      created
 +   ├─ kafka-app                                    kafka-3                      created
 +   │  ├─ kubernetes:core/v1:Service                kafka-3                      created
 +   │  └─ kubernetes:core/v1:ReplicationController  kafka-3                      created
 +   ├─ kubernetes:core/v1:Service                   kafka                        created
 +   ├─ kubernetes:apps/v1:Deployment                flaskdepdep                  created
 +   └─ kubernetes:core/v1:Service                   flaskappservice              created

[ ... snip ... ]
Outputs:
    Node Port     : "flaskappservice-vw0gvovs"
    name flask pod: "flaskdepdep-cqlpe724"

Resources:
    + 22 created

Duration: 19s

```

Check for created resources:

```
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming# kubectl get svc
NAME                       TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                      AGE
flaskappservice-vw0gvovs   NodePort    10.110.222.123   <none>        5000:31624/TCP               102s
kafka                      ClusterIP   10.105.56.69     <none>        9092/TCP                     102s
kafka-1                    ClusterIP   10.105.38.244    <none>        9092/TCP                     102s
kafka-2                    ClusterIP   10.101.209.114   <none>        9092/TCP                     98s
kafka-3                    ClusterIP   10.100.230.0     <none>        9092/TCP                     97s
kubernetes                 ClusterIP   10.96.0.1        <none>        443/TCP                      19d
zoo1                       ClusterIP   10.111.65.129    <none>        2181/TCP,2888/TCP,3888/TCP   102s
zoo2                       ClusterIP   10.97.69.76      <none>        2181/TCP,2888/TCP,3888/TCP   99s
zoo3                       ClusterIP   10.102.37.107    <none>        2181/TCP,2888/TCP,3888/TCP   102s
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming# kubectl get pods
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming#
NAME                                    READY   STATUS    RESTARTS      AGE
flaskdepdep-cqlpe724-746c65fb89-hk4cd   1/1     Running   0             104s
kafka-1-7br6b                           1/1     Running   2 (85s ago)   99s
kafka-2-jz9l2                           1/1     Running   2 (90s ago)   101s
kafka-3-fx7zk                           1/1     Running   2 (87s ago)   99s
zoo1-5c47c64bb8-bdmsw                   1/1     Running   0             104s
zoo2-5f6bfdc6b7-cq889                   1/1     Running   0             99s
zoo3-54cd886746-6ht9v                   1/1     Running   0             104s

root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming#
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming# kubectl get rc
NAME      DESIRED   CURRENT   READY   AGE
kafka-1   1         1         1       2m29s
kafka-2   1         1         1       2m31s
kafka-3   1         1         1       2m29s
root@tron-VirtualBox:/home/tron/Desktop/DemoModelStreaming# kubectl get deployment
NAME                   READY   UP-TO-DATE   AVAILABLE   AGE
flaskdepdep-cqlpe724   1/1     1            1           2m37s
zoo1                   1/1     1            1           2m37s
zoo2                   1/1     1            1           2m32s
zoo3                   1/1     1            1           2m37s

```
