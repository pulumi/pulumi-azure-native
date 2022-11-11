


package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkloadNetworkVMGroup struct {
	pulumi.CustomResourceState

	DisplayName       pulumi.StringPtrOutput   `pulumi:"displayName"`
	Members           pulumi.StringArrayOutput `pulumi:"members"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	Revision          pulumi.Float64PtrOutput  `pulumi:"revision"`
	Status            pulumi.StringOutput      `pulumi:"status"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewWorkloadNetworkVMGroup(ctx *pulumi.Context,
	name string, args *WorkloadNetworkVMGroupArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkVMGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs:WorkloadNetworkVMGroup"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:WorkloadNetworkVMGroup"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkVMGroup"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkVMGroup"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:WorkloadNetworkVMGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkloadNetworkVMGroup
	err := ctx.RegisterResource("azure-native:avs/v20220501:WorkloadNetworkVMGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkloadNetworkVMGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkVMGroupState, opts ...pulumi.ResourceOption) (*WorkloadNetworkVMGroup, error) {
	var resource WorkloadNetworkVMGroup
	err := ctx.ReadResource("azure-native:avs/v20220501:WorkloadNetworkVMGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workloadNetworkVMGroupState struct {
}

type WorkloadNetworkVMGroupState struct {
}

func (WorkloadNetworkVMGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkVMGroupState)(nil)).Elem()
}

type workloadNetworkVMGroupArgs struct {
	DisplayName       *string  `pulumi:"displayName"`
	Members           []string `pulumi:"members"`
	PrivateCloudName  string   `pulumi:"privateCloudName"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	Revision          *float64 `pulumi:"revision"`
	VmGroupId         *string  `pulumi:"vmGroupId"`
}


type WorkloadNetworkVMGroupArgs struct {
	DisplayName       pulumi.StringPtrInput
	Members           pulumi.StringArrayInput
	PrivateCloudName  pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	Revision          pulumi.Float64PtrInput
	VmGroupId         pulumi.StringPtrInput
}

func (WorkloadNetworkVMGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkVMGroupArgs)(nil)).Elem()
}

type WorkloadNetworkVMGroupInput interface {
	pulumi.Input

	ToWorkloadNetworkVMGroupOutput() WorkloadNetworkVMGroupOutput
	ToWorkloadNetworkVMGroupOutputWithContext(ctx context.Context) WorkloadNetworkVMGroupOutput
}

func (*WorkloadNetworkVMGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkVMGroup)(nil)).Elem()
}

func (i *WorkloadNetworkVMGroup) ToWorkloadNetworkVMGroupOutput() WorkloadNetworkVMGroupOutput {
	return i.ToWorkloadNetworkVMGroupOutputWithContext(context.Background())
}

func (i *WorkloadNetworkVMGroup) ToWorkloadNetworkVMGroupOutputWithContext(ctx context.Context) WorkloadNetworkVMGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkVMGroupOutput)
}

type WorkloadNetworkVMGroupOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkVMGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkVMGroup)(nil)).Elem()
}

func (o WorkloadNetworkVMGroupOutput) ToWorkloadNetworkVMGroupOutput() WorkloadNetworkVMGroupOutput {
	return o
}

func (o WorkloadNetworkVMGroupOutput) ToWorkloadNetworkVMGroupOutputWithContext(ctx context.Context) WorkloadNetworkVMGroupOutput {
	return o
}

func (o WorkloadNetworkVMGroupOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkVMGroup) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o WorkloadNetworkVMGroupOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkloadNetworkVMGroup) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o WorkloadNetworkVMGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkVMGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkloadNetworkVMGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkVMGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WorkloadNetworkVMGroupOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkVMGroup) pulumi.Float64PtrOutput { return v.Revision }).(pulumi.Float64PtrOutput)
}

func (o WorkloadNetworkVMGroupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkVMGroup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func (o WorkloadNetworkVMGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkVMGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkVMGroupOutput{})
}
