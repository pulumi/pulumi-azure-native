


package v20201201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLedger(ctx *pulumi.Context, args *LookupLedgerArgs, opts ...pulumi.InvokeOption) (*LookupLedgerResult, error) {
	var rv LookupLedgerResult
	err := ctx.Invoke("azure-native:confidentialledger/v20201201preview:getLedger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLedgerArgs struct {
	LedgerName        string `pulumi:"ledgerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLedgerResult struct {
	Id         string                   `pulumi:"id"`
	Location   *string                  `pulumi:"location"`
	Name       string                   `pulumi:"name"`
	Properties LedgerPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse       `pulumi:"systemData"`
	Tags       map[string]string        `pulumi:"tags"`
	Type       string                   `pulumi:"type"`
}

func LookupLedgerOutput(ctx *pulumi.Context, args LookupLedgerOutputArgs, opts ...pulumi.InvokeOption) LookupLedgerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLedgerResult, error) {
			args := v.(LookupLedgerArgs)
			r, err := LookupLedger(ctx, &args, opts...)
			var s LookupLedgerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLedgerResultOutput)
}

type LookupLedgerOutputArgs struct {
	LedgerName        pulumi.StringInput `pulumi:"ledgerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLedgerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLedgerArgs)(nil)).Elem()
}


type LookupLedgerResultOutput struct{ *pulumi.OutputState }

func (LookupLedgerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLedgerResult)(nil)).Elem()
}

func (o LookupLedgerResultOutput) ToLookupLedgerResultOutput() LookupLedgerResultOutput {
	return o
}

func (o LookupLedgerResultOutput) ToLookupLedgerResultOutputWithContext(ctx context.Context) LookupLedgerResultOutput {
	return o
}

func (o LookupLedgerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLedgerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLedgerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLedgerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupLedgerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLedgerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLedgerResultOutput) Properties() LedgerPropertiesResponseOutput {
	return o.ApplyT(func(v LookupLedgerResult) LedgerPropertiesResponse { return v.Properties }).(LedgerPropertiesResponseOutput)
}

func (o LookupLedgerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLedgerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupLedgerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLedgerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupLedgerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLedgerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLedgerResultOutput{})
}
