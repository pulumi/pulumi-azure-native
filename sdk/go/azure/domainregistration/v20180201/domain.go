


package v20180201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Domain struct {
	pulumi.CustomResourceState

	AuthCode                    pulumi.StringPtrOutput      `pulumi:"authCode"`
	AutoRenew                   pulumi.BoolPtrOutput        `pulumi:"autoRenew"`
	CreatedTime                 pulumi.StringOutput         `pulumi:"createdTime"`
	DnsType                     pulumi.StringPtrOutput      `pulumi:"dnsType"`
	DnsZoneId                   pulumi.StringPtrOutput      `pulumi:"dnsZoneId"`
	DomainNotRenewableReasons   pulumi.StringArrayOutput    `pulumi:"domainNotRenewableReasons"`
	ExpirationTime              pulumi.StringOutput         `pulumi:"expirationTime"`
	Kind                        pulumi.StringPtrOutput      `pulumi:"kind"`
	LastRenewedTime             pulumi.StringOutput         `pulumi:"lastRenewedTime"`
	Location                    pulumi.StringOutput         `pulumi:"location"`
	ManagedHostNames            HostNameResponseArrayOutput `pulumi:"managedHostNames"`
	Name                        pulumi.StringOutput         `pulumi:"name"`
	NameServers                 pulumi.StringArrayOutput    `pulumi:"nameServers"`
	Privacy                     pulumi.BoolPtrOutput        `pulumi:"privacy"`
	ProvisioningState           pulumi.StringOutput         `pulumi:"provisioningState"`
	ReadyForDnsRecordManagement pulumi.BoolOutput           `pulumi:"readyForDnsRecordManagement"`
	RegistrationStatus          pulumi.StringOutput         `pulumi:"registrationStatus"`
	Tags                        pulumi.StringMapOutput      `pulumi:"tags"`
	TargetDnsType               pulumi.StringPtrOutput      `pulumi:"targetDnsType"`
	Type                        pulumi.StringOutput         `pulumi:"type"`
}


func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Consent == nil {
		return nil, errors.New("invalid value for required argument 'Consent'")
	}
	if args.ContactAdmin == nil {
		return nil, errors.New("invalid value for required argument 'ContactAdmin'")
	}
	if args.ContactBilling == nil {
		return nil, errors.New("invalid value for required argument 'ContactBilling'")
	}
	if args.ContactRegistrant == nil {
		return nil, errors.New("invalid value for required argument 'ContactRegistrant'")
	}
	if args.ContactTech == nil {
		return nil, errors.New("invalid value for required argument 'ContactTech'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.AutoRenew) {
		args.AutoRenew = pulumi.BoolPtr(true)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:domainregistration:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20150401:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20190801:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20200601:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20200901:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20201001:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20201201:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20210101:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20210115:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20210201:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20210301:Domain"),
		},
		{
			Type: pulumi.String("azure-native:domainregistration/v20220301:Domain"),
		},
	})
	opts = append(opts, aliases)
	var resource Domain
	err := ctx.RegisterResource("azure-native:domainregistration/v20180201:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("azure-native:domainregistration/v20180201:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type domainState struct {
}

type DomainState struct {
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	AuthCode          *string               `pulumi:"authCode"`
	AutoRenew         *bool                 `pulumi:"autoRenew"`
	Consent           DomainPurchaseConsent `pulumi:"consent"`
	ContactAdmin      Contact               `pulumi:"contactAdmin"`
	ContactBilling    Contact               `pulumi:"contactBilling"`
	ContactRegistrant Contact               `pulumi:"contactRegistrant"`
	ContactTech       Contact               `pulumi:"contactTech"`
	DnsType           *DnsType              `pulumi:"dnsType"`
	DnsZoneId         *string               `pulumi:"dnsZoneId"`
	DomainName        *string               `pulumi:"domainName"`
	Kind              *string               `pulumi:"kind"`
	Location          *string               `pulumi:"location"`
	Privacy           *bool                 `pulumi:"privacy"`
	ResourceGroupName string                `pulumi:"resourceGroupName"`
	Tags              map[string]string     `pulumi:"tags"`
	TargetDnsType     *DnsType              `pulumi:"targetDnsType"`
}


type DomainArgs struct {
	AuthCode          pulumi.StringPtrInput
	AutoRenew         pulumi.BoolPtrInput
	Consent           DomainPurchaseConsentInput
	ContactAdmin      ContactInput
	ContactBilling    ContactInput
	ContactRegistrant ContactInput
	ContactTech       ContactInput
	DnsType           DnsTypePtrInput
	DnsZoneId         pulumi.StringPtrInput
	DomainName        pulumi.StringPtrInput
	Kind              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Privacy           pulumi.BoolPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	TargetDnsType     DnsTypePtrInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

func (o DomainOutput) AuthCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.AuthCode }).(pulumi.StringPtrOutput)
}

func (o DomainOutput) AutoRenew() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.BoolPtrOutput { return v.AutoRenew }).(pulumi.BoolPtrOutput)
}

func (o DomainOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o DomainOutput) DnsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.DnsType }).(pulumi.StringPtrOutput)
}

func (o DomainOutput) DnsZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.DnsZoneId }).(pulumi.StringPtrOutput)
}

func (o DomainOutput) DomainNotRenewableReasons() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringArrayOutput { return v.DomainNotRenewableReasons }).(pulumi.StringArrayOutput)
}

func (o DomainOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.ExpirationTime }).(pulumi.StringOutput)
}

func (o DomainOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o DomainOutput) LastRenewedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.LastRenewedTime }).(pulumi.StringOutput)
}

func (o DomainOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DomainOutput) ManagedHostNames() HostNameResponseArrayOutput {
	return o.ApplyT(func(v *Domain) HostNameResponseArrayOutput { return v.ManagedHostNames }).(HostNameResponseArrayOutput)
}

func (o DomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DomainOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringArrayOutput { return v.NameServers }).(pulumi.StringArrayOutput)
}

func (o DomainOutput) Privacy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.BoolPtrOutput { return v.Privacy }).(pulumi.BoolPtrOutput)
}

func (o DomainOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DomainOutput) ReadyForDnsRecordManagement() pulumi.BoolOutput {
	return o.ApplyT(func(v *Domain) pulumi.BoolOutput { return v.ReadyForDnsRecordManagement }).(pulumi.BoolOutput)
}

func (o DomainOutput) RegistrationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.RegistrationStatus }).(pulumi.StringOutput)
}

func (o DomainOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DomainOutput) TargetDnsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.TargetDnsType }).(pulumi.StringPtrOutput)
}

func (o DomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
}
