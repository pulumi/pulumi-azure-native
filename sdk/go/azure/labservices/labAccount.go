


package labservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LabAccount struct {
	pulumi.CustomResourceState

	EnabledRegionSelection pulumi.BoolPtrOutput                      `pulumi:"enabledRegionSelection"`
	LatestOperationResult  LatestOperationResultResponseOutput       `pulumi:"latestOperationResult"`
	Location               pulumi.StringPtrOutput                    `pulumi:"location"`
	Name                   pulumi.StringOutput                       `pulumi:"name"`
	ProvisioningState      pulumi.StringPtrOutput                    `pulumi:"provisioningState"`
	SizeConfiguration      SizeConfigurationPropertiesResponseOutput `pulumi:"sizeConfiguration"`
	Tags                   pulumi.StringMapOutput                    `pulumi:"tags"`
	Type                   pulumi.StringOutput                       `pulumi:"type"`
	UniqueIdentifier       pulumi.StringPtrOutput                    `pulumi:"uniqueIdentifier"`
}


func NewLabAccount(ctx *pulumi.Context,
	name string, args *LabAccountArgs, opts ...pulumi.ResourceOption) (*LabAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:labservices/v20181015:LabAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource LabAccount
	err := ctx.RegisterResource("azure-native:labservices:LabAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLabAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabAccountState, opts ...pulumi.ResourceOption) (*LabAccount, error) {
	var resource LabAccount
	err := ctx.ReadResource("azure-native:labservices:LabAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type labAccountState struct {
}

type LabAccountState struct {
}

func (LabAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*labAccountState)(nil)).Elem()
}

type labAccountArgs struct {
	EnabledRegionSelection *bool             `pulumi:"enabledRegionSelection"`
	LabAccountName         *string           `pulumi:"labAccountName"`
	Location               *string           `pulumi:"location"`
	ProvisioningState      *string           `pulumi:"provisioningState"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	Tags                   map[string]string `pulumi:"tags"`
	UniqueIdentifier       *string           `pulumi:"uniqueIdentifier"`
}


type LabAccountArgs struct {
	EnabledRegionSelection pulumi.BoolPtrInput
	LabAccountName         pulumi.StringPtrInput
	Location               pulumi.StringPtrInput
	ProvisioningState      pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
	UniqueIdentifier       pulumi.StringPtrInput
}

func (LabAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labAccountArgs)(nil)).Elem()
}

type LabAccountInput interface {
	pulumi.Input

	ToLabAccountOutput() LabAccountOutput
	ToLabAccountOutputWithContext(ctx context.Context) LabAccountOutput
}

func (*LabAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**LabAccount)(nil)).Elem()
}

func (i *LabAccount) ToLabAccountOutput() LabAccountOutput {
	return i.ToLabAccountOutputWithContext(context.Background())
}

func (i *LabAccount) ToLabAccountOutputWithContext(ctx context.Context) LabAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabAccountOutput)
}

type LabAccountOutput struct{ *pulumi.OutputState }

func (LabAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabAccount)(nil)).Elem()
}

func (o LabAccountOutput) ToLabAccountOutput() LabAccountOutput {
	return o
}

func (o LabAccountOutput) ToLabAccountOutputWithContext(ctx context.Context) LabAccountOutput {
	return o
}

func (o LabAccountOutput) EnabledRegionSelection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LabAccount) pulumi.BoolPtrOutput { return v.EnabledRegionSelection }).(pulumi.BoolPtrOutput)
}

func (o LabAccountOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v *LabAccount) LatestOperationResultResponseOutput { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

func (o LabAccountOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabAccount) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LabAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LabAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LabAccountOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabAccount) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LabAccountOutput) SizeConfiguration() SizeConfigurationPropertiesResponseOutput {
	return o.ApplyT(func(v *LabAccount) SizeConfigurationPropertiesResponseOutput { return v.SizeConfiguration }).(SizeConfigurationPropertiesResponseOutput)
}

func (o LabAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LabAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LabAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LabAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o LabAccountOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LabAccount) pulumi.StringPtrOutput { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LabAccountOutput{})
}
