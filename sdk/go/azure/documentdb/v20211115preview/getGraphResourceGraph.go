


package v20211115preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGraphResourceGraph(ctx *pulumi.Context, args *LookupGraphResourceGraphArgs, opts ...pulumi.InvokeOption) (*LookupGraphResourceGraphResult, error) {
	var rv LookupGraphResourceGraphResult
	err := ctx.Invoke("azure-native:documentdb/v20211115preview:getGraphResourceGraph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGraphResourceGraphArgs struct {
	AccountName       string `pulumi:"accountName"`
	GraphName         string `pulumi:"graphName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupGraphResourceGraphResult struct {
	Id       string                                      `pulumi:"id"`
	Identity *ManagedServiceIdentityResponse             `pulumi:"identity"`
	Location *string                                     `pulumi:"location"`
	Name     string                                      `pulumi:"name"`
	Options  *GraphResourceGetPropertiesResponseOptions  `pulumi:"options"`
	Resource *GraphResourceGetPropertiesResponseResource `pulumi:"resource"`
	Tags     map[string]string                           `pulumi:"tags"`
	Type     string                                      `pulumi:"type"`
}

func LookupGraphResourceGraphOutput(ctx *pulumi.Context, args LookupGraphResourceGraphOutputArgs, opts ...pulumi.InvokeOption) LookupGraphResourceGraphResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGraphResourceGraphResult, error) {
			args := v.(LookupGraphResourceGraphArgs)
			r, err := LookupGraphResourceGraph(ctx, &args, opts...)
			var s LookupGraphResourceGraphResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGraphResourceGraphResultOutput)
}

type LookupGraphResourceGraphOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	GraphName         pulumi.StringInput `pulumi:"graphName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGraphResourceGraphOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphResourceGraphArgs)(nil)).Elem()
}


type LookupGraphResourceGraphResultOutput struct{ *pulumi.OutputState }

func (LookupGraphResourceGraphResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGraphResourceGraphResult)(nil)).Elem()
}

func (o LookupGraphResourceGraphResultOutput) ToLookupGraphResourceGraphResultOutput() LookupGraphResourceGraphResultOutput {
	return o
}

func (o LookupGraphResourceGraphResultOutput) ToLookupGraphResourceGraphResultOutputWithContext(ctx context.Context) LookupGraphResourceGraphResultOutput {
	return o
}

func (o LookupGraphResourceGraphResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphResourceGraphResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGraphResourceGraphResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupGraphResourceGraphResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupGraphResourceGraphResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGraphResourceGraphResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupGraphResourceGraphResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphResourceGraphResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGraphResourceGraphResultOutput) Options() GraphResourceGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v LookupGraphResourceGraphResult) *GraphResourceGetPropertiesResponseOptions { return v.Options }).(GraphResourceGetPropertiesResponseOptionsPtrOutput)
}

func (o LookupGraphResourceGraphResultOutput) Resource() GraphResourceGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v LookupGraphResourceGraphResult) *GraphResourceGetPropertiesResponseResource { return v.Resource }).(GraphResourceGetPropertiesResponseResourcePtrOutput)
}

func (o LookupGraphResourceGraphResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGraphResourceGraphResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGraphResourceGraphResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGraphResourceGraphResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGraphResourceGraphResultOutput{})
}
