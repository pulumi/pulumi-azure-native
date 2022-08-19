


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkGroup(ctx *pulumi.Context, args *LookupNetworkGroupArgs, opts ...pulumi.InvokeOption) (*LookupNetworkGroupResult, error) {
	var rv LookupNetworkGroupResult
	err := ctx.Invoke("azure-native:network/v20220401preview:getNetworkGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkGroupArgs struct {
	NetworkGroupName   string `pulumi:"networkGroupName"`
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupNetworkGroupResult struct {
	Description       *string            `pulumi:"description"`
	Etag              string             `pulumi:"etag"`
	Id                string             `pulumi:"id"`
	MemberType        string             `pulumi:"memberType"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}

func LookupNetworkGroupOutput(ctx *pulumi.Context, args LookupNetworkGroupOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkGroupResult, error) {
			args := v.(LookupNetworkGroupArgs)
			r, err := LookupNetworkGroup(ctx, &args, opts...)
			var s LookupNetworkGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkGroupResultOutput)
}

type LookupNetworkGroupOutputArgs struct {
	NetworkGroupName   pulumi.StringInput `pulumi:"networkGroupName"`
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkGroupArgs)(nil)).Elem()
}


type LookupNetworkGroupResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkGroupResult)(nil)).Elem()
}

func (o LookupNetworkGroupResultOutput) ToLookupNetworkGroupResultOutput() LookupNetworkGroupResultOutput {
	return o
}

func (o LookupNetworkGroupResultOutput) ToLookupNetworkGroupResultOutputWithContext(ctx context.Context) LookupNetworkGroupResultOutput {
	return o
}

func (o LookupNetworkGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupNetworkGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkGroupResultOutput) MemberType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) string { return v.MemberType }).(pulumi.StringOutput)
}

func (o LookupNetworkGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNetworkGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupNetworkGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkGroupResultOutput{})
}
