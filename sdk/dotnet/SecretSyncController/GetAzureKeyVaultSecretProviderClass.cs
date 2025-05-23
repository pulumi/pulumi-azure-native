// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SecretSyncController
{
    public static class GetAzureKeyVaultSecretProviderClass
    {
        /// <summary>
        /// Gets the properties of an AzureKeyVaultSecretProviderClass instance.
        /// 
        /// Uses Azure REST API version 2024-08-21-preview.
        /// </summary>
        public static Task<GetAzureKeyVaultSecretProviderClassResult> InvokeAsync(GetAzureKeyVaultSecretProviderClassArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAzureKeyVaultSecretProviderClassResult>("azure-native:secretsynccontroller:getAzureKeyVaultSecretProviderClass", args ?? new GetAzureKeyVaultSecretProviderClassArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the properties of an AzureKeyVaultSecretProviderClass instance.
        /// 
        /// Uses Azure REST API version 2024-08-21-preview.
        /// </summary>
        public static Output<GetAzureKeyVaultSecretProviderClassResult> Invoke(GetAzureKeyVaultSecretProviderClassInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAzureKeyVaultSecretProviderClassResult>("azure-native:secretsynccontroller:getAzureKeyVaultSecretProviderClass", args ?? new GetAzureKeyVaultSecretProviderClassInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the properties of an AzureKeyVaultSecretProviderClass instance.
        /// 
        /// Uses Azure REST API version 2024-08-21-preview.
        /// </summary>
        public static Output<GetAzureKeyVaultSecretProviderClassResult> Invoke(GetAzureKeyVaultSecretProviderClassInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAzureKeyVaultSecretProviderClassResult>("azure-native:secretsynccontroller:getAzureKeyVaultSecretProviderClass", args ?? new GetAzureKeyVaultSecretProviderClassInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAzureKeyVaultSecretProviderClassArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the AzureKeyVaultSecretProviderClass
        /// </summary>
        [Input("azureKeyVaultSecretProviderClassName", required: true)]
        public string AzureKeyVaultSecretProviderClassName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAzureKeyVaultSecretProviderClassArgs()
        {
        }
        public static new GetAzureKeyVaultSecretProviderClassArgs Empty => new GetAzureKeyVaultSecretProviderClassArgs();
    }

    public sealed class GetAzureKeyVaultSecretProviderClassInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the AzureKeyVaultSecretProviderClass
        /// </summary>
        [Input("azureKeyVaultSecretProviderClassName", required: true)]
        public Input<string> AzureKeyVaultSecretProviderClassName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAzureKeyVaultSecretProviderClassInvokeArgs()
        {
        }
        public static new GetAzureKeyVaultSecretProviderClassInvokeArgs Empty => new GetAzureKeyVaultSecretProviderClassInvokeArgs();
    }


    [OutputType]
    public sealed class GetAzureKeyVaultSecretProviderClassResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The user assigned managed identity client ID that should be used to access the Azure Key Vault.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// The complex type of the extended location.
        /// </summary>
        public readonly Outputs.AzureResourceManagerCommonTypesExtendedLocationResponse? ExtendedLocation;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the Azure Key Vault to sync secrets from.
        /// </summary>
        public readonly string KeyvaultName;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Objects defines the desired state of synced K8s secret objects
        /// </summary>
        public readonly string? Objects;
        /// <summary>
        /// Provisioning state of the AzureKeyVaultSecretProviderClass instance.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The Azure Active Directory tenant ID that should be used for authenticating requests to the Azure Key Vault.
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAzureKeyVaultSecretProviderClassResult(
            string azureApiVersion,

            string clientId,

            Outputs.AzureResourceManagerCommonTypesExtendedLocationResponse? extendedLocation,

            string id,

            string keyvaultName,

            string location,

            string name,

            string? objects,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string tenantId,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            ClientId = clientId;
            ExtendedLocation = extendedLocation;
            Id = id;
            KeyvaultName = keyvaultName;
            Location = location;
            Name = name;
            Objects = objects;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Tags = tags;
            TenantId = tenantId;
            Type = type;
        }
    }
}
