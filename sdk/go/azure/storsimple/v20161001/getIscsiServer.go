


package v20161001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupIscsiServer(ctx *pulumi.Context, args *LookupIscsiServerArgs, opts ...pulumi.InvokeOption) (*LookupIscsiServerResult, error) {
	var rv LookupIscsiServerResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getIscsiServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIscsiServerArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	IscsiServerName   string `pulumi:"iscsiServerName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupIscsiServerResult struct {
	BackupScheduleGroupId string  `pulumi:"backupScheduleGroupId"`
	ChapId                *string `pulumi:"chapId"`
	Description           *string `pulumi:"description"`
	Id                    string  `pulumi:"id"`
	Name                  string  `pulumi:"name"`
	ReverseChapId         *string `pulumi:"reverseChapId"`
	StorageDomainId       string  `pulumi:"storageDomainId"`
	Type                  string  `pulumi:"type"`
}

func LookupIscsiServerOutput(ctx *pulumi.Context, args LookupIscsiServerOutputArgs, opts ...pulumi.InvokeOption) LookupIscsiServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIscsiServerResult, error) {
			args := v.(LookupIscsiServerArgs)
			r, err := LookupIscsiServer(ctx, &args, opts...)
			var s LookupIscsiServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIscsiServerResultOutput)
}

type LookupIscsiServerOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	IscsiServerName   pulumi.StringInput `pulumi:"iscsiServerName"`
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupIscsiServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIscsiServerArgs)(nil)).Elem()
}


type LookupIscsiServerResultOutput struct{ *pulumi.OutputState }

func (LookupIscsiServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIscsiServerResult)(nil)).Elem()
}

func (o LookupIscsiServerResultOutput) ToLookupIscsiServerResultOutput() LookupIscsiServerResultOutput {
	return o
}

func (o LookupIscsiServerResultOutput) ToLookupIscsiServerResultOutputWithContext(ctx context.Context) LookupIscsiServerResultOutput {
	return o
}

func (o LookupIscsiServerResultOutput) BackupScheduleGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiServerResult) string { return v.BackupScheduleGroupId }).(pulumi.StringOutput)
}

func (o LookupIscsiServerResultOutput) ChapId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIscsiServerResult) *string { return v.ChapId }).(pulumi.StringPtrOutput)
}

func (o LookupIscsiServerResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIscsiServerResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupIscsiServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIscsiServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiServerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIscsiServerResultOutput) ReverseChapId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIscsiServerResult) *string { return v.ReverseChapId }).(pulumi.StringPtrOutput)
}

func (o LookupIscsiServerResultOutput) StorageDomainId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiServerResult) string { return v.StorageDomainId }).(pulumi.StringOutput)
}

func (o LookupIscsiServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIscsiServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIscsiServerResultOutput{})
}
