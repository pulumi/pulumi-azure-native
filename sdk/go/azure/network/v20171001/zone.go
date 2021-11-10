


package v20171001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Zone struct {
	pulumi.CustomResourceState

	Etag                           pulumi.StringPtrOutput   `pulumi:"etag"`
	Location                       pulumi.StringOutput      `pulumi:"location"`
	MaxNumberOfRecordSets          pulumi.Float64Output     `pulumi:"maxNumberOfRecordSets"`
	MaxNumberOfRecordsPerRecordSet pulumi.Float64Output     `pulumi:"maxNumberOfRecordsPerRecordSet"`
	Name                           pulumi.StringOutput      `pulumi:"name"`
	NameServers                    pulumi.StringArrayOutput `pulumi:"nameServers"`
	NumberOfRecordSets             pulumi.Float64Output     `pulumi:"numberOfRecordSets"`
	Tags                           pulumi.StringMapOutput   `pulumi:"tags"`
	Type                           pulumi.StringOutput      `pulumi:"type"`
	ZoneType                       pulumi.StringPtrOutput   `pulumi:"zoneType"`
}


func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOption) (*Zone, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ZoneType == nil {
		args.ZoneType = ZoneType("Public")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:Zone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150504preview:Zone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160401:Zone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:Zone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180301preview:Zone"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180501:Zone"),
		},
	})
	opts = append(opts, aliases)
	var resource Zone
	err := ctx.RegisterResource("azure-native:network/v20171001:Zone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneState, opts ...pulumi.ResourceOption) (*Zone, error) {
	var resource Zone
	err := ctx.ReadResource("azure-native:network/v20171001:Zone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type zoneState struct {
}

type ZoneState struct {
}

func (ZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneState)(nil)).Elem()
}

type zoneArgs struct {
	Etag              *string           `pulumi:"etag"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	ZoneName          *string           `pulumi:"zoneName"`
	ZoneType          *ZoneType         `pulumi:"zoneType"`
}


type ZoneArgs struct {
	Etag              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	ZoneName          pulumi.StringPtrInput
	ZoneType          ZoneTypePtrInput
}

func (ZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneArgs)(nil)).Elem()
}

type ZoneInput interface {
	pulumi.Input

	ToZoneOutput() ZoneOutput
	ToZoneOutputWithContext(ctx context.Context) ZoneOutput
}

func (*Zone) ElementType() reflect.Type {
	return reflect.TypeOf((*Zone)(nil))
}

func (i *Zone) ToZoneOutput() ZoneOutput {
	return i.ToZoneOutputWithContext(context.Background())
}

func (i *Zone) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneOutput)
}

type ZoneOutput struct{ *pulumi.OutputState }

func (ZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Zone)(nil))
}

func (o ZoneOutput) ToZoneOutput() ZoneOutput {
	return o
}

func (o ZoneOutput) ToZoneOutputWithContext(ctx context.Context) ZoneOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ZoneOutput{})
}
