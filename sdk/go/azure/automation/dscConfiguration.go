


package automation

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
			Type: pulumi.String("azure-native:automation/v20151031:DscConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:DscConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource DscConfiguration
	err := ctx.RegisterResource("azure-native:automation:DscConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDscConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DscConfigurationState, opts ...pulumi.ResourceOption) (*DscConfiguration, error) {
	var resource DscConfiguration
	err := ctx.ReadResource("azure-native:automation:DscConfiguration", name, id, state, &resource, opts...)
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

func init() {
	pulumi.RegisterOutputType(DscConfigurationOutput{})
}
