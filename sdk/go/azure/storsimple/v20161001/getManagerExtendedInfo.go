


package v20161001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupManagerExtendedInfo(ctx *pulumi.Context, args *LookupManagerExtendedInfoArgs, opts ...pulumi.InvokeOption) (*LookupManagerExtendedInfoResult, error) {
	var rv LookupManagerExtendedInfoResult
	err := ctx.Invoke("azure-native:storsimple/v20161001:getManagerExtendedInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagerExtendedInfoArgs struct {
	ManagerName       string `pulumi:"managerName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupManagerExtendedInfoResult struct {
	Algorithm                   string  `pulumi:"algorithm"`
	EncryptionKey               *string `pulumi:"encryptionKey"`
	EncryptionKeyThumbprint     *string `pulumi:"encryptionKeyThumbprint"`
	Etag                        *string `pulumi:"etag"`
	Id                          string  `pulumi:"id"`
	IntegrityKey                string  `pulumi:"integrityKey"`
	Name                        string  `pulumi:"name"`
	PortalCertificateThumbprint *string `pulumi:"portalCertificateThumbprint"`
	Type                        string  `pulumi:"type"`
	Version                     *string `pulumi:"version"`
}

func LookupManagerExtendedInfoOutput(ctx *pulumi.Context, args LookupManagerExtendedInfoOutputArgs, opts ...pulumi.InvokeOption) LookupManagerExtendedInfoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagerExtendedInfoResult, error) {
			args := v.(LookupManagerExtendedInfoArgs)
			r, err := LookupManagerExtendedInfo(ctx, &args, opts...)
			var s LookupManagerExtendedInfoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagerExtendedInfoResultOutput)
}

type LookupManagerExtendedInfoOutputArgs struct {
	ManagerName       pulumi.StringInput `pulumi:"managerName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupManagerExtendedInfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagerExtendedInfoArgs)(nil)).Elem()
}


type LookupManagerExtendedInfoResultOutput struct{ *pulumi.OutputState }

func (LookupManagerExtendedInfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagerExtendedInfoResult)(nil)).Elem()
}

func (o LookupManagerExtendedInfoResultOutput) ToLookupManagerExtendedInfoResultOutput() LookupManagerExtendedInfoResultOutput {
	return o
}

func (o LookupManagerExtendedInfoResultOutput) ToLookupManagerExtendedInfoResultOutputWithContext(ctx context.Context) LookupManagerExtendedInfoResultOutput {
	return o
}

func (o LookupManagerExtendedInfoResultOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o LookupManagerExtendedInfoResultOutput) EncryptionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) *string { return v.EncryptionKey }).(pulumi.StringPtrOutput)
}

func (o LookupManagerExtendedInfoResultOutput) EncryptionKeyThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) *string { return v.EncryptionKeyThumbprint }).(pulumi.StringPtrOutput)
}

func (o LookupManagerExtendedInfoResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupManagerExtendedInfoResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagerExtendedInfoResultOutput) IntegrityKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) string { return v.IntegrityKey }).(pulumi.StringOutput)
}

func (o LookupManagerExtendedInfoResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagerExtendedInfoResultOutput) PortalCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) *string { return v.PortalCertificateThumbprint }).(pulumi.StringPtrOutput)
}

func (o LookupManagerExtendedInfoResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupManagerExtendedInfoResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagerExtendedInfoResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagerExtendedInfoResultOutput{})
}
