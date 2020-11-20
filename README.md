# pncounter-crdt

PNCounter CRDT Cluster implemented in Go & Docker

## Introduction

CRDTs (Commutative Replicated Data Types) are a certain form of data types that when replicated across several nodes over a network achieve eventual consistency without the need for a consensus round. PNCounters abbreviated as positive-negative counters are CRDT counters modified to increment as well as decrement the count in it and becomes consistent across nodes in a cluster having replicated the counter.

## Example

After building a cluster of PNCounter nodes we can now increment counts to either one or many nodes in the cluster.

```
$ curl -i -X GET localhost:8080/pncounter/increment
$ curl -i -X GET localhost:8081/pncounter/decrement
```

When looking up the total count in the counter they then sync up with each other and thus return consistent values every time from any node in the cluster

```
$ curl -i -X GET localhost:8081/pncounter/count
{
    count: 2
}
```

The values remain consistent for nodes in the cluster that have never incremented the count value in it

```
$ curl -i -X GET localhost:8082/pncounter/count
{
    count: 2
}
```

## Steps

After cloning the repo. To provision the cluster:

```
$ make provision
```

This creates a 3 node PNCounter cluster established in their own docker network.

To view the status of the cluster

```
$ make info
```

Now we can send requests to increment, decrement, and get the total cluster count of any peer node using its port allocated.

```
$ curl -i -X GET localhost:<peer-port>/pncounter/increment
$ curl -i -X GET localhost:<peer-port>/pncounter/decrement
$ curl -i -X GET localhost:<peer-port>/pncounter/count
```

In the logs for each peer docker container, we can see the logs of the peer nodes getting in sync during read operations.

To tear down the cluster and remove the built docker images:

```
$ make clean
```

This is not certain to clean up all the locally created docker images at times. You can do a docker rmi to delete them.

## References

- [A CRDT Primer: Defanging Order Theory](https://www.youtube.com/watch?v=OOlnp2bZVRs) [John Mumm]
- [A comprehensive study of Convergent and Commutative Replicated Data Types](https://hal.inria.fr/inria-00555588/document) [Marc Shapiro et al]
