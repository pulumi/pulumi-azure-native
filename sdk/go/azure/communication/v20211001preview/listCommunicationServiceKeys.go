


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListCommunicationServiceKeys(ctx *pulumi.Context, args *ListCommunicationServiceKeysArgs, opts ...pulumi.InvokeOption) (*ListCommunicationServiceKeysResult, error) {
	var rv ListCommunicationServiceKeysResult
	err := ctx.Invoke("azure-native:communication/v20211001preview:listCommunicationServiceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListCommunicationServiceKeysArgs struct {
	CommunicationServiceName string `pulumi:"communicationServiceName"`
	ResourceGroupName        string `pulumi:"resourceGroupName"`
}


type ListCommunicationServiceKeysResult struct {
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	PrimaryKey                *string `pulumi:"primaryKey"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	SecondaryKey              *string `pulumi:"secondaryKey"`
}

func ListCommunicationServiceKeysOutput(ctx *pulumi.Context, args ListCommunicationServiceKeysOutputArgs, opts ...pulumi.InvokeOption) ListCommunicationServiceKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListCommunicationServiceKeysResult, error) {
			args := v.(ListCommunicationServiceKeysArgs)
			r, err := ListCommunicationServiceKeys(ctx, &args, opts...)
			var s ListCommunicationServiceKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListCommunicationServiceKeysResultOutput)
}

type ListCommunicationServiceKeysOutputArgs struct {
	CommunicationServiceName pulumi.StringInput `pulumi:"communicationServiceName"`
	ResourceGroupName        pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListCommunicationServiceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCommunicationServiceKeysArgs)(nil)).Elem()
}


type ListCommunicationServiceKeysResultOutput struct{ *pulumi.OutputState }

func (ListCommunicationServiceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListCommunicationServiceKeysResult)(nil)).Elem()
}

func (o ListCommunicationServiceKeysResultOutput) ToListCommunicationServiceKeysResultOutput() ListCommunicationServiceKeysResultOutput {
	return o
}

func (o ListCommunicationServiceKeysResultOutput) ToListCommunicationServiceKeysResultOutputWithContext(ctx context.Context) ListCommunicationServiceKeysResultOutput {
	return o
}

func (o ListCommunicationServiceKeysResultOutput) PrimaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListCommunicationServiceKeysResult) *string { return v.PrimaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ListCommunicationServiceKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListCommunicationServiceKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o ListCommunicationServiceKeysResultOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListCommunicationServiceKeysResult) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ListCommunicationServiceKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListCommunicationServiceKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListCommunicationServiceKeysResultOutput{})
}
