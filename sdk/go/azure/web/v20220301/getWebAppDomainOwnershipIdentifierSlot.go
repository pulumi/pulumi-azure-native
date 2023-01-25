


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppDomainOwnershipIdentifierSlot(ctx *pulumi.Context, args *LookupWebAppDomainOwnershipIdentifierSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppDomainOwnershipIdentifierSlotResult, error) {
	var rv LookupWebAppDomainOwnershipIdentifierSlotResult
	err := ctx.Invoke("azure-native:web/v20220301:getWebAppDomainOwnershipIdentifierSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppDomainOwnershipIdentifierSlotArgs struct {
	DomainOwnershipIdentifierName string `pulumi:"domainOwnershipIdentifierName"`
	Name                          string `pulumi:"name"`
	ResourceGroupName             string `pulumi:"resourceGroupName"`
	Slot                          string `pulumi:"slot"`
}


type LookupWebAppDomainOwnershipIdentifierSlotResult struct {
	Id    string  `pulumi:"id"`
	Kind  *string `pulumi:"kind"`
	Name  string  `pulumi:"name"`
	Type  string  `pulumi:"type"`
	Value *string `pulumi:"value"`
}

func LookupWebAppDomainOwnershipIdentifierSlotOutput(ctx *pulumi.Context, args LookupWebAppDomainOwnershipIdentifierSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppDomainOwnershipIdentifierSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppDomainOwnershipIdentifierSlotResult, error) {
			args := v.(LookupWebAppDomainOwnershipIdentifierSlotArgs)
			r, err := LookupWebAppDomainOwnershipIdentifierSlot(ctx, &args, opts...)
			var s LookupWebAppDomainOwnershipIdentifierSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppDomainOwnershipIdentifierSlotResultOutput)
}

type LookupWebAppDomainOwnershipIdentifierSlotOutputArgs struct {
	DomainOwnershipIdentifierName pulumi.StringInput `pulumi:"domainOwnershipIdentifierName"`
	Name                          pulumi.StringInput `pulumi:"name"`
	ResourceGroupName             pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot                          pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppDomainOwnershipIdentifierSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDomainOwnershipIdentifierSlotArgs)(nil)).Elem()
}


type LookupWebAppDomainOwnershipIdentifierSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppDomainOwnershipIdentifierSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDomainOwnershipIdentifierSlotResult)(nil)).Elem()
}

func (o LookupWebAppDomainOwnershipIdentifierSlotResultOutput) ToLookupWebAppDomainOwnershipIdentifierSlotResultOutput() LookupWebAppDomainOwnershipIdentifierSlotResultOutput {
	return o
}

func (o LookupWebAppDomainOwnershipIdentifierSlotResultOutput) ToLookupWebAppDomainOwnershipIdentifierSlotResultOutputWithContext(ctx context.Context) LookupWebAppDomainOwnershipIdentifierSlotResultOutput {
	return o
}

func (o LookupWebAppDomainOwnershipIdentifierSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppDomainOwnershipIdentifierSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDomainOwnershipIdentifierSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppDomainOwnershipIdentifierSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWebAppDomainOwnershipIdentifierSlotResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDomainOwnershipIdentifierSlotResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppDomainOwnershipIdentifierSlotResultOutput{})
}
