


package v20150501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupStorageAccount(ctx *pulumi.Context, args *LookupStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountResult, error) {
	var rv LookupStorageAccountResult
	err := ctx.Invoke("azure-native:storage/v20150501preview:getStorageAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStorageAccountResult struct {
	AccountType         *string               `pulumi:"accountType"`
	CreationTime        *string               `pulumi:"creationTime"`
	CustomDomain        *CustomDomainResponse `pulumi:"customDomain"`
	Id                  string                `pulumi:"id"`
	LastGeoFailoverTime *string               `pulumi:"lastGeoFailoverTime"`
	Location            string                `pulumi:"location"`
	Name                string                `pulumi:"name"`
	PrimaryEndpoints    *EndpointsResponse    `pulumi:"primaryEndpoints"`
	PrimaryLocation     *string               `pulumi:"primaryLocation"`
	ProvisioningState   *string               `pulumi:"provisioningState"`
	SecondaryEndpoints  *EndpointsResponse    `pulumi:"secondaryEndpoints"`
	SecondaryLocation   *string               `pulumi:"secondaryLocation"`
	StatusOfPrimary     *string               `pulumi:"statusOfPrimary"`
	StatusOfSecondary   *string               `pulumi:"statusOfSecondary"`
	Tags                map[string]string     `pulumi:"tags"`
	Type                string                `pulumi:"type"`
}

func LookupStorageAccountOutput(ctx *pulumi.Context, args LookupStorageAccountOutputArgs, opts ...pulumi.InvokeOption) LookupStorageAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStorageAccountResult, error) {
			args := v.(LookupStorageAccountArgs)
			r, err := LookupStorageAccount(ctx, &args, opts...)
			var s LookupStorageAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStorageAccountResultOutput)
}

type LookupStorageAccountOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStorageAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountArgs)(nil)).Elem()
}


type LookupStorageAccountResultOutput struct{ *pulumi.OutputState }

func (LookupStorageAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageAccountResult)(nil)).Elem()
}

func (o LookupStorageAccountResultOutput) ToLookupStorageAccountResultOutput() LookupStorageAccountResultOutput {
	return o
}

func (o LookupStorageAccountResultOutput) ToLookupStorageAccountResultOutputWithContext(ctx context.Context) LookupStorageAccountResultOutput {
	return o
}

func (o LookupStorageAccountResultOutput) AccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.AccountType }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) CustomDomain() CustomDomainResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *CustomDomainResponse { return v.CustomDomain }).(CustomDomainResponsePtrOutput)
}

func (o LookupStorageAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) LastGeoFailoverTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.LastGeoFailoverTime }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStorageAccountResultOutput) PrimaryEndpoints() EndpointsResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *EndpointsResponse { return v.PrimaryEndpoints }).(EndpointsResponsePtrOutput)
}

func (o LookupStorageAccountResultOutput) PrimaryLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.PrimaryLocation }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) SecondaryEndpoints() EndpointsResponsePtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *EndpointsResponse { return v.SecondaryEndpoints }).(EndpointsResponsePtrOutput)
}

func (o LookupStorageAccountResultOutput) SecondaryLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.SecondaryLocation }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) StatusOfPrimary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.StatusOfPrimary }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) StatusOfSecondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) *string { return v.StatusOfSecondary }).(pulumi.StringPtrOutput)
}

func (o LookupStorageAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupStorageAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageAccountResultOutput{})
}
