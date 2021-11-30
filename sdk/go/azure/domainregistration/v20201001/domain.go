


package v20201001

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
	SystemData                  SystemDataResponseOutput    `pulumi:"systemData"`
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
	if args.AutoRenew == nil {
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
			Type: pulumi.String("azure-native:domainregistration/v20180201:Domain"),
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
	})
	opts = append(opts, aliases)
	var resource Domain
	err := ctx.RegisterResource("azure-native:domainregistration/v20201001:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("azure-native:domainregistration/v20201001:Domain", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*Domain)(nil))
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Domain)(nil))
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
}
