


package sqlvirtualmachine

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SqlVirtualMachineGroup struct {
	pulumi.CustomResourceState

	ClusterConfiguration pulumi.StringOutput                `pulumi:"clusterConfiguration"`
	ClusterManagerType   pulumi.StringOutput                `pulumi:"clusterManagerType"`
	Location             pulumi.StringOutput                `pulumi:"location"`
	Name                 pulumi.StringOutput                `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput                `pulumi:"provisioningState"`
	ScaleType            pulumi.StringOutput                `pulumi:"scaleType"`
	SqlImageOffer        pulumi.StringPtrOutput             `pulumi:"sqlImageOffer"`
	SqlImageSku          pulumi.StringPtrOutput             `pulumi:"sqlImageSku"`
	Tags                 pulumi.StringMapOutput             `pulumi:"tags"`
	Type                 pulumi.StringOutput                `pulumi:"type"`
	WsfcDomainProfile    WsfcDomainProfileResponsePtrOutput `pulumi:"wsfcDomainProfile"`
}


func NewSqlVirtualMachineGroup(ctx *pulumi.Context,
	name string, args *SqlVirtualMachineGroupArgs, opts ...pulumi.ResourceOption) (*SqlVirtualMachineGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20170301preview:SqlVirtualMachineGroup"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20211101preview:SqlVirtualMachineGroup"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20220201:SqlVirtualMachineGroup"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine/v20220201preview:SqlVirtualMachineGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlVirtualMachineGroup
	err := ctx.RegisterResource("azure-native:sqlvirtualmachine:SqlVirtualMachineGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlVirtualMachineGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlVirtualMachineGroupState, opts ...pulumi.ResourceOption) (*SqlVirtualMachineGroup, error) {
	var resource SqlVirtualMachineGroup
	err := ctx.ReadResource("azure-native:sqlvirtualmachine:SqlVirtualMachineGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type sqlVirtualMachineGroupState struct {
}

type SqlVirtualMachineGroupState struct {
}

func (SqlVirtualMachineGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlVirtualMachineGroupState)(nil)).Elem()
}

type sqlVirtualMachineGroupArgs struct {
	Location                   *string            `pulumi:"location"`
	ResourceGroupName          string             `pulumi:"resourceGroupName"`
	SqlImageOffer              *string            `pulumi:"sqlImageOffer"`
	SqlImageSku                *string            `pulumi:"sqlImageSku"`
	SqlVirtualMachineGroupName *string            `pulumi:"sqlVirtualMachineGroupName"`
	Tags                       map[string]string  `pulumi:"tags"`
	WsfcDomainProfile          *WsfcDomainProfile `pulumi:"wsfcDomainProfile"`
}


type SqlVirtualMachineGroupArgs struct {
	Location                   pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	SqlImageOffer              pulumi.StringPtrInput
	SqlImageSku                pulumi.StringPtrInput
	SqlVirtualMachineGroupName pulumi.StringPtrInput
	Tags                       pulumi.StringMapInput
	WsfcDomainProfile          WsfcDomainProfilePtrInput
}

func (SqlVirtualMachineGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlVirtualMachineGroupArgs)(nil)).Elem()
}

type SqlVirtualMachineGroupInput interface {
	pulumi.Input

	ToSqlVirtualMachineGroupOutput() SqlVirtualMachineGroupOutput
	ToSqlVirtualMachineGroupOutputWithContext(ctx context.Context) SqlVirtualMachineGroupOutput
}

func (*SqlVirtualMachineGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlVirtualMachineGroup)(nil)).Elem()
}

func (i *SqlVirtualMachineGroup) ToSqlVirtualMachineGroupOutput() SqlVirtualMachineGroupOutput {
	return i.ToSqlVirtualMachineGroupOutputWithContext(context.Background())
}

func (i *SqlVirtualMachineGroup) ToSqlVirtualMachineGroupOutputWithContext(ctx context.Context) SqlVirtualMachineGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlVirtualMachineGroupOutput)
}

type SqlVirtualMachineGroupOutput struct{ *pulumi.OutputState }

func (SqlVirtualMachineGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlVirtualMachineGroup)(nil)).Elem()
}

func (o SqlVirtualMachineGroupOutput) ToSqlVirtualMachineGroupOutput() SqlVirtualMachineGroupOutput {
	return o
}

func (o SqlVirtualMachineGroupOutput) ToSqlVirtualMachineGroupOutputWithContext(ctx context.Context) SqlVirtualMachineGroupOutput {
	return o
}

func (o SqlVirtualMachineGroupOutput) ClusterConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringOutput { return v.ClusterConfiguration }).(pulumi.StringOutput)
}

func (o SqlVirtualMachineGroupOutput) ClusterManagerType() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringOutput { return v.ClusterManagerType }).(pulumi.StringOutput)
}

func (o SqlVirtualMachineGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SqlVirtualMachineGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SqlVirtualMachineGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SqlVirtualMachineGroupOutput) ScaleType() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringOutput { return v.ScaleType }).(pulumi.StringOutput)
}

func (o SqlVirtualMachineGroupOutput) SqlImageOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringPtrOutput { return v.SqlImageOffer }).(pulumi.StringPtrOutput)
}

func (o SqlVirtualMachineGroupOutput) SqlImageSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringPtrOutput { return v.SqlImageSku }).(pulumi.StringPtrOutput)
}

func (o SqlVirtualMachineGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SqlVirtualMachineGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SqlVirtualMachineGroupOutput) WsfcDomainProfile() WsfcDomainProfileResponsePtrOutput {
	return o.ApplyT(func(v *SqlVirtualMachineGroup) WsfcDomainProfileResponsePtrOutput { return v.WsfcDomainProfile }).(WsfcDomainProfileResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlVirtualMachineGroupOutput{})
}
