package provider

import (
	"fmt"
	"github.com/anexia-it/anxcloud-cloud-controller-manager/anx/provider/mocks"
	"github.com/anexia-it/go-anxcloud/pkg/ipam"
	"github.com/anexia-it/go-anxcloud/pkg/test"
	"github.com/anexia-it/go-anxcloud/pkg/vlan"
	"github.com/anexia-it/go-anxcloud/pkg/vsphere"
	v1 "k8s.io/api/core/v1"
)

//go:generate mockery --srcpkg github.com/anexia-it/go-anxcloud/pkg --name API --structname API --filename api.go
//go:generate mockery --srcpkg github.com/anexia-it/go-anxcloud/pkg/vsphere/powercontrol --name API --structname PowerControl --filename powercontrol.go
//go:generate mockery --srcpkg github.com/anexia-it/go-anxcloud/pkg/vsphere --name API --structname VSphere --filename vsphere.go
//go:generate mockery --srcpkg github.com/anexia-it/go-anxcloud/pkg/vsphere/search --name API --structname Search --filename search.go
//go:generate mockery --srcpkg github.com/anexia-it/go-anxcloud/pkg/vsphere/info --name API --structname Info --filename info.go

type mockedProvider struct {
	apiMock          *mocks.API
	vsphereMock      *mocks.VSphere
	powerControlMock *mocks.PowerControl
	searchMock       *mocks.Search
	infoMock         *mocks.Info
	config           *providerConfig
}

func getMockedAnxProvider() mockedProvider {
	apiMock := &mocks.API{}
	vsphereMock := &mocks.VSphere{}
	powerControlMock := &mocks.PowerControl{}
	searchMock := &mocks.Search{}
	infoMock := &mocks.Info{}

	vsphereMock.On("PowerControl").Return(powerControlMock)
	vsphereMock.On("Search").Return(searchMock)
	vsphereMock.On("Info").Return(infoMock)
	apiMock.On("VSphere").Return(vsphereMock)

	return mockedProvider{
		apiMock:          apiMock,
		vsphereMock:      vsphereMock,
		powerControlMock: powerControlMock,
		searchMock:       searchMock,
		infoMock:         infoMock,
	}
}

func (m mockedProvider) IPAM() ipam.API {
	panic("implement me")
}

func (m mockedProvider) Test() test.API {
	panic("implement me")
}

func (m mockedProvider) VLAN() vlan.API {
	panic("implement me")
}

func (m mockedProvider) VSphere() vsphere.API {
	return m.vsphereMock
}

func (m mockedProvider) Config() *providerConfig {
	if m.config == nil {
		return &providerConfig{
			CustomerPrefix: "no-set",
		}
	}
	return m.config
}

func providerManagedNode() v1.Node {
	return v1.Node{
		Spec: v1.NodeSpec{
			ProviderID: fmt.Sprintf("%s%s", cloudProviderScheme, nodeIdentifier),
		},
	}
}
