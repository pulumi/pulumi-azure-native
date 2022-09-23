


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Api struct {
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
	IsCurrent                     pulumi.BoolOutput                                      `pulumi:"isCurrent"`
	IsOnline                      pulumi.BoolOutput                                      `pulumi:"isOnline"`
	Name                          pulumi.StringOutput                                    `pulumi:"name"`
	Path                          pulumi.StringOutput                                    `pulumi:"path"`
	Protocols                     pulumi.StringArrayOutput                               `pulumi:"protocols"`
	ServiceUrl                    pulumi.StringPtrOutput                                 `pulumi:"serviceUrl"`
	SubscriptionKeyParameterNames SubscriptionKeyParameterNamesContractResponsePtrOutput `pulumi:"subscriptionKeyParameterNames"`
	SubscriptionRequired          pulumi.BoolPtrOutput                                   `pulumi:"subscriptionRequired"`
	Type                          pulumi.StringOutput                                    `pulumi:"type"`
}


func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Api"),
		},
	})
	opts = append(opts, aliases)
	var resource Api
	err := ctx.RegisterResource("azure-native:apimanagement/v20180601preview:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("azure-native:apimanagement/v20180601preview:Api", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiState struct {
}

type ApiState struct {
}

func (ApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiState)(nil)).Elem()
}

type apiArgs struct {
	ApiId                         *string                                  `pulumi:"apiId"`
	ApiRevision                   *string                                  `pulumi:"apiRevision"`
	ApiRevisionDescription        *string                                  `pulumi:"apiRevisionDescription"`
	ApiType                       *string                                  `pulumi:"apiType"`
	ApiVersion                    *string                                  `pulumi:"apiVersion"`
	ApiVersionDescription         *string                                  `pulumi:"apiVersionDescription"`
	ApiVersionSet                 *ApiVersionSetContractDetails            `pulumi:"apiVersionSet"`
	ApiVersionSetId               *string                                  `pulumi:"apiVersionSetId"`
	AuthenticationSettings        *AuthenticationSettingsContract          `pulumi:"authenticationSettings"`
	ContentFormat                 *string                                  `pulumi:"contentFormat"`
	ContentValue                  *string                                  `pulumi:"contentValue"`
	Description                   *string                                  `pulumi:"description"`
	DisplayName                   *string                                  `pulumi:"displayName"`
	Path                          string                                   `pulumi:"path"`
	Protocols                     []Protocol                               `pulumi:"protocols"`
	ResourceGroupName             string                                   `pulumi:"resourceGroupName"`
	ServiceName                   string                                   `pulumi:"serviceName"`
	ServiceUrl                    *string                                  `pulumi:"serviceUrl"`
	SoapApiType                   *string                                  `pulumi:"soapApiType"`
	SubscriptionKeyParameterNames *SubscriptionKeyParameterNamesContract   `pulumi:"subscriptionKeyParameterNames"`
	SubscriptionRequired          *bool                                    `pulumi:"subscriptionRequired"`
	WsdlSelector                  *ApiCreateOrUpdatePropertiesWsdlSelector `pulumi:"wsdlSelector"`
}


type ApiArgs struct {
	ApiId                         pulumi.StringPtrInput
	ApiRevision                   pulumi.StringPtrInput
	ApiRevisionDescription        pulumi.StringPtrInput
	ApiType                       pulumi.StringPtrInput
	ApiVersion                    pulumi.StringPtrInput
	ApiVersionDescription         pulumi.StringPtrInput
	ApiVersionSet                 ApiVersionSetContractDetailsPtrInput
	ApiVersionSetId               pulumi.StringPtrInput
	AuthenticationSettings        AuthenticationSettingsContractPtrInput
	ContentFormat                 pulumi.StringPtrInput
	ContentValue                  pulumi.StringPtrInput
	Description                   pulumi.StringPtrInput
	DisplayName                   pulumi.StringPtrInput
	Path                          pulumi.StringInput
	Protocols                     ProtocolArrayInput
	ResourceGroupName             pulumi.StringInput
	ServiceName                   pulumi.StringInput
	ServiceUrl                    pulumi.StringPtrInput
	SoapApiType                   pulumi.StringPtrInput
	SubscriptionKeyParameterNames SubscriptionKeyParameterNamesContractPtrInput
	SubscriptionRequired          pulumi.BoolPtrInput
	WsdlSelector                  ApiCreateOrUpdatePropertiesWsdlSelectorPtrInput
}

func (ApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiArgs)(nil)).Elem()
}

type ApiInput interface {
	pulumi.Input

	ToApiOutput() ApiOutput
	ToApiOutputWithContext(ctx context.Context) ApiOutput
}

func (*Api) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (i *Api) ToApiOutput() ApiOutput {
	return i.ToApiOutputWithContext(context.Background())
}

func (i *Api) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOutput)
}

type ApiOutput struct{ *pulumi.OutputState }

func (ApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (o ApiOutput) ToApiOutput() ApiOutput {
	return o
}

func (o ApiOutput) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return o
}

func (o ApiOutput) ApiRevision() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ApiRevision }).(pulumi.StringPtrOutput)
}

func (o ApiOutput) ApiRevisionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ApiRevisionDescription }).(pulumi.StringPtrOutput)
}

func (o ApiOutput) ApiType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ApiType }).(pulumi.StringPtrOutput)
}

func (o ApiOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o ApiOutput) ApiVersionDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ApiVersionDescription }).(pulumi.StringPtrOutput)
}

func (o ApiOutput) ApiVersionSet() ApiVersionSetContractDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Api) ApiVersionSetContractDetailsResponsePtrOutput { return v.ApiVersionSet }).(ApiVersionSetContractDetailsResponsePtrOutput)
}

func (o ApiOutput) ApiVersionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ApiVersionSetId }).(pulumi.StringPtrOutput)
}

func (o ApiOutput) AuthenticationSettings() AuthenticationSettingsContractResponsePtrOutput {
	return o.ApplyT(func(v *Api) AuthenticationSettingsContractResponsePtrOutput { return v.AuthenticationSettings }).(AuthenticationSettingsContractResponsePtrOutput)
}

func (o ApiOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ApiOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ApiOutput) IsCurrent() pulumi.BoolOutput {
	return o.ApplyT(func(v *Api) pulumi.BoolOutput { return v.IsCurrent }).(pulumi.BoolOutput)
}

func (o ApiOutput) IsOnline() pulumi.BoolOutput {
	return o.ApplyT(func(v *Api) pulumi.BoolOutput { return v.IsOnline }).(pulumi.BoolOutput)
}

func (o ApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApiOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

func (o ApiOutput) Protocols() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Api) pulumi.StringArrayOutput { return v.Protocols }).(pulumi.StringArrayOutput)
}

func (o ApiOutput) ServiceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.StringPtrOutput { return v.ServiceUrl }).(pulumi.StringPtrOutput)
}

func (o ApiOutput) SubscriptionKeyParameterNames() SubscriptionKeyParameterNamesContractResponsePtrOutput {
	return o.ApplyT(func(v *Api) SubscriptionKeyParameterNamesContractResponsePtrOutput {
		return v.SubscriptionKeyParameterNames
	}).(SubscriptionKeyParameterNamesContractResponsePtrOutput)
}

func (o ApiOutput) SubscriptionRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Api) pulumi.BoolPtrOutput { return v.SubscriptionRequired }).(pulumi.BoolPtrOutput)
}

func (o ApiOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiOutput{})
}
