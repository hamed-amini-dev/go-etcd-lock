In distributed systems, ensuring that only one process can access a shared resource at a time is crucial. A common solution to this problem is implementing a distributed lock system. In this article, we'll explore how to build a distributed lock system using etcd and Golang.

## What is etcd?

etcd is a distributed key-value store that provides a reliable way to store data across a cluster of machines. It's commonly used for service discovery, configuration management, and distributed coordination, including distributed locking.

## Why Use a Distributed Lock?

A distributed lock ensures that multiple processes or nodes do not access the same resource simultaneously, which is essential for maintaining data integrity and consistency in distributed systems.

## Setting Up etcd

First, ensure you have etcd installed and running. You can download it from the official etcd website.

Installing etcd Client in Golang

To interact with etcd from a Go application, we need the go.etcd.io/etcd/client/v3 package. Install it using:

```
go get go.etcd.io/etcd/client/v3
```

Running the Example

To run the example, use the following commands:

First, run this command:

```
go run main.go test1
```

Then, run this command:

```
go run main.go test2
```

The first command will acquire a lock from etcd, perform some operations, and then release the lock. After the first command completes, the second command will acquire the lock and proceed with its operations.
