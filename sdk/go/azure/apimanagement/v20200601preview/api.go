


package v20200601preview

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
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement:Api"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Api"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20160707:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Api"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20161010:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Api"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Api"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Api"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180601preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Api"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Api"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Api"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Api"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Api"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Api"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:Api"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Api"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210801:Api"),
		},
	})
	opts = append(opts, aliases)
	var resource Api
	err := ctx.RegisterResource("azure-native:apimanagement/v20200601preview:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("azure-native:apimanagement/v20200601preview:Api", name, id, state, &resource, opts...)
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
	Description                   *string                                  `pulumi:"description"`
	DisplayName                   *string                                  `pulumi:"displayName"`
	Format                        *string                                  `pulumi:"format"`
	IsCurrent                     *bool                                    `pulumi:"isCurrent"`
	Path                          string                                   `pulumi:"path"`
	Protocols                     []Protocol                               `pulumi:"protocols"`
	ResourceGroupName             string                                   `pulumi:"resourceGroupName"`
	ServiceName                   string                                   `pulumi:"serviceName"`
	ServiceUrl                    *string                                  `pulumi:"serviceUrl"`
	SoapApiType                   *string                                  `pulumi:"soapApiType"`
	SourceApiId                   *string                                  `pulumi:"sourceApiId"`
	SubscriptionKeyParameterNames *SubscriptionKeyParameterNamesContract   `pulumi:"subscriptionKeyParameterNames"`
	SubscriptionRequired          *bool                                    `pulumi:"subscriptionRequired"`
	Value                         *string                                  `pulumi:"value"`
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
	Description                   pulumi.StringPtrInput
	DisplayName                   pulumi.StringPtrInput
	Format                        pulumi.StringPtrInput
	IsCurrent                     pulumi.BoolPtrInput
	Path                          pulumi.StringInput
	Protocols                     ProtocolArrayInput
	ResourceGroupName             pulumi.StringInput
	ServiceName                   pulumi.StringInput
	ServiceUrl                    pulumi.StringPtrInput
	SoapApiType                   pulumi.StringPtrInput
	SourceApiId                   pulumi.StringPtrInput
	SubscriptionKeyParameterNames SubscriptionKeyParameterNamesContractPtrInput
	SubscriptionRequired          pulumi.BoolPtrInput
	Value                         pulumi.StringPtrInput
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
	return reflect.TypeOf((*Api)(nil))
}

func (i *Api) ToApiOutput() ApiOutput {
	return i.ToApiOutputWithContext(context.Background())
}

func (i *Api) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOutput)
}

type ApiOutput struct{ *pulumi.OutputState }

func (ApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Api)(nil))
}

func (o ApiOutput) ToApiOutput() ApiOutput {
	return o
}

func (o ApiOutput) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiInput)(nil)).Elem(), &Api{})
	pulumi.RegisterOutputType(ApiOutput{})
}
