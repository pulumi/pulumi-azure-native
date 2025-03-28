// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataReplication.V20240901
{
    public static class GetVault
    {
        /// <summary>
        /// Gets the details of the vault.
        /// </summary>
        public static Task<GetVaultResult> InvokeAsync(GetVaultArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVaultResult>("azure-native:datareplication/v20240901:getVault", args ?? new GetVaultArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the vault.
        /// </summary>
        public static Output<GetVaultResult> Invoke(GetVaultInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVaultResult>("azure-native:datareplication/v20240901:getVault", args ?? new GetVaultInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the vault.
        /// </summary>
        public static Output<GetVaultResult> Invoke(GetVaultInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVaultResult>("azure-native:datareplication/v20240901:getVault", args ?? new GetVaultInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVaultArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The vault name.
        /// </summary>
        [Input("vaultName", required: true)]
        public string VaultName { get; set; } = null!;

        public GetVaultArgs()
        {
        }
        public static new GetVaultArgs Empty => new GetVaultArgs();
    }

    public sealed class GetVaultInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The vault name.
        /// </summary>
        [Input("vaultName", required: true)]
        public Input<string> VaultName { get; set; } = null!;

        public GetVaultInvokeArgs()
        {
        }
        public static new GetVaultInvokeArgs Empty => new GetVaultInvokeArgs();
    }


    [OutputType]
    public sealed class GetVaultResult
    {
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The managed service identities assigned to this resource.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource-specific properties for this resource.
        /// </summary>
        public readonly Outputs.VaultModelPropertiesResponse Properties;
        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetVaultResult(
            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            string location,

            string name,

            Outputs.VaultModelPropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Id = id;
            Identity = identity;
            Location = location;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
