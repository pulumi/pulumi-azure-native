


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkManager(ctx *pulumi.Context, args *LookupNetworkManagerArgs, opts ...pulumi.InvokeOption) (*LookupNetworkManagerResult, error) {
	var rv LookupNetworkManagerResult
	err := ctx.Invoke("azure-native:network/v20210201preview:getNetworkManager", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkManagerArgs struct {
	NetworkManagerName string `pulumi:"networkManagerName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupNetworkManagerResult struct {
	Description                 *string                                               `pulumi:"description"`
	DisplayName                 *string                                               `pulumi:"displayName"`
	Etag                        string                                                `pulumi:"etag"`
	Id                          *string                                               `pulumi:"id"`
	Location                    *string                                               `pulumi:"location"`
	Name                        string                                                `pulumi:"name"`
	NetworkManagerScopeAccesses []string                                              `pulumi:"networkManagerScopeAccesses"`
	NetworkManagerScopes        *NetworkManagerPropertiesResponseNetworkManagerScopes `pulumi:"networkManagerScopes"`
	ProvisioningState           string                                                `pulumi:"provisioningState"`
	SystemData                  SystemDataResponse                                    `pulumi:"systemData"`
	Tags                        map[string]string                                     `pulumi:"tags"`
	Type                        string                                                `pulumi:"type"`
}

func LookupNetworkManagerOutput(ctx *pulumi.Context, args LookupNetworkManagerOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkManagerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkManagerResult, error) {
			args := v.(LookupNetworkManagerArgs)
			r, err := LookupNetworkManager(ctx, &args, opts...)
			var s LookupNetworkManagerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkManagerResultOutput)
}

type LookupNetworkManagerOutputArgs struct {
	NetworkManagerName pulumi.StringInput `pulumi:"networkManagerName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNetworkManagerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkManagerArgs)(nil)).Elem()
}


type LookupNetworkManagerResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkManagerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkManagerResult)(nil)).Elem()
}

func (o LookupNetworkManagerResultOutput) ToLookupNetworkManagerResultOutput() LookupNetworkManagerResultOutput {
	return o
}

func (o LookupNetworkManagerResultOutput) ToLookupNetworkManagerResultOutputWithContext(ctx context.Context) LookupNetworkManagerResultOutput {
	return o
}

func (o LookupNetworkManagerResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkManagerResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkManagerResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupNetworkManagerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkManagerResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkManagerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNetworkManagerResultOutput) NetworkManagerScopeAccesses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) []string { return v.NetworkManagerScopeAccesses }).(pulumi.StringArrayOutput)
}

func (o LookupNetworkManagerResultOutput) NetworkManagerScopes() NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) *NetworkManagerPropertiesResponseNetworkManagerScopes {
		return v.NetworkManagerScopes
	}).(NetworkManagerPropertiesResponseNetworkManagerScopesPtrOutput)
}

func (o LookupNetworkManagerResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNetworkManagerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupNetworkManagerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupNetworkManagerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkManagerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkManagerResultOutput{})
}
