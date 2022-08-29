


package v20160515

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("azure-native:devtestlab/v20160515:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserArgs struct {
	Expand            *string `pulumi:"expand"`
	LabName           string  `pulumi:"labName"`
	Name              string  `pulumi:"name"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupUserResult struct {
	CreatedDate       string                   `pulumi:"createdDate"`
	Id                string                   `pulumi:"id"`
	Identity          *UserIdentityResponse    `pulumi:"identity"`
	Location          *string                  `pulumi:"location"`
	Name              string                   `pulumi:"name"`
	ProvisioningState *string                  `pulumi:"provisioningState"`
	SecretStore       *UserSecretStoreResponse `pulumi:"secretStore"`
	Tags              map[string]string        `pulumi:"tags"`
	Type              string                   `pulumi:"type"`
	UniqueIdentifier  *string                  `pulumi:"uniqueIdentifier"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

type LookupUserOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	LabName           pulumi.StringInput    `pulumi:"labName"`
	Name              pulumi.StringInput    `pulumi:"name"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}


type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) Identity() UserIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupUserResult) *UserIdentityResponse { return v.Identity }).(UserIdentityResponsePtrOutput)
}

func (o LookupUserResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupUserResultOutput) SecretStore() UserSecretStoreResponsePtrOutput {
	return o.ApplyT(func(v LookupUserResult) *UserSecretStoreResponse { return v.SecretStore }).(UserSecretStoreResponsePtrOutput)
}

func (o LookupUserResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupUserResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupUserResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
