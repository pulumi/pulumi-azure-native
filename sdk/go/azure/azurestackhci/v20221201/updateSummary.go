


package v20221201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type UpdateSummary struct {
	pulumi.CustomResourceState

	CurrentVersion    pulumi.StringPtrOutput   `pulumi:"currentVersion"`
	HardwareModel     pulumi.StringPtrOutput   `pulumi:"hardwareModel"`
	HealthCheckDate   pulumi.StringPtrOutput   `pulumi:"healthCheckDate"`
	LastChecked       pulumi.StringPtrOutput   `pulumi:"lastChecked"`
	LastUpdated       pulumi.StringPtrOutput   `pulumi:"lastUpdated"`
	Location          pulumi.StringPtrOutput   `pulumi:"location"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	OemFamily         pulumi.StringPtrOutput   `pulumi:"oemFamily"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	State             pulumi.StringPtrOutput   `pulumi:"state"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewUpdateSummary(ctx *pulumi.Context,
	name string, args *UpdateSummaryArgs, opts ...pulumi.ResourceOption) (*UpdateSummary, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource UpdateSummary
	err := ctx.RegisterResource("azure-native:azurestackhci/v20221201:UpdateSummary", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetUpdateSummary(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UpdateSummaryState, opts ...pulumi.ResourceOption) (*UpdateSummary, error) {
	var resource UpdateSummary
	err := ctx.ReadResource("azure-native:azurestackhci/v20221201:UpdateSummary", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type updateSummaryState struct {
}

type UpdateSummaryState struct {
}

func (UpdateSummaryState) ElementType() reflect.Type {
	return reflect.TypeOf((*updateSummaryState)(nil)).Elem()
}

type updateSummaryArgs struct {
	ClusterName       string  `pulumi:"clusterName"`
	CurrentVersion    *string `pulumi:"currentVersion"`
	HardwareModel     *string `pulumi:"hardwareModel"`
	HealthCheckDate   *string `pulumi:"healthCheckDate"`
	LastChecked       *string `pulumi:"lastChecked"`
	LastUpdated       *string `pulumi:"lastUpdated"`
	Location          *string `pulumi:"location"`
	OemFamily         *string `pulumi:"oemFamily"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	State             *string `pulumi:"state"`
}


type UpdateSummaryArgs struct {
	ClusterName       pulumi.StringInput
	CurrentVersion    pulumi.StringPtrInput
	HardwareModel     pulumi.StringPtrInput
	HealthCheckDate   pulumi.StringPtrInput
	LastChecked       pulumi.StringPtrInput
	LastUpdated       pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	OemFamily         pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	State             pulumi.StringPtrInput
}

func (UpdateSummaryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*updateSummaryArgs)(nil)).Elem()
}

type UpdateSummaryInput interface {
	pulumi.Input

	ToUpdateSummaryOutput() UpdateSummaryOutput
	ToUpdateSummaryOutputWithContext(ctx context.Context) UpdateSummaryOutput
}

func (*UpdateSummary) ElementType() reflect.Type {
	return reflect.TypeOf((**UpdateSummary)(nil)).Elem()
}

func (i *UpdateSummary) ToUpdateSummaryOutput() UpdateSummaryOutput {
	return i.ToUpdateSummaryOutputWithContext(context.Background())
}

func (i *UpdateSummary) ToUpdateSummaryOutputWithContext(ctx context.Context) UpdateSummaryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdateSummaryOutput)
}

type UpdateSummaryOutput struct{ *pulumi.OutputState }

func (UpdateSummaryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpdateSummary)(nil)).Elem()
}

func (o UpdateSummaryOutput) ToUpdateSummaryOutput() UpdateSummaryOutput {
	return o
}

func (o UpdateSummaryOutput) ToUpdateSummaryOutputWithContext(ctx context.Context) UpdateSummaryOutput {
	return o
}

func (o UpdateSummaryOutput) CurrentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.CurrentVersion }).(pulumi.StringPtrOutput)
}

func (o UpdateSummaryOutput) HardwareModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.HardwareModel }).(pulumi.StringPtrOutput)
}

func (o UpdateSummaryOutput) HealthCheckDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.HealthCheckDate }).(pulumi.StringPtrOutput)
}

func (o UpdateSummaryOutput) LastChecked() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.LastChecked }).(pulumi.StringPtrOutput)
}

func (o UpdateSummaryOutput) LastUpdated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.LastUpdated }).(pulumi.StringPtrOutput)
}

func (o UpdateSummaryOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o UpdateSummaryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UpdateSummaryOutput) OemFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.OemFamily }).(pulumi.StringPtrOutput)
}

func (o UpdateSummaryOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o UpdateSummaryOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o UpdateSummaryOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *UpdateSummary) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o UpdateSummaryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateSummary) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UpdateSummaryOutput{})
}
