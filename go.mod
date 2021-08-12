module github.com/anexia-it/anxcloud-cloud-controller-manager

go 1.16

require (
	github.com/anexia-it/go-anxcloud v0.3.21
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/api v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/cloud-provider v0.21.2
	k8s.io/component-base v0.21.2
	k8s.io/klog/v2 v2.9.0
)

replace github.com/anexia-it/go-anxcloud => github.com/kstiehl/go-anxcloud v0.3.23-0.20210810173316-acfdc680f0b4
