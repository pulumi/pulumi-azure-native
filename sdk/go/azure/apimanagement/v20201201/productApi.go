


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProductApi struct {
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


func NewProductApi(ctx *pulumi.Context,
	name string, args *ProductApiArgs, opts ...pulumi.ResourceOption) (*ProductApi, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ProductApi"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ProductApi"),
		},
	})
	opts = append(opts, aliases)
	var resource ProductApi
	err := ctx.RegisterResource("azure-native:apimanagement/v20201201:ProductApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProductApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductApiState, opts ...pulumi.ResourceOption) (*ProductApi, error) {
	var resource ProductApi
	err := ctx.ReadResource("azure-native:apimanagement/v20201201:ProductApi", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type productApiState struct {
}

type ProductApiState struct {
}

func (ProductApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*productApiState)(nil)).Elem()
}

type productApiArgs struct {
	ApiId             *string `pulumi:"apiId"`
	ProductId         string  `pulumi:"productId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type ProductApiArgs struct {
	ApiId             pulumi.StringPtrInput
	ProductId         pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (ProductApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productApiArgs)(nil)).Elem()
}

type ProductApiInput interface {
	pulumi.Input

	ToProductApiOutput() ProductApiOutput
	ToProductApiOutputWithContext(ctx context.Context) ProductApiOutput
}

func (*ProductApi) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductApi)(nil)).Elem()
}

func (i *ProductApi) ToProductApiOutput() ProductApiOutput {
	return i.ToProductApiOutputWithContext(context.Background())
}

func (i *ProductApi) ToProductApiOutputWithContext(ctx context.Context) ProductApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductApiOutput)
}

type ProductApiOutput struct{ *pulumi.OutputState }

func (ProductApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductApi)(nil)).Elem()
}

func (o ProductApiOutput) ToProductApiOutput() ProductApiOutput {
	return o
}

func (o ProductApiOutput) ToProductApiOutputWithContext(ctx context.Context) ProductApiOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProductApiOutput{})
}
