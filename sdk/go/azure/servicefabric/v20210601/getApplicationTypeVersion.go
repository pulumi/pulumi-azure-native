


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApplicationTypeVersion(ctx *pulumi.Context, args *LookupApplicationTypeVersionArgs, opts ...pulumi.InvokeOption) (*LookupApplicationTypeVersionResult, error) {
	var rv LookupApplicationTypeVersionResult
	err := ctx.Invoke("azure-native:servicefabric/v20210601:getApplicationTypeVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationTypeVersionArgs struct {
	ApplicationTypeName string `pulumi:"applicationTypeName"`
	ClusterName         string `pulumi:"clusterName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	Version             string `pulumi:"version"`
}


type LookupApplicationTypeVersionResult struct {
	AppPackageUrl        string             `pulumi:"appPackageUrl"`
	DefaultParameterList map[string]string  `pulumi:"defaultParameterList"`
	Etag                 string             `pulumi:"etag"`
	Id                   string             `pulumi:"id"`
	Location             *string            `pulumi:"location"`
	Name                 string             `pulumi:"name"`
	ProvisioningState    string             `pulumi:"provisioningState"`
	SystemData           SystemDataResponse `pulumi:"systemData"`
	Tags                 map[string]string  `pulumi:"tags"`
	Type                 string             `pulumi:"type"`
}

func LookupApplicationTypeVersionOutput(ctx *pulumi.Context, args LookupApplicationTypeVersionOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationTypeVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationTypeVersionResult, error) {
			args := v.(LookupApplicationTypeVersionArgs)
			r, err := LookupApplicationTypeVersion(ctx, &args, opts...)
			var s LookupApplicationTypeVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationTypeVersionResultOutput)
}

type LookupApplicationTypeVersionOutputArgs struct {
	ApplicationTypeName pulumi.StringInput `pulumi:"applicationTypeName"`
	ClusterName         pulumi.StringInput `pulumi:"clusterName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	Version             pulumi.StringInput `pulumi:"version"`
}

func (LookupApplicationTypeVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationTypeVersionArgs)(nil)).Elem()
}


type LookupApplicationTypeVersionResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationTypeVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationTypeVersionResult)(nil)).Elem()
}

func (o LookupApplicationTypeVersionResultOutput) ToLookupApplicationTypeVersionResultOutput() LookupApplicationTypeVersionResultOutput {
	return o
}

func (o LookupApplicationTypeVersionResultOutput) ToLookupApplicationTypeVersionResultOutputWithContext(ctx context.Context) LookupApplicationTypeVersionResultOutput {
	return o
}

func (o LookupApplicationTypeVersionResultOutput) AppPackageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) string { return v.AppPackageUrl }).(pulumi.StringOutput)
}

func (o LookupApplicationTypeVersionResultOutput) DefaultParameterList() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) map[string]string { return v.DefaultParameterList }).(pulumi.StringMapOutput)
}

func (o LookupApplicationTypeVersionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupApplicationTypeVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationTypeVersionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationTypeVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApplicationTypeVersionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupApplicationTypeVersionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupApplicationTypeVersionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupApplicationTypeVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationTypeVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationTypeVersionResultOutput{})
}
