


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSiteDeployment(ctx *pulumi.Context, args *LookupSiteDeploymentArgs, opts ...pulumi.InvokeOption) (*LookupSiteDeploymentResult, error) {
	var rv LookupSiteDeploymentResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteDeployment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteDeploymentArgs struct {
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSiteDeploymentResult struct {
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

func LookupSiteDeploymentOutput(ctx *pulumi.Context, args LookupSiteDeploymentOutputArgs, opts ...pulumi.InvokeOption) LookupSiteDeploymentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteDeploymentResult, error) {
			args := v.(LookupSiteDeploymentArgs)
			r, err := LookupSiteDeployment(ctx, &args, opts...)
			var s LookupSiteDeploymentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteDeploymentResultOutput)
}

type LookupSiteDeploymentOutputArgs struct {
	Id                pulumi.StringInput `pulumi:"id"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSiteDeploymentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteDeploymentArgs)(nil)).Elem()
}


type LookupSiteDeploymentResultOutput struct{ *pulumi.OutputState }

func (LookupSiteDeploymentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteDeploymentResult)(nil)).Elem()
}

func (o LookupSiteDeploymentResultOutput) ToLookupSiteDeploymentResultOutput() LookupSiteDeploymentResultOutput {
	return o
}

func (o LookupSiteDeploymentResultOutput) ToLookupSiteDeploymentResultOutputWithContext(ctx context.Context) LookupSiteDeploymentResultOutput {
	return o
}

func (o LookupSiteDeploymentResultOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentResult) *bool { return v.Active }).(pulumi.BoolPtrOutput)
}

func (o LookupSiteDeploymentResultOutput) Author() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentResult) *string { return v.Author }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentResultOutput) AuthorEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentResult) *string { return v.AuthorEmail }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentResultOutput) Deployer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentResult) *string { return v.Deployer }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentResultOutput) Details() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentResult) *string { return v.Details }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentResultOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentResult) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteDeploymentResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteDeploymentResultOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentResult) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentResultOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentResult) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o LookupSiteDeploymentResultOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentResult) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func (o LookupSiteDeploymentResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteDeploymentResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteDeploymentResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteDeploymentResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteDeploymentResultOutput{})
}
