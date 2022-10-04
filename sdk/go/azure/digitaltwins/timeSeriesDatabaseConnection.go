


package digitaltwins

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TimeSeriesDatabaseConnection struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                                 `pulumi:"name"`
	Properties AzureDataExplorerConnectionPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                            `pulumi:"systemData"`
	Type       pulumi.StringOutput                                 `pulumi:"type"`
}


func NewTimeSeriesDatabaseConnection(ctx *pulumi.Context,
	name string, args *TimeSeriesDatabaseConnectionArgs, opts ...pulumi.ResourceOption) (*TimeSeriesDatabaseConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToAzureDataExplorerConnectionPropertiesPtrOutput().ApplyT(func(v *AzureDataExplorerConnectionProperties) *AzureDataExplorerConnectionProperties {
			return v.Defaults()
		}).(AzureDataExplorerConnectionPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:digitaltwins/v20210630preview:TimeSeriesDatabaseConnection"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20220531:TimeSeriesDatabaseConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource TimeSeriesDatabaseConnection
	err := ctx.RegisterResource("azure-native:digitaltwins:TimeSeriesDatabaseConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTimeSeriesDatabaseConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TimeSeriesDatabaseConnectionState, opts ...pulumi.ResourceOption) (*TimeSeriesDatabaseConnection, error) {
	var resource TimeSeriesDatabaseConnection
	err := ctx.ReadResource("azure-native:digitaltwins:TimeSeriesDatabaseConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type timeSeriesDatabaseConnectionState struct {
}

type TimeSeriesDatabaseConnectionState struct {
}

func (TimeSeriesDatabaseConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*timeSeriesDatabaseConnectionState)(nil)).Elem()
}

type timeSeriesDatabaseConnectionArgs struct {
	Properties                       *AzureDataExplorerConnectionProperties `pulumi:"properties"`
	ResourceGroupName                string                                 `pulumi:"resourceGroupName"`
	ResourceName                     string                                 `pulumi:"resourceName"`
	TimeSeriesDatabaseConnectionName *string                                `pulumi:"timeSeriesDatabaseConnectionName"`
}


type TimeSeriesDatabaseConnectionArgs struct {
	Properties                       AzureDataExplorerConnectionPropertiesPtrInput
	ResourceGroupName                pulumi.StringInput
	ResourceName                     pulumi.StringInput
	TimeSeriesDatabaseConnectionName pulumi.StringPtrInput
}

func (TimeSeriesDatabaseConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*timeSeriesDatabaseConnectionArgs)(nil)).Elem()
}

type TimeSeriesDatabaseConnectionInput interface {
	pulumi.Input

	ToTimeSeriesDatabaseConnectionOutput() TimeSeriesDatabaseConnectionOutput
	ToTimeSeriesDatabaseConnectionOutputWithContext(ctx context.Context) TimeSeriesDatabaseConnectionOutput
}

func (*TimeSeriesDatabaseConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeSeriesDatabaseConnection)(nil)).Elem()
}

func (i *TimeSeriesDatabaseConnection) ToTimeSeriesDatabaseConnectionOutput() TimeSeriesDatabaseConnectionOutput {
	return i.ToTimeSeriesDatabaseConnectionOutputWithContext(context.Background())
}

func (i *TimeSeriesDatabaseConnection) ToTimeSeriesDatabaseConnectionOutputWithContext(ctx context.Context) TimeSeriesDatabaseConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeSeriesDatabaseConnectionOutput)
}

type TimeSeriesDatabaseConnectionOutput struct{ *pulumi.OutputState }

func (TimeSeriesDatabaseConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeSeriesDatabaseConnection)(nil)).Elem()
}

func (o TimeSeriesDatabaseConnectionOutput) ToTimeSeriesDatabaseConnectionOutput() TimeSeriesDatabaseConnectionOutput {
	return o
}

func (o TimeSeriesDatabaseConnectionOutput) ToTimeSeriesDatabaseConnectionOutputWithContext(ctx context.Context) TimeSeriesDatabaseConnectionOutput {
	return o
}

func (o TimeSeriesDatabaseConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TimeSeriesDatabaseConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TimeSeriesDatabaseConnectionOutput) Properties() AzureDataExplorerConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v *TimeSeriesDatabaseConnection) AzureDataExplorerConnectionPropertiesResponseOutput {
		return v.Properties
	}).(AzureDataExplorerConnectionPropertiesResponseOutput)
}

func (o TimeSeriesDatabaseConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TimeSeriesDatabaseConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o TimeSeriesDatabaseConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TimeSeriesDatabaseConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TimeSeriesDatabaseConnectionOutput{})
}
