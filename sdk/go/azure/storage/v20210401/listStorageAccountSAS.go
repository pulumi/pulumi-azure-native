


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStorageAccountSAS(ctx *pulumi.Context, args *ListStorageAccountSASArgs, opts ...pulumi.InvokeOption) (*ListStorageAccountSASResult, error) {
	var rv ListStorageAccountSASResult
	err := ctx.Invoke("azure-native:storage/v20210401:listStorageAccountSAS", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStorageAccountSASArgs struct {
	AccountName            string        `pulumi:"accountName"`
	IPAddressOrRange       *string       `pulumi:"iPAddressOrRange"`
	KeyToSign              *string       `pulumi:"keyToSign"`
	Permissions            string        `pulumi:"permissions"`
	Protocols              *HttpProtocol `pulumi:"protocols"`
	ResourceGroupName      string        `pulumi:"resourceGroupName"`
	ResourceTypes          string        `pulumi:"resourceTypes"`
	Services               string        `pulumi:"services"`
	SharedAccessExpiryTime string        `pulumi:"sharedAccessExpiryTime"`
	SharedAccessStartTime  *string       `pulumi:"sharedAccessStartTime"`
}


type ListStorageAccountSASResult struct {
	AccountSasToken string `pulumi:"accountSasToken"`
}

func ListStorageAccountSASOutput(ctx *pulumi.Context, args ListStorageAccountSASOutputArgs, opts ...pulumi.InvokeOption) ListStorageAccountSASResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStorageAccountSASResult, error) {
			args := v.(ListStorageAccountSASArgs)
			r, err := ListStorageAccountSAS(ctx, &args, opts...)
			var s ListStorageAccountSASResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStorageAccountSASResultOutput)
}

type ListStorageAccountSASOutputArgs struct {
	AccountName            pulumi.StringInput    `pulumi:"accountName"`
	IPAddressOrRange       pulumi.StringPtrInput `pulumi:"iPAddressOrRange"`
	KeyToSign              pulumi.StringPtrInput `pulumi:"keyToSign"`
	Permissions            pulumi.StringInput    `pulumi:"permissions"`
	Protocols              HttpProtocolPtrInput  `pulumi:"protocols"`
	ResourceGroupName      pulumi.StringInput    `pulumi:"resourceGroupName"`
	ResourceTypes          pulumi.StringInput    `pulumi:"resourceTypes"`
	Services               pulumi.StringInput    `pulumi:"services"`
	SharedAccessExpiryTime pulumi.StringInput    `pulumi:"sharedAccessExpiryTime"`
	SharedAccessStartTime  pulumi.StringPtrInput `pulumi:"sharedAccessStartTime"`
}

func (ListStorageAccountSASOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStorageAccountSASArgs)(nil)).Elem()
}


type ListStorageAccountSASResultOutput struct{ *pulumi.OutputState }

func (ListStorageAccountSASResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStorageAccountSASResult)(nil)).Elem()
}

func (o ListStorageAccountSASResultOutput) ToListStorageAccountSASResultOutput() ListStorageAccountSASResultOutput {
	return o
}

func (o ListStorageAccountSASResultOutput) ToListStorageAccountSASResultOutputWithContext(ctx context.Context) ListStorageAccountSASResultOutput {
	return o
}

func (o ListStorageAccountSASResultOutput) AccountSasToken() pulumi.StringOutput {
	return o.ApplyT(func(v ListStorageAccountSASResult) string { return v.AccountSasToken }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStorageAccountSASResultOutput{})
}
