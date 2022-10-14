


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupShare(ctx *pulumi.Context, args *LookupShareArgs, opts ...pulumi.InvokeOption) (*LookupShareResult, error) {
	var rv LookupShareResult
	err := ctx.Invoke("azure-native:datashare/v20200901:getShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupShareArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ShareName         string `pulumi:"shareName"`
}


type LookupShareResult struct {
	CreatedAt         string             `pulumi:"createdAt"`
	Description       *string            `pulumi:"description"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	ShareKind         *string            `pulumi:"shareKind"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Terms             *string            `pulumi:"terms"`
	Type              string             `pulumi:"type"`
	UserEmail         string             `pulumi:"userEmail"`
	UserName          string             `pulumi:"userName"`
}

func LookupShareOutput(ctx *pulumi.Context, args LookupShareOutputArgs, opts ...pulumi.InvokeOption) LookupShareResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupShareResult, error) {
			args := v.(LookupShareArgs)
			r, err := LookupShare(ctx, &args, opts...)
			var s LookupShareResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupShareResultOutput)
}

type LookupShareOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ShareName         pulumi.StringInput `pulumi:"shareName"`
}

func (LookupShareOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareArgs)(nil)).Elem()
}


type LookupShareResultOutput struct{ *pulumi.OutputState }

func (LookupShareResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupShareResult)(nil)).Elem()
}

func (o LookupShareResultOutput) ToLookupShareResultOutput() LookupShareResultOutput {
	return o
}

func (o LookupShareResultOutput) ToLookupShareResultOutputWithContext(ctx context.Context) LookupShareResultOutput {
	return o
}

func (o LookupShareResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupShareResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupShareResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupShareResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupShareResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupShareResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupShareResultOutput) ShareKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupShareResult) *string { return v.ShareKind }).(pulumi.StringPtrOutput)
}

func (o LookupShareResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupShareResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupShareResultOutput) Terms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupShareResult) *string { return v.Terms }).(pulumi.StringPtrOutput)
}

func (o LookupShareResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupShareResultOutput) UserEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.UserEmail }).(pulumi.StringOutput)
}

func (o LookupShareResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupShareResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupShareResultOutput{})
}
