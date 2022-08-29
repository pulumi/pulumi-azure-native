


package v20191201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSubscriptionSecrets(ctx *pulumi.Context, args *ListSubscriptionSecretsArgs, opts ...pulumi.InvokeOption) (*ListSubscriptionSecretsResult, error) {
	var rv ListSubscriptionSecretsResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201preview:listSubscriptionSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSubscriptionSecretsArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
	Sid               string `pulumi:"sid"`
}


type ListSubscriptionSecretsResult struct {
	PrimaryKey   *string `pulumi:"primaryKey"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}

func ListSubscriptionSecretsOutput(ctx *pulumi.Context, args ListSubscriptionSecretsOutputArgs, opts ...pulumi.InvokeOption) ListSubscriptionSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSubscriptionSecretsResult, error) {
			args := v.(ListSubscriptionSecretsArgs)
			r, err := ListSubscriptionSecrets(ctx, &args, opts...)
			var s ListSubscriptionSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSubscriptionSecretsResultOutput)
}

type ListSubscriptionSecretsOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
	Sid               pulumi.StringInput `pulumi:"sid"`
}

func (ListSubscriptionSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSubscriptionSecretsArgs)(nil)).Elem()
}


type ListSubscriptionSecretsResultOutput struct{ *pulumi.OutputState }

func (ListSubscriptionSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSubscriptionSecretsResult)(nil)).Elem()
}

func (o ListSubscriptionSecretsResultOutput) ToListSubscriptionSecretsResultOutput() ListSubscriptionSecretsResultOutput {
	return o
}

func (o ListSubscriptionSecretsResultOutput) ToListSubscriptionSecretsResultOutputWithContext(ctx context.Context) ListSubscriptionSecretsResultOutput {
	return o
}

func (o ListSubscriptionSecretsResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSubscriptionSecretsResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListSubscriptionSecretsResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSubscriptionSecretsResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSubscriptionSecretsResultOutput{})
}
