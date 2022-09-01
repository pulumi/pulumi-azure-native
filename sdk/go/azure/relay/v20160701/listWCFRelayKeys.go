


package v20160701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWCFRelayKeys(ctx *pulumi.Context, args *ListWCFRelayKeysArgs, opts ...pulumi.InvokeOption) (*ListWCFRelayKeysResult, error) {
	var rv ListWCFRelayKeysResult
	err := ctx.Invoke("azure-native:relay/v20160701:listWCFRelayKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWCFRelayKeysArgs struct {
	AuthorizationRuleName string `pulumi:"authorizationRuleName"`
	NamespaceName         string `pulumi:"namespaceName"`
	RelayName             string `pulumi:"relayName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type ListWCFRelayKeysResult struct {
	KeyName                   *string `pulumi:"keyName"`
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	PrimaryKey                *string `pulumi:"primaryKey"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	SecondaryKey              *string `pulumi:"secondaryKey"`
}

func ListWCFRelayKeysOutput(ctx *pulumi.Context, args ListWCFRelayKeysOutputArgs, opts ...pulumi.InvokeOption) ListWCFRelayKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWCFRelayKeysResult, error) {
			args := v.(ListWCFRelayKeysArgs)
			r, err := ListWCFRelayKeys(ctx, &args, opts...)
			var s ListWCFRelayKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWCFRelayKeysResultOutput)
}

type ListWCFRelayKeysOutputArgs struct {
	AuthorizationRuleName pulumi.StringInput `pulumi:"authorizationRuleName"`
	NamespaceName         pulumi.StringInput `pulumi:"namespaceName"`
	RelayName             pulumi.StringInput `pulumi:"relayName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWCFRelayKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWCFRelayKeysArgs)(nil)).Elem()
}


type ListWCFRelayKeysResultOutput struct{ *pulumi.OutputState }

func (ListWCFRelayKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWCFRelayKeysResult)(nil)).Elem()
}

func (o ListWCFRelayKeysResultOutput) ToListWCFRelayKeysResultOutput() ListWCFRelayKeysResultOutput {
	return o
}

func (o ListWCFRelayKeysResultOutput) ToListWCFRelayKeysResultOutputWithContext(ctx context.Context) ListWCFRelayKeysResultOutput {
	return o
}

func (o ListWCFRelayKeysResultOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWCFRelayKeysResult) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o ListWCFRelayKeysResultOutput) PrimaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWCFRelayKeysResult) *string { return v.PrimaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ListWCFRelayKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWCFRelayKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListWCFRelayKeysResultOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWCFRelayKeysResult) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ListWCFRelayKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWCFRelayKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWCFRelayKeysResultOutput{})
}
