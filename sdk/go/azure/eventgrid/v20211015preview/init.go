


package v20211015preview

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:eventgrid/v20211015preview:Channel":
		r = &Channel{}
	case "azure-native:eventgrid/v20211015preview:Domain":
		r = &Domain{}
	case "azure-native:eventgrid/v20211015preview:DomainEventSubscription":
		r = &DomainEventSubscription{}
	case "azure-native:eventgrid/v20211015preview:DomainTopic":
		r = &DomainTopic{}
	case "azure-native:eventgrid/v20211015preview:DomainTopicEventSubscription":
		r = &DomainTopicEventSubscription{}
	case "azure-native:eventgrid/v20211015preview:EventChannel":
		r = &EventChannel{}
	case "azure-native:eventgrid/v20211015preview:EventSubscription":
		r = &EventSubscription{}
	case "azure-native:eventgrid/v20211015preview:PartnerConfiguration":
		r = &PartnerConfiguration{}
	case "azure-native:eventgrid/v20211015preview:PartnerDestination":
		r = &PartnerDestination{}
	case "azure-native:eventgrid/v20211015preview:PartnerNamespace":
		r = &PartnerNamespace{}
	case "azure-native:eventgrid/v20211015preview:PartnerRegistration":
		r = &PartnerRegistration{}
	case "azure-native:eventgrid/v20211015preview:PartnerTopic":
		r = &PartnerTopic{}
	case "azure-native:eventgrid/v20211015preview:PartnerTopicEventSubscription":
		r = &PartnerTopicEventSubscription{}
	case "azure-native:eventgrid/v20211015preview:PrivateEndpointConnection":
		r = &PrivateEndpointConnection{}
	case "azure-native:eventgrid/v20211015preview:SystemTopic":
		r = &SystemTopic{}
	case "azure-native:eventgrid/v20211015preview:SystemTopicEventSubscription":
		r = &SystemTopicEventSubscription{}
	case "azure-native:eventgrid/v20211015preview:Topic":
		r = &Topic{}
	case "azure-native:eventgrid/v20211015preview:TopicEventSubscription":
		r = &TopicEventSubscription{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"eventgrid/v20211015preview",
		&module{version},
	)
}
