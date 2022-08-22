


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStaticSiteUsers(ctx *pulumi.Context, args *ListStaticSiteUsersArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteUsersResult, error) {
	var rv ListStaticSiteUsersResult
	err := ctx.Invoke("azure-native:web/v20190801:listStaticSiteUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteUsersArgs struct {
	Authprovider      string `pulumi:"authprovider"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListStaticSiteUsersResult struct {
	NextLink string                              `pulumi:"nextLink"`
	Value    []StaticSiteUserARMResourceResponse `pulumi:"value"`
}

func ListStaticSiteUsersOutput(ctx *pulumi.Context, args ListStaticSiteUsersOutputArgs, opts ...pulumi.InvokeOption) ListStaticSiteUsersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStaticSiteUsersResult, error) {
			args := v.(ListStaticSiteUsersArgs)
			r, err := ListStaticSiteUsers(ctx, &args, opts...)
			var s ListStaticSiteUsersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStaticSiteUsersResultOutput)
}

type ListStaticSiteUsersOutputArgs struct {
	Authprovider      pulumi.StringInput `pulumi:"authprovider"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListStaticSiteUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteUsersArgs)(nil)).Elem()
}


type ListStaticSiteUsersResultOutput struct{ *pulumi.OutputState }

func (ListStaticSiteUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteUsersResult)(nil)).Elem()
}

func (o ListStaticSiteUsersResultOutput) ToListStaticSiteUsersResultOutput() ListStaticSiteUsersResultOutput {
	return o
}

func (o ListStaticSiteUsersResultOutput) ToListStaticSiteUsersResultOutputWithContext(ctx context.Context) ListStaticSiteUsersResultOutput {
	return o
}

func (o ListStaticSiteUsersResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteUsersResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListStaticSiteUsersResultOutput) Value() StaticSiteUserARMResourceResponseArrayOutput {
	return o.ApplyT(func(v ListStaticSiteUsersResult) []StaticSiteUserARMResourceResponse { return v.Value }).(StaticSiteUserARMResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStaticSiteUsersResultOutput{})
}
