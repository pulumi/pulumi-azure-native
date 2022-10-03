


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSiteDeploymentSlot(ctx *pulumi.Context, args *LookupSiteDeploymentSlotArgs, opts ...pulumi.InvokeOption) (*LookupSiteDeploymentSlotResult, error) {
	var rv LookupSiteDeploymentSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteDeploymentSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteDeploymentSlotArgs struct {
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupSiteDeploymentSlotResult struct {
	Active      *bool             `pulumi:"active"`
	Author      *string           `pulumi:"author"`
	AuthorEmail *string           `pulumi:"authorEmail"`
	Deployer    *string           `pulumi:"deployer"`
	Details     *string           `pulumi:"details"`
	EndTime     *string           `pulumi:"endTime"`
	Id          *string           `pulumi:"id"`
	Kind        *string           `pulumi:"kind"`
	Location    string            `pulumi:"location"`
	Message     *string           `pulumi:"message"`
	Name        *string           `pulumi:"name"`
	StartTime   *string           `pulumi:"startTime"`
	Status      *int              `pulumi:"status"`
	Tags        map[string]string `pulumi:"tags"`
	Type        *string           `pulumi:"type"`
}

func LookupSiteDeploymentSlotOutput(ctx *pulumi.Context, args LookupSiteDeploymentSlotOutputArgs, opts ...pulumi.InvokeOption) LookupSiteDeploymentSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteDeploymentSlotResult, error) {
			args := v.(LookupSiteDeploymentSlotArgs)
			r, err := LookupSiteDeploymentSlot(ctx, &args, opts...)
			var s LookupSiteDeploymentSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteDeploymentSlotResultOutput)
}

type LookupSiteDeploymentSlotOutputArgs struct {
	Id                pulumi.StringInput `pulumi:"id"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupSiteDeploymentSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteDeploymentSlotArgs)(nil)).Elem()
}


type LookupSiteDeploymentSlotResultOutput struct{ *pulumi.OutputState }

func (LookupSiteDeploymentSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteDeploymentSlotResult)(nil)).Elem()
}

func (o LookupSiteDeploymentSlotResultOutput) ToLookupSiteDeploymentSlotResultOutput() LookupSiteDeploymentSlotResultOutput {
	return o
}

func (o LookupSiteDeploymentSlotResultOutput) ToLookupSiteDeploymentSlotResultOutputWithContext(ctx context.Context) LookupSiteDeploymentSlotResultOutput {
	return o
}

func (o LookupSiteDeploymentSlotResultOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentSlotResult) *bool { return v.Active }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteDeploymentSlotResultOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentSlotResult) *string { return v.Author }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentSlotResultOutput) AuthorEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentSlotResult) *string { return v.AuthorEmail }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentSlotResultOutput) Deployer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentSlotResult) *string { return v.Deployer }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentSlotResultOutput) Details() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentSlotResult) *string { return v.Details }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentSlotResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentSlotResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteDeploymentSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteDeploymentSlotResultOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentSlotResult) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentSlotResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentSlotResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentSlotResultOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentSlotResult) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func (o LookupSiteDeploymentSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteDeploymentSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteDeploymentSlotResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentSlotResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteDeploymentSlotResultOutput{})
}
