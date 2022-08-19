


package v20200601preview

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
			Type: pulumi.String("azure-native:apimanagement/v20201201:ProductApi"),
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
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ProductApi"),
		},
	})
	opts = append(opts, aliases)
	var resource ProductApi
	err := ctx.RegisterResource("azure-native:apimanagement/v20200601preview:ProductApi", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProductApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductApiState, opts ...pulumi.ResourceOption) (*ProductApi, error) {
	var resource ProductApi
	err := ctx.ReadResource("azure-native:apimanagement/v20200601preview:ProductApi", name, id, state, &resource, opts...)
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

func (o ProductApiOutput) ApiRevision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.ApiRevision }).(pulumi.StringPtrOutput)
}

func (o ProductApiOutput) ApiRevisionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.ApiRevisionDescription }).(pulumi.StringPtrOutput)
}

func (o ProductApiOutput) ApiType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.ApiType }).(pulumi.StringPtrOutput)
}

func (o ProductApiOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o ProductApiOutput) ApiVersionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.ApiVersionDescription }).(pulumi.StringPtrOutput)
}

func (o ProductApiOutput) ApiVersionSet() ApiVersionSetContractDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ProductApi) ApiVersionSetContractDetailsResponsePtrOutput { return v.ApiVersionSet }).(ApiVersionSetContractDetailsResponsePtrOutput)
}

func (o ProductApiOutput) ApiVersionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.ApiVersionSetId }).(pulumi.StringPtrOutput)
}

func (o ProductApiOutput) AuthenticationSettings() AuthenticationSettingsContractResponsePtrOutput {
	return o.ApplyT(func(v *ProductApi) AuthenticationSettingsContractResponsePtrOutput { return v.AuthenticationSettings }).(AuthenticationSettingsContractResponsePtrOutput)
}

func (o ProductApiOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ProductApiOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ProductApiOutput) IsCurrent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.BoolPtrOutput { return v.IsCurrent }).(pulumi.BoolPtrOutput)
}

func (o ProductApiOutput) IsOnline() pulumi.BoolOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.BoolOutput { return v.IsOnline }).(pulumi.BoolOutput)
}

func (o ProductApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProductApiOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

func (o ProductApiOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringArrayOutput { return v.Protocols }).(pulumi.StringArrayOutput)
}

func (o ProductApiOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

func (o ProductApiOutput) SourceApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringPtrOutput { return v.SourceApiId }).(pulumi.StringPtrOutput)
}

func (o ProductApiOutput) SubscriptionKeyParameterNames() SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return o.ApplyT(func(v *ProductApi) SubscriptionKeyParameterNamesContractResponsePtrOutput {
		return v.SubscriptionKeyParameterNames
	}).(SubscriptionKeyParameterNamesContractResponsePtrOutput)
}

func (o ProductApiOutput) SubscriptionRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.BoolPtrOutput { return v.SubscriptionRequired }).(pulumi.BoolPtrOutput)
}

func (o ProductApiOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProductApi) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProductApiOutput{})
}
