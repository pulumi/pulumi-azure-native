


package v20200320

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListPrivateCloudAdminCredentials(ctx *pulumi.Context, args *ListPrivateCloudAdminCredentialsArgs, opts ...pulumi.InvokeOption) (*ListPrivateCloudAdminCredentialsResult, error) {
	var rv ListPrivateCloudAdminCredentialsResult
	err := ctx.Invoke("azure-native:avs/v20200320:listPrivateCloudAdminCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPrivateCloudAdminCredentialsArgs struct {
	PrivateCloudName  string `pulumi:"privateCloudName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListPrivateCloudAdminCredentialsResult struct {
	NsxtPassword    string `pulumi:"nsxtPassword"`
	NsxtUsername    string `pulumi:"nsxtUsername"`
	VcenterPassword string `pulumi:"vcenterPassword"`
	VcenterUsername string `pulumi:"vcenterUsername"`
}

func ListPrivateCloudAdminCredentialsOutput(ctx *pulumi.Context, args ListPrivateCloudAdminCredentialsOutputArgs, opts ...pulumi.InvokeOption) ListPrivateCloudAdminCredentialsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListPrivateCloudAdminCredentialsResult, error) {
			args := v.(ListPrivateCloudAdminCredentialsArgs)
			r, err := ListPrivateCloudAdminCredentials(ctx, &args, opts...)
			var s ListPrivateCloudAdminCredentialsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListPrivateCloudAdminCredentialsResultOutput)
}

type ListPrivateCloudAdminCredentialsOutputArgs struct {
	PrivateCloudName  pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListPrivateCloudAdminCredentialsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPrivateCloudAdminCredentialsArgs)(nil)).Elem()
}


type ListPrivateCloudAdminCredentialsResultOutput struct{ *pulumi.OutputState }

func (ListPrivateCloudAdminCredentialsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPrivateCloudAdminCredentialsResult)(nil)).Elem()
}

func (o ListPrivateCloudAdminCredentialsResultOutput) ToListPrivateCloudAdminCredentialsResultOutput() ListPrivateCloudAdminCredentialsResultOutput {
	return o
}

func (o ListPrivateCloudAdminCredentialsResultOutput) ToListPrivateCloudAdminCredentialsResultOutputWithContext(ctx context.Context) ListPrivateCloudAdminCredentialsResultOutput {
	return o
}

func (o ListPrivateCloudAdminCredentialsResultOutput) NsxtPassword() pulumi.StringOutput {
	return o.ApplyT(func(v ListPrivateCloudAdminCredentialsResult) string { return v.NsxtPassword }).(pulumi.StringOutput)
}

func (o ListPrivateCloudAdminCredentialsResultOutput) NsxtUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ListPrivateCloudAdminCredentialsResult) string { return v.NsxtUsername }).(pulumi.StringOutput)
}

func (o ListPrivateCloudAdminCredentialsResultOutput) VcenterPassword() pulumi.StringOutput {
	return o.ApplyT(func(v ListPrivateCloudAdminCredentialsResult) string { return v.VcenterPassword }).(pulumi.StringOutput)
}

func (o ListPrivateCloudAdminCredentialsResultOutput) VcenterUsername() pulumi.StringOutput {
	return o.ApplyT(func(v ListPrivateCloudAdminCredentialsResult) string { return v.VcenterUsername }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPrivateCloudAdminCredentialsResultOutput{})
}
