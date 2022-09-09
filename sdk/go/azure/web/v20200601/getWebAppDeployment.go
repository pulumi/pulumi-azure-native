


package v20200601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppDeployment(ctx *pulumi.Context, args *LookupWebAppDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupWebAppDeploymentResult, error) {
	var rv LookupWebAppDeploymentResult
	err := ctx.Invoke("azure-native:web/v20200601:getWebAppDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppDeploymentArgs struct {
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppDeploymentResult struct {
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

func LookupWebAppDeploymentOutput(ctx *pulumi.Context, args LookupWebAppDeploymentOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppDeploymentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppDeploymentResult, error) {
			args := v.(LookupWebAppDeploymentArgs)
			r, err := LookupWebAppDeployment(ctx, &args, opts...)
			var s LookupWebAppDeploymentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppDeploymentResultOutput)
}

type LookupWebAppDeploymentOutputArgs struct {
	Id                pulumi.StringInput `pulumi:"id"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppDeploymentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDeploymentArgs)(nil)).Elem()
}


type LookupWebAppDeploymentResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppDeploymentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDeploymentResult)(nil)).Elem()
}

func (o LookupWebAppDeploymentResultOutput) ToLookupWebAppDeploymentResultOutput() LookupWebAppDeploymentResultOutput {
	return o
}

func (o LookupWebAppDeploymentResultOutput) ToLookupWebAppDeploymentResultOutputWithContext(ctx context.Context) LookupWebAppDeploymentResultOutput {
	return o
}

func (o LookupWebAppDeploymentResultOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *bool { return v.Active }).(pulumi.BoolPtrOutput)
}

func (o LookupWebAppDeploymentResultOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.Author }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentResultOutput) AuthorEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.AuthorEmail }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentResultOutput) Deployer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.Deployer }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentResultOutput) Details() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.Details }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppDeploymentResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentResultOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppDeploymentResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDeploymentResultOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func (o LookupWebAppDeploymentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDeploymentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppDeploymentResultOutput{})
}
