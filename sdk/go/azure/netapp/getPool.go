


package netapp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPool(ctx *pulumi.Context, args *LookupPoolArgs, opts ...pulumi.InvokeOption) (*LookupPoolResult, error) {
	var rv LookupPoolResult
	err := ctx.Invoke("azure-native:netapp:getPool", args, &rv, opts...)
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
	Id                      string            `pulumi:"id"`
	Location                string            `pulumi:"location"`
	Name                    string            `pulumi:"name"`
	PoolId                  string            `pulumi:"poolId"`
	ProvisioningState       string            `pulumi:"provisioningState"`
	QosType                 *string           `pulumi:"qosType"`
	ServiceLevel            string            `pulumi:"serviceLevel"`
	Size                    float64           `pulumi:"size"`
	Tags                    map[string]string `pulumi:"tags"`
	TotalThroughputMibps    float64           `pulumi:"totalThroughputMibps"`
	Type                    string            `pulumi:"type"`
	UtilizedThroughputMibps float64           `pulumi:"utilizedThroughputMibps"`
}


func (val *LookupPoolResult) Defaults() *LookupPoolResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.QosType) {
		qosType_ := "Auto"
		tmp.QosType = &qosType_
	}
	if isZero(tmp.ServiceLevel) {
		tmp.ServiceLevel = "Premium"
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

func (o LookupPoolResultOutput) QosType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPoolResult) *string { return v.QosType }).(pulumi.StringPtrOutput)
}

func (o LookupPoolResultOutput) ServiceLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.ServiceLevel }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPoolResult) float64 { return v.Size }).(pulumi.Float64Output)
}

func (o LookupPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPoolResultOutput) TotalThroughputMibps() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPoolResult) float64 { return v.TotalThroughputMibps }).(pulumi.Float64Output)
}

func (o LookupPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPoolResultOutput) UtilizedThroughputMibps() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPoolResult) float64 { return v.UtilizedThroughputMibps }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupPoolResultOutput{})
}
