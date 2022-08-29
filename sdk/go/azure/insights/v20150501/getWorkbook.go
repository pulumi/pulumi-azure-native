


package v20150501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkbook(ctx *pulumi.Context, args *LookupWorkbookArgs, opts ...pulumi.InvokeOption) (*LookupWorkbookResult, error) {
	var rv LookupWorkbookResult
	err := ctx.Invoke("azure-native:insights/v20150501:getWorkbook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWorkbookArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupWorkbookResult struct {
	Category         string            `pulumi:"category"`
	Id               string            `pulumi:"id"`
	Kind             *string           `pulumi:"kind"`
	Location         *string           `pulumi:"location"`
	Name             string            `pulumi:"name"`
	SerializedData   string            `pulumi:"serializedData"`
	SharedTypeKind   string            `pulumi:"sharedTypeKind"`
	SourceResourceId *string           `pulumi:"sourceResourceId"`
	Tags             map[string]string `pulumi:"tags"`
	TimeModified     string            `pulumi:"timeModified"`
	Type             string            `pulumi:"type"`
	UserId           string            `pulumi:"userId"`
	Version          *string           `pulumi:"version"`
	WorkbookId       string            `pulumi:"workbookId"`
}


func (val *LookupWorkbookResult) Defaults() *LookupWorkbookResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.SharedTypeKind) {
		tmp.SharedTypeKind = "shared"
	}
	return &tmp
}

func LookupWorkbookOutput(ctx *pulumi.Context, args LookupWorkbookOutputArgs, opts ...pulumi.InvokeOption) LookupWorkbookResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkbookResult, error) {
			args := v.(LookupWorkbookArgs)
			r, err := LookupWorkbook(ctx, &args, opts...)
			var s LookupWorkbookResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkbookResultOutput)
}

type LookupWorkbookOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupWorkbookOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkbookArgs)(nil)).Elem()
}


type LookupWorkbookResultOutput struct{ *pulumi.OutputState }

func (LookupWorkbookResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkbookResult)(nil)).Elem()
}

func (o LookupWorkbookResultOutput) ToLookupWorkbookResultOutput() LookupWorkbookResultOutput {
	return o
}

func (o LookupWorkbookResultOutput) ToLookupWorkbookResultOutputWithContext(ctx context.Context) LookupWorkbookResultOutput {
	return o
}

func (o LookupWorkbookResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWorkbookResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupWorkbookResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) SerializedData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.SerializedData }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) SharedTypeKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.SharedTypeKind }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupWorkbookResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkbookResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWorkbookResultOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.TimeModified }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.UserId }).(pulumi.StringOutput)
}

func (o LookupWorkbookResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkbookResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func (o LookupWorkbookResultOutput) WorkbookId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkbookResult) string { return v.WorkbookId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkbookResultOutput{})
}
