


package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIntegrationAccountSession(ctx *pulumi.Context, args *LookupIntegrationAccountSessionArgs, opts ...pulumi.InvokeOption) (*LookupIntegrationAccountSessionResult, error) {
	var rv LookupIntegrationAccountSessionResult
	err := ctx.Invoke("azure-native:logic:getIntegrationAccountSession", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIntegrationAccountSessionArgs struct {
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	SessionName            string `pulumi:"sessionName"`
}


type LookupIntegrationAccountSessionResult struct {
	ChangedTime string            `pulumi:"changedTime"`
	Content     interface{}       `pulumi:"content"`
	CreatedTime string            `pulumi:"createdTime"`
	Id          string            `pulumi:"id"`
	Location    *string           `pulumi:"location"`
	Name        string            `pulumi:"name"`
	Tags        map[string]string `pulumi:"tags"`
	Type        string            `pulumi:"type"`
}

func LookupIntegrationAccountSessionOutput(ctx *pulumi.Context, args LookupIntegrationAccountSessionOutputArgs, opts ...pulumi.InvokeOption) LookupIntegrationAccountSessionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIntegrationAccountSessionResult, error) {
			args := v.(LookupIntegrationAccountSessionArgs)
			r, err := LookupIntegrationAccountSession(ctx, &args, opts...)
			var s LookupIntegrationAccountSessionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIntegrationAccountSessionResultOutput)
}

type LookupIntegrationAccountSessionOutputArgs struct {
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
	SessionName            pulumi.StringInput `pulumi:"sessionName"`
}

func (LookupIntegrationAccountSessionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountSessionArgs)(nil)).Elem()
}


type LookupIntegrationAccountSessionResultOutput struct{ *pulumi.OutputState }

func (LookupIntegrationAccountSessionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIntegrationAccountSessionResult)(nil)).Elem()
}

func (o LookupIntegrationAccountSessionResultOutput) ToLookupIntegrationAccountSessionResultOutput() LookupIntegrationAccountSessionResultOutput {
	return o
}

func (o LookupIntegrationAccountSessionResultOutput) ToLookupIntegrationAccountSessionResultOutputWithContext(ctx context.Context) LookupIntegrationAccountSessionResultOutput {
	return o
}

func (o LookupIntegrationAccountSessionResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSessionResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountSessionResultOutput) Content() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSessionResult) interface{} { return v.Content }).(pulumi.AnyOutput)
}

func (o LookupIntegrationAccountSessionResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSessionResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountSessionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSessionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountSessionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSessionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupIntegrationAccountSessionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSessionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIntegrationAccountSessionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSessionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupIntegrationAccountSessionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIntegrationAccountSessionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIntegrationAccountSessionResultOutput{})
}
