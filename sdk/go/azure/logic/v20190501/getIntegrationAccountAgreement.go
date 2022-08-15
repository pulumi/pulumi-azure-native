


package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountAgreement(ctx *pulumi.Context, args *LookupIntegrationAccountAgreementArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountAgreementResult, error) {
	var rv LookupIntegrationAccountAgreementResult
	err := ctx.Invoke("azure-native:logic/v20190501:getIntegrationAccountAgreement", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountAgreementArgs struct {
	AgreementName          string `pulumi:"agreementName"`
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupIntegrationAccountAgreementResult struct {
	AgreementType string                   `pulumi:"agreementType"`
	ChangedTime   string                   `pulumi:"changedTime"`
	Content       AgreementContentResponse `pulumi:"content"`
	CreatedTime   string                   `pulumi:"createdTime"`
	GuestIdentity BusinessIdentityResponse `pulumi:"guestIdentity"`
	GuestPartner  string                   `pulumi:"guestPartner"`
	HostIdentity  BusinessIdentityResponse `pulumi:"hostIdentity"`
	HostPartner   string                   `pulumi:"hostPartner"`
	Id            string                   `pulumi:"id"`
	Location      *string                  `pulumi:"location"`
	Metadata      interface{}              `pulumi:"metadata"`
	Name          string                   `pulumi:"name"`
	Tags          map[string]string        `pulumi:"tags"`
	Type          string                   `pulumi:"type"`
}

func LookupIntegrationAccountAgreementOutput(ctx *pulumi.Context, args LookupIntegrationAccountAgreementOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationAccountAgreementResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationAccountAgreementResult, error) {
			args := v.(LookupIntegrationAccountAgreementArgs)
			r, err := LookupIntegrationAccountAgreement(ctx, &args, opts...)
			var s LookupIntegrationAccountAgreementResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationAccountAgreementResultOutput)
}

type LookupIntegrationAccountAgreementOutputArgs struct {
	AgreementName          pulumi.StringInput `pulumi:"agreementName"`
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIntegrationAccountAgreementOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountAgreementArgs)(nil)).Elem()
}


type LookupIntegrationAccountAgreementResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationAccountAgreementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountAgreementResult)(nil)).Elem()
}

func (o LookupIntegrationAccountAgreementResultOutput) ToLookupIntegrationAccountAgreementResultOutput() LookupIntegrationAccountAgreementResultOutput {
	return o
}

func (o LookupIntegrationAccountAgreementResultOutput) ToLookupIntegrationAccountAgreementResultOutputWithContext(ctx context.Context) LookupIntegrationAccountAgreementResultOutput {
	return o
}

func (o LookupIntegrationAccountAgreementResultOutput) AgreementType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAgreementResult) string { return v.AgreementType }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountAgreementResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAgreementResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountAgreementResultOutput) Content() AgreementContentResponseOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAgreementResult) AgreementContentResponse { return v.Content }).(AgreementContentResponseOutput)
}

func (o LookupIntegrationAccountAgreementResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAgreementResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountAgreementResultOutput) GuestIdentity() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAgreementResult) BusinessIdentityResponse { return v.GuestIdentity }).(BusinessIdentityResponseOutput)
}

func (o LookupIntegrationAccountAgreementResultOutput) GuestPartner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAgreementResult) string { return v.GuestPartner }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountAgreementResultOutput) HostIdentity() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAgreementResult) BusinessIdentityResponse { return v.HostIdentity }).(BusinessIdentityResponseOutput)
}

func (o LookupIntegrationAccountAgreementResultOutput) HostPartner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAgreementResult) string { return v.HostPartner }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountAgreementResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAgreementResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountAgreementResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAgreementResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountAgreementResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAgreementResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupIntegrationAccountAgreementResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAgreementResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountAgreementResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAgreementResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIntegrationAccountAgreementResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountAgreementResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationAccountAgreementResultOutput{})
}
