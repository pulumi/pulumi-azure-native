


package v20151031

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DscConfiguration struct {
	pulumi.CustomResourceState

	CreationTime           pulumi.StringPtrOutput                     `pulumi:"creationTime"`
	Description            pulumi.StringPtrOutput                     `pulumi:"description"`
	Etag                   pulumi.StringPtrOutput                     `pulumi:"etag"`
	JobCount               pulumi.IntPtrOutput                        `pulumi:"jobCount"`
	LastModifiedTime       pulumi.StringPtrOutput                     `pulumi:"lastModifiedTime"`
	Location               pulumi.StringPtrOutput                     `pulumi:"location"`
	LogVerbose             pulumi.BoolPtrOutput                       `pulumi:"logVerbose"`
	Name                   pulumi.StringOutput                        `pulumi:"name"`
	NodeConfigurationCount pulumi.IntPtrOutput                        `pulumi:"nodeConfigurationCount"`
	Parameters             DscConfigurationParameterResponseMapOutput `pulumi:"parameters"`
	ProvisioningState      pulumi.StringPtrOutput                     `pulumi:"provisioningState"`
	Source                 ContentSourceResponsePtrOutput             `pulumi:"source"`
	State                  pulumi.StringPtrOutput                     `pulumi:"state"`
	Tags                   pulumi.StringMapOutput                     `pulumi:"tags"`
	Type                   pulumi.StringOutput                        `pulumi:"type"`
}


func NewDscConfiguration(ctx *pulumi.Context,
	name string, args *DscConfigurationArgs, opts ...pulumi.ResourceOption) (*DscConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:DscConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:DscConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20220808:DscConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource DscConfiguration
	err := ctx.RegisterResource("azure-native:automation/v20151031:DscConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDscConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DscConfigurationState, opts ...pulumi.ResourceOption) (*DscConfiguration, error) {
	var resource DscConfiguration
	err := ctx.ReadResource("azure-native:automation/v20151031:DscConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dscConfigurationState struct {
}

type DscConfigurationState struct {
}

func (DscConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dscConfigurationState)(nil)).Elem()
}

type dscConfigurationArgs struct {
	AutomationAccountName string                               `pulumi:"automationAccountName"`
	ConfigurationName     *string                              `pulumi:"configurationName"`
	Description           *string                              `pulumi:"description"`
	Location              *string                              `pulumi:"location"`
	LogProgress           *bool                                `pulumi:"logProgress"`
	LogVerbose            *bool                                `pulumi:"logVerbose"`
	Name                  *string                              `pulumi:"name"`
	Parameters            map[string]DscConfigurationParameter `pulumi:"parameters"`
	ResourceGroupName     string                               `pulumi:"resourceGroupName"`
	Source                ContentSource                        `pulumi:"source"`
	Tags                  map[string]string                    `pulumi:"tags"`
}


type DscConfigurationArgs struct {
	AutomationAccountName pulumi.StringInput
	ConfigurationName     pulumi.StringPtrInput
	Description           pulumi.StringPtrInput
	Location              pulumi.StringPtrInput
	LogProgress           pulumi.BoolPtrInput
	LogVerbose            pulumi.BoolPtrInput
	Name                  pulumi.StringPtrInput
	Parameters            DscConfigurationParameterMapInput
	ResourceGroupName     pulumi.StringInput
	Source                ContentSourceInput
	Tags                  pulumi.StringMapInput
}

func (DscConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dscConfigurationArgs)(nil)).Elem()
}

type DscConfigurationInput interface {
	pulumi.Input

	ToDscConfigurationOutput() DscConfigurationOutput
	ToDscConfigurationOutputWithContext(ctx context.Context) DscConfigurationOutput
}

func (*DscConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**DscConfiguration)(nil)).Elem()
}

func (i *DscConfiguration) ToDscConfigurationOutput() DscConfigurationOutput {
	return i.ToDscConfigurationOutputWithContext(context.Background())
}

func (i *DscConfiguration) ToDscConfigurationOutputWithContext(ctx context.Context) DscConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationOutput)
}

type DscConfigurationOutput struct{ *pulumi.OutputState }

func (DscConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DscConfiguration)(nil)).Elem()
}

func (o DscConfigurationOutput) ToDscConfigurationOutput() DscConfigurationOutput {
	return o
}

func (o DscConfigurationOutput) ToDscConfigurationOutputWithContext(ctx context.Context) DscConfigurationOutput {
	return o
}

func (o DscConfigurationOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscConfiguration) pulumi.StringPtrOutput { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o DscConfigurationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscConfiguration) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DscConfigurationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscConfiguration) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o DscConfigurationOutput) JobCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DscConfiguration) pulumi.IntPtrOutput { return v.JobCount }).(pulumi.IntPtrOutput)
}

func (o DscConfigurationOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscConfiguration) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o DscConfigurationOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscConfiguration) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o DscConfigurationOutput) LogVerbose() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DscConfiguration) pulumi.BoolPtrOutput { return v.LogVerbose }).(pulumi.BoolPtrOutput)
}

func (o DscConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DscConfiguration) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DscConfigurationOutput) NodeConfigurationCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DscConfiguration) pulumi.IntPtrOutput { return v.NodeConfigurationCount }).(pulumi.IntPtrOutput)
}

func (o DscConfigurationOutput) Parameters() DscConfigurationParameterResponseMapOutput {
	return o.ApplyT(func(v *DscConfiguration) DscConfigurationParameterResponseMapOutput { return v.Parameters }).(DscConfigurationParameterResponseMapOutput)
}

func (o DscConfigurationOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscConfiguration) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o DscConfigurationOutput) Source() ContentSourceResponsePtrOutput {
	return o.ApplyT(func(v *DscConfiguration) ContentSourceResponsePtrOutput { return v.Source }).(ContentSourceResponsePtrOutput)
}

func (o DscConfigurationOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscConfiguration) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o DscConfigurationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DscConfiguration) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DscConfigurationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DscConfiguration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DscConfigurationOutput{})
}
