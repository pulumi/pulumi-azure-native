


package v20221201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type UpdateRun struct {
	pulumi.CustomResourceState

	Description        pulumi.StringPtrOutput   `pulumi:"description"`
	Duration           pulumi.StringPtrOutput   `pulumi:"duration"`
	EndTimeUtc         pulumi.StringPtrOutput   `pulumi:"endTimeUtc"`
	ErrorMessage       pulumi.StringPtrOutput   `pulumi:"errorMessage"`
	LastUpdatedTime    pulumi.StringPtrOutput   `pulumi:"lastUpdatedTime"`
	LastUpdatedTimeUtc pulumi.StringPtrOutput   `pulumi:"lastUpdatedTimeUtc"`
	Location           pulumi.StringPtrOutput   `pulumi:"location"`
	Name               pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState  pulumi.StringOutput      `pulumi:"provisioningState"`
	StartTimeUtc       pulumi.StringPtrOutput   `pulumi:"startTimeUtc"`
	State              pulumi.StringPtrOutput   `pulumi:"state"`
	Status             pulumi.StringPtrOutput   `pulumi:"status"`
	Steps              StepResponseArrayOutput  `pulumi:"steps"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	TimeStarted        pulumi.StringPtrOutput   `pulumi:"timeStarted"`
	Type               pulumi.StringOutput      `pulumi:"type"`
}


func NewUpdateRun(ctx *pulumi.Context,
	name string, args *UpdateRunArgs, opts ...pulumi.ResourceOption) (*UpdateRun, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UpdateName == nil {
		return nil, errors.New("invalid value for required argument 'UpdateName'")
	}
	var resource UpdateRun
	err := ctx.RegisterResource("azure-native:azurestackhci/v20221201:UpdateRun", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetUpdateRun(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UpdateRunState, opts ...pulumi.ResourceOption) (*UpdateRun, error) {
	var resource UpdateRun
	err := ctx.ReadResource("azure-native:azurestackhci/v20221201:UpdateRun", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type updateRunState struct {
}

type UpdateRunState struct {
}

func (UpdateRunState) ElementType() reflect.Type {
	return reflect.TypeOf((*updateRunState)(nil)).Elem()
}

type updateRunArgs struct {
	ClusterName        string  `pulumi:"clusterName"`
	Description        *string `pulumi:"description"`
	Duration           *string `pulumi:"duration"`
	EndTimeUtc         *string `pulumi:"endTimeUtc"`
	ErrorMessage       *string `pulumi:"errorMessage"`
	LastUpdatedTime    *string `pulumi:"lastUpdatedTime"`
	LastUpdatedTimeUtc *string `pulumi:"lastUpdatedTimeUtc"`
	Location           *string `pulumi:"location"`
	Name               *string `pulumi:"name"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	StartTimeUtc       *string `pulumi:"startTimeUtc"`
	State              *string `pulumi:"state"`
	Status             *string `pulumi:"status"`
	Steps              []Step  `pulumi:"steps"`
	TimeStarted        *string `pulumi:"timeStarted"`
	UpdateName         string  `pulumi:"updateName"`
	UpdateRunName      *string `pulumi:"updateRunName"`
}


type UpdateRunArgs struct {
	ClusterName        pulumi.StringInput
	Description        pulumi.StringPtrInput
	Duration           pulumi.StringPtrInput
	EndTimeUtc         pulumi.StringPtrInput
	ErrorMessage       pulumi.StringPtrInput
	LastUpdatedTime    pulumi.StringPtrInput
	LastUpdatedTimeUtc pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	Name               pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	StartTimeUtc       pulumi.StringPtrInput
	State              pulumi.StringPtrInput
	Status             pulumi.StringPtrInput
	Steps              StepArrayInput
	TimeStarted        pulumi.StringPtrInput
	UpdateName         pulumi.StringInput
	UpdateRunName      pulumi.StringPtrInput
}

func (UpdateRunArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*updateRunArgs)(nil)).Elem()
}

type UpdateRunInput interface {
	pulumi.Input

	ToUpdateRunOutput() UpdateRunOutput
	ToUpdateRunOutputWithContext(ctx context.Context) UpdateRunOutput
}

func (*UpdateRun) ElementType() reflect.Type {
	return reflect.TypeOf((**UpdateRun)(nil)).Elem()
}

func (i *UpdateRun) ToUpdateRunOutput() UpdateRunOutput {
	return i.ToUpdateRunOutputWithContext(context.Background())
}

func (i *UpdateRun) ToUpdateRunOutputWithContext(ctx context.Context) UpdateRunOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdateRunOutput)
}

type UpdateRunOutput struct{ *pulumi.OutputState }

func (UpdateRunOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpdateRun)(nil)).Elem()
}

func (o UpdateRunOutput) ToUpdateRunOutput() UpdateRunOutput {
	return o
}

func (o UpdateRunOutput) ToUpdateRunOutputWithContext(ctx context.Context) UpdateRunOutput {
	return o
}

func (o UpdateRunOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o UpdateRunOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.Duration }).(pulumi.StringPtrOutput)
}

func (o UpdateRunOutput) EndTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.EndTimeUtc }).(pulumi.StringPtrOutput)
}

func (o UpdateRunOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o UpdateRunOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

func (o UpdateRunOutput) LastUpdatedTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.LastUpdatedTimeUtc }).(pulumi.StringPtrOutput)
}

func (o UpdateRunOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o UpdateRunOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o UpdateRunOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o UpdateRunOutput) StartTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.StartTimeUtc }).(pulumi.StringPtrOutput)
}

func (o UpdateRunOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o UpdateRunOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o UpdateRunOutput) Steps() StepResponseArrayOutput {
	return o.ApplyT(func(v *UpdateRun) StepResponseArrayOutput { return v.Steps }).(StepResponseArrayOutput)
}

func (o UpdateRunOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *UpdateRun) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o UpdateRunOutput) TimeStarted() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringPtrOutput { return v.TimeStarted }).(pulumi.StringPtrOutput)
}

func (o UpdateRunOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *UpdateRun) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(UpdateRunOutput{})
}
