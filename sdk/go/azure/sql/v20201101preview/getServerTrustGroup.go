


package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServerTrustGroup(ctx *pulumi.Context, args *LookupServerTrustGroupArgs, opts ...pulumi.InvokeOption) (*LookupServerTrustGroupResult, error) {
	var rv LookupServerTrustGroupResult
	err := ctx.Invoke("azure-native:sql/v20201101preview:getServerTrustGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerTrustGroupArgs struct {
	LocationName         string `pulumi:"locationName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	ServerTrustGroupName string `pulumi:"serverTrustGroupName"`
}


type LookupServerTrustGroupResult struct {
	GroupMembers []ServerInfoResponse `pulumi:"groupMembers"`
	Id           string               `pulumi:"id"`
	Name         string               `pulumi:"name"`
	TrustScopes  []string             `pulumi:"trustScopes"`
	Type         string               `pulumi:"type"`
}

func LookupServerTrustGroupOutput(ctx *pulumi.Context, args LookupServerTrustGroupOutputArgs, opts ...pulumi.InvokeOption) LookupServerTrustGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerTrustGroupResult, error) {
			args := v.(LookupServerTrustGroupArgs)
			r, err := LookupServerTrustGroup(ctx, &args, opts...)
			var s LookupServerTrustGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerTrustGroupResultOutput)
}

type LookupServerTrustGroupOutputArgs struct {
	LocationName         pulumi.StringInput `pulumi:"locationName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerTrustGroupName pulumi.StringInput `pulumi:"serverTrustGroupName"`
}

func (LookupServerTrustGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerTrustGroupArgs)(nil)).Elem()
}


type LookupServerTrustGroupResultOutput struct{ *pulumi.OutputState }

func (LookupServerTrustGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerTrustGroupResult)(nil)).Elem()
}

func (o LookupServerTrustGroupResultOutput) ToLookupServerTrustGroupResultOutput() LookupServerTrustGroupResultOutput {
	return o
}

func (o LookupServerTrustGroupResultOutput) ToLookupServerTrustGroupResultOutputWithContext(ctx context.Context) LookupServerTrustGroupResultOutput {
	return o
}

func (o LookupServerTrustGroupResultOutput) GroupMembers() ServerInfoResponseArrayOutput {
	return o.ApplyT(func(v LookupServerTrustGroupResult) []ServerInfoResponse { return v.GroupMembers }).(ServerInfoResponseArrayOutput)
}

func (o LookupServerTrustGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerTrustGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServerTrustGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerTrustGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServerTrustGroupResultOutput) TrustScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupServerTrustGroupResult) []string { return v.TrustScopes }).(pulumi.StringArrayOutput)
}

func (o LookupServerTrustGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerTrustGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerTrustGroupResultOutput{})
}
