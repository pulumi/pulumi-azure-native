


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStorageAccountServiceSAS(ctx *pulumi.Context, args *ListStorageAccountServiceSASArgs, opts ...pulumi.InvokeOption) (*ListStorageAccountServiceSASResult, error) {
	var rv ListStorageAccountServiceSASResult
	err := ctx.Invoke("azure-native:storage/v20200801preview:listStorageAccountServiceSAS", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStorageAccountServiceSASArgs struct {
	AccountName            string        `pulumi:"accountName"`
	CacheControl           *string       `pulumi:"cacheControl"`
	CanonicalizedResource  string        `pulumi:"canonicalizedResource"`
	ContentDisposition     *string       `pulumi:"contentDisposition"`
	ContentEncoding        *string       `pulumi:"contentEncoding"`
	ContentLanguage        *string       `pulumi:"contentLanguage"`
	ContentType            *string       `pulumi:"contentType"`
	IPAddressOrRange       *string       `pulumi:"iPAddressOrRange"`
	Identifier             *string       `pulumi:"identifier"`
	KeyToSign              *string       `pulumi:"keyToSign"`
	PartitionKeyEnd        *string       `pulumi:"partitionKeyEnd"`
	PartitionKeyStart      *string       `pulumi:"partitionKeyStart"`
	Permissions            *string       `pulumi:"permissions"`
	Protocols              *HttpProtocol `pulumi:"protocols"`
	Resource               *string       `pulumi:"resource"`
	ResourceGroupName      string        `pulumi:"resourceGroupName"`
	RowKeyEnd              *string       `pulumi:"rowKeyEnd"`
	RowKeyStart            *string       `pulumi:"rowKeyStart"`
	SharedAccessExpiryTime *string       `pulumi:"sharedAccessExpiryTime"`
	SharedAccessStartTime  *string       `pulumi:"sharedAccessStartTime"`
}


type ListStorageAccountServiceSASResult struct {
	ServiceSasToken string `pulumi:"serviceSasToken"`
}

func ListStorageAccountServiceSASOutput(ctx *pulumi.Context, args ListStorageAccountServiceSASOutputArgs, opts ...pulumi.InvokeOption) ListStorageAccountServiceSASResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStorageAccountServiceSASResult, error) {
			args := v.(ListStorageAccountServiceSASArgs)
			r, err := ListStorageAccountServiceSAS(ctx, &args, opts...)
			var s ListStorageAccountServiceSASResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStorageAccountServiceSASResultOutput)
}

type ListStorageAccountServiceSASOutputArgs struct {
	AccountName            pulumi.StringInput    `pulumi:"accountName"`
	CacheControl           pulumi.StringPtrInput `pulumi:"cacheControl"`
	CanonicalizedResource  pulumi.StringInput    `pulumi:"canonicalizedResource"`
	ContentDisposition     pulumi.StringPtrInput `pulumi:"contentDisposition"`
	ContentEncoding        pulumi.StringPtrInput `pulumi:"contentEncoding"`
	ContentLanguage        pulumi.StringPtrInput `pulumi:"contentLanguage"`
	ContentType            pulumi.StringPtrInput `pulumi:"contentType"`
	IPAddressOrRange       pulumi.StringPtrInput `pulumi:"iPAddressOrRange"`
	Identifier             pulumi.StringPtrInput `pulumi:"identifier"`
	KeyToSign              pulumi.StringPtrInput `pulumi:"keyToSign"`
	PartitionKeyEnd        pulumi.StringPtrInput `pulumi:"partitionKeyEnd"`
	PartitionKeyStart      pulumi.StringPtrInput `pulumi:"partitionKeyStart"`
	Permissions            pulumi.StringPtrInput `pulumi:"permissions"`
	Protocols              HttpProtocolPtrInput  `pulumi:"protocols"`
	Resource               pulumi.StringPtrInput `pulumi:"resource"`
	ResourceGroupName      pulumi.StringInput    `pulumi:"resourceGroupName"`
	RowKeyEnd              pulumi.StringPtrInput `pulumi:"rowKeyEnd"`
	RowKeyStart            pulumi.StringPtrInput `pulumi:"rowKeyStart"`
	SharedAccessExpiryTime pulumi.StringPtrInput `pulumi:"sharedAccessExpiryTime"`
	SharedAccessStartTime  pulumi.StringPtrInput `pulumi:"sharedAccessStartTime"`
}

func (ListStorageAccountServiceSASOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStorageAccountServiceSASArgs)(nil)).Elem()
}


type ListStorageAccountServiceSASResultOutput struct{ *pulumi.OutputState }

func (ListStorageAccountServiceSASResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStorageAccountServiceSASResult)(nil)).Elem()
}

func (o ListStorageAccountServiceSASResultOutput) ToListStorageAccountServiceSASResultOutput() ListStorageAccountServiceSASResultOutput {
	return o
}

func (o ListStorageAccountServiceSASResultOutput) ToListStorageAccountServiceSASResultOutputWithContext(ctx context.Context) ListStorageAccountServiceSASResultOutput {
	return o
}

func (o ListStorageAccountServiceSASResultOutput) ServiceSasToken() pulumi.StringOutput {
	return o.ApplyT(func(v ListStorageAccountServiceSASResult) string { return v.ServiceSasToken }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStorageAccountServiceSASResultOutput{})
}
