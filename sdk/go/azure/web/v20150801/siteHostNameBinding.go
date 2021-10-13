


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteHostNameBinding struct {
	pulumi.CustomResourceState

	AzureResourceName           pulumi.StringPtrOutput `pulumi:"azureResourceName"`
	AzureResourceType           pulumi.StringPtrOutput `pulumi:"azureResourceType"`
	CustomHostNameDnsRecordType pulumi.StringPtrOutput `pulumi:"customHostNameDnsRecordType"`
	DomainId                    pulumi.StringPtrOutput `pulumi:"domainId"`
	HostNameType                pulumi.StringPtrOutput `pulumi:"hostNameType"`
	Kind                        pulumi.StringPtrOutput `pulumi:"kind"`
	Location                    pulumi.StringOutput    `pulumi:"location"`
	Name                        pulumi.StringPtrOutput `pulumi:"name"`
	SiteName                    pulumi.StringPtrOutput `pulumi:"siteName"`
	Tags                        pulumi.StringMapOutput `pulumi:"tags"`
	Type                        pulumi.StringPtrOutput `pulumi:"type"`
}


func NewSiteHostNameBinding(ctx *pulumi.Context,
	name string, args *SiteHostNameBindingArgs, opts ...pulumi.ResourceOption) (*SiteHostNameBinding, error) {
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
			Type: pulumi.String("azure-nextgen:web/v20150801:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200901:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201001:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteHostNameBinding"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:SiteHostNameBinding"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteHostNameBinding
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteHostNameBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteHostNameBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteHostNameBindingState, opts ...pulumi.ResourceOption) (*SiteHostNameBinding, error) {
	var resource SiteHostNameBinding
	err := ctx.ReadResource("azure-native:web/v20150801:SiteHostNameBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteHostNameBindingState struct {
}

type SiteHostNameBindingState struct {
}

func (SiteHostNameBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteHostNameBindingState)(nil)).Elem()
}

type siteHostNameBindingArgs struct {
	AzureResourceName           *string                      `pulumi:"azureResourceName"`
	AzureResourceType           *AzureResourceType           `pulumi:"azureResourceType"`
	CustomHostNameDnsRecordType *CustomHostNameDnsRecordType `pulumi:"customHostNameDnsRecordType"`
	DomainId                    *string                      `pulumi:"domainId"`
	HostName                    *string                      `pulumi:"hostName"`
	HostNameType                *HostNameType                `pulumi:"hostNameType"`
	Id                          *string                      `pulumi:"id"`
	Kind                        *string                      `pulumi:"kind"`
	Location                    *string                      `pulumi:"location"`
	Name                        string                       `pulumi:"name"`
	ResourceGroupName           string                       `pulumi:"resourceGroupName"`
	SiteName                    *string                      `pulumi:"siteName"`
	Tags                        map[string]string            `pulumi:"tags"`
	Type                        *string                      `pulumi:"type"`
}


type SiteHostNameBindingArgs struct {
	AzureResourceName           pulumi.StringPtrInput
	AzureResourceType           AzureResourceTypePtrInput
	CustomHostNameDnsRecordType CustomHostNameDnsRecordTypePtrInput
	DomainId                    pulumi.StringPtrInput
	HostName                    pulumi.StringPtrInput
	HostNameType                HostNameTypePtrInput
	Id                          pulumi.StringPtrInput
	Kind                        pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	Name                        pulumi.StringInput
	ResourceGroupName           pulumi.StringInput
	SiteName                    pulumi.StringPtrInput
	Tags                        pulumi.StringMapInput
	Type                        pulumi.StringPtrInput
}

func (SiteHostNameBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteHostNameBindingArgs)(nil)).Elem()
}

type SiteHostNameBindingInput interface {
	pulumi.Input

	ToSiteHostNameBindingOutput() SiteHostNameBindingOutput
	ToSiteHostNameBindingOutputWithContext(ctx context.Context) SiteHostNameBindingOutput
}

func (*SiteHostNameBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteHostNameBinding)(nil))
}

func (i *SiteHostNameBinding) ToSiteHostNameBindingOutput() SiteHostNameBindingOutput {
	return i.ToSiteHostNameBindingOutputWithContext(context.Background())
}

func (i *SiteHostNameBinding) ToSiteHostNameBindingOutputWithContext(ctx context.Context) SiteHostNameBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteHostNameBindingOutput)
}

type SiteHostNameBindingOutput struct{ *pulumi.OutputState }

func (SiteHostNameBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteHostNameBinding)(nil))
}

func (o SiteHostNameBindingOutput) ToSiteHostNameBindingOutput() SiteHostNameBindingOutput {
	return o
}

func (o SiteHostNameBindingOutput) ToSiteHostNameBindingOutputWithContext(ctx context.Context) SiteHostNameBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SiteHostNameBindingOutput{})
}
