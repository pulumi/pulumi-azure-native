


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSiteInstanceDeployment(ctx *pulumi.Context, args *LookupSiteInstanceDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupSiteInstanceDeploymentResult, error) {
	var rv LookupSiteInstanceDeploymentResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteInstanceDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteInstanceDeploymentArgs struct {
	Id                string `pulumi:"id"`
	InstanceId        string `pulumi:"instanceId"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSiteInstanceDeploymentResult struct {
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

func LookupSiteInstanceDeploymentOutput(ctx *pulumi.Context, args LookupSiteInstanceDeploymentOutputArgs, opts ...pulumi.InvokeOption) LookupSiteInstanceDeploymentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteInstanceDeploymentResult, error) {
			args := v.(LookupSiteInstanceDeploymentArgs)
			r, err := LookupSiteInstanceDeployment(ctx, &args, opts...)
			var s LookupSiteInstanceDeploymentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteInstanceDeploymentResultOutput)
}

type LookupSiteInstanceDeploymentOutputArgs struct {
	Id                pulumi.StringInput `pulumi:"id"`
	InstanceId        pulumi.StringInput `pulumi:"instanceId"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSiteInstanceDeploymentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteInstanceDeploymentArgs)(nil)).Elem()
}


type LookupSiteInstanceDeploymentResultOutput struct{ *pulumi.OutputState }

func (LookupSiteInstanceDeploymentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteInstanceDeploymentResult)(nil)).Elem()
}

func (o LookupSiteInstanceDeploymentResultOutput) ToLookupSiteInstanceDeploymentResultOutput() LookupSiteInstanceDeploymentResultOutput {
	return o
}

func (o LookupSiteInstanceDeploymentResultOutput) ToLookupSiteInstanceDeploymentResultOutputWithContext(ctx context.Context) LookupSiteInstanceDeploymentResultOutput {
	return o
}

func (o LookupSiteInstanceDeploymentResultOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentResult) *bool { return v.Active }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteInstanceDeploymentResultOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentResult) *string { return v.Author }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentResultOutput) AuthorEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentResult) *string { return v.AuthorEmail }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentResultOutput) Deployer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentResult) *string { return v.Deployer }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentResultOutput) Details() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentResult) *string { return v.Details }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteInstanceDeploymentResultOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentResult) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o LookupSiteInstanceDeploymentResultOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentResult) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func (o LookupSiteInstanceDeploymentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteInstanceDeploymentResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteInstanceDeploymentResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteInstanceDeploymentResultOutput{})
}
