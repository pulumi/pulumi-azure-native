


package v20201201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSshPublicKey(ctx *pulumi.Context, args *LookupSshPublicKeyArgs, opts ...pulumi.InvokeOption) (*LookupSshPublicKeyResult, error) {
	var rv LookupSshPublicKeyResult
	err := ctx.Invoke("azure-native:compute/v20201201:getSshPublicKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSshPublicKeyArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SshPublicKeyName  string `pulumi:"sshPublicKeyName"`
}


type LookupSshPublicKeyResult struct {
	Id        string            `pulumi:"id"`
	Location  string            `pulumi:"location"`
	Name      string            `pulumi:"name"`
	PublicKey *string           `pulumi:"publicKey"`
	Tags      map[string]string `pulumi:"tags"`
	Type      string            `pulumi:"type"`
}

func LookupSshPublicKeyOutput(ctx *pulumi.Context, args LookupSshPublicKeyOutputArgs, opts ...pulumi.InvokeOption) LookupSshPublicKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSshPublicKeyResult, error) {
			args := v.(LookupSshPublicKeyArgs)
			r, err := LookupSshPublicKey(ctx, &args, opts...)
			var s LookupSshPublicKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSshPublicKeyResultOutput)
}

type LookupSshPublicKeyOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SshPublicKeyName  pulumi.StringInput `pulumi:"sshPublicKeyName"`
}

func (LookupSshPublicKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSshPublicKeyArgs)(nil)).Elem()
}


type LookupSshPublicKeyResultOutput struct{ *pulumi.OutputState }

func (LookupSshPublicKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSshPublicKeyResult)(nil)).Elem()
}

func (o LookupSshPublicKeyResultOutput) ToLookupSshPublicKeyResultOutput() LookupSshPublicKeyResultOutput {
	return o
}

func (o LookupSshPublicKeyResultOutput) ToLookupSshPublicKeyResultOutputWithContext(ctx context.Context) LookupSshPublicKeyResultOutput {
	return o
}

func (o LookupSshPublicKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSshPublicKeyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSshPublicKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSshPublicKeyResultOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) *string { return v.PublicKey }).(pulumi.StringPtrOutput)
}

func (o LookupSshPublicKeyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSshPublicKeyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSshPublicKeyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSshPublicKeyResultOutput{})
}
