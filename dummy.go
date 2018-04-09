package main

import (
	"fmt"
	"github.com/golang/glog"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1alpha"
)

// DummyDeviceManager manages our dummy devices
type DummyDeviceManager struct {
	devices     map[string]*pluginapi.Device
	deviceFiles []string
}

// Init function for our dummy devices
func (dummydev *DummyDeviceManager) Init() error {
	glog.Info("Initializing dummy device plugin...\n")
	return nil
}

// Register with kubelet
func Register() error {
	conn, err := grpc.Dial(pluginapi.KubeletSocket, grpc.WithInsecure(),
		grpc.WithDialer(func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", addr, timeout)
		}))
	defer conn.Close()
	if err != nil {
		return fmt.Errorf("device-plugin: cannot connect to kubelet service: %v", err)
	}
	client := pluginapi.NewRegistrationClient(conn)
	reqt := &pluginapi.RegisterRequest{
		Version:      pluginapi.Version,
		// Name of the unix socket the device plugin is listening on
		// PATH = path.Join(DevicePluginPath, endpoint)
		Endpoint:     "dummy",
		// Schedulable resource name.
		ResourceName: "dummy/dummyDev",
	}

	_, err = client.Register(context.Background(), reqt)
	if err != nil {
		return fmt.Errorf("device-plugin: cannot register to kubelet service: %v", err)
	}
	return nil
}



func main(){
	// Registers with Kubelet.
	err := Register()
	if err != nil {
		glog.Fatal(err)
	}
	fmt.Printf("device-plugin registered\n")

}