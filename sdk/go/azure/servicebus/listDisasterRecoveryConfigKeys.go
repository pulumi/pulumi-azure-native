


package servicebus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDisasterRecoveryConfigKeys(ctx *pulumi.Context, args *ListDisasterRecoveryConfigKeysArgs, opts ...pulumi.InvokeOption) (*ListDisasterRecoveryConfigKeysResult, error) {
	var rv ListDisasterRecoveryConfigKeysResult
	err := ctx.Invoke("azure-native:servicebus:listDisasterRecoveryConfigKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDisasterRecoveryConfigKeysArgs struct {
	Alias                 string `pulumi:"alias"`
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type ListDisasterRecoveryConfigKeysResult struct {
	AliasPrimaryConnectionString   string `pulumi:"aliasPrimaryConnectionString"`
	AliasSecondaryConnectionString string `pulumi:"aliasSecondaryConnectionString"`
	KeyName                        string `pulumi:"keyName"`
	PrimaryConnectionString        string `pulumi:"primaryConnectionString"`
	PrimaryKey                     string `pulumi:"primaryKey"`
	SecondaryConnectionString      string `pulumi:"secondaryConnectionString"`
	SecondaryKey                   string `pulumi:"secondaryKey"`
}

func ListDisasterRecoveryConfigKeysOutput(ctx *pulumi.Context, args ListDisasterRecoveryConfigKeysOutputArgs, opts ...pulumi.InvokeOption) ListDisasterRecoveryConfigKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDisasterRecoveryConfigKeysResult, error) {
			args := v.(ListDisasterRecoveryConfigKeysArgs)
			r, err := ListDisasterRecoveryConfigKeys(ctx, &args, opts...)
			var s ListDisasterRecoveryConfigKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDisasterRecoveryConfigKeysResultOutput)
}

type ListDisasterRecoveryConfigKeysOutputArgs struct {
	Alias                 pulumi.StringInput `pulumi:"alias"`
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListDisasterRecoveryConfigKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDisasterRecoveryConfigKeysArgs)(nil)).Elem()
}


type ListDisasterRecoveryConfigKeysResultOutput struct{ *pulumi.OutputState }

func (ListDisasterRecoveryConfigKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDisasterRecoveryConfigKeysResult)(nil)).Elem()
}

func (o ListDisasterRecoveryConfigKeysResultOutput) ToListDisasterRecoveryConfigKeysResultOutput() ListDisasterRecoveryConfigKeysResultOutput {
	return o
}

func (o ListDisasterRecoveryConfigKeysResultOutput) ToListDisasterRecoveryConfigKeysResultOutputWithContext(ctx context.Context) ListDisasterRecoveryConfigKeysResultOutput {
	return o
}

func (o ListDisasterRecoveryConfigKeysResultOutput) AliasPrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListDisasterRecoveryConfigKeysResult) string { return v.AliasPrimaryConnectionString }).(pulumi.StringOutput)
}

func (o ListDisasterRecoveryConfigKeysResultOutput) AliasSecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListDisasterRecoveryConfigKeysResult) string { return v.AliasSecondaryConnectionString }).(pulumi.StringOutput)
}

func (o ListDisasterRecoveryConfigKeysResultOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v ListDisasterRecoveryConfigKeysResult) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o ListDisasterRecoveryConfigKeysResultOutput) PrimaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListDisasterRecoveryConfigKeysResult) string { return v.PrimaryConnectionString }).(pulumi.StringOutput)
}

func (o ListDisasterRecoveryConfigKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDisasterRecoveryConfigKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o ListDisasterRecoveryConfigKeysResultOutput) SecondaryConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v ListDisasterRecoveryConfigKeysResult) string { return v.SecondaryConnectionString }).(pulumi.StringOutput)
}

func (o ListDisasterRecoveryConfigKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListDisasterRecoveryConfigKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDisasterRecoveryConfigKeysResultOutput{})
}
