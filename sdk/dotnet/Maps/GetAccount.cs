// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Maps
{
    public static class GetAccount
    {
        /// <summary>
        /// Get a Maps Account.
        /// 
        /// Uses Azure REST API version 2024-07-01-preview.
        /// 
        /// Other available API versions: 2020-02-01-preview, 2021-02-01, 2021-07-01-preview, 2021-12-01-preview, 2023-06-01, 2023-08-01-preview, 2023-12-01-preview, 2024-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maps [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetAccountResult> InvokeAsync(GetAccountArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAccountResult>("azure-native:maps:getAccount", args ?? new GetAccountArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Maps Account.
        /// 
        /// Uses Azure REST API version 2024-07-01-preview.
        /// 
        /// Other available API versions: 2020-02-01-preview, 2021-02-01, 2021-07-01-preview, 2021-12-01-preview, 2023-06-01, 2023-08-01-preview, 2023-12-01-preview, 2024-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maps [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAccountResult> Invoke(GetAccountInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountResult>("azure-native:maps:getAccount", args ?? new GetAccountInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a Maps Account.
        /// 
        /// Uses Azure REST API version 2024-07-01-preview.
        /// 
        /// Other available API versions: 2020-02-01-preview, 2021-02-01, 2021-07-01-preview, 2021-12-01-preview, 2023-06-01, 2023-08-01-preview, 2023-12-01-preview, 2024-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maps [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetAccountResult> Invoke(GetAccountInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAccountResult>("azure-native:maps:getAccount", args ?? new GetAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAccountArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Maps Account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAccountArgs()
        {
        }
        public static new GetAccountArgs Empty => new GetAccountArgs();
    }

    public sealed class GetAccountInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Maps Account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAccountInvokeArgs()
        {
        }
        public static new GetAccountInvokeArgs Empty => new GetAccountInvokeArgs();
    }


    [OutputType]
    public sealed class GetAccountResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Managed service identity (system assigned and/or user assigned identities)
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// Get or Set Kind property.
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The map account properties.
        /// </summary>
        public readonly Outputs.MapsAccountPropertiesResponse Properties;
        /// <summary>
        /// The SKU of this account.
        /// </summary>
        public readonly Outputs.SkuResponse Sku;
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
        private GetAccountResult(
            string azureApiVersion,

            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            string? kind,

            string location,

            string name,

            Outputs.MapsAccountPropertiesResponse properties,

            Outputs.SkuResponse sku,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Identity = identity;
            Kind = kind;
            Location = location;
            Name = name;
            Properties = properties;
            Sku = sku;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
