


package v20220401preview

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
	Contact                       ApiContactInformationResponsePtrOutput                 `pulumi:"contact"`
	Description                   pulumi.StringPtrOutput                                 `pulumi:"description"`
	DisplayName                   pulumi.StringPtrOutput                                 `pulumi:"displayName"`
	IsCurrent                     pulumi.BoolPtrOutput                                   `pulumi:"isCurrent"`
	IsOnline                      pulumi.BoolOutput                                      `pulumi:"isOnline"`
	License                       ApiLicenseInformationResponsePtrOutput                 `pulumi:"license"`
	Name                          pulumi.StringOutput                                    `pulumi:"name"`
	Path                          pulumi.StringOutput                                    `pulumi:"path"`
	Protocols                     pulumi.StringArrayOutput                               `pulumi:"protocols"`
	ServiceUrl                    pulumi.StringPtrOutput                                 `pulumi:"serviceUrl"`
	SourceApiId                   pulumi.StringPtrOutput                                 `pulumi:"sourceApiId"`
	SubscriptionKeyParameterNames SubscriptionKeyParameterNamesContractResponsePtrOutput `pulumi:"subscriptionKeyParameterNames"`
	SubscriptionRequired          pulumi.BoolPtrOutput                                   `pulumi:"subscriptionRequired"`
	TermsOfServiceUrl             pulumi.StringPtrOutput                                 `pulumi:"termsOfServiceUrl"`
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
			Type: pulumi.String("azure-native:apimanagement:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:GatewayApiEntityTag"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:GatewayApiEntityTag"),
		},
	})
	opts = append(opts, aliases)
	var resource GatewayApiEntityTag
	err := ctx.RegisterResource("azure-native:apimanagement/v20220401preview:GatewayApiEntityTag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGatewayApiEntityTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayApiEntityTagState, opts ...pulumi.ResourceOption) (*GatewayApiEntityTag, error) {
	var resource GatewayApiEntityTag
	err := ctx.ReadResource("azure-native:apimanagement/v20220401preview:GatewayApiEntityTag", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((**GatewayApiEntityTag)(nil)).Elem()
}

func (i *GatewayApiEntityTag) ToGatewayApiEntityTagOutput() GatewayApiEntityTagOutput {
	return i.ToGatewayApiEntityTagOutputWithContext(context.Background())
}

func (i *GatewayApiEntityTag) ToGatewayApiEntityTagOutputWithContext(ctx context.Context) GatewayApiEntityTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayApiEntityTagOutput)
}

type GatewayApiEntityTagOutput struct{ *pulumi.OutputState }

func (GatewayApiEntityTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayApiEntityTag)(nil)).Elem()
}

func (o GatewayApiEntityTagOutput) ToGatewayApiEntityTagOutput() GatewayApiEntityTagOutput {
	return o
}

func (o GatewayApiEntityTagOutput) ToGatewayApiEntityTagOutputWithContext(ctx context.Context) GatewayApiEntityTagOutput {
	return o
}

func (o GatewayApiEntityTagOutput) ApiRevision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.StringPtrOutput { return v.ApiRevision }).(pulumi.StringPtrOutput)
}

func (o GatewayApiEntityTagOutput) ApiRevisionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.StringPtrOutput { return v.ApiRevisionDescription }).(pulumi.StringPtrOutput)
}

func (o GatewayApiEntityTagOutput) ApiType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.StringPtrOutput { return v.ApiType }).(pulumi.StringPtrOutput)
}

func (o GatewayApiEntityTagOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o GatewayApiEntityTagOutput) ApiVersionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.StringPtrOutput { return v.ApiVersionDescription }).(pulumi.StringPtrOutput)
}

func (o GatewayApiEntityTagOutput) ApiVersionSet() ApiVersionSetContractDetailsResponsePtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) ApiVersionSetContractDetailsResponsePtrOutput { return v.ApiVersionSet }).(ApiVersionSetContractDetailsResponsePtrOutput)
}

func (o GatewayApiEntityTagOutput) ApiVersionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.StringPtrOutput { return v.ApiVersionSetId }).(pulumi.StringPtrOutput)
}

func (o GatewayApiEntityTagOutput) AuthenticationSettings() AuthenticationSettingsContractResponsePtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) AuthenticationSettingsContractResponsePtrOutput {
		return v.AuthenticationSettings
	}).(AuthenticationSettingsContractResponsePtrOutput)
}

func (o GatewayApiEntityTagOutput) Contact() ApiContactInformationResponsePtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) ApiContactInformationResponsePtrOutput { return v.Contact }).(ApiContactInformationResponsePtrOutput)
}

func (o GatewayApiEntityTagOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GatewayApiEntityTagOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o GatewayApiEntityTagOutput) IsCurrent() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.BoolPtrOutput { return v.IsCurrent }).(pulumi.BoolPtrOutput)
}

func (o GatewayApiEntityTagOutput) IsOnline() pulumi.BoolOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.BoolOutput { return v.IsOnline }).(pulumi.BoolOutput)
}

func (o GatewayApiEntityTagOutput) License() ApiLicenseInformationResponsePtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) ApiLicenseInformationResponsePtrOutput { return v.License }).(ApiLicenseInformationResponsePtrOutput)
}

func (o GatewayApiEntityTagOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GatewayApiEntityTagOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

func (o GatewayApiEntityTagOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.StringArrayOutput { return v.Protocols }).(pulumi.StringArrayOutput)
}

func (o GatewayApiEntityTagOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.StringPtrOutput { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

func (o GatewayApiEntityTagOutput) SourceApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.StringPtrOutput { return v.SourceApiId }).(pulumi.StringPtrOutput)
}

func (o GatewayApiEntityTagOutput) SubscriptionKeyParameterNames() SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) SubscriptionKeyParameterNamesContractResponsePtrOutput {
		return v.SubscriptionKeyParameterNames
	}).(SubscriptionKeyParameterNamesContractResponsePtrOutput)
}

func (o GatewayApiEntityTagOutput) SubscriptionRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.BoolPtrOutput { return v.SubscriptionRequired }).(pulumi.BoolPtrOutput)
}

func (o GatewayApiEntityTagOutput) TermsOfServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.StringPtrOutput { return v.TermsOfServiceUrl }).(pulumi.StringPtrOutput)
}

func (o GatewayApiEntityTagOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayApiEntityTag) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayApiEntityTagOutput{})
}
