


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAgreement(ctx *pulumi.Context, args *LookupAgreementArgs, opts ...pulumi.InvokeOption) (*LookupAgreementResult, error) {
	var rv LookupAgreementResult
	err := ctx.Invoke("azure-native:logic/v20160601:getAgreement", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAgreementArgs struct {
	AgreementName          string `pulumi:"agreementName"`
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupAgreementResult struct {
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

func LookupAgreementOutput(ctx *pulumi.Context, args LookupAgreementOutputArgs, opts ...pulumi.InvokeOption) LookupAgreementResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAgreementResult, error) {
			args := v.(LookupAgreementArgs)
			r, err := LookupAgreement(ctx, &args, opts...)
			var s LookupAgreementResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAgreementResultOutput)
}

type LookupAgreementOutputArgs struct {
	AgreementName          pulumi.StringInput `pulumi:"agreementName"`
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAgreementOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgreementArgs)(nil)).Elem()
}


type LookupAgreementResultOutput struct{ *pulumi.OutputState }

func (LookupAgreementResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgreementResult)(nil)).Elem()
}

func (o LookupAgreementResultOutput) ToLookupAgreementResultOutput() LookupAgreementResultOutput {
	return o
}

func (o LookupAgreementResultOutput) ToLookupAgreementResultOutputWithContext(ctx context.Context) LookupAgreementResultOutput {
	return o
}

func (o LookupAgreementResultOutput) AgreementType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgreementResult) string { return v.AgreementType }).(pulumi.StringOutput)
}

func (o LookupAgreementResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgreementResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o LookupAgreementResultOutput) Content() AgreementContentResponseOutput {
	return o.ApplyT(func(v LookupAgreementResult) AgreementContentResponse { return v.Content }).(AgreementContentResponseOutput)
}

func (o LookupAgreementResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgreementResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupAgreementResultOutput) GuestIdentity() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v LookupAgreementResult) BusinessIdentityResponse { return v.GuestIdentity }).(BusinessIdentityResponseOutput)
}

func (o LookupAgreementResultOutput) GuestPartner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgreementResult) string { return v.GuestPartner }).(pulumi.StringOutput)
}

func (o LookupAgreementResultOutput) HostIdentity() BusinessIdentityResponseOutput {
	return o.ApplyT(func(v LookupAgreementResult) BusinessIdentityResponse { return v.HostIdentity }).(BusinessIdentityResponseOutput)
}

func (o LookupAgreementResultOutput) HostPartner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgreementResult) string { return v.HostPartner }).(pulumi.StringOutput)
}

func (o LookupAgreementResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgreementResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAgreementResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAgreementResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupAgreementResultOutput) Metadata() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupAgreementResult) interface{} { return v.Metadata }).(pulumi.AnyOutput)
}

func (o LookupAgreementResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgreementResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAgreementResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAgreementResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAgreementResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgreementResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAgreementResultOutput{})
}
