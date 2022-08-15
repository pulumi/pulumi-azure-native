


package v20170815

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	var rv LookupPoolResult
	err := ctx.Invoke("azure-native:netapp/v20170815:getPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPoolArgs struct {
	AccountName       string `pulumi:"accountName"`
	PoolName          string `pulumi:"poolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPoolResult struct {
	Id                string      `pulumi:"id"`
	Location          string      `pulumi:"location"`
	Name              string      `pulumi:"name"`
	PoolId            string      `pulumi:"poolId"`
	ProvisioningState string      `pulumi:"provisioningState"`
	ServiceLevel      *string     `pulumi:"serviceLevel"`
	Size              *float64    `pulumi:"size"`
	Tags              interface{} `pulumi:"tags"`
	Type              string      `pulumi:"type"`
}


func (val *LookupPoolResult) Defaults() *LookupPoolResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceLevel) {
		serviceLevel_ := "Premium"
		tmp.ServiceLevel = &serviceLevel_
	}
	if isZero(tmp.Size) {
		size_ := 4398046511104.0
		tmp.Size = &size_
	}
	return &tmp
}

func LookupPoolOutput(ctx *pulumi.Context, args LookupPoolOutputArgs, opts ...pulumi.InvokeOption) LookupPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPoolResult, error) {
			args := v.(LookupPoolArgs)
			r, err := LookupPool(ctx, &args, opts...)
			var s LookupPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPoolResultOutput)
}

type LookupPoolOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	PoolName          pulumi.StringInput `pulumi:"poolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolArgs)(nil)).Elem()
}


type LookupPoolResultOutput struct{ *pulumi.OutputState }

func (LookupPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPoolResult)(nil)).Elem()
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutput() LookupPoolResultOutput {
	return o
}

func (o LookupPoolResultOutput) ToLookupPoolResultOutputWithContext(ctx context.Context) LookupPoolResultOutput {
	return o
}

func (o LookupPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.PoolId }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) ServiceLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *string { return v.ServiceLevel }).(pulumi.StringPtrOutput)
}

func (o LookupPoolResultOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *float64 { return v.Size }).(pulumi.Float64PtrOutput)
}

func (o LookupPoolResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPoolResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func (o LookupPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPoolResultOutput{})
}
