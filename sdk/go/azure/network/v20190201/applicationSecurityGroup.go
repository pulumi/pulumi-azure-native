


package v20190201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationSecurityGroup struct {
	pulumi.CustomResourceState

	Etag              pulumi.StringOutput    `pulumi:"etag"`
	Location          pulumi.StringPtrOutput `pulumi:"location"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringOutput    `pulumi:"provisioningState"`
	ResourceGuid      pulumi.StringOutput    `pulumi:"resourceGuid"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewApplicationSecurityGroup(ctx *pulumi.Context,
	name string, args *ApplicationSecurityGroupArgs, opts ...pulumi.ResourceOption) (*ApplicationSecurityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ApplicationSecurityGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationSecurityGroup
	err := ctx.RegisterResource("azure-native:network/v20190201:ApplicationSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplicationSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationSecurityGroupState, opts ...pulumi.ResourceOption) (*ApplicationSecurityGroup, error) {
	var resource ApplicationSecurityGroup
	err := ctx.ReadResource("azure-native:network/v20190201:ApplicationSecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationSecurityGroupState struct {
}

type ApplicationSecurityGroupState struct {
}

func (ApplicationSecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSecurityGroupState)(nil)).Elem()
}

type applicationSecurityGroupArgs struct {
	ApplicationSecurityGroupName *string           `pulumi:"applicationSecurityGroupName"`
	Id                           *string           `pulumi:"id"`
	Location                     *string           `pulumi:"location"`
	ResourceGroupName            string            `pulumi:"resourceGroupName"`
	Tags                         map[string]string `pulumi:"tags"`
}


type ApplicationSecurityGroupArgs struct {
	ApplicationSecurityGroupName pulumi.StringPtrInput
	Id                           pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	Tags                         pulumi.StringMapInput
}

func (ApplicationSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSecurityGroupArgs)(nil)).Elem()
}

type ApplicationSecurityGroupInput interface {
	pulumi.Input

	ToApplicationSecurityGroupOutput() ApplicationSecurityGroupOutput
	ToApplicationSecurityGroupOutputWithContext(ctx context.Context) ApplicationSecurityGroupOutput
}

func (*ApplicationSecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationSecurityGroup)(nil)).Elem()
}

func (i *ApplicationSecurityGroup) ToApplicationSecurityGroupOutput() ApplicationSecurityGroupOutput {
	return i.ToApplicationSecurityGroupOutputWithContext(context.Background())
}

func (i *ApplicationSecurityGroup) ToApplicationSecurityGroupOutputWithContext(ctx context.Context) ApplicationSecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSecurityGroupOutput)
}

type ApplicationSecurityGroupOutput struct{ *pulumi.OutputState }

func (ApplicationSecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationSecurityGroup)(nil)).Elem()
}

func (o ApplicationSecurityGroupOutput) ToApplicationSecurityGroupOutput() ApplicationSecurityGroupOutput {
	return o
}

func (o ApplicationSecurityGroupOutput) ToApplicationSecurityGroupOutputWithContext(ctx context.Context) ApplicationSecurityGroupOutput {
	return o
}

func (o ApplicationSecurityGroupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSecurityGroup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ApplicationSecurityGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationSecurityGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ApplicationSecurityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSecurityGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApplicationSecurityGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSecurityGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ApplicationSecurityGroupOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSecurityGroup) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o ApplicationSecurityGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApplicationSecurityGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ApplicationSecurityGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSecurityGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationSecurityGroupOutput{})
}
