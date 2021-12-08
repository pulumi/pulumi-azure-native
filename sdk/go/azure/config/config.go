


package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func GetAuxiliaryTenantIds(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-native:auxiliaryTenantIds")
}


func GetClientCertificatePassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-native:clientCertificatePassword")
}


func GetClientCertificatePath(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-native:clientCertificatePath")
}


func GetClientId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-native:clientId")
}


func GetClientSecret(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-native:clientSecret")
}


func GetDisablePulumiPartnerId(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "azure-native:disablePulumiPartnerId")
}


func GetEnvironment(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-native:environment")
}


func GetMsiEndpoint(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-native:msiEndpoint")
}


func GetPartnerId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-native:partnerId")
}


func GetSubscriptionId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-native:subscriptionId")
}


func GetTenantId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-native:tenantId")
}


func GetUseMsi(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "azure-native:useMsi")
	if err == nil {
		return v
	}
	return false
}
