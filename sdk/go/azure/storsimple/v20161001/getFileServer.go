


package v20161001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupFileServer(ctx *pulumi.Context, args *LookupFileServerArgs, opts ...pulumi.InvokeOption) (*LookupFileServerResult, error) {
	var rv LookupFileServerResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getFileServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileServerArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	FileServerName    string `pulumi:"fileServerName"`
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupFileServerResult struct {
	BackupScheduleGroupId string  `pulumi:"backupScheduleGroupId"`
	Description           *string `pulumi:"description"`
	DomainName            string  `pulumi:"domainName"`
	Id                    string  `pulumi:"id"`
	Name                  string  `pulumi:"name"`
	StorageDomainId       string  `pulumi:"storageDomainId"`
	Type                  string  `pulumi:"type"`
}

func LookupFileServerOutput(ctx *pulumi.Context, args LookupFileServerOutputArgs, opts ...pulumi.InvokeOption) LookupFileServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFileServerResult, error) {
			args := v.(LookupFileServerArgs)
			r, err := LookupFileServer(ctx, &args, opts...)
			var s LookupFileServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFileServerResultOutput)
}

type LookupFileServerOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	FileServerName    pulumi.StringInput `pulumi:"fileServerName"`
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFileServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileServerArgs)(nil)).Elem()
}


type LookupFileServerResultOutput struct{ *pulumi.OutputState }

func (LookupFileServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileServerResult)(nil)).Elem()
}

func (o LookupFileServerResultOutput) ToLookupFileServerResultOutput() LookupFileServerResultOutput {
	return o
}

func (o LookupFileServerResultOutput) ToLookupFileServerResultOutputWithContext(ctx context.Context) LookupFileServerResultOutput {
	return o
}

func (o LookupFileServerResultOutput) BackupScheduleGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileServerResult) string { return v.BackupScheduleGroupId }).(pulumi.StringOutput)
}

func (o LookupFileServerResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileServerResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFileServerResultOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileServerResult) string { return v.DomainName }).(pulumi.StringOutput)
}

func (o LookupFileServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileServerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFileServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileServerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFileServerResultOutput) StorageDomainId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileServerResult) string { return v.StorageDomainId }).(pulumi.StringOutput)
}

func (o LookupFileServerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileServerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFileServerResultOutput{})
}
