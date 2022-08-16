


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	var rv LookupDomainResult
	err := ctx.Invoke("azure-native:communication/v20211001preview:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainArgs struct {
	DomainName        string `pulumi:"domainName"`
	EmailServiceName  string `pulumi:"emailServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDomainResult struct {
	DataLocation           string                                      `pulumi:"dataLocation"`
	DomainManagement       string                                      `pulumi:"domainManagement"`
	FromSenderDomain       string                                      `pulumi:"fromSenderDomain"`
	Id                     string                                      `pulumi:"id"`
	Location               string                                      `pulumi:"location"`
	MailFromSenderDomain   string                                      `pulumi:"mailFromSenderDomain"`
	Name                   string                                      `pulumi:"name"`
	ProvisioningState      string                                      `pulumi:"provisioningState"`
	SystemData             SystemDataResponse                          `pulumi:"systemData"`
	Tags                   map[string]string                           `pulumi:"tags"`
	Type                   string                                      `pulumi:"type"`
	UserEngagementTracking *string                                     `pulumi:"userEngagementTracking"`
	ValidSenderUsernames   map[string]string                           `pulumi:"validSenderUsernames"`
	VerificationRecords    DomainPropertiesResponseVerificationRecords `pulumi:"verificationRecords"`
	VerificationStates     DomainPropertiesResponseVerificationStates  `pulumi:"verificationStates"`
}

func LookupDomainOutput(ctx *pulumi.Context, args LookupDomainOutputArgs, opts ...pulumi.InvokeOption) LookupDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainResult, error) {
			args := v.(LookupDomainArgs)
			r, err := LookupDomain(ctx, &args, opts...)
			var s LookupDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainResultOutput)
}

type LookupDomainOutputArgs struct {
	DomainName        pulumi.StringInput `pulumi:"domainName"`
	EmailServiceName  pulumi.StringInput `pulumi:"emailServiceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainArgs)(nil)).Elem()
}


type LookupDomainResultOutput struct{ *pulumi.OutputState }

func (LookupDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainResult)(nil)).Elem()
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutput() LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutputWithContext(ctx context.Context) LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) DataLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.DataLocation }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) DomainManagement() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.DomainManagement }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) FromSenderDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.FromSenderDomain }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) MailFromSenderDomain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.MailFromSenderDomain }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDomainResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDomainResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDomainResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupDomainResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDomainResultOutput) UserEngagementTracking() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.UserEngagementTracking }).(pulumi.StringPtrOutput)
}

func (o LookupDomainResultOutput) ValidSenderUsernames() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDomainResult) map[string]string { return v.ValidSenderUsernames }).(pulumi.StringMapOutput)
}

func (o LookupDomainResultOutput) VerificationRecords() DomainPropertiesResponseVerificationRecordsOutput {
	return o.ApplyT(func(v LookupDomainResult) DomainPropertiesResponseVerificationRecords { return v.VerificationRecords }).(DomainPropertiesResponseVerificationRecordsOutput)
}

func (o LookupDomainResultOutput) VerificationStates() DomainPropertiesResponseVerificationStatesOutput {
	return o.ApplyT(func(v LookupDomainResult) DomainPropertiesResponseVerificationStates { return v.VerificationStates }).(DomainPropertiesResponseVerificationStatesOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainResultOutput{})
}
