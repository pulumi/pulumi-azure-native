


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSiteInstanceDeploymentSlot(ctx *pulumi.Context, args *LookupSiteInstanceDeploymentSlotArgs, opts ...pulumi.InvokeOption) (*LookupSiteInstanceDeploymentSlotResult, error) {
	var rv LookupSiteInstanceDeploymentSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteInstanceDeploymentSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteInstanceDeploymentSlotArgs struct {
	Id                string `pulumi:"id"`
	InstanceId        string `pulumi:"instanceId"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupSiteInstanceDeploymentSlotResult struct {
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

func LookupSiteInstanceDeploymentSlotOutput(ctx *pulumi.Context, args LookupSiteInstanceDeploymentSlotOutputArgs, opts ...pulumi.InvokeOption) LookupSiteInstanceDeploymentSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteInstanceDeploymentSlotResult, error) {
			args := v.(LookupSiteInstanceDeploymentSlotArgs)
			r, err := LookupSiteInstanceDeploymentSlot(ctx, &args, opts...)
			var s LookupSiteInstanceDeploymentSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteInstanceDeploymentSlotResultOutput)
}

type LookupSiteInstanceDeploymentSlotOutputArgs struct {
	Id                pulumi.StringInput `pulumi:"id"`
	InstanceId        pulumi.StringInput `pulumi:"instanceId"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupSiteInstanceDeploymentSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteInstanceDeploymentSlotArgs)(nil)).Elem()
}


type LookupSiteInstanceDeploymentSlotResultOutput struct{ *pulumi.OutputState }

func (LookupSiteInstanceDeploymentSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteInstanceDeploymentSlotResult)(nil)).Elem()
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) ToLookupSiteInstanceDeploymentSlotResultOutput() LookupSiteInstanceDeploymentSlotResultOutput {
	return o
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) ToLookupSiteInstanceDeploymentSlotResultOutputWithContext(ctx context.Context) LookupSiteInstanceDeploymentSlotResultOutput {
	return o
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *bool { return v.Active }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Author }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) AuthorEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.AuthorEmail }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) Deployer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Deployer }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) Details() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Details }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteInstanceDeploymentSlotResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentSlotResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteInstanceDeploymentSlotResultOutput{})
}
