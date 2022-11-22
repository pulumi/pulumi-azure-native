


package testbase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTestBaseAccount(ctx *pulumi.Context, args *LookupTestBaseAccountArgs, opts ...pulumi.InvokeOption) (*LookupTestBaseAccountResult, error) {
	var rv LookupTestBaseAccountResult
	err := ctx.Invoke("azure-native:testbase:getTestBaseAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTestBaseAccountArgs struct {
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	TestBaseAccountName string `pulumi:"testBaseAccountName"`
}


type LookupTestBaseAccountResult struct {
	AccessLevel       string                     `pulumi:"accessLevel"`
	Etag              string                     `pulumi:"etag"`
	Id                string                     `pulumi:"id"`
	Location          string                     `pulumi:"location"`
	Name              string                     `pulumi:"name"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	Sku               TestBaseAccountSKUResponse `pulumi:"sku"`
	SystemData        SystemDataResponse         `pulumi:"systemData"`
	Tags              map[string]string          `pulumi:"tags"`
	Type              string                     `pulumi:"type"`
}

func LookupTestBaseAccountOutput(ctx *pulumi.Context, args LookupTestBaseAccountOutputArgs, opts ...pulumi.InvokeOption) LookupTestBaseAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTestBaseAccountResult, error) {
			args := v.(LookupTestBaseAccountArgs)
			r, err := LookupTestBaseAccount(ctx, &args, opts...)
			var s LookupTestBaseAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTestBaseAccountResultOutput)
}

type LookupTestBaseAccountOutputArgs struct {
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	TestBaseAccountName pulumi.StringInput `pulumi:"testBaseAccountName"`
}

func (LookupTestBaseAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTestBaseAccountArgs)(nil)).Elem()
}


type LookupTestBaseAccountResultOutput struct{ *pulumi.OutputState }

func (LookupTestBaseAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTestBaseAccountResult)(nil)).Elem()
}

func (o LookupTestBaseAccountResultOutput) ToLookupTestBaseAccountResultOutput() LookupTestBaseAccountResultOutput {
	return o
}

func (o LookupTestBaseAccountResultOutput) ToLookupTestBaseAccountResultOutputWithContext(ctx context.Context) LookupTestBaseAccountResultOutput {
	return o
}

func (o LookupTestBaseAccountResultOutput) AccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.AccessLevel }).(pulumi.StringOutput)
}

func (o LookupTestBaseAccountResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupTestBaseAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTestBaseAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupTestBaseAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTestBaseAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupTestBaseAccountResultOutput) Sku() TestBaseAccountSKUResponseOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) TestBaseAccountSKUResponse { return v.Sku }).(TestBaseAccountSKUResponseOutput)
}

func (o LookupTestBaseAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTestBaseAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupTestBaseAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestBaseAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTestBaseAccountResultOutput{})
}
