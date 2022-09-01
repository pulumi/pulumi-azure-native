


package v20201110preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupHostPool(ctx *pulumi.Context, args *LookupHostPoolArgs, opts ...pulumi.InvokeOption) (*LookupHostPoolResult, error) {
	var rv LookupHostPoolResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20201110preview:getHostPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHostPoolArgs struct {
	HostPoolName      string `pulumi:"hostPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupHostPoolResult struct {
	ApplicationGroupReferences    []string                  `pulumi:"applicationGroupReferences"`
	CustomRdpProperty             *string                   `pulumi:"customRdpProperty"`
	Description                   *string                   `pulumi:"description"`
	FriendlyName                  *string                   `pulumi:"friendlyName"`
	HostPoolType                  string                    `pulumi:"hostPoolType"`
	Id                            string                    `pulumi:"id"`
	LoadBalancerType              string                    `pulumi:"loadBalancerType"`
	Location                      string                    `pulumi:"location"`
	MaxSessionLimit               *int                      `pulumi:"maxSessionLimit"`
	Name                          string                    `pulumi:"name"`
	PersonalDesktopAssignmentType *string                   `pulumi:"personalDesktopAssignmentType"`
	PreferredAppGroupType         string                    `pulumi:"preferredAppGroupType"`
	RegistrationInfo              *RegistrationInfoResponse `pulumi:"registrationInfo"`
	Ring                          *int                      `pulumi:"ring"`
	SsoClientId                   *string                   `pulumi:"ssoClientId"`
	SsoClientSecretKeyVaultPath   *string                   `pulumi:"ssoClientSecretKeyVaultPath"`
	SsoSecretType                 *string                   `pulumi:"ssoSecretType"`
	SsoadfsAuthority              *string                   `pulumi:"ssoadfsAuthority"`
	StartVMOnConnect              *bool                     `pulumi:"startVMOnConnect"`
	Tags                          map[string]string         `pulumi:"tags"`
	Type                          string                    `pulumi:"type"`
	ValidationEnvironment         *bool                     `pulumi:"validationEnvironment"`
	VmTemplate                    *string                   `pulumi:"vmTemplate"`
}

func LookupHostPoolOutput(ctx *pulumi.Context, args LookupHostPoolOutputArgs, opts ...pulumi.InvokeOption) LookupHostPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHostPoolResult, error) {
			args := v.(LookupHostPoolArgs)
			r, err := LookupHostPool(ctx, &args, opts...)
			var s LookupHostPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHostPoolResultOutput)
}

type LookupHostPoolOutputArgs struct {
	HostPoolName      pulumi.StringInput `pulumi:"hostPoolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupHostPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostPoolArgs)(nil)).Elem()
}


type LookupHostPoolResultOutput struct{ *pulumi.OutputState }

func (LookupHostPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHostPoolResult)(nil)).Elem()
}

func (o LookupHostPoolResultOutput) ToLookupHostPoolResultOutput() LookupHostPoolResultOutput {
	return o
}

func (o LookupHostPoolResultOutput) ToLookupHostPoolResultOutputWithContext(ctx context.Context) LookupHostPoolResultOutput {
	return o
}

func (o LookupHostPoolResultOutput) ApplicationGroupReferences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHostPoolResult) []string { return v.ApplicationGroupReferences }).(pulumi.StringArrayOutput)
}

func (o LookupHostPoolResultOutput) CustomRdpProperty() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.CustomRdpProperty }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) HostPoolType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.HostPoolType }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) LoadBalancerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.LoadBalancerType }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) MaxSessionLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *int { return v.MaxSessionLimit }).(pulumi.IntPtrOutput)
}

func (o LookupHostPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) PersonalDesktopAssignmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.PersonalDesktopAssignmentType }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) PreferredAppGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.PreferredAppGroupType }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) RegistrationInfo() RegistrationInfoResponsePtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *RegistrationInfoResponse { return v.RegistrationInfo }).(RegistrationInfoResponsePtrOutput)
}

func (o LookupHostPoolResultOutput) Ring() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *int { return v.Ring }).(pulumi.IntPtrOutput)
}

func (o LookupHostPoolResultOutput) SsoClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.SsoClientId }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) SsoClientSecretKeyVaultPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.SsoClientSecretKeyVaultPath }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) SsoSecretType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.SsoSecretType }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) SsoadfsAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.SsoadfsAuthority }).(pulumi.StringPtrOutput)
}

func (o LookupHostPoolResultOutput) StartVMOnConnect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *bool { return v.StartVMOnConnect }).(pulumi.BoolPtrOutput)
}

func (o LookupHostPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupHostPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupHostPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHostPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupHostPoolResultOutput) ValidationEnvironment() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *bool { return v.ValidationEnvironment }).(pulumi.BoolPtrOutput)
}

func (o LookupHostPoolResultOutput) VmTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHostPoolResult) *string { return v.VmTemplate }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHostPoolResultOutput{})
}
