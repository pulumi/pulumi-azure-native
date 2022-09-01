


package v20220301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiPortalCustomDomain struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                           `pulumi:"name"`
	Properties ApiPortalCustomDomainPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                      `pulumi:"systemData"`
	Type       pulumi.StringOutput                           `pulumi:"type"`
}


func NewApiPortalCustomDomain(ctx *pulumi.Context,
	name string, args *ApiPortalCustomDomainArgs, opts ...pulumi.ResourceOption) (*ApiPortalCustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiPortalName == nil {
		return nil, errors.New("invalid value for required argument 'ApiPortalName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:ApiPortalCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:ApiPortalCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:ApiPortalCustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:ApiPortalCustomDomain"),
		},
	})
	opts = append(opts, aliases)
	var resource ApiPortalCustomDomain
	err := ctx.RegisterResource("azure-native:appplatform/v20220301preview:ApiPortalCustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApiPortalCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiPortalCustomDomainState, opts ...pulumi.ResourceOption) (*ApiPortalCustomDomain, error) {
	var resource ApiPortalCustomDomain
	err := ctx.ReadResource("azure-native:appplatform/v20220301preview:ApiPortalCustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type apiPortalCustomDomainState struct {
}

type ApiPortalCustomDomainState struct {
}

func (ApiPortalCustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiPortalCustomDomainState)(nil)).Elem()
}

type apiPortalCustomDomainArgs struct {
	ApiPortalName     string                           `pulumi:"apiPortalName"`
	DomainName        *string                          `pulumi:"domainName"`
	Properties        *ApiPortalCustomDomainProperties `pulumi:"properties"`
	ResourceGroupName string                           `pulumi:"resourceGroupName"`
	ServiceName       string                           `pulumi:"serviceName"`
}


type ApiPortalCustomDomainArgs struct {
	ApiPortalName     pulumi.StringInput
	DomainName        pulumi.StringPtrInput
	Properties        ApiPortalCustomDomainPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
}

func (ApiPortalCustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiPortalCustomDomainArgs)(nil)).Elem()
}

type ApiPortalCustomDomainInput interface {
	pulumi.Input

	ToApiPortalCustomDomainOutput() ApiPortalCustomDomainOutput
	ToApiPortalCustomDomainOutputWithContext(ctx context.Context) ApiPortalCustomDomainOutput
}

func (*ApiPortalCustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPortalCustomDomain)(nil)).Elem()
}

func (i *ApiPortalCustomDomain) ToApiPortalCustomDomainOutput() ApiPortalCustomDomainOutput {
	return i.ToApiPortalCustomDomainOutputWithContext(context.Background())
}

func (i *ApiPortalCustomDomain) ToApiPortalCustomDomainOutputWithContext(ctx context.Context) ApiPortalCustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPortalCustomDomainOutput)
}

type ApiPortalCustomDomainOutput struct{ *pulumi.OutputState }

func (ApiPortalCustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPortalCustomDomain)(nil)).Elem()
}

func (o ApiPortalCustomDomainOutput) ToApiPortalCustomDomainOutput() ApiPortalCustomDomainOutput {
	return o
}

func (o ApiPortalCustomDomainOutput) ToApiPortalCustomDomainOutputWithContext(ctx context.Context) ApiPortalCustomDomainOutput {
	return o
}

func (o ApiPortalCustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiPortalCustomDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApiPortalCustomDomainOutput) Properties() ApiPortalCustomDomainPropertiesResponseOutput {
	return o.ApplyT(func(v *ApiPortalCustomDomain) ApiPortalCustomDomainPropertiesResponseOutput { return v.Properties }).(ApiPortalCustomDomainPropertiesResponseOutput)
}

func (o ApiPortalCustomDomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ApiPortalCustomDomain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ApiPortalCustomDomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiPortalCustomDomain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiPortalCustomDomainOutput{})
}
