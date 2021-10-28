


package v20170301preview

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
			Type: pulumi.String("azure-nextgen:sqlvirtualmachine/v20170301preview:SqlVirtualMachineGroup"),
		},
		{
			Type: pulumi.String("azure-native:sqlvirtualmachine:SqlVirtualMachineGroup"),
		},
		{
			Type: pulumi.String("azure-nextgen:sqlvirtualmachine:SqlVirtualMachineGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlVirtualMachineGroup
	err := ctx.RegisterResource("azure-native:sqlvirtualmachine/v20170301preview:SqlVirtualMachineGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSqlVirtualMachineGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlVirtualMachineGroupState, opts ...pulumi.ResourceOption) (*SqlVirtualMachineGroup, error) {
	var resource SqlVirtualMachineGroup
	err := ctx.ReadResource("azure-native:sqlvirtualmachine/v20170301preview:SqlVirtualMachineGroup", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*SqlVirtualMachineGroup)(nil))
}

func (i *SqlVirtualMachineGroup) ToSqlVirtualMachineGroupOutput() SqlVirtualMachineGroupOutput {
	return i.ToSqlVirtualMachineGroupOutputWithContext(context.Background())
}

func (i *SqlVirtualMachineGroup) ToSqlVirtualMachineGroupOutputWithContext(ctx context.Context) SqlVirtualMachineGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlVirtualMachineGroupOutput)
}

type SqlVirtualMachineGroupOutput struct{ *pulumi.OutputState }

func (SqlVirtualMachineGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlVirtualMachineGroup)(nil))
}

func (o SqlVirtualMachineGroupOutput) ToSqlVirtualMachineGroupOutput() SqlVirtualMachineGroupOutput {
	return o
}

func (o SqlVirtualMachineGroupOutput) ToSqlVirtualMachineGroupOutputWithContext(ctx context.Context) SqlVirtualMachineGroupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SqlVirtualMachineGroupInput)(nil)).Elem(), &SqlVirtualMachineGroup{})
	pulumi.RegisterOutputType(SqlVirtualMachineGroupOutput{})
}
