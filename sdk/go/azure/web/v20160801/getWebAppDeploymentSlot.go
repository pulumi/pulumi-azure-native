


package v20160801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppDeploymentSlot(ctx *pulumi.Context, args *LookupWebAppDeploymentSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppDeploymentSlotResult, error) {
	var rv LookupWebAppDeploymentSlotResult
	err := ctx.Invoke("azure-native:web/v20160801:getWebAppDeploymentSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppDeploymentSlotArgs struct {
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupWebAppDeploymentSlotResult struct {
	Active      *bool   `pulumi:"active"`
	Author      *string `pulumi:"author"`
	AuthorEmail *string `pulumi:"authorEmail"`
	Deployer    *string `pulumi:"deployer"`
	Details     *string `pulumi:"details"`
	EndTime     *string `pulumi:"endTime"`
	Id          string  `pulumi:"id"`
	Kind        *string `pulumi:"kind"`
	Message     *string `pulumi:"message"`
	Name        string  `pulumi:"name"`
	StartTime   *string `pulumi:"startTime"`
	Status      *int    `pulumi:"status"`
	Type        string  `pulumi:"type"`
}

func LookupWebAppDeploymentSlotOutput(ctx *pulumi.Context, args LookupWebAppDeploymentSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppDeploymentSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppDeploymentSlotResult, error) {
			args := v.(LookupWebAppDeploymentSlotArgs)
			r, err := LookupWebAppDeploymentSlot(ctx, &args, opts...)
			var s LookupWebAppDeploymentSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppDeploymentSlotResultOutput)
}

type LookupWebAppDeploymentSlotOutputArgs struct {
	Id                pulumi.StringInput `pulumi:"id"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppDeploymentSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDeploymentSlotArgs)(nil)).Elem()
}


type LookupWebAppDeploymentSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppDeploymentSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDeploymentSlotResult)(nil)).Elem()
}

func (o LookupWebAppDeploymentSlotResultOutput) ToLookupWebAppDeploymentSlotResultOutput() LookupWebAppDeploymentSlotResultOutput {
	return o
}

func (o LookupWebAppDeploymentSlotResultOutput) ToLookupWebAppDeploymentSlotResultOutputWithContext(ctx context.Context) LookupWebAppDeploymentSlotResultOutput {
	return o
}

func (o LookupWebAppDeploymentSlotResultOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *bool { return v.Active }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppDeploymentSlotResultOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.Author }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentSlotResultOutput) AuthorEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.AuthorEmail }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentSlotResultOutput) Deployer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.Deployer }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentSlotResultOutput) Details() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.Details }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentSlotResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppDeploymentSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentSlotResultOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppDeploymentSlotResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentSlotResultOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func (o LookupWebAppDeploymentSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppDeploymentSlotResultOutput{})
}
