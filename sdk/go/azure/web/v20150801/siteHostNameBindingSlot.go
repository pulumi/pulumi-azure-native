


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SiteHostNameBindingSlot struct {
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


func NewSiteHostNameBindingSlot(ctx *pulumi.Context,
	name string, args *SiteHostNameBindingSlotArgs, opts ...pulumi.ResourceOption) (*SiteHostNameBindingSlot, error) {
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
			Type: pulumi.String("azure-native:web:SiteHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:SiteHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:SiteHostNameBindingSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteHostNameBindingSlot
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteHostNameBindingSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteHostNameBindingSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteHostNameBindingSlotState, opts ...pulumi.ResourceOption) (*SiteHostNameBindingSlot, error) {
	var resource SiteHostNameBindingSlot
	err := ctx.ReadResource("azure-native:web/v20150801:SiteHostNameBindingSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteHostNameBindingSlotState struct {
}

type SiteHostNameBindingSlotState struct {
}

func (SiteHostNameBindingSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteHostNameBindingSlotState)(nil)).Elem()
}

type siteHostNameBindingSlotArgs struct {
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
	Slot                        string                       `pulumi:"slot"`
	Tags                        map[string]string            `pulumi:"tags"`
	Type                        *string                      `pulumi:"type"`
}


type SiteHostNameBindingSlotArgs struct {
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
	Slot                        pulumi.StringInput
	Tags                        pulumi.StringMapInput
	Type                        pulumi.StringPtrInput
}

func (SiteHostNameBindingSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteHostNameBindingSlotArgs)(nil)).Elem()
}

type SiteHostNameBindingSlotInput interface {
	pulumi.Input

	ToSiteHostNameBindingSlotOutput() SiteHostNameBindingSlotOutput
	ToSiteHostNameBindingSlotOutputWithContext(ctx context.Context) SiteHostNameBindingSlotOutput
}

func (*SiteHostNameBindingSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteHostNameBindingSlot)(nil)).Elem()
}

func (i *SiteHostNameBindingSlot) ToSiteHostNameBindingSlotOutput() SiteHostNameBindingSlotOutput {
	return i.ToSiteHostNameBindingSlotOutputWithContext(context.Background())
}

func (i *SiteHostNameBindingSlot) ToSiteHostNameBindingSlotOutputWithContext(ctx context.Context) SiteHostNameBindingSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteHostNameBindingSlotOutput)
}

type SiteHostNameBindingSlotOutput struct{ *pulumi.OutputState }

func (SiteHostNameBindingSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteHostNameBindingSlot)(nil)).Elem()
}

func (o SiteHostNameBindingSlotOutput) ToSiteHostNameBindingSlotOutput() SiteHostNameBindingSlotOutput {
	return o
}

func (o SiteHostNameBindingSlotOutput) ToSiteHostNameBindingSlotOutputWithContext(ctx context.Context) SiteHostNameBindingSlotOutput {
	return o
}

func (o SiteHostNameBindingSlotOutput) AzureResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteHostNameBindingSlot) pulumi.StringPtrOutput { return v.AzureResourceName }).(pulumi.StringPtrOutput)
}

func (o SiteHostNameBindingSlotOutput) AzureResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteHostNameBindingSlot) pulumi.StringPtrOutput { return v.AzureResourceType }).(pulumi.StringPtrOutput)
}

func (o SiteHostNameBindingSlotOutput) CustomHostNameDnsRecordType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteHostNameBindingSlot) pulumi.StringPtrOutput { return v.CustomHostNameDnsRecordType }).(pulumi.StringPtrOutput)
}

func (o SiteHostNameBindingSlotOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteHostNameBindingSlot) pulumi.StringPtrOutput { return v.DomainId }).(pulumi.StringPtrOutput)
}

func (o SiteHostNameBindingSlotOutput) HostNameType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteHostNameBindingSlot) pulumi.StringPtrOutput { return v.HostNameType }).(pulumi.StringPtrOutput)
}

func (o SiteHostNameBindingSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteHostNameBindingSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o SiteHostNameBindingSlotOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SiteHostNameBindingSlot) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SiteHostNameBindingSlotOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteHostNameBindingSlot) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SiteHostNameBindingSlotOutput) SiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteHostNameBindingSlot) pulumi.StringPtrOutput { return v.SiteName }).(pulumi.StringPtrOutput)
}

func (o SiteHostNameBindingSlotOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SiteHostNameBindingSlot) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o SiteHostNameBindingSlotOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteHostNameBindingSlot) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteHostNameBindingSlotOutput{})
}
