
package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/spf13/pflag"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/cloud-provider"
	"k8s.io/cloud-provider/app"
	cloudcontrollerconfig "k8s.io/cloud-provider/app/config"
	"k8s.io/cloud-provider/options"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/logs"
	_ "k8s.io/component-base/metrics/prometheus/clientgo" // load all the prometheus client-go plugins
	_ "k8s.io/component-base/metrics/prometheus/version"  // for version metric registration
	"k8s.io/klog/v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	pflag.CommandLine.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)

	ccmOptions, err := options.NewCloudControllerManagerOptions()
	if err != nil {
		klog.Fatalf("unable to initialize command options: %v", err)
	}

	controllerInitializers := app.DefaultInitFuncConstructors
	// Here is an example to remove the controller which is not needed.
	// e.g. remove the cloud-node-lifecycle controller which current cloud provider does not need.
	//delete(controllerInitializers, "cloud-node-lifecycle")

	// Here is an example to add an controller(NodeIpamController) which will be used by cloud provider
	// generate nodeIPAMConfig. Here is an sample code.
	// If you do not need additional controller, please ignore.

	//nodeIpamController := nodeIPAMController{}
	//nodeIpamController.nodeIPAMControllerOptions.NodeIPAMControllerConfiguration = &nodeIpamController.nodeIPAMControllerConfiguration
	fss := cliflag.NamedFlagSets{}
	//nodeIpamController.nodeIPAMControllerOptions.AddFlags(fss.FlagSet("nodeipam controller"))
	//controllerInitializers["nodeipam"] = nodeIpamController.startNodeIpamControllerWrapper

	command := app.NewCloudControllerManagerCommand(ccmOptions, cloudInitializer, controllerInitializers, fss, wait.NeverStop)

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

func cloudInitializer(config *cloudcontrollerconfig.CompletedConfig) cloudprovider.Interface {
	cloudConfig := config.ComponentConfig.KubeCloudShared.CloudProvider
	// initialize cloud provider with the cloud provider name and config file provided
	cloud, err := cloudprovider.InitCloudProvider(cloudConfig.Name, cloudConfig.CloudConfigFile)
	if err != nil {
		klog.Fatalf("Cloud provider could not be initialized: %v", err)
	}
	if cloud == nil {
		klog.Fatalf("Cloud provider is nil")
	}

	if !cloud.HasClusterID() {
		if config.ComponentConfig.KubeCloudShared.AllowUntaggedCloud {
			klog.Warning("detected a cluster without a ClusterID.  A ClusterID will be required in the future.  Please tag your cluster to avoid any future issues")
		} else {
			klog.Fatalf("no ClusterID found.  A ClusterID is required for the cloud provider to function properly.  This check can be bypassed by setting the allow-untagged-cloud option")
		}
	}

	return cloud
}