


package v20180201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppHostNameBinding struct {
	pulumi.CustomResourceState

	AzureResourceName           pulumi.StringPtrOutput `pulumi:"azureResourceName"`
	AzureResourceType           pulumi.StringPtrOutput `pulumi:"azureResourceType"`
	CustomHostNameDnsRecordType pulumi.StringPtrOutput `pulumi:"customHostNameDnsRecordType"`
	DomainId                    pulumi.StringPtrOutput `pulumi:"domainId"`
	HostNameType                pulumi.StringPtrOutput `pulumi:"hostNameType"`
	Kind                        pulumi.StringPtrOutput `pulumi:"kind"`
	Name                        pulumi.StringOutput    `pulumi:"name"`
	SiteName                    pulumi.StringPtrOutput `pulumi:"siteName"`
	SslState                    pulumi.StringPtrOutput `pulumi:"sslState"`
	Thumbprint                  pulumi.StringPtrOutput `pulumi:"thumbprint"`
	Type                        pulumi.StringOutput    `pulumi:"type"`
	VirtualIP                   pulumi.StringOutput    `pulumi:"virtualIP"`
}


func NewWebAppHostNameBinding(ctx *pulumi.Context,
	name string, args *WebAppHostNameBindingArgs, opts ...pulumi.ResourceOption) (*WebAppHostNameBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppHostNameBinding"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppHostNameBinding
	err := ctx.RegisterResource("azure-native:web/v20180201:WebAppHostNameBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppHostNameBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppHostNameBindingState, opts ...pulumi.ResourceOption) (*WebAppHostNameBinding, error) {
	var resource WebAppHostNameBinding
	err := ctx.ReadResource("azure-native:web/v20180201:WebAppHostNameBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppHostNameBindingState struct {
}

type WebAppHostNameBindingState struct {
}

func (WebAppHostNameBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppHostNameBindingState)(nil)).Elem()
}

type webAppHostNameBindingArgs struct {
	AzureResourceName           *string                      `pulumi:"azureResourceName"`
	AzureResourceType           *AzureResourceType           `pulumi:"azureResourceType"`
	CustomHostNameDnsRecordType *CustomHostNameDnsRecordType `pulumi:"customHostNameDnsRecordType"`
	DomainId                    *string                      `pulumi:"domainId"`
	HostName                    *string                      `pulumi:"hostName"`
	HostNameType                *HostNameType                `pulumi:"hostNameType"`
	Kind                        *string                      `pulumi:"kind"`
	Name                        string                       `pulumi:"name"`
	ResourceGroupName           string                       `pulumi:"resourceGroupName"`
	SiteName                    *string                      `pulumi:"siteName"`
	SslState                    *SslState                    `pulumi:"sslState"`
	Thumbprint                  *string                      `pulumi:"thumbprint"`
}


type WebAppHostNameBindingArgs struct {
	AzureResourceName           pulumi.StringPtrInput
	AzureResourceType           AzureResourceTypePtrInput
	CustomHostNameDnsRecordType CustomHostNameDnsRecordTypePtrInput
	DomainId                    pulumi.StringPtrInput
	HostName                    pulumi.StringPtrInput
	HostNameType                HostNameTypePtrInput
	Kind                        pulumi.StringPtrInput
	Name                        pulumi.StringInput
	ResourceGroupName           pulumi.StringInput
	SiteName                    pulumi.StringPtrInput
	SslState                    SslStatePtrInput
	Thumbprint                  pulumi.StringPtrInput
}

func (WebAppHostNameBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppHostNameBindingArgs)(nil)).Elem()
}

type WebAppHostNameBindingInput interface {
	pulumi.Input

	ToWebAppHostNameBindingOutput() WebAppHostNameBindingOutput
	ToWebAppHostNameBindingOutputWithContext(ctx context.Context) WebAppHostNameBindingOutput
}

func (*WebAppHostNameBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppHostNameBinding)(nil))
}

func (i *WebAppHostNameBinding) ToWebAppHostNameBindingOutput() WebAppHostNameBindingOutput {
	return i.ToWebAppHostNameBindingOutputWithContext(context.Background())
}

func (i *WebAppHostNameBinding) ToWebAppHostNameBindingOutputWithContext(ctx context.Context) WebAppHostNameBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppHostNameBindingOutput)
}

type WebAppHostNameBindingOutput struct{ *pulumi.OutputState }

func (WebAppHostNameBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppHostNameBinding)(nil))
}

func (o WebAppHostNameBindingOutput) ToWebAppHostNameBindingOutput() WebAppHostNameBindingOutput {
	return o
}

func (o WebAppHostNameBindingOutput) ToWebAppHostNameBindingOutputWithContext(ctx context.Context) WebAppHostNameBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppHostNameBindingOutput{})
}
