


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Domain struct {
	pulumi.CustomResourceState

	DataLocation           pulumi.StringOutput                               `pulumi:"dataLocation"`
	DomainManagement       pulumi.StringOutput                               `pulumi:"domainManagement"`
	FromSenderDomain       pulumi.StringOutput                               `pulumi:"fromSenderDomain"`
	Location               pulumi.StringOutput                               `pulumi:"location"`
	MailFromSenderDomain   pulumi.StringOutput                               `pulumi:"mailFromSenderDomain"`
	Name                   pulumi.StringOutput                               `pulumi:"name"`
	ProvisioningState      pulumi.StringOutput                               `pulumi:"provisioningState"`
	SystemData             SystemDataResponseOutput                          `pulumi:"systemData"`
	Tags                   pulumi.StringMapOutput                            `pulumi:"tags"`
	Type                   pulumi.StringOutput                               `pulumi:"type"`
	UserEngagementTracking pulumi.StringPtrOutput                            `pulumi:"userEngagementTracking"`
	ValidSenderUsernames   pulumi.StringMapOutput                            `pulumi:"validSenderUsernames"`
	VerificationRecords    DomainPropertiesResponseVerificationRecordsOutput `pulumi:"verificationRecords"`
	VerificationStates     DomainPropertiesResponseVerificationStatesOutput  `pulumi:"verificationStates"`
}


func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainManagement == nil {
		return nil, errors.New("invalid value for required argument 'DomainManagement'")
	}
	if args.EmailServiceName == nil {
		return nil, errors.New("invalid value for required argument 'EmailServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:communication:Domain"),
		},
		{
			Type: pulumi.String("azure-native:communication/v20220701preview:Domain"),
		},
	})
	opts = append(opts, aliases)
	var resource Domain
	err := ctx.RegisterResource("azure-native:communication/v20211001preview:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("azure-native:communication/v20211001preview:Domain", name, id, state, &resource, opts...)
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
	DomainManagement       string            `pulumi:"domainManagement"`
	DomainName             *string           `pulumi:"domainName"`
	EmailServiceName       string            `pulumi:"emailServiceName"`
	Location               *string           `pulumi:"location"`
	ResourceGroupName      string            `pulumi:"resourceGroupName"`
	Tags                   map[string]string `pulumi:"tags"`
	UserEngagementTracking *string           `pulumi:"userEngagementTracking"`
	ValidSenderUsernames   map[string]string `pulumi:"validSenderUsernames"`
}


type DomainArgs struct {
	DomainManagement       pulumi.StringInput
	DomainName             pulumi.StringPtrInput
	EmailServiceName       pulumi.StringInput
	Location               pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
	UserEngagementTracking pulumi.StringPtrInput
	ValidSenderUsernames   pulumi.StringMapInput
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

func (o DomainOutput) DataLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.DataLocation }).(pulumi.StringOutput)
}

func (o DomainOutput) DomainManagement() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.DomainManagement }).(pulumi.StringOutput)
}

func (o DomainOutput) FromSenderDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.FromSenderDomain }).(pulumi.StringOutput)
}

func (o DomainOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DomainOutput) MailFromSenderDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.MailFromSenderDomain }).(pulumi.StringOutput)
}

func (o DomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DomainOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Domain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DomainOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o DomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o DomainOutput) UserEngagementTracking() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.UserEngagementTracking }).(pulumi.StringPtrOutput)
}

func (o DomainOutput) ValidSenderUsernames() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringMapOutput { return v.ValidSenderUsernames }).(pulumi.StringMapOutput)
}

func (o DomainOutput) VerificationRecords() DomainPropertiesResponseVerificationRecordsOutput {
	return o.ApplyT(func(v *Domain) DomainPropertiesResponseVerificationRecordsOutput { return v.VerificationRecords }).(DomainPropertiesResponseVerificationRecordsOutput)
}

func (o DomainOutput) VerificationStates() DomainPropertiesResponseVerificationStatesOutput {
	return o.ApplyT(func(v *Domain) DomainPropertiesResponseVerificationStatesOutput { return v.VerificationStates }).(DomainPropertiesResponseVerificationStatesOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainOutput{})
}
