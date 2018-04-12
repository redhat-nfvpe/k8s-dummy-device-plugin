# k8s-dummy-device-plugin

K8s Dummy Device Plugin *(for testing purpose only)*

This is a plugin that's used for testing and exploring [Kubernetes Device Plugins](https://kubernetes.io/docs/concepts/cluster-administration/device-plugins/).

In essence, it works as a kind of echo device. One specifies the (albeit pretend) devices in a JSON file, and the plugin operates on those, and allocates the devices to containers that request them -- it does this by setting those devices into environment variables in those containers.

## Building

This plugin is built by simply building the `dummy.go` file, such as:

```
go build dummy.go
```

Dependencies are managed and versioned internally with [dep](https://github.com/golang/dep).

## Deployment via daemonSet



## Usage

Currently, no arguments are honored at runtime.

## Configuration

Configuration of the "pretend" devices are in the `./dummyResources.json` file.