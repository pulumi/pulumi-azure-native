// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevCenter
{
    public static class GetDevCenter
    {
        /// <summary>
        /// Gets a devcenter.
        /// 
        /// Uses Azure REST API version 2024-02-01.
        /// 
        /// Other available API versions: 2023-04-01, 2023-08-01-preview, 2023-10-01-preview, 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-02-01, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetDevCenterResult> InvokeAsync(GetDevCenterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDevCenterResult>("azure-native:devcenter:getDevCenter", args ?? new GetDevCenterArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a devcenter.
        /// 
        /// Uses Azure REST API version 2024-02-01.
        /// 
        /// Other available API versions: 2023-04-01, 2023-08-01-preview, 2023-10-01-preview, 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-02-01, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDevCenterResult> Invoke(GetDevCenterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDevCenterResult>("azure-native:devcenter:getDevCenter", args ?? new GetDevCenterInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a devcenter.
        /// 
        /// Uses Azure REST API version 2024-02-01.
        /// 
        /// Other available API versions: 2023-04-01, 2023-08-01-preview, 2023-10-01-preview, 2024-05-01-preview, 2024-06-01-preview, 2024-07-01-preview, 2024-08-01-preview, 2024-10-01-preview, 2025-02-01, 2025-04-01-preview, 2025-07-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native devcenter [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDevCenterResult> Invoke(GetDevCenterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDevCenterResult>("azure-native:devcenter:getDevCenter", args ?? new GetDevCenterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDevCenterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the devcenter.
        /// </summary>
        [Input("devCenterName", required: true)]
        public string DevCenterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDevCenterArgs()
        {
        }
        public static new GetDevCenterArgs Empty => new GetDevCenterArgs();
    }

    public sealed class GetDevCenterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the devcenter.
        /// </summary>
        [Input("devCenterName", required: true)]
        public Input<string> DevCenterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetDevCenterInvokeArgs()
        {
        }
        public static new GetDevCenterInvokeArgs Empty => new GetDevCenterInvokeArgs();
    }


    [OutputType]
    public sealed class GetDevCenterResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The URI of the Dev Center.
        /// </summary>
        public readonly string DevCenterUri;
        /// <summary>
        /// The display name of the devcenter.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// Encryption settings to be used for server-side encryption for proprietary content (such as catalogs, logs, customizations).
        /// </summary>
        public readonly Outputs.EncryptionResponse? Encryption;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Managed identity properties
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
        /// Dev Center settings to be used when associating a project with a catalog.
        /// </summary>
        public readonly Outputs.DevCenterProjectCatalogSettingsResponse? ProjectCatalogSettings;
        /// <summary>
        /// The provisioning state of the resource.
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
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDevCenterResult(
            string azureApiVersion,

            string devCenterUri,

            string? displayName,

            Outputs.EncryptionResponse? encryption,

            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            string location,

            string name,

            Outputs.DevCenterProjectCatalogSettingsResponse? projectCatalogSettings,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            DevCenterUri = devCenterUri;
            DisplayName = displayName;
            Encryption = encryption;
            Id = id;
            Identity = identity;
            Location = location;
            Name = name;
            ProjectCatalogSettings = projectCatalogSettings;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
