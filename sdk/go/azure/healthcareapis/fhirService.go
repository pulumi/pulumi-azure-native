


package healthcareapis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FhirService struct {
	pulumi.CustomResourceState

	AccessPolicies                     FhirServiceAccessPolicyEntryResponseArrayOutput         `pulumi:"accessPolicies"`
	AcrConfiguration                   FhirServiceAcrConfigurationResponsePtrOutput            `pulumi:"acrConfiguration"`
	AuthenticationConfiguration        FhirServiceAuthenticationConfigurationResponsePtrOutput `pulumi:"authenticationConfiguration"`
	CorsConfiguration                  FhirServiceCorsConfigurationResponsePtrOutput           `pulumi:"corsConfiguration"`
	Etag                               pulumi.StringPtrOutput                                  `pulumi:"etag"`
	EventState                         pulumi.StringOutput                                     `pulumi:"eventState"`
	ExportConfiguration                FhirServiceExportConfigurationResponsePtrOutput         `pulumi:"exportConfiguration"`
	Identity                           ServiceManagedIdentityResponseIdentityPtrOutput         `pulumi:"identity"`
	Kind                               pulumi.StringPtrOutput                                  `pulumi:"kind"`
	Location                           pulumi.StringPtrOutput                                  `pulumi:"location"`
	Name                               pulumi.StringOutput                                     `pulumi:"name"`
	PrivateEndpointConnections         PrivateEndpointConnectionResponseArrayOutput            `pulumi:"privateEndpointConnections"`
	ProvisioningState                  pulumi.StringOutput                                     `pulumi:"provisioningState"`
	PublicNetworkAccess                pulumi.StringOutput                                     `pulumi:"publicNetworkAccess"`
	ResourceVersionPolicyConfiguration ResourceVersionPolicyConfigurationResponsePtrOutput     `pulumi:"resourceVersionPolicyConfiguration"`
	SystemData                         SystemDataResponseOutput                                `pulumi:"systemData"`
	Tags                               pulumi.StringMapOutput                                  `pulumi:"tags"`
	Type                               pulumi.StringOutput                                     `pulumi:"type"`
}


func NewFhirService(ctx *pulumi.Context,
	name string, args *FhirServiceArgs, opts ...pulumi.ResourceOption) (*FhirService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:healthcareapis/v20210601preview:FhirService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20211101:FhirService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220131preview:FhirService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220515:FhirService"),
		},
		{
			Type: pulumi.String("azure-native:healthcareapis/v20220601:FhirService"),
		},
	})
	opts = append(opts, aliases)
	var resource FhirService
	err := ctx.RegisterResource("azure-native:healthcareapis:FhirService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFhirService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FhirServiceState, opts ...pulumi.ResourceOption) (*FhirService, error) {
	var resource FhirService
	err := ctx.ReadResource("azure-native:healthcareapis:FhirService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type fhirServiceState struct {
}

type FhirServiceState struct {
}

func (FhirServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirServiceState)(nil)).Elem()
}

type fhirServiceArgs struct {
	AccessPolicies                     []FhirServiceAccessPolicyEntry          `pulumi:"accessPolicies"`
	AcrConfiguration                   *FhirServiceAcrConfiguration            `pulumi:"acrConfiguration"`
	AuthenticationConfiguration        *FhirServiceAuthenticationConfiguration `pulumi:"authenticationConfiguration"`
	CorsConfiguration                  *FhirServiceCorsConfiguration           `pulumi:"corsConfiguration"`
	ExportConfiguration                *FhirServiceExportConfiguration         `pulumi:"exportConfiguration"`
	FhirServiceName                    *string                                 `pulumi:"fhirServiceName"`
	Identity                           *ServiceManagedIdentityIdentity         `pulumi:"identity"`
	Kind                               *string                                 `pulumi:"kind"`
	Location                           *string                                 `pulumi:"location"`
	ResourceGroupName                  string                                  `pulumi:"resourceGroupName"`
	ResourceVersionPolicyConfiguration *ResourceVersionPolicyConfiguration     `pulumi:"resourceVersionPolicyConfiguration"`
	Tags                               map[string]string                       `pulumi:"tags"`
	WorkspaceName                      string                                  `pulumi:"workspaceName"`
}


type FhirServiceArgs struct {
	AccessPolicies                     FhirServiceAccessPolicyEntryArrayInput
	AcrConfiguration                   FhirServiceAcrConfigurationPtrInput
	AuthenticationConfiguration        FhirServiceAuthenticationConfigurationPtrInput
	CorsConfiguration                  FhirServiceCorsConfigurationPtrInput
	ExportConfiguration                FhirServiceExportConfigurationPtrInput
	FhirServiceName                    pulumi.StringPtrInput
	Identity                           ServiceManagedIdentityIdentityPtrInput
	Kind                               pulumi.StringPtrInput
	Location                           pulumi.StringPtrInput
	ResourceGroupName                  pulumi.StringInput
	ResourceVersionPolicyConfiguration ResourceVersionPolicyConfigurationPtrInput
	Tags                               pulumi.StringMapInput
	WorkspaceName                      pulumi.StringInput
}

func (FhirServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirServiceArgs)(nil)).Elem()
}

type FhirServiceInput interface {
	pulumi.Input

	ToFhirServiceOutput() FhirServiceOutput
	ToFhirServiceOutputWithContext(ctx context.Context) FhirServiceOutput
}

func (*FhirService) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirService)(nil)).Elem()
}

func (i *FhirService) ToFhirServiceOutput() FhirServiceOutput {
	return i.ToFhirServiceOutputWithContext(context.Background())
}

func (i *FhirService) ToFhirServiceOutputWithContext(ctx context.Context) FhirServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FhirServiceOutput)
}

type FhirServiceOutput struct{ *pulumi.OutputState }

func (FhirServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FhirService)(nil)).Elem()
}

func (o FhirServiceOutput) ToFhirServiceOutput() FhirServiceOutput {
	return o
}

func (o FhirServiceOutput) ToFhirServiceOutputWithContext(ctx context.Context) FhirServiceOutput {
	return o
}

func (o FhirServiceOutput) AccessPolicies() FhirServiceAccessPolicyEntryResponseArrayOutput {
	return o.ApplyT(func(v *FhirService) FhirServiceAccessPolicyEntryResponseArrayOutput { return v.AccessPolicies }).(FhirServiceAccessPolicyEntryResponseArrayOutput)
}

func (o FhirServiceOutput) AcrConfiguration() FhirServiceAcrConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *FhirService) FhirServiceAcrConfigurationResponsePtrOutput { return v.AcrConfiguration }).(FhirServiceAcrConfigurationResponsePtrOutput)
}

func (o FhirServiceOutput) AuthenticationConfiguration() FhirServiceAuthenticationConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *FhirService) FhirServiceAuthenticationConfigurationResponsePtrOutput {
		return v.AuthenticationConfiguration
	}).(FhirServiceAuthenticationConfigurationResponsePtrOutput)
}

func (o FhirServiceOutput) CorsConfiguration() FhirServiceCorsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *FhirService) FhirServiceCorsConfigurationResponsePtrOutput { return v.CorsConfiguration }).(FhirServiceCorsConfigurationResponsePtrOutput)
}

func (o FhirServiceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o FhirServiceOutput) EventState() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringOutput { return v.EventState }).(pulumi.StringOutput)
}

func (o FhirServiceOutput) ExportConfiguration() FhirServiceExportConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *FhirService) FhirServiceExportConfigurationResponsePtrOutput { return v.ExportConfiguration }).(FhirServiceExportConfigurationResponsePtrOutput)
}

func (o FhirServiceOutput) Identity() ServiceManagedIdentityResponseIdentityPtrOutput {
	return o.ApplyT(func(v *FhirService) ServiceManagedIdentityResponseIdentityPtrOutput { return v.Identity }).(ServiceManagedIdentityResponseIdentityPtrOutput)
}

func (o FhirServiceOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o FhirServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o FhirServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FhirServiceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *FhirService) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o FhirServiceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o FhirServiceOutput) PublicNetworkAccess() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringOutput { return v.PublicNetworkAccess }).(pulumi.StringOutput)
}

func (o FhirServiceOutput) ResourceVersionPolicyConfiguration() ResourceVersionPolicyConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *FhirService) ResourceVersionPolicyConfigurationResponsePtrOutput {
		return v.ResourceVersionPolicyConfiguration
	}).(ResourceVersionPolicyConfigurationResponsePtrOutput)
}

func (o FhirServiceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FhirService) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o FhirServiceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o FhirServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FhirService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FhirServiceOutput{})
}
