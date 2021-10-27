


package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GatewayApiEntityTag struct {
	pulumi.CustomResourceState

	ApiRevision                   pulumi.StringPtrOutput                                 `pulumi:"apiRevision"`
	ApiRevisionDescription        pulumi.StringPtrOutput                                 `pulumi:"apiRevisionDescription"`
	ApiType                       pulumi.StringPtrOutput                                 `pulumi:"apiType"`
	ApiVersion                    pulumi.StringPtrOutput                                 `pulumi:"apiVersion"`
	ApiVersionDescription         pulumi.StringPtrOutput                                 `pulumi:"apiVersionDescription"`
	ApiVersionSet                 ApiVersionSetContractDetailsResponsePtrOutput          `pulumi:"apiVersionSet"`
	ApiVersionSetId               pulumi.StringPtrOutput                                 `pulumi:"apiVersionSetId"`
	AuthenticationSettings        AuthenticationSettingsContractResponsePtrOutput        `pulumi:"authenticationSettings"`
	Description                   pulumi.StringPtrOutput                                 `pulumi:"description"`
	DisplayName                   pulumi.StringPtrOutput                                 `pulumi:"displayName"`
	IsCurrent                     pulumi.BoolPtrOutput                                   `pulumi:"isCurrent"`
	IsOnline                      pulumi.BoolOutput                                      `pulumi:"isOnline"`
	Name                          pulumi.StringOutput                                    `pulumi:"name"`
	Path                          pulumi.StringOutput                                    `pulumi:"path"`
	Protocols                     pulumi.StringArrayOutput                               `pulumi:"protocols"`
	ServiceUrl                    pulumi.StringPtrOutput                                 `pulumi:"serviceUrl"`
	SourceApiId                   pulumi.StringPtrOutput                                 `pulumi:"sourceApiId"`
	SubscriptionKeyParameterNames SubscriptionKeyParameterNamesContractResponsePtrOutput `pulumi:"subscriptionKeyParameterNames"`
	SubscriptionRequired          pulumi.BoolPtrOutput                                   `pulumi:"subscriptionRequired"`
	Type                          pulumi.StringOutput                                    `pulumi:"type"`
}


func NewGatewayApiEntityTag(ctx *pulumi.Context,
	name string, args *GatewayApiEntityTagArgs, opts ...pulumi.ResourceOption) (*GatewayApiEntityTag, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210801:GatewayApiEntityTag"),
		},
	})
	opts = append(opts, aliases)
	var resource GatewayApiEntityTag
	err := ctx.RegisterResource("azure-native:apimanagement:GatewayApiEntityTag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGatewayApiEntityTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayApiEntityTagState, opts ...pulumi.ResourceOption) (*GatewayApiEntityTag, error) {
	var resource GatewayApiEntityTag
	err := ctx.ReadResource("azure-native:apimanagement:GatewayApiEntityTag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type gatewayApiEntityTagState struct {
}

type GatewayApiEntityTagState struct {
}

func (GatewayApiEntityTagState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayApiEntityTagState)(nil)).Elem()
}

type gatewayApiEntityTagArgs struct {
	ApiId             *string            `pulumi:"apiId"`
	GatewayId         string             `pulumi:"gatewayId"`
	ProvisioningState *ProvisioningState `pulumi:"provisioningState"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	ServiceName       string             `pulumi:"serviceName"`
}


type GatewayApiEntityTagArgs struct {
	ApiId             pulumi.StringPtrInput
	GatewayId         pulumi.StringInput
	ProvisioningState ProvisioningStatePtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (GatewayApiEntityTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayApiEntityTagArgs)(nil)).Elem()
}

type GatewayApiEntityTagInput interface {
	pulumi.Input

	ToGatewayApiEntityTagOutput() GatewayApiEntityTagOutput
	ToGatewayApiEntityTagOutputWithContext(ctx context.Context) GatewayApiEntityTagOutput
}

func (*GatewayApiEntityTag) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayApiEntityTag)(nil))
}

func (i *GatewayApiEntityTag) ToGatewayApiEntityTagOutput() GatewayApiEntityTagOutput {
	return i.ToGatewayApiEntityTagOutputWithContext(context.Background())
}

func (i *GatewayApiEntityTag) ToGatewayApiEntityTagOutputWithContext(ctx context.Context) GatewayApiEntityTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayApiEntityTagOutput)
}

type GatewayApiEntityTagOutput struct{ *pulumi.OutputState }

func (GatewayApiEntityTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayApiEntityTag)(nil))
}

func (o GatewayApiEntityTagOutput) ToGatewayApiEntityTagOutput() GatewayApiEntityTagOutput {
	return o
}

func (o GatewayApiEntityTagOutput) ToGatewayApiEntityTagOutputWithContext(ctx context.Context) GatewayApiEntityTagOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayApiEntityTagInput)(nil)).Elem(), &GatewayApiEntityTag{})
	pulumi.RegisterOutputType(GatewayApiEntityTagOutput{})
}
