


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDomainOwnershipIdentifier(ctx *pulumi.Context, args *LookupDomainOwnershipIdentifierArgs, opts ...pulumi.InvokeOption) (*LookupDomainOwnershipIdentifierResult, error) {
	var rv LookupDomainOwnershipIdentifierResult
	err := ctx.Invoke("azure-native:domainregistration/v20190801:getDomainOwnershipIdentifier", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainOwnershipIdentifierArgs struct {
	DomainName        string `pulumi:"domainName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDomainOwnershipIdentifierResult struct {
	Id          string  `pulumi:"id"`
	Kind        *string `pulumi:"kind"`
	Name        string  `pulumi:"name"`
	OwnershipId *string `pulumi:"ownershipId"`
	Type        string  `pulumi:"type"`
}

func LookupDomainOwnershipIdentifierOutput(ctx *pulumi.Context, args LookupDomainOwnershipIdentifierOutputArgs, opts ...pulumi.InvokeOption) LookupDomainOwnershipIdentifierResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainOwnershipIdentifierResult, error) {
			args := v.(LookupDomainOwnershipIdentifierArgs)
			r, err := LookupDomainOwnershipIdentifier(ctx, &args, opts...)
			var s LookupDomainOwnershipIdentifierResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainOwnershipIdentifierResultOutput)
}

type LookupDomainOwnershipIdentifierOutputArgs struct {
	DomainName        pulumi.StringInput `pulumi:"domainName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDomainOwnershipIdentifierOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainOwnershipIdentifierArgs)(nil)).Elem()
}


type LookupDomainOwnershipIdentifierResultOutput struct{ *pulumi.OutputState }

func (LookupDomainOwnershipIdentifierResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainOwnershipIdentifierResult)(nil)).Elem()
}

func (o LookupDomainOwnershipIdentifierResultOutput) ToLookupDomainOwnershipIdentifierResultOutput() LookupDomainOwnershipIdentifierResultOutput {
	return o
}

func (o LookupDomainOwnershipIdentifierResultOutput) ToLookupDomainOwnershipIdentifierResultOutputWithContext(ctx context.Context) LookupDomainOwnershipIdentifierResultOutput {
	return o
}

func (o LookupDomainOwnershipIdentifierResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainOwnershipIdentifierResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDomainOwnershipIdentifierResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainOwnershipIdentifierResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupDomainOwnershipIdentifierResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainOwnershipIdentifierResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDomainOwnershipIdentifierResultOutput) OwnershipId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainOwnershipIdentifierResult) *string { return v.OwnershipId }).(pulumi.StringPtrOutput)
}

func (o LookupDomainOwnershipIdentifierResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainOwnershipIdentifierResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainOwnershipIdentifierResultOutput{})
}
