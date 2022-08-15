


package v20211110preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	var rv LookupAccountResult
	err := ctx.Invoke("azure-native:videoindexer/v20211110preview:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAccountResult struct {
	AccountId           *string                             `pulumi:"accountId"`
	AccountName         string                              `pulumi:"accountName"`
	Id                  string                              `pulumi:"id"`
	Identity            *ManagedServiceIdentityResponse     `pulumi:"identity"`
	Location            string                              `pulumi:"location"`
	MediaServices       *MediaServicesForPutRequestResponse `pulumi:"mediaServices"`
	Name                string                              `pulumi:"name"`
	ProvisioningState   string                              `pulumi:"provisioningState"`
	SystemData          SystemDataResponse                  `pulumi:"systemData"`
	Tags                map[string]string                   `pulumi:"tags"`
	TenantId            string                              `pulumi:"tenantId"`
	TotalSecondsIndexed int                                 `pulumi:"totalSecondsIndexed"`
	Type                string                              `pulumi:"type"`
}


func (val *LookupAccountResult) Defaults() *LookupAccountResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AccountId) {
		accountId_ := "00000000-0000-0000-0000-000000000000"
		tmp.AccountId = &accountId_
	}
	return &tmp
}

func LookupAccountOutput(ctx *pulumi.Context, args LookupAccountOutputArgs, opts ...pulumi.InvokeOption) LookupAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccountResult, error) {
			args := v.(LookupAccountArgs)
			r, err := LookupAccount(ctx, &args, opts...)
			var s LookupAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccountResultOutput)
}

type LookupAccountOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountArgs)(nil)).Elem()
}


type LookupAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountResult)(nil)).Elem()
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutput() LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutputWithContext(ctx context.Context) LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

func (o LookupAccountResultOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) MediaServices() MediaServicesForPutRequestResponsePtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *MediaServicesForPutRequestResponse { return v.MediaServices }).(MediaServicesForPutRequestResponsePtrOutput)
}

func (o LookupAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAccountResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupAccountResultOutput) TotalSecondsIndexed() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAccountResult) int { return v.TotalSecondsIndexed }).(pulumi.IntOutput)
}

func (o LookupAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountResultOutput{})
}
