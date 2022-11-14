


package automation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SoftwareUpdateConfigurationByName struct {
	pulumi.CustomResourceState

	CreatedBy           pulumi.StringOutput                               `pulumi:"createdBy"`
	CreationTime        pulumi.StringOutput                               `pulumi:"creationTime"`
	Error               ErrorResponseResponsePtrOutput                    `pulumi:"error"`
	LastModifiedBy      pulumi.StringOutput                               `pulumi:"lastModifiedBy"`
	LastModifiedTime    pulumi.StringOutput                               `pulumi:"lastModifiedTime"`
	Name                pulumi.StringOutput                               `pulumi:"name"`
	ProvisioningState   pulumi.StringOutput                               `pulumi:"provisioningState"`
	ScheduleInfo        SUCSchedulePropertiesResponseOutput               `pulumi:"scheduleInfo"`
	Tasks               SoftwareUpdateConfigurationTasksResponsePtrOutput `pulumi:"tasks"`
	Type                pulumi.StringOutput                               `pulumi:"type"`
	UpdateConfiguration UpdateConfigurationResponseOutput                 `pulumi:"updateConfiguration"`
}


func NewSoftwareUpdateConfigurationByName(ctx *pulumi.Context,
	name string, args *SoftwareUpdateConfigurationByNameArgs, opts ...pulumi.ResourceOption) (*SoftwareUpdateConfigurationByName, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScheduleInfo == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleInfo'")
	}
	if args.UpdateConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'UpdateConfiguration'")
	}
	args.ScheduleInfo = args.ScheduleInfo.ToSUCSchedulePropertiesOutput().ApplyT(func(v SUCScheduleProperties) SUCScheduleProperties { return *v.Defaults() }).(SUCSchedulePropertiesOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation/v20170515preview:SoftwareUpdateConfigurationByName"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:SoftwareUpdateConfigurationByName"),
		},
	})
	opts = append(opts, aliases)
	var resource SoftwareUpdateConfigurationByName
	err := ctx.RegisterResource("azure-native:automation:SoftwareUpdateConfigurationByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSoftwareUpdateConfigurationByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SoftwareUpdateConfigurationByNameState, opts ...pulumi.ResourceOption) (*SoftwareUpdateConfigurationByName, error) {
	var resource SoftwareUpdateConfigurationByName
	err := ctx.ReadResource("azure-native:automation:SoftwareUpdateConfigurationByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type softwareUpdateConfigurationByNameState struct {
}

type SoftwareUpdateConfigurationByNameState struct {
}

func (SoftwareUpdateConfigurationByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*softwareUpdateConfigurationByNameState)(nil)).Elem()
}

type softwareUpdateConfigurationByNameArgs struct {
	AutomationAccountName           string                            `pulumi:"automationAccountName"`
	Error                           *ErrorResponse                    `pulumi:"error"`
	ResourceGroupName               string                            `pulumi:"resourceGroupName"`
	ScheduleInfo                    SUCScheduleProperties             `pulumi:"scheduleInfo"`
	SoftwareUpdateConfigurationName *string                           `pulumi:"softwareUpdateConfigurationName"`
	Tasks                           *SoftwareUpdateConfigurationTasks `pulumi:"tasks"`
	UpdateConfiguration             UpdateConfiguration               `pulumi:"updateConfiguration"`
}


type SoftwareUpdateConfigurationByNameArgs struct {
	AutomationAccountName           pulumi.StringInput
	Error                           ErrorResponsePtrInput
	ResourceGroupName               pulumi.StringInput
	ScheduleInfo                    SUCSchedulePropertiesInput
	SoftwareUpdateConfigurationName pulumi.StringPtrInput
	Tasks                           SoftwareUpdateConfigurationTasksPtrInput
	UpdateConfiguration             UpdateConfigurationInput
}

func (SoftwareUpdateConfigurationByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*softwareUpdateConfigurationByNameArgs)(nil)).Elem()
}

type SoftwareUpdateConfigurationByNameInput interface {
	pulumi.Input

	ToSoftwareUpdateConfigurationByNameOutput() SoftwareUpdateConfigurationByNameOutput
	ToSoftwareUpdateConfigurationByNameOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationByNameOutput
}

func (*SoftwareUpdateConfigurationByName) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftwareUpdateConfigurationByName)(nil)).Elem()
}

func (i *SoftwareUpdateConfigurationByName) ToSoftwareUpdateConfigurationByNameOutput() SoftwareUpdateConfigurationByNameOutput {
	return i.ToSoftwareUpdateConfigurationByNameOutputWithContext(context.Background())
}

func (i *SoftwareUpdateConfigurationByName) ToSoftwareUpdateConfigurationByNameOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftwareUpdateConfigurationByNameOutput)
}

type SoftwareUpdateConfigurationByNameOutput struct{ *pulumi.OutputState }

func (SoftwareUpdateConfigurationByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftwareUpdateConfigurationByName)(nil)).Elem()
}

func (o SoftwareUpdateConfigurationByNameOutput) ToSoftwareUpdateConfigurationByNameOutput() SoftwareUpdateConfigurationByNameOutput {
	return o
}

func (o SoftwareUpdateConfigurationByNameOutput) ToSoftwareUpdateConfigurationByNameOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationByNameOutput {
	return o
}

func (o SoftwareUpdateConfigurationByNameOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) pulumi.StringOutput { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o SoftwareUpdateConfigurationByNameOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o SoftwareUpdateConfigurationByNameOutput) Error() ErrorResponseResponsePtrOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) ErrorResponseResponsePtrOutput { return v.Error }).(ErrorResponseResponsePtrOutput)
}

func (o SoftwareUpdateConfigurationByNameOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) pulumi.StringOutput { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o SoftwareUpdateConfigurationByNameOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o SoftwareUpdateConfigurationByNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SoftwareUpdateConfigurationByNameOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SoftwareUpdateConfigurationByNameOutput) ScheduleInfo() SUCSchedulePropertiesResponseOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) SUCSchedulePropertiesResponseOutput { return v.ScheduleInfo }).(SUCSchedulePropertiesResponseOutput)
}

func (o SoftwareUpdateConfigurationByNameOutput) Tasks() SoftwareUpdateConfigurationTasksResponsePtrOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) SoftwareUpdateConfigurationTasksResponsePtrOutput {
		return v.Tasks
	}).(SoftwareUpdateConfigurationTasksResponsePtrOutput)
}

func (o SoftwareUpdateConfigurationByNameOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o SoftwareUpdateConfigurationByNameOutput) UpdateConfiguration() UpdateConfigurationResponseOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationByName) UpdateConfigurationResponseOutput {
		return v.UpdateConfiguration
	}).(UpdateConfigurationResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(SoftwareUpdateConfigurationByNameOutput{})
}
