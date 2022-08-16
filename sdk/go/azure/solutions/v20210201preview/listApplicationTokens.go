


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListApplicationTokens(ctx *pulumi.Context, args *ListApplicationTokensArgs, opts ...pulumi.InvokeOption) (*ListApplicationTokensResult, error) {
	var rv ListApplicationTokensResult
	err := ctx.Invoke("azure-native:solutions/v20210201preview:listApplicationTokens", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListApplicationTokensArgs struct {
	ApplicationName        string   `pulumi:"applicationName"`
	AuthorizationAudience  *string  `pulumi:"authorizationAudience"`
	ResourceGroupName      string   `pulumi:"resourceGroupName"`
	UserAssignedIdentities []string `pulumi:"userAssignedIdentities"`
}


type ListApplicationTokensResult struct {
	Value []ManagedIdentityTokenResponse `pulumi:"value"`
}

func ListApplicationTokensOutput(ctx *pulumi.Context, args ListApplicationTokensOutputArgs, opts ...pulumi.InvokeOption) ListApplicationTokensResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListApplicationTokensResult, error) {
			args := v.(ListApplicationTokensArgs)
			r, err := ListApplicationTokens(ctx, &args, opts...)
			var s ListApplicationTokensResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListApplicationTokensResultOutput)
}

type ListApplicationTokensOutputArgs struct {
	ApplicationName        pulumi.StringInput      `pulumi:"applicationName"`
	AuthorizationAudience  pulumi.StringPtrInput   `pulumi:"authorizationAudience"`
	ResourceGroupName      pulumi.StringInput      `pulumi:"resourceGroupName"`
	UserAssignedIdentities pulumi.StringArrayInput `pulumi:"userAssignedIdentities"`
}

func (ListApplicationTokensOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplicationTokensArgs)(nil)).Elem()
}


type ListApplicationTokensResultOutput struct{ *pulumi.OutputState }

func (ListApplicationTokensResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListApplicationTokensResult)(nil)).Elem()
}

func (o ListApplicationTokensResultOutput) ToListApplicationTokensResultOutput() ListApplicationTokensResultOutput {
	return o
}

func (o ListApplicationTokensResultOutput) ToListApplicationTokensResultOutputWithContext(ctx context.Context) ListApplicationTokensResultOutput {
	return o
}

func (o ListApplicationTokensResultOutput) Value() ManagedIdentityTokenResponseArrayOutput {
	return o.ApplyT(func(v ListApplicationTokensResult) []ManagedIdentityTokenResponse { return v.Value }).(ManagedIdentityTokenResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListApplicationTokensResultOutput{})
}
