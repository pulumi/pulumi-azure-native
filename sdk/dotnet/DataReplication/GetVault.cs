// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataReplication
{
    public static class GetVault
    {
        /// <summary>
        /// Gets the details of the vault.
        /// 
        /// Uses Azure REST API version 2021-02-16-preview.
        /// 
        /// Other available API versions: 2024-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datareplication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetVaultResult> InvokeAsync(GetVaultArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVaultResult>("azure-native:datareplication:getVault", args ?? new GetVaultArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the vault.
        /// 
        /// Uses Azure REST API version 2021-02-16-preview.
        /// 
        /// Other available API versions: 2024-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datareplication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetVaultResult> Invoke(GetVaultInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVaultResult>("azure-native:datareplication:getVault", args ?? new GetVaultInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the vault.
        /// 
        /// Uses Azure REST API version 2021-02-16-preview.
        /// 
        /// Other available API versions: 2024-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datareplication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetVaultResult> Invoke(GetVaultInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetVaultResult>("azure-native:datareplication:getVault", args ?? new GetVaultInvokeArgs(), options.WithDefaults());
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
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Gets or sets the Id of the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Gets or sets the location of the vault.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Gets or sets the name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Vault properties.
        /// </summary>
        public readonly Outputs.VaultModelPropertiesResponse Properties;
        public readonly Outputs.VaultModelResponseSystemData SystemData;
        /// <summary>
        /// Gets or sets the resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Gets or sets the type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetVaultResult(
            string azureApiVersion,

            string id,

            string location,

            string name,

            Outputs.VaultModelPropertiesResponse properties,

            Outputs.VaultModelResponseSystemData systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Location = location;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
