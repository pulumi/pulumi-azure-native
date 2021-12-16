


package logic

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IntegrationAccountMap struct {
	pulumi.CustomResourceState

	ChangedTime      pulumi.StringOutput                                              `pulumi:"changedTime"`
	Content          pulumi.StringPtrOutput                                           `pulumi:"content"`
	ContentLink      ContentLinkResponseOutput                                        `pulumi:"contentLink"`
	ContentType      pulumi.StringPtrOutput                                           `pulumi:"contentType"`
	CreatedTime      pulumi.StringOutput                                              `pulumi:"createdTime"`
	Location         pulumi.StringPtrOutput                                           `pulumi:"location"`
	MapType          pulumi.StringOutput                                              `pulumi:"mapType"`
	Metadata         pulumi.AnyOutput                                                 `pulumi:"metadata"`
	Name             pulumi.StringOutput                                              `pulumi:"name"`
	ParametersSchema IntegrationAccountMapPropertiesResponseParametersSchemaPtrOutput `pulumi:"parametersSchema"`
	Tags             pulumi.StringMapOutput                                           `pulumi:"tags"`
	Type             pulumi.StringOutput                                              `pulumi:"type"`
}


func NewIntegrationAccountMap(ctx *pulumi.Context,
	name string, args *IntegrationAccountMapArgs, opts ...pulumi.ResourceOption) (*IntegrationAccountMap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationAccountName'")
	}
	if args.MapType == nil {
		return nil, errors.New("invalid value for required argument 'MapType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic/v20150801preview:IntegrationAccountMap"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:IntegrationAccountMap"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:IntegrationAccountMap"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationAccountMap"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationAccountMap
	err := ctx.RegisterResource("azure-native:logic:IntegrationAccountMap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIntegrationAccountMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationAccountMapState, opts ...pulumi.ResourceOption) (*IntegrationAccountMap, error) {
	var resource IntegrationAccountMap
	err := ctx.ReadResource("azure-native:logic:IntegrationAccountMap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type integrationAccountMapState struct {
}

type IntegrationAccountMapState struct {
}

func (IntegrationAccountMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountMapState)(nil)).Elem()
}

type integrationAccountMapArgs struct {
	Content                *string                                          `pulumi:"content"`
	ContentType            *string                                          `pulumi:"contentType"`
	IntegrationAccountName string                                           `pulumi:"integrationAccountName"`
	Location               *string                                          `pulumi:"location"`
	MapName                *string                                          `pulumi:"mapName"`
	MapType                string                                           `pulumi:"mapType"`
	Metadata               interface{}                                      `pulumi:"metadata"`
	ParametersSchema       *IntegrationAccountMapPropertiesParametersSchema `pulumi:"parametersSchema"`
	ResourceGroupName      string                                           `pulumi:"resourceGroupName"`
	Tags                   map[string]string                                `pulumi:"tags"`
}


type IntegrationAccountMapArgs struct {
	Content                pulumi.StringPtrInput
	ContentType            pulumi.StringPtrInput
	IntegrationAccountName pulumi.StringInput
	Location               pulumi.StringPtrInput
	MapName                pulumi.StringPtrInput
	MapType                pulumi.StringInput
	Metadata               pulumi.Input
	ParametersSchema       IntegrationAccountMapPropertiesParametersSchemaPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (IntegrationAccountMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationAccountMapArgs)(nil)).Elem()
}

type IntegrationAccountMapInput interface {
	pulumi.Input

	ToIntegrationAccountMapOutput() IntegrationAccountMapOutput
	ToIntegrationAccountMapOutputWithContext(ctx context.Context) IntegrationAccountMapOutput
}

func (*IntegrationAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountMap)(nil)).Elem()
}

func (i *IntegrationAccountMap) ToIntegrationAccountMapOutput() IntegrationAccountMapOutput {
	return i.ToIntegrationAccountMapOutputWithContext(context.Background())
}

func (i *IntegrationAccountMap) ToIntegrationAccountMapOutputWithContext(ctx context.Context) IntegrationAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationAccountMapOutput)
}

type IntegrationAccountMapOutput struct{ *pulumi.OutputState }

func (IntegrationAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationAccountMap)(nil)).Elem()
}

func (o IntegrationAccountMapOutput) ToIntegrationAccountMapOutput() IntegrationAccountMapOutput {
	return o
}

func (o IntegrationAccountMapOutput) ToIntegrationAccountMapOutputWithContext(ctx context.Context) IntegrationAccountMapOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntegrationAccountMapOutput{})
}
