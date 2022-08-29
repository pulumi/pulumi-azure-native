


package v20210315preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomLocation(ctx *pulumi.Context, args *LookupCustomLocationArgs, opts ...pulumi.InvokeOption) (*LookupCustomLocationResult, error) {
	var rv LookupCustomLocationResult
	err := ctx.Invoke("azure-native:extendedlocation/v20210315preview:getCustomLocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomLocationArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupCustomLocationResult struct {
	Authentication      *CustomLocationPropertiesResponseAuthentication `pulumi:"authentication"`
	ClusterExtensionIds []string                                        `pulumi:"clusterExtensionIds"`
	DisplayName         *string                                         `pulumi:"displayName"`
	HostResourceId      *string                                         `pulumi:"hostResourceId"`
	HostType            *string                                         `pulumi:"hostType"`
	Id                  string                                          `pulumi:"id"`
	Location            string                                          `pulumi:"location"`
	Name                string                                          `pulumi:"name"`
	Namespace           *string                                         `pulumi:"namespace"`
	ProvisioningState   *string                                         `pulumi:"provisioningState"`
	SystemData          SystemDataResponse                              `pulumi:"systemData"`
	Tags                map[string]string                               `pulumi:"tags"`
	Type                string                                          `pulumi:"type"`
}

func LookupCustomLocationOutput(ctx *pulumi.Context, args LookupCustomLocationOutputArgs, opts ...pulumi.InvokeOption) LookupCustomLocationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomLocationResult, error) {
			args := v.(LookupCustomLocationArgs)
			r, err := LookupCustomLocation(ctx, &args, opts...)
			var s LookupCustomLocationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomLocationResultOutput)
}

type LookupCustomLocationOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ResourceName      pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupCustomLocationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomLocationArgs)(nil)).Elem()
}


type LookupCustomLocationResultOutput struct{ *pulumi.OutputState }

func (LookupCustomLocationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomLocationResult)(nil)).Elem()
}

func (o LookupCustomLocationResultOutput) ToLookupCustomLocationResultOutput() LookupCustomLocationResultOutput {
	return o
}

func (o LookupCustomLocationResultOutput) ToLookupCustomLocationResultOutputWithContext(ctx context.Context) LookupCustomLocationResultOutput {
	return o
}

func (o LookupCustomLocationResultOutput) Authentication() CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) *CustomLocationPropertiesResponseAuthentication {
		return v.Authentication
	}).(CustomLocationPropertiesResponseAuthenticationPtrOutput)
}

func (o LookupCustomLocationResultOutput) ClusterExtensionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) []string { return v.ClusterExtensionIds }).(pulumi.StringArrayOutput)
}

func (o LookupCustomLocationResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupCustomLocationResultOutput) HostResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) *string { return v.HostResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupCustomLocationResultOutput) HostType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) *string { return v.HostType }).(pulumi.StringPtrOutput)
}

func (o LookupCustomLocationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCustomLocationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupCustomLocationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCustomLocationResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o LookupCustomLocationResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupCustomLocationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCustomLocationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupCustomLocationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomLocationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomLocationResultOutput{})
}
