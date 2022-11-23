


package v20170601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAccessControlRecord(ctx *pulumi.Context, args *LookupAccessControlRecordArgs, opts ...pulumi.InvokeOption) (*LookupAccessControlRecordResult, error) {
	var rv LookupAccessControlRecordResult
	err := ctx.Invoke("azure-native:storsimple/v20170601:getAccessControlRecord", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccessControlRecordArgs struct {
	AccessControlRecordName string `pulumi:"accessControlRecordName"`
	ManagerName             string `pulumi:"managerName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
}


type LookupAccessControlRecordResult struct {
	Id            string  `pulumi:"id"`
	InitiatorName string  `pulumi:"initiatorName"`
	Kind          *string `pulumi:"kind"`
	Name          string  `pulumi:"name"`
	Type          string  `pulumi:"type"`
	VolumeCount   int     `pulumi:"volumeCount"`
}

func LookupAccessControlRecordOutput(ctx *pulumi.Context, args LookupAccessControlRecordOutputArgs, opts ...pulumi.InvokeOption) LookupAccessControlRecordResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessControlRecordResult, error) {
			args := v.(LookupAccessControlRecordArgs)
			r, err := LookupAccessControlRecord(ctx, &args, opts...)
			var s LookupAccessControlRecordResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessControlRecordResultOutput)
}

type LookupAccessControlRecordOutputArgs struct {
	AccessControlRecordName pulumi.StringInput `pulumi:"accessControlRecordName"`
	ManagerName             pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAccessControlRecordOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessControlRecordArgs)(nil)).Elem()
}


type LookupAccessControlRecordResultOutput struct{ *pulumi.OutputState }

func (LookupAccessControlRecordResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessControlRecordResult)(nil)).Elem()
}

func (o LookupAccessControlRecordResultOutput) ToLookupAccessControlRecordResultOutput() LookupAccessControlRecordResultOutput {
	return o
}

func (o LookupAccessControlRecordResultOutput) ToLookupAccessControlRecordResultOutputWithContext(ctx context.Context) LookupAccessControlRecordResultOutput {
	return o
}

func (o LookupAccessControlRecordResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessControlRecordResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccessControlRecordResultOutput) InitiatorName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessControlRecordResult) string { return v.InitiatorName }).(pulumi.StringOutput)
}

func (o LookupAccessControlRecordResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccessControlRecordResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupAccessControlRecordResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessControlRecordResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAccessControlRecordResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessControlRecordResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAccessControlRecordResultOutput) VolumeCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAccessControlRecordResult) int { return v.VolumeCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessControlRecordResultOutput{})
}
