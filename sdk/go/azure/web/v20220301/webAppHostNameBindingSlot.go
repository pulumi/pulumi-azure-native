


package v20220301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppHostNameBindingSlot struct {
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


func NewWebAppHostNameBindingSlot(ctx *pulumi.Context,
	name string, args *WebAppHostNameBindingSlotArgs, opts ...pulumi.ResourceOption) (*WebAppHostNameBindingSlot, error) {
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
			Type: pulumi.String("azure-native:web:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppHostNameBindingSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppHostNameBindingSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppHostNameBindingSlot
	err := ctx.RegisterResource("azure-native:web/v20220301:WebAppHostNameBindingSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppHostNameBindingSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppHostNameBindingSlotState, opts ...pulumi.ResourceOption) (*WebAppHostNameBindingSlot, error) {
	var resource WebAppHostNameBindingSlot
	err := ctx.ReadResource("azure-native:web/v20220301:WebAppHostNameBindingSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppHostNameBindingSlotState struct {
}

type WebAppHostNameBindingSlotState struct {
}

func (WebAppHostNameBindingSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppHostNameBindingSlotState)(nil)).Elem()
}

type webAppHostNameBindingSlotArgs struct {
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
	Slot                        string                       `pulumi:"slot"`
	SslState                    *SslState                    `pulumi:"sslState"`
	Thumbprint                  *string                      `pulumi:"thumbprint"`
}


type WebAppHostNameBindingSlotArgs struct {
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
	Slot                        pulumi.StringInput
	SslState                    SslStatePtrInput
	Thumbprint                  pulumi.StringPtrInput
}

func (WebAppHostNameBindingSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppHostNameBindingSlotArgs)(nil)).Elem()
}

type WebAppHostNameBindingSlotInput interface {
	pulumi.Input

	ToWebAppHostNameBindingSlotOutput() WebAppHostNameBindingSlotOutput
	ToWebAppHostNameBindingSlotOutputWithContext(ctx context.Context) WebAppHostNameBindingSlotOutput
}

func (*WebAppHostNameBindingSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppHostNameBindingSlot)(nil)).Elem()
}

func (i *WebAppHostNameBindingSlot) ToWebAppHostNameBindingSlotOutput() WebAppHostNameBindingSlotOutput {
	return i.ToWebAppHostNameBindingSlotOutputWithContext(context.Background())
}

func (i *WebAppHostNameBindingSlot) ToWebAppHostNameBindingSlotOutputWithContext(ctx context.Context) WebAppHostNameBindingSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppHostNameBindingSlotOutput)
}

type WebAppHostNameBindingSlotOutput struct{ *pulumi.OutputState }

func (WebAppHostNameBindingSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppHostNameBindingSlot)(nil)).Elem()
}

func (o WebAppHostNameBindingSlotOutput) ToWebAppHostNameBindingSlotOutput() WebAppHostNameBindingSlotOutput {
	return o
}

func (o WebAppHostNameBindingSlotOutput) ToWebAppHostNameBindingSlotOutputWithContext(ctx context.Context) WebAppHostNameBindingSlotOutput {
	return o
}

func (o WebAppHostNameBindingSlotOutput) AzureResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.AzureResourceName }).(pulumi.StringPtrOutput)
}

func (o WebAppHostNameBindingSlotOutput) AzureResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.AzureResourceType }).(pulumi.StringPtrOutput)
}

func (o WebAppHostNameBindingSlotOutput) CustomHostNameDnsRecordType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.CustomHostNameDnsRecordType }).(pulumi.StringPtrOutput)
}

func (o WebAppHostNameBindingSlotOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.DomainId }).(pulumi.StringPtrOutput)
}

func (o WebAppHostNameBindingSlotOutput) HostNameType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.HostNameType }).(pulumi.StringPtrOutput)
}

func (o WebAppHostNameBindingSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppHostNameBindingSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppHostNameBindingSlotOutput) SiteName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.SiteName }).(pulumi.StringPtrOutput)
}

func (o WebAppHostNameBindingSlotOutput) SslState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.SslState }).(pulumi.StringPtrOutput)
}

func (o WebAppHostNameBindingSlotOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringPtrOutput { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func (o WebAppHostNameBindingSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WebAppHostNameBindingSlotOutput) VirtualIP() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppHostNameBindingSlot) pulumi.StringOutput { return v.VirtualIP }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppHostNameBindingSlotOutput{})
}
