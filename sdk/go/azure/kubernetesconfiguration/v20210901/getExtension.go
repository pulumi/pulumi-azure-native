


package v20210901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExtension(ctx *pulumi.Context, args *LookupExtensionArgs, opts ...pulumi.InvokeOption) (*LookupExtensionResult, error) {
	var rv LookupExtensionResult
	err := ctx.Invoke("azure-native:kubernetesconfiguration/v20210901:getExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupExtensionArgs struct {
	ClusterName         string `pulumi:"clusterName"`
	ClusterResourceName string `pulumi:"clusterResourceName"`
	ClusterRp           string `pulumi:"clusterRp"`
	ExtensionName       string `pulumi:"extensionName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupExtensionResult struct {
	AksAssignedIdentity            *ExtensionResponseAksAssignedIdentity `pulumi:"aksAssignedIdentity"`
	AutoUpgradeMinorVersion        *bool                                 `pulumi:"autoUpgradeMinorVersion"`
	ConfigurationProtectedSettings map[string]string                     `pulumi:"configurationProtectedSettings"`
	ConfigurationSettings          map[string]string                     `pulumi:"configurationSettings"`
	CustomLocationSettings         map[string]string                     `pulumi:"customLocationSettings"`
	ErrorInfo                      *ErrorDetailResponse                  `pulumi:"errorInfo"`
	ExtensionType                  *string                               `pulumi:"extensionType"`
	Id                             string                                `pulumi:"id"`
	Identity                       *IdentityResponse                     `pulumi:"identity"`
	Name                           string                                `pulumi:"name"`
	PackageUri                     string                                `pulumi:"packageUri"`
	ProvisioningState              string                                `pulumi:"provisioningState"`
	ReleaseTrain                   *string                               `pulumi:"releaseTrain"`
	Scope                          *ScopeResponse                        `pulumi:"scope"`
	Statuses                       []ExtensionStatusResponse             `pulumi:"statuses"`
	SystemData                     SystemDataResponse                    `pulumi:"systemData"`
	Type                           string                                `pulumi:"type"`
	Version                        *string                               `pulumi:"version"`
}


func (val *LookupExtensionResult) Defaults() *LookupExtensionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AutoUpgradeMinorVersion) {
		autoUpgradeMinorVersion_ := true
		tmp.AutoUpgradeMinorVersion = &autoUpgradeMinorVersion_
	}
	if isZero(tmp.ReleaseTrain) {
		releaseTrain_ := "Stable"
		tmp.ReleaseTrain = &releaseTrain_
	}
	return &tmp
}
