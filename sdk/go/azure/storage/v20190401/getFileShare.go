


package v20190401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupFileShare(ctx *pulumi.Context, args *LookupFileShareArgs, opts ...pulumi.InvokeOption) (*LookupFileShareResult, error) {
	var rv LookupFileShareResult
	err := ctx.Invoke("azure-native:storage/v20190401:getFileShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileShareArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupFileShareResult struct {
	Etag             string            `pulumi:"etag"`
	Id               string            `pulumi:"id"`
	LastModifiedTime string            `pulumi:"lastModifiedTime"`
	Metadata         map[string]string `pulumi:"metadata"`
	Name             string            `pulumi:"name"`
	ShareQuota       *int              `pulumi:"shareQuota"`
	Type             string            `pulumi:"type"`
}

func LookupFileShareOutput(ctx *pulumi.Context, args LookupFileShareOutputArgs, opts ...pulumi.InvokeOption) LookupFileShareResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFileShareResult, error) {
			args := v.(LookupFileShareArgs)
			r, err := LookupFileShare(ctx, &args, opts...)
			var s LookupFileShareResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFileShareResultOutput)
}

type LookupFileShareOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupFileShareOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileShareArgs)(nil)).Elem()
}


type LookupFileShareResultOutput struct{ *pulumi.OutputState }

func (LookupFileShareResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileShareResult)(nil)).Elem()
}

func (o LookupFileShareResultOutput) ToLookupFileShareResultOutput() LookupFileShareResultOutput {
	return o
}

func (o LookupFileShareResultOutput) ToLookupFileShareResultOutputWithContext(ctx context.Context) LookupFileShareResultOutput {
	return o
}

func (o LookupFileShareResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFileShareResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o LookupFileShareResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFileShareResultOutput) ShareQuota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFileShareResult) *int { return v.ShareQuota }).(pulumi.IntPtrOutput)
}

func (o LookupFileShareResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFileShareResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFileShareResultOutput{})
}
