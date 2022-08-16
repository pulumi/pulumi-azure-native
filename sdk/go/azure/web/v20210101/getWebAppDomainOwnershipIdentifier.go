


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppDomainOwnershipIdentifier(ctx *pulumi.Context, args *LookupWebAppDomainOwnershipIdentifierArgs, opts ...pulumi.InvokeOption) (*LookupWebAppDomainOwnershipIdentifierResult, error) {
	var rv LookupWebAppDomainOwnershipIdentifierResult
	err := ctx.Invoke("azure-native:web/v20210101:getWebAppDomainOwnershipIdentifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppDomainOwnershipIdentifierArgs struct {
	DomainOwnershipIdentifierName string `pulumi:"domainOwnershipIdentifierName"`
	Name                          string `pulumi:"name"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
}


type LookupWebAppDomainOwnershipIdentifierResult struct {
	Id    string  `pulumi:"id"`
	Kind  *string `pulumi:"kind"`
	Name  string  `pulumi:"name"`
	Type  string  `pulumi:"type"`
	Value *string `pulumi:"value"`
}

func LookupWebAppDomainOwnershipIdentifierOutput(ctx *pulumi.Context, args LookupWebAppDomainOwnershipIdentifierOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppDomainOwnershipIdentifierResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppDomainOwnershipIdentifierResult, error) {
			args := v.(LookupWebAppDomainOwnershipIdentifierArgs)
			r, err := LookupWebAppDomainOwnershipIdentifier(ctx, &args, opts...)
			var s LookupWebAppDomainOwnershipIdentifierResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppDomainOwnershipIdentifierResultOutput)
}

type LookupWebAppDomainOwnershipIdentifierOutputArgs struct {
	DomainOwnershipIdentifierName pulumi.StringInput `pulumi:"domainOwnershipIdentifierName"`
	Name                          pulumi.StringInput `pulumi:"name"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppDomainOwnershipIdentifierOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDomainOwnershipIdentifierArgs)(nil)).Elem()
}


type LookupWebAppDomainOwnershipIdentifierResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppDomainOwnershipIdentifierResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDomainOwnershipIdentifierResult)(nil)).Elem()
}

func (o LookupWebAppDomainOwnershipIdentifierResultOutput) ToLookupWebAppDomainOwnershipIdentifierResultOutput() LookupWebAppDomainOwnershipIdentifierResultOutput {
	return o
}

func (o LookupWebAppDomainOwnershipIdentifierResultOutput) ToLookupWebAppDomainOwnershipIdentifierResultOutputWithContext(ctx context.Context) LookupWebAppDomainOwnershipIdentifierResultOutput {
	return o
}

func (o LookupWebAppDomainOwnershipIdentifierResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppDomainOwnershipIdentifierResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDomainOwnershipIdentifierResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppDomainOwnershipIdentifierResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWebAppDomainOwnershipIdentifierResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppDomainOwnershipIdentifierResultOutput{})
}
