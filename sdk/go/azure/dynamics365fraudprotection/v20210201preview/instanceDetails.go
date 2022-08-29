


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type InstanceDetails struct {
	pulumi.CustomResourceState

	Administration    DFPInstanceAdministratorsResponsePtrOutput `pulumi:"administration"`
	Location          pulumi.StringOutput                        `pulumi:"location"`
	Name              pulumi.StringOutput                        `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                        `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput                   `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                     `pulumi:"tags"`
	Type              pulumi.StringOutput                        `pulumi:"type"`
}


func NewInstanceDetails(ctx *pulumi.Context,
	name string, args *InstanceDetailsArgs, opts ...pulumi.ResourceOption) (*InstanceDetails, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:dynamics365fraudprotection:InstanceDetails"),
		},
	})
	opts = append(opts, aliases)
	var resource InstanceDetails
	err := ctx.RegisterResource("azure-native:dynamics365fraudprotection/v20210201preview:InstanceDetails", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetInstanceDetails(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceDetailsState, opts ...pulumi.ResourceOption) (*InstanceDetails, error) {
	var resource InstanceDetails
	err := ctx.ReadResource("azure-native:dynamics365fraudprotection/v20210201preview:InstanceDetails", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type instanceDetailsState struct {
}

type InstanceDetailsState struct {
}

func (InstanceDetailsState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceDetailsState)(nil)).Elem()
}

type instanceDetailsArgs struct {
	Administration    *DFPInstanceAdministrators `pulumi:"administration"`
	InstanceName      *string                    `pulumi:"instanceName"`
	Location          *string                    `pulumi:"location"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
	Tags              map[string]string          `pulumi:"tags"`
}


type InstanceDetailsArgs struct {
	Administration    DFPInstanceAdministratorsPtrInput
	InstanceName      pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (InstanceDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceDetailsArgs)(nil)).Elem()
}

type InstanceDetailsInput interface {
	pulumi.Input

	ToInstanceDetailsOutput() InstanceDetailsOutput
	ToInstanceDetailsOutputWithContext(ctx context.Context) InstanceDetailsOutput
}

func (*InstanceDetails) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceDetails)(nil)).Elem()
}

func (i *InstanceDetails) ToInstanceDetailsOutput() InstanceDetailsOutput {
	return i.ToInstanceDetailsOutputWithContext(context.Background())
}

func (i *InstanceDetails) ToInstanceDetailsOutputWithContext(ctx context.Context) InstanceDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceDetailsOutput)
}

type InstanceDetailsOutput struct{ *pulumi.OutputState }

func (InstanceDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceDetails)(nil)).Elem()
}

func (o InstanceDetailsOutput) ToInstanceDetailsOutput() InstanceDetailsOutput {
	return o
}

func (o InstanceDetailsOutput) ToInstanceDetailsOutputWithContext(ctx context.Context) InstanceDetailsOutput {
	return o
}

func (o InstanceDetailsOutput) Administration() DFPInstanceAdministratorsResponsePtrOutput {
	return o.ApplyT(func(v *InstanceDetails) DFPInstanceAdministratorsResponsePtrOutput { return v.Administration }).(DFPInstanceAdministratorsResponsePtrOutput)
}

func (o InstanceDetailsOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceDetails) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o InstanceDetailsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceDetails) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o InstanceDetailsOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceDetails) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o InstanceDetailsOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *InstanceDetails) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o InstanceDetailsOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InstanceDetails) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o InstanceDetailsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceDetails) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceDetailsOutput{})
}
