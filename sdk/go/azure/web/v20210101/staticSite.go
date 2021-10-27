


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StaticSite struct {
	pulumi.CustomResourceState

	AllowConfigFileUpdates      pulumi.BoolPtrOutput                                                      `pulumi:"allowConfigFileUpdates"`
	Branch                      pulumi.StringPtrOutput                                                    `pulumi:"branch"`
	BuildProperties             StaticSiteBuildPropertiesResponsePtrOutput                                `pulumi:"buildProperties"`
	ContentDistributionEndpoint pulumi.StringOutput                                                       `pulumi:"contentDistributionEndpoint"`
	CustomDomains               pulumi.StringArrayOutput                                                  `pulumi:"customDomains"`
	DefaultHostname             pulumi.StringOutput                                                       `pulumi:"defaultHostname"`
	Identity                    ManagedServiceIdentityResponsePtrOutput                                   `pulumi:"identity"`
	KeyVaultReferenceIdentity   pulumi.StringOutput                                                       `pulumi:"keyVaultReferenceIdentity"`
	Kind                        pulumi.StringPtrOutput                                                    `pulumi:"kind"`
	Location                    pulumi.StringOutput                                                       `pulumi:"location"`
	Name                        pulumi.StringOutput                                                       `pulumi:"name"`
	PrivateEndpointConnections  ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	Provider                    pulumi.StringOutput                                                       `pulumi:"provider"`
	RepositoryToken             pulumi.StringPtrOutput                                                    `pulumi:"repositoryToken"`
	RepositoryUrl               pulumi.StringPtrOutput                                                    `pulumi:"repositoryUrl"`
	Sku                         SkuDescriptionResponsePtrOutput                                           `pulumi:"sku"`
	StagingEnvironmentPolicy    pulumi.StringPtrOutput                                                    `pulumi:"stagingEnvironmentPolicy"`
	Tags                        pulumi.StringMapOutput                                                    `pulumi:"tags"`
	TemplateProperties          StaticSiteTemplateOptionsResponsePtrOutput                                `pulumi:"templateProperties"`
	Type                        pulumi.StringOutput                                                       `pulumi:"type"`
	UserProvidedFunctionApps    StaticSiteUserProvidedFunctionAppResponseArrayOutput                      `pulumi:"userProvidedFunctionApps"`
}


func NewStaticSite(ctx *pulumi.Context,
	name string, args *StaticSiteArgs, opts ...pulumi.ResourceOption) (*StaticSite, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web:StaticSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:StaticSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:StaticSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:StaticSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:StaticSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:StaticSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:StaticSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:StaticSite"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:StaticSite"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:StaticSite"),
		},
	})
	opts = append(opts, aliases)
	var resource StaticSite
	err := ctx.RegisterResource("azure-native:web/v20210101:StaticSite", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStaticSite(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticSiteState, opts ...pulumi.ResourceOption) (*StaticSite, error) {
	var resource StaticSite
	err := ctx.ReadResource("azure-native:web/v20210101:StaticSite", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type staticSiteState struct {
}

type StaticSiteState struct {
}

func (StaticSiteState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteState)(nil)).Elem()
}

type staticSiteArgs struct {
	AllowConfigFileUpdates   *bool                      `pulumi:"allowConfigFileUpdates"`
	Branch                   *string                    `pulumi:"branch"`
	BuildProperties          *StaticSiteBuildProperties `pulumi:"buildProperties"`
	Identity                 *ManagedServiceIdentity    `pulumi:"identity"`
	Kind                     *string                    `pulumi:"kind"`
	Location                 *string                    `pulumi:"location"`
	Name                     *string                    `pulumi:"name"`
	RepositoryToken          *string                    `pulumi:"repositoryToken"`
	RepositoryUrl            *string                    `pulumi:"repositoryUrl"`
	ResourceGroupName        string                     `pulumi:"resourceGroupName"`
	Sku                      *SkuDescription            `pulumi:"sku"`
	StagingEnvironmentPolicy *StagingEnvironmentPolicy  `pulumi:"stagingEnvironmentPolicy"`
	Tags                     map[string]string          `pulumi:"tags"`
	TemplateProperties       *StaticSiteTemplateOptions `pulumi:"templateProperties"`
}


type StaticSiteArgs struct {
	AllowConfigFileUpdates   pulumi.BoolPtrInput
	Branch                   pulumi.StringPtrInput
	BuildProperties          StaticSiteBuildPropertiesPtrInput
	Identity                 ManagedServiceIdentityPtrInput
	Kind                     pulumi.StringPtrInput
	Location                 pulumi.StringPtrInput
	Name                     pulumi.StringPtrInput
	RepositoryToken          pulumi.StringPtrInput
	RepositoryUrl            pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	Sku                      SkuDescriptionPtrInput
	StagingEnvironmentPolicy StagingEnvironmentPolicyPtrInput
	Tags                     pulumi.StringMapInput
	TemplateProperties       StaticSiteTemplateOptionsPtrInput
}

func (StaticSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSiteArgs)(nil)).Elem()
}

type StaticSiteInput interface {
	pulumi.Input

	ToStaticSiteOutput() StaticSiteOutput
	ToStaticSiteOutputWithContext(ctx context.Context) StaticSiteOutput
}

func (*StaticSite) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSite)(nil))
}

func (i *StaticSite) ToStaticSiteOutput() StaticSiteOutput {
	return i.ToStaticSiteOutputWithContext(context.Background())
}

func (i *StaticSite) ToStaticSiteOutputWithContext(ctx context.Context) StaticSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteOutput)
}

type StaticSiteOutput struct{ *pulumi.OutputState }

func (StaticSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSite)(nil))
}

func (o StaticSiteOutput) ToStaticSiteOutput() StaticSiteOutput {
	return o
}

func (o StaticSiteOutput) ToStaticSiteOutputWithContext(ctx context.Context) StaticSiteOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StaticSiteInput)(nil)).Elem(), &StaticSite{})
	pulumi.RegisterOutputType(StaticSiteOutput{})
}
