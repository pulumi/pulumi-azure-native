


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteRelayServiceConnection struct {
	pulumi.CustomResourceState

	BiztalkUri               pulumi.StringPtrOutput `pulumi:"biztalkUri"`
	EntityConnectionString   pulumi.StringPtrOutput `pulumi:"entityConnectionString"`
	EntityName               pulumi.StringPtrOutput `pulumi:"entityName"`
	Hostname                 pulumi.StringPtrOutput `pulumi:"hostname"`
	Kind                     pulumi.StringPtrOutput `pulumi:"kind"`
	Location                 pulumi.StringOutput    `pulumi:"location"`
	Name                     pulumi.StringPtrOutput `pulumi:"name"`
	Port                     pulumi.IntPtrOutput    `pulumi:"port"`
	ResourceConnectionString pulumi.StringPtrOutput `pulumi:"resourceConnectionString"`
	ResourceType             pulumi.StringPtrOutput `pulumi:"resourceType"`
	Tags                     pulumi.StringMapOutput `pulumi:"tags"`
	Type                     pulumi.StringPtrOutput `pulumi:"type"`
}


func NewSiteRelayServiceConnection(ctx *pulumi.Context,
	name string, args *SiteRelayServiceConnectionArgs, opts ...pulumi.ResourceOption) (*SiteRelayServiceConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:SiteRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteRelayServiceConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteRelayServiceConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteRelayServiceConnection
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteRelayServiceConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteRelayServiceConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteRelayServiceConnectionState, opts ...pulumi.ResourceOption) (*SiteRelayServiceConnection, error) {
	var resource SiteRelayServiceConnection
	err := ctx.ReadResource("azure-native:web/v20150801:SiteRelayServiceConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteRelayServiceConnectionState struct {
}

type SiteRelayServiceConnectionState struct {
}

func (SiteRelayServiceConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteRelayServiceConnectionState)(nil)).Elem()
}

type siteRelayServiceConnectionArgs struct {
	BiztalkUri               *string           `pulumi:"biztalkUri"`
	EntityConnectionString   *string           `pulumi:"entityConnectionString"`
	EntityName               *string           `pulumi:"entityName"`
	Hostname                 *string           `pulumi:"hostname"`
	Id                       *string           `pulumi:"id"`
	Kind                     *string           `pulumi:"kind"`
	Location                 *string           `pulumi:"location"`
	Name                     string            `pulumi:"name"`
	Port                     *int              `pulumi:"port"`
	ResourceConnectionString *string           `pulumi:"resourceConnectionString"`
	ResourceGroupName        string            `pulumi:"resourceGroupName"`
	ResourceType             *string           `pulumi:"resourceType"`
	Tags                     map[string]string `pulumi:"tags"`
	Type                     *string           `pulumi:"type"`
}


type SiteRelayServiceConnectionArgs struct {
	BiztalkUri               pulumi.StringPtrInput
	EntityConnectionString   pulumi.StringPtrInput
	EntityName               pulumi.StringPtrInput
	Hostname                 pulumi.StringPtrInput
	Id                       pulumi.StringPtrInput
	Kind                     pulumi.StringPtrInput
	Location                 pulumi.StringPtrInput
	Name                     pulumi.StringInput
	Port                     pulumi.IntPtrInput
	ResourceConnectionString pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	ResourceType             pulumi.StringPtrInput
	Tags                     pulumi.StringMapInput
	Type                     pulumi.StringPtrInput
}

func (SiteRelayServiceConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteRelayServiceConnectionArgs)(nil)).Elem()
}

type SiteRelayServiceConnectionInput interface {
	pulumi.Input

	ToSiteRelayServiceConnectionOutput() SiteRelayServiceConnectionOutput
	ToSiteRelayServiceConnectionOutputWithContext(ctx context.Context) SiteRelayServiceConnectionOutput
}

func (*SiteRelayServiceConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteRelayServiceConnection)(nil)).Elem()
}

func (i *SiteRelayServiceConnection) ToSiteRelayServiceConnectionOutput() SiteRelayServiceConnectionOutput {
	return i.ToSiteRelayServiceConnectionOutputWithContext(context.Background())
}

func (i *SiteRelayServiceConnection) ToSiteRelayServiceConnectionOutputWithContext(ctx context.Context) SiteRelayServiceConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteRelayServiceConnectionOutput)
}

type SiteRelayServiceConnectionOutput struct{ *pulumi.OutputState }

func (SiteRelayServiceConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteRelayServiceConnection)(nil)).Elem()
}

func (o SiteRelayServiceConnectionOutput) ToSiteRelayServiceConnectionOutput() SiteRelayServiceConnectionOutput {
	return o
}

func (o SiteRelayServiceConnectionOutput) ToSiteRelayServiceConnectionOutputWithContext(ctx context.Context) SiteRelayServiceConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SiteRelayServiceConnectionOutput{})
}
