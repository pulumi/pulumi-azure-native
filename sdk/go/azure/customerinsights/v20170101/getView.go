


package v20170101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupView(ctx *pulumi.Context, args *LookupViewArgs, opts ...pulumi.InvokeOption) (*LookupViewResult, error) {
	var rv LookupViewResult
	err := ctx.Invoke("azure-native:customerinsights/v20170101:getView", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupViewArgs struct {
	HubName           string `pulumi:"hubName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	UserId            string `pulumi:"userId"`
	ViewName          string `pulumi:"viewName"`
}


type LookupViewResult struct {
	Changed     string            `pulumi:"changed"`
	Created     string            `pulumi:"created"`
	Definition  string            `pulumi:"definition"`
	DisplayName map[string]string `pulumi:"displayName"`
	Id          string            `pulumi:"id"`
	Name        string            `pulumi:"name"`
	TenantId    string            `pulumi:"tenantId"`
	Type        string            `pulumi:"type"`
	UserId      *string           `pulumi:"userId"`
	ViewName    string            `pulumi:"viewName"`
}

func LookupViewOutput(ctx *pulumi.Context, args LookupViewOutputArgs, opts ...pulumi.InvokeOption) LookupViewResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupViewResult, error) {
			args := v.(LookupViewArgs)
			r, err := LookupView(ctx, &args, opts...)
			var s LookupViewResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupViewResultOutput)
}

type LookupViewOutputArgs struct {
	HubName           pulumi.StringInput `pulumi:"hubName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	UserId            pulumi.StringInput `pulumi:"userId"`
	ViewName          pulumi.StringInput `pulumi:"viewName"`
}

func (LookupViewOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewArgs)(nil)).Elem()
}


type LookupViewResultOutput struct{ *pulumi.OutputState }

func (LookupViewResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewResult)(nil)).Elem()
}

func (o LookupViewResultOutput) ToLookupViewResultOutput() LookupViewResultOutput {
	return o
}

func (o LookupViewResultOutput) ToLookupViewResultOutputWithContext(ctx context.Context) LookupViewResultOutput {
	return o
}

func (o LookupViewResultOutput) Changed() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Changed }).(pulumi.StringOutput)
}

func (o LookupViewResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Created }).(pulumi.StringOutput)
}

func (o LookupViewResultOutput) Definition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Definition }).(pulumi.StringOutput)
}

func (o LookupViewResultOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupViewResult) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o LookupViewResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupViewResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupViewResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupViewResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupViewResultOutput) UserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.UserId }).(pulumi.StringPtrOutput)
}

func (o LookupViewResultOutput) ViewName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupViewResult) string { return v.ViewName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupViewResultOutput{})
}
