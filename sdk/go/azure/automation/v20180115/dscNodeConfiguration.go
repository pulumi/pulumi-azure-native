


package v20180115

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DscNodeConfiguration struct {
	pulumi.CustomResourceState

	Configuration                   DscConfigurationAssociationPropertyResponsePtrOutput `pulumi:"configuration"`
	CreationTime                    pulumi.StringPtrOutput                               `pulumi:"creationTime"`
	IncrementNodeConfigurationBuild pulumi.BoolPtrOutput                                 `pulumi:"incrementNodeConfigurationBuild"`
	LastModifiedTime                pulumi.StringPtrOutput                               `pulumi:"lastModifiedTime"`
	Name                            pulumi.StringOutput                                  `pulumi:"name"`
	NodeCount                       pulumi.Float64PtrOutput                              `pulumi:"nodeCount"`
	Source                          pulumi.StringPtrOutput                               `pulumi:"source"`
	Type                            pulumi.StringOutput                                  `pulumi:"type"`
}


func NewDscNodeConfiguration(ctx *pulumi.Context,
	name string, args *DscNodeConfigurationArgs, opts ...pulumi.ResourceOption) (*DscNodeConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:DscNodeConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:DscNodeConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:DscNodeConfiguration"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:DscNodeConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource DscNodeConfiguration
	err := ctx.RegisterResource("azure-native:automation/v20180115:DscNodeConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDscNodeConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DscNodeConfigurationState, opts ...pulumi.ResourceOption) (*DscNodeConfiguration, error) {
	var resource DscNodeConfiguration
	err := ctx.ReadResource("azure-native:automation/v20180115:DscNodeConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dscNodeConfigurationState struct {
}

type DscNodeConfigurationState struct {
}

func (DscNodeConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*dscNodeConfigurationState)(nil)).Elem()
}

type dscNodeConfigurationArgs struct {
	AutomationAccountName           string                              `pulumi:"automationAccountName"`
	Configuration                   DscConfigurationAssociationProperty `pulumi:"configuration"`
	IncrementNodeConfigurationBuild *bool                               `pulumi:"incrementNodeConfigurationBuild"`
	Name                            *string                             `pulumi:"name"`
	NodeConfigurationName           *string                             `pulumi:"nodeConfigurationName"`
	ResourceGroupName               string                              `pulumi:"resourceGroupName"`
	Source                          ContentSource                       `pulumi:"source"`
	Tags                            map[string]string                   `pulumi:"tags"`
}


type DscNodeConfigurationArgs struct {
	AutomationAccountName           pulumi.StringInput
	Configuration                   DscConfigurationAssociationPropertyInput
	IncrementNodeConfigurationBuild pulumi.BoolPtrInput
	Name                            pulumi.StringPtrInput
	NodeConfigurationName           pulumi.StringPtrInput
	ResourceGroupName               pulumi.StringInput
	Source                          ContentSourceInput
	Tags                            pulumi.StringMapInput
}

func (DscNodeConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dscNodeConfigurationArgs)(nil)).Elem()
}

type DscNodeConfigurationInput interface {
	pulumi.Input

	ToDscNodeConfigurationOutput() DscNodeConfigurationOutput
	ToDscNodeConfigurationOutputWithContext(ctx context.Context) DscNodeConfigurationOutput
}

func (*DscNodeConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*DscNodeConfiguration)(nil))
}

func (i *DscNodeConfiguration) ToDscNodeConfigurationOutput() DscNodeConfigurationOutput {
	return i.ToDscNodeConfigurationOutputWithContext(context.Background())
}

func (i *DscNodeConfiguration) ToDscNodeConfigurationOutputWithContext(ctx context.Context) DscNodeConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscNodeConfigurationOutput)
}

type DscNodeConfigurationOutput struct{ *pulumi.OutputState }

func (DscNodeConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DscNodeConfiguration)(nil))
}

func (o DscNodeConfigurationOutput) ToDscNodeConfigurationOutput() DscNodeConfigurationOutput {
	return o
}

func (o DscNodeConfigurationOutput) ToDscNodeConfigurationOutputWithContext(ctx context.Context) DscNodeConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DscNodeConfigurationOutput{})
}
