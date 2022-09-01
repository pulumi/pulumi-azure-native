


package v20150601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Origin struct {
	pulumi.CustomResourceState

	HostName          pulumi.StringOutput `pulumi:"hostName"`
	HttpPort          pulumi.IntPtrOutput `pulumi:"httpPort"`
	HttpsPort         pulumi.IntPtrOutput `pulumi:"httpsPort"`
	Name              pulumi.StringOutput `pulumi:"name"`
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	ResourceState     pulumi.StringOutput `pulumi:"resourceState"`
	Type              pulumi.StringOutput `pulumi:"type"`
}


func NewOrigin(ctx *pulumi.Context,
	name string, args *OriginArgs, opts ...pulumi.ResourceOption) (*Origin, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointName == nil {
		return nil, errors.New("invalid value for required argument 'EndpointName'")
	}
	if args.HostName == nil {
		return nil, errors.New("invalid value for required argument 'HostName'")
	}
	if args.ProfileName == nil {
		return nil, errors.New("invalid value for required argument 'ProfileName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cdn:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20160402:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20191231:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200331:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200415:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20200901:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20210601:Origin"),
		},
		{
			Type: pulumi.String("azure-native:cdn/v20220501preview:Origin"),
		},
	})
	opts = append(opts, aliases)
	var resource Origin
	err := ctx.RegisterResource("azure-native:cdn/v20150601:Origin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetOrigin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OriginState, opts ...pulumi.ResourceOption) (*Origin, error) {
	var resource Origin
	err := ctx.ReadResource("azure-native:cdn/v20150601:Origin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type originState struct {
}

type OriginState struct {
}

func (OriginState) ElementType() reflect.Type {
	return reflect.TypeOf((*originState)(nil)).Elem()
}

type originArgs struct {
	EndpointName      string  `pulumi:"endpointName"`
	HostName          string  `pulumi:"hostName"`
	HttpPort          *int    `pulumi:"httpPort"`
	HttpsPort         *int    `pulumi:"httpsPort"`
	OriginName        *string `pulumi:"originName"`
	ProfileName       string  `pulumi:"profileName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type OriginArgs struct {
	EndpointName      pulumi.StringInput
	HostName          pulumi.StringInput
	HttpPort          pulumi.IntPtrInput
	HttpsPort         pulumi.IntPtrInput
	OriginName        pulumi.StringPtrInput
	ProfileName       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
}

func (OriginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*originArgs)(nil)).Elem()
}

type OriginInput interface {
	pulumi.Input

	ToOriginOutput() OriginOutput
	ToOriginOutputWithContext(ctx context.Context) OriginOutput
}

func (*Origin) ElementType() reflect.Type {
	return reflect.TypeOf((**Origin)(nil)).Elem()
}

func (i *Origin) ToOriginOutput() OriginOutput {
	return i.ToOriginOutputWithContext(context.Background())
}

func (i *Origin) ToOriginOutputWithContext(ctx context.Context) OriginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginOutput)
}

type OriginOutput struct{ *pulumi.OutputState }

func (OriginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Origin)(nil)).Elem()
}

func (o OriginOutput) ToOriginOutput() OriginOutput {
	return o
}

func (o OriginOutput) ToOriginOutputWithContext(ctx context.Context) OriginOutput {
	return o
}

func (o OriginOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.HostName }).(pulumi.StringOutput)
}

func (o OriginOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.IntPtrOutput { return v.HttpPort }).(pulumi.IntPtrOutput)
}

func (o OriginOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Origin) pulumi.IntPtrOutput { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

func (o OriginOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OriginOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o OriginOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.ResourceState }).(pulumi.StringOutput)
}

func (o OriginOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Origin) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OriginOutput{})
}
