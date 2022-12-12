


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSession(ctx *pulumi.Context, args *LookupSessionArgs, opts ...pulumi.InvokeOption) (*LookupSessionResult, error) {
	var rv LookupSessionResult
	err := ctx.Invoke("azure-native:logic/v20160601:getSession", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSessionArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	SessionName            string `pulumi:"sessionName"`
}


type LookupSessionResult struct {
	ChangedTime string            `pulumi:"changedTime"`
	Content     interface{}       `pulumi:"content"`
	CreatedTime string            `pulumi:"createdTime"`
	Id          string            `pulumi:"id"`
	Location    *string           `pulumi:"location"`
	Name        string            `pulumi:"name"`
	Tags        map[string]string `pulumi:"tags"`
	Type        string            `pulumi:"type"`
}

func LookupSessionOutput(ctx *pulumi.Context, args LookupSessionOutputArgs, opts ...pulumi.InvokeOption) LookupSessionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSessionResult, error) {
			args := v.(LookupSessionArgs)
			r, err := LookupSession(ctx, &args, opts...)
			var s LookupSessionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSessionResultOutput)
}

type LookupSessionOutputArgs struct {
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	SessionName            pulumi.StringInput `pulumi:"sessionName"`
}

func (LookupSessionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSessionArgs)(nil)).Elem()
}


type LookupSessionResultOutput struct{ *pulumi.OutputState }

func (LookupSessionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSessionResult)(nil)).Elem()
}

func (o LookupSessionResultOutput) ToLookupSessionResultOutput() LookupSessionResultOutput {
	return o
}

func (o LookupSessionResultOutput) ToLookupSessionResultOutputWithContext(ctx context.Context) LookupSessionResultOutput {
	return o
}

func (o LookupSessionResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSessionResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o LookupSessionResultOutput) Content() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSessionResult) interface{} { return v.Content }).(pulumi.AnyOutput)
}

func (o LookupSessionResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSessionResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupSessionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSessionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSessionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSessionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSessionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSessionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSessionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSessionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSessionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSessionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSessionResultOutput{})
}
