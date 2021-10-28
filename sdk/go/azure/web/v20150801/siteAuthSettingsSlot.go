


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteAuthSettingsSlot struct {
	pulumi.CustomResourceState

	AadClientId                  pulumi.StringPtrOutput   `pulumi:"aadClientId"`
	AdditionalLoginParams        pulumi.StringArrayOutput `pulumi:"additionalLoginParams"`
	AllowedAudiences             pulumi.StringArrayOutput `pulumi:"allowedAudiences"`
	AllowedExternalRedirectUrls  pulumi.StringArrayOutput `pulumi:"allowedExternalRedirectUrls"`
	ClientId                     pulumi.StringPtrOutput   `pulumi:"clientId"`
	ClientSecret                 pulumi.StringPtrOutput   `pulumi:"clientSecret"`
	DefaultProvider              pulumi.StringPtrOutput   `pulumi:"defaultProvider"`
	Enabled                      pulumi.BoolPtrOutput     `pulumi:"enabled"`
	FacebookAppId                pulumi.StringPtrOutput   `pulumi:"facebookAppId"`
	FacebookAppSecret            pulumi.StringPtrOutput   `pulumi:"facebookAppSecret"`
	FacebookOAuthScopes          pulumi.StringArrayOutput `pulumi:"facebookOAuthScopes"`
	GoogleClientId               pulumi.StringPtrOutput   `pulumi:"googleClientId"`
	GoogleClientSecret           pulumi.StringPtrOutput   `pulumi:"googleClientSecret"`
	GoogleOAuthScopes            pulumi.StringArrayOutput `pulumi:"googleOAuthScopes"`
	HttpApiPrefixPath            pulumi.StringPtrOutput   `pulumi:"httpApiPrefixPath"`
	Issuer                       pulumi.StringPtrOutput   `pulumi:"issuer"`
	MicrosoftAccountClientId     pulumi.StringPtrOutput   `pulumi:"microsoftAccountClientId"`
	MicrosoftAccountClientSecret pulumi.StringPtrOutput   `pulumi:"microsoftAccountClientSecret"`
	MicrosoftAccountOAuthScopes  pulumi.StringArrayOutput `pulumi:"microsoftAccountOAuthScopes"`
	OpenIdIssuer                 pulumi.StringPtrOutput   `pulumi:"openIdIssuer"`
	TokenRefreshExtensionHours   pulumi.Float64PtrOutput  `pulumi:"tokenRefreshExtensionHours"`
	TokenStoreEnabled            pulumi.BoolPtrOutput     `pulumi:"tokenStoreEnabled"`
	TwitterConsumerKey           pulumi.StringPtrOutput   `pulumi:"twitterConsumerKey"`
	TwitterConsumerSecret        pulumi.StringPtrOutput   `pulumi:"twitterConsumerSecret"`
	UnauthenticatedClientAction  pulumi.StringPtrOutput   `pulumi:"unauthenticatedClientAction"`
}


func NewSiteAuthSettingsSlot(ctx *pulumi.Context,
	name string, args *SiteAuthSettingsSlotArgs, opts ...pulumi.ResourceOption) (*SiteAuthSettingsSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:SiteAuthSettingsSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteAuthSettingsSlot
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteAuthSettingsSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteAuthSettingsSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteAuthSettingsSlotState, opts ...pulumi.ResourceOption) (*SiteAuthSettingsSlot, error) {
	var resource SiteAuthSettingsSlot
	err := ctx.ReadResource("azure-native:web/v20150801:SiteAuthSettingsSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteAuthSettingsSlotState struct {
}

type SiteAuthSettingsSlotState struct {
}

func (SiteAuthSettingsSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteAuthSettingsSlotState)(nil)).Elem()
}

type siteAuthSettingsSlotArgs struct {
	AadClientId                  *string                        `pulumi:"aadClientId"`
	AdditionalLoginParams        []string                       `pulumi:"additionalLoginParams"`
	AllowedAudiences             []string                       `pulumi:"allowedAudiences"`
	AllowedExternalRedirectUrls  []string                       `pulumi:"allowedExternalRedirectUrls"`
	ClientId                     *string                        `pulumi:"clientId"`
	ClientSecret                 *string                        `pulumi:"clientSecret"`
	DefaultProvider              *BuiltInAuthenticationProvider `pulumi:"defaultProvider"`
	Enabled                      *bool                          `pulumi:"enabled"`
	FacebookAppId                *string                        `pulumi:"facebookAppId"`
	FacebookAppSecret            *string                        `pulumi:"facebookAppSecret"`
	FacebookOAuthScopes          []string                       `pulumi:"facebookOAuthScopes"`
	GoogleClientId               *string                        `pulumi:"googleClientId"`
	GoogleClientSecret           *string                        `pulumi:"googleClientSecret"`
	GoogleOAuthScopes            []string                       `pulumi:"googleOAuthScopes"`
	HttpApiPrefixPath            *string                        `pulumi:"httpApiPrefixPath"`
	Issuer                       *string                        `pulumi:"issuer"`
	MicrosoftAccountClientId     *string                        `pulumi:"microsoftAccountClientId"`
	MicrosoftAccountClientSecret *string                        `pulumi:"microsoftAccountClientSecret"`
	MicrosoftAccountOAuthScopes  []string                       `pulumi:"microsoftAccountOAuthScopes"`
	Name                         string                         `pulumi:"name"`
	OpenIdIssuer                 *string                        `pulumi:"openIdIssuer"`
	ResourceGroupName            string                         `pulumi:"resourceGroupName"`
	Slot                         string                         `pulumi:"slot"`
	TokenRefreshExtensionHours   *float64                       `pulumi:"tokenRefreshExtensionHours"`
	TokenStoreEnabled            *bool                          `pulumi:"tokenStoreEnabled"`
	TwitterConsumerKey           *string                        `pulumi:"twitterConsumerKey"`
	TwitterConsumerSecret        *string                        `pulumi:"twitterConsumerSecret"`
	UnauthenticatedClientAction  *UnauthenticatedClientAction   `pulumi:"unauthenticatedClientAction"`
}


type SiteAuthSettingsSlotArgs struct {
	AadClientId                  pulumi.StringPtrInput
	AdditionalLoginParams        pulumi.StringArrayInput
	AllowedAudiences             pulumi.StringArrayInput
	AllowedExternalRedirectUrls  pulumi.StringArrayInput
	ClientId                     pulumi.StringPtrInput
	ClientSecret                 pulumi.StringPtrInput
	DefaultProvider              BuiltInAuthenticationProviderPtrInput
	Enabled                      pulumi.BoolPtrInput
	FacebookAppId                pulumi.StringPtrInput
	FacebookAppSecret            pulumi.StringPtrInput
	FacebookOAuthScopes          pulumi.StringArrayInput
	GoogleClientId               pulumi.StringPtrInput
	GoogleClientSecret           pulumi.StringPtrInput
	GoogleOAuthScopes            pulumi.StringArrayInput
	HttpApiPrefixPath            pulumi.StringPtrInput
	Issuer                       pulumi.StringPtrInput
	MicrosoftAccountClientId     pulumi.StringPtrInput
	MicrosoftAccountClientSecret pulumi.StringPtrInput
	MicrosoftAccountOAuthScopes  pulumi.StringArrayInput
	Name                         pulumi.StringInput
	OpenIdIssuer                 pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	Slot                         pulumi.StringInput
	TokenRefreshExtensionHours   pulumi.Float64PtrInput
	TokenStoreEnabled            pulumi.BoolPtrInput
	TwitterConsumerKey           pulumi.StringPtrInput
	TwitterConsumerSecret        pulumi.StringPtrInput
	UnauthenticatedClientAction  UnauthenticatedClientActionPtrInput
}

func (SiteAuthSettingsSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteAuthSettingsSlotArgs)(nil)).Elem()
}

type SiteAuthSettingsSlotInput interface {
	pulumi.Input

	ToSiteAuthSettingsSlotOutput() SiteAuthSettingsSlotOutput
	ToSiteAuthSettingsSlotOutputWithContext(ctx context.Context) SiteAuthSettingsSlotOutput
}

func (*SiteAuthSettingsSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteAuthSettingsSlot)(nil))
}

func (i *SiteAuthSettingsSlot) ToSiteAuthSettingsSlotOutput() SiteAuthSettingsSlotOutput {
	return i.ToSiteAuthSettingsSlotOutputWithContext(context.Background())
}

func (i *SiteAuthSettingsSlot) ToSiteAuthSettingsSlotOutputWithContext(ctx context.Context) SiteAuthSettingsSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteAuthSettingsSlotOutput)
}

type SiteAuthSettingsSlotOutput struct{ *pulumi.OutputState }

func (SiteAuthSettingsSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteAuthSettingsSlot)(nil))
}

func (o SiteAuthSettingsSlotOutput) ToSiteAuthSettingsSlotOutput() SiteAuthSettingsSlotOutput {
	return o
}

func (o SiteAuthSettingsSlotOutput) ToSiteAuthSettingsSlotOutputWithContext(ctx context.Context) SiteAuthSettingsSlotOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SiteAuthSettingsSlotInput)(nil)).Elem(), &SiteAuthSettingsSlot{})
	pulumi.RegisterOutputType(SiteAuthSettingsSlotOutput{})
}
