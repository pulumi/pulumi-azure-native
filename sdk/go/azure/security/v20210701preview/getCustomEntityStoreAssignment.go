


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomEntityStoreAssignment(ctx *pulumi.Context, args *LookupCustomEntityStoreAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupCustomEntityStoreAssignmentResult, error) {
	var rv LookupCustomEntityStoreAssignmentResult
	err := ctx.Invoke("azure-native:security/v20210701preview:getCustomEntityStoreAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomEntityStoreAssignmentArgs struct {
	CustomEntityStoreAssignmentName string `pulumi:"customEntityStoreAssignmentName"`
	ResourceGroupName               string `pulumi:"resourceGroupName"`
}


type LookupCustomEntityStoreAssignmentResult struct {
	EntityStoreDatabaseLink *string            `pulumi:"entityStoreDatabaseLink"`
	Id                      string             `pulumi:"id"`
	Name                    string             `pulumi:"name"`
	Principal               *string            `pulumi:"principal"`
	SystemData              SystemDataResponse `pulumi:"systemData"`
	Type                    string             `pulumi:"type"`
}

func LookupCustomEntityStoreAssignmentOutput(ctx *pulumi.Context, args LookupCustomEntityStoreAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupCustomEntityStoreAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomEntityStoreAssignmentResult, error) {
			args := v.(LookupCustomEntityStoreAssignmentArgs)
			r, err := LookupCustomEntityStoreAssignment(ctx, &args, opts...)
			var s LookupCustomEntityStoreAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomEntityStoreAssignmentResultOutput)
}

type LookupCustomEntityStoreAssignmentOutputArgs struct {
	CustomEntityStoreAssignmentName pulumi.StringInput `pulumi:"customEntityStoreAssignmentName"`
	ResourceGroupName               pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCustomEntityStoreAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomEntityStoreAssignmentArgs)(nil)).Elem()
}


type LookupCustomEntityStoreAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupCustomEntityStoreAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomEntityStoreAssignmentResult)(nil)).Elem()
}

func (o LookupCustomEntityStoreAssignmentResultOutput) ToLookupCustomEntityStoreAssignmentResultOutput() LookupCustomEntityStoreAssignmentResultOutput {
	return o
}

func (o LookupCustomEntityStoreAssignmentResultOutput) ToLookupCustomEntityStoreAssignmentResultOutputWithContext(ctx context.Context) LookupCustomEntityStoreAssignmentResultOutput {
	return o
}

func (o LookupCustomEntityStoreAssignmentResultOutput) EntityStoreDatabaseLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomEntityStoreAssignmentResult) *string { return v.EntityStoreDatabaseLink }).(pulumi.StringPtrOutput)
}

func (o LookupCustomEntityStoreAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomEntityStoreAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCustomEntityStoreAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomEntityStoreAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCustomEntityStoreAssignmentResultOutput) Principal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomEntityStoreAssignmentResult) *string { return v.Principal }).(pulumi.StringPtrOutput)
}

func (o LookupCustomEntityStoreAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCustomEntityStoreAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCustomEntityStoreAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomEntityStoreAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomEntityStoreAssignmentResultOutput{})
}
