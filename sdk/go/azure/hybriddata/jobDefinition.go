


package hybriddata

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type JobDefinition struct {
	pulumi.CustomResourceState

	CustomerSecrets  CustomerSecretResponseArrayOutput `pulumi:"customerSecrets"`
	DataServiceInput pulumi.AnyOutput                  `pulumi:"dataServiceInput"`
	DataSinkId       pulumi.StringOutput               `pulumi:"dataSinkId"`
	DataSourceId     pulumi.StringOutput               `pulumi:"dataSourceId"`
	LastModifiedTime pulumi.StringPtrOutput            `pulumi:"lastModifiedTime"`
	Name             pulumi.StringOutput               `pulumi:"name"`
	RunLocation      pulumi.StringPtrOutput            `pulumi:"runLocation"`
	Schedules        ScheduleResponseArrayOutput       `pulumi:"schedules"`
	State            pulumi.StringOutput               `pulumi:"state"`
	Type             pulumi.StringOutput               `pulumi:"type"`
	UserConfirmation pulumi.StringPtrOutput            `pulumi:"userConfirmation"`
}


func NewJobDefinition(ctx *pulumi.Context,
	name string, args *JobDefinitionArgs, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataManagerName == nil {
		return nil, errors.New("invalid value for required argument 'DataManagerName'")
	}
	if args.DataServiceName == nil {
		return nil, errors.New("invalid value for required argument 'DataServiceName'")
	}
	if args.DataSinkId == nil {
		return nil, errors.New("invalid value for required argument 'DataSinkId'")
	}
	if args.DataSourceId == nil {
		return nil, errors.New("invalid value for required argument 'DataSourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.State == nil {
		return nil, errors.New("invalid value for required argument 'State'")
	}
	if isZero(args.UserConfirmation) {
		args.UserConfirmation = UserConfirmation("NotRequired")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybriddata/v20160601:JobDefinition"),
		},
		{
			Type: pulumi.String("azure-native:hybriddata/v20190601:JobDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource JobDefinition
	err := ctx.RegisterResource("azure-native:hybriddata:JobDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJobDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobDefinitionState, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	var resource JobDefinition
	err := ctx.ReadResource("azure-native:hybriddata:JobDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type jobDefinitionState struct {
}

type JobDefinitionState struct {
}

func (JobDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionState)(nil)).Elem()
}

type jobDefinitionArgs struct {
	CustomerSecrets   []CustomerSecret  `pulumi:"customerSecrets"`
	DataManagerName   string            `pulumi:"dataManagerName"`
	DataServiceInput  interface{}       `pulumi:"dataServiceInput"`
	DataServiceName   string            `pulumi:"dataServiceName"`
	DataSinkId        string            `pulumi:"dataSinkId"`
	DataSourceId      string            `pulumi:"dataSourceId"`
	JobDefinitionName *string           `pulumi:"jobDefinitionName"`
	LastModifiedTime  *string           `pulumi:"lastModifiedTime"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	RunLocation       *RunLocation      `pulumi:"runLocation"`
	Schedules         []Schedule        `pulumi:"schedules"`
	State             State             `pulumi:"state"`
	UserConfirmation  *UserConfirmation `pulumi:"userConfirmation"`
}


type JobDefinitionArgs struct {
	CustomerSecrets   CustomerSecretArrayInput
	DataManagerName   pulumi.StringInput
	DataServiceInput  pulumi.Input
	DataServiceName   pulumi.StringInput
	DataSinkId        pulumi.StringInput
	DataSourceId      pulumi.StringInput
	JobDefinitionName pulumi.StringPtrInput
	LastModifiedTime  pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	RunLocation       RunLocationPtrInput
	Schedules         ScheduleArrayInput
	State             StateInput
	UserConfirmation  UserConfirmationPtrInput
}

func (JobDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionArgs)(nil)).Elem()
}

type JobDefinitionInput interface {
	pulumi.Input

	ToJobDefinitionOutput() JobDefinitionOutput
	ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput
}

func (*JobDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinition)(nil)).Elem()
}

func (i *JobDefinition) ToJobDefinitionOutput() JobDefinitionOutput {
	return i.ToJobDefinitionOutputWithContext(context.Background())
}

func (i *JobDefinition) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionOutput)
}

type JobDefinitionOutput struct{ *pulumi.OutputState }

func (JobDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinition)(nil)).Elem()
}

func (o JobDefinitionOutput) ToJobDefinitionOutput() JobDefinitionOutput {
	return o
}

func (o JobDefinitionOutput) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return o
}

func (o JobDefinitionOutput) CustomerSecrets() CustomerSecretResponseArrayOutput {
	return o.ApplyT(func(v *JobDefinition) CustomerSecretResponseArrayOutput { return v.CustomerSecrets }).(CustomerSecretResponseArrayOutput)
}

func (o JobDefinitionOutput) DataServiceInput() pulumi.AnyOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.AnyOutput { return v.DataServiceInput }).(pulumi.AnyOutput)
}

func (o JobDefinitionOutput) DataSinkId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.DataSinkId }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) DataSourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.DataSourceId }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o JobDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) RunLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.RunLocation }).(pulumi.StringPtrOutput)
}

func (o JobDefinitionOutput) Schedules() ScheduleResponseArrayOutput {
	return o.ApplyT(func(v *JobDefinition) ScheduleResponseArrayOutput { return v.Schedules }).(ScheduleResponseArrayOutput)
}

func (o JobDefinitionOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) UserConfirmation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.UserConfirmation }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(JobDefinitionOutput{})
}
