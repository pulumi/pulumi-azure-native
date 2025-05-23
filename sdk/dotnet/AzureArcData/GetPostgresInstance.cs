// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData
{
    public static class GetPostgresInstance
    {
        /// <summary>
        /// Retrieves a postgres Instance resource
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2023-01-15-preview, 2024-05-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurearcdata [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetPostgresInstanceResult> InvokeAsync(GetPostgresInstanceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPostgresInstanceResult>("azure-native:azurearcdata:getPostgresInstance", args ?? new GetPostgresInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a postgres Instance resource
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2023-01-15-preview, 2024-05-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurearcdata [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetPostgresInstanceResult> Invoke(GetPostgresInstanceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPostgresInstanceResult>("azure-native:azurearcdata:getPostgresInstance", args ?? new GetPostgresInstanceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a postgres Instance resource
        /// 
        /// Uses Azure REST API version 2024-01-01.
        /// 
        /// Other available API versions: 2023-01-15-preview, 2024-05-01-preview, 2025-03-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native azurearcdata [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetPostgresInstanceResult> Invoke(GetPostgresInstanceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPostgresInstanceResult>("azure-native:azurearcdata:getPostgresInstance", args ?? new GetPostgresInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPostgresInstanceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of Postgres Instance
        /// </summary>
        [Input("postgresInstanceName", required: true)]
        public string PostgresInstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPostgresInstanceArgs()
        {
        }
        public static new GetPostgresInstanceArgs Empty => new GetPostgresInstanceArgs();
    }

    public sealed class GetPostgresInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of Postgres Instance
        /// </summary>
        [Input("postgresInstanceName", required: true)]
        public Input<string> PostgresInstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetPostgresInstanceInvokeArgs()
        {
        }
        public static new GetPostgresInstanceInvokeArgs Empty => new GetPostgresInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetPostgresInstanceResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// The extendedLocation of the resource.
        /// </summary>
        public readonly Outputs.ExtendedLocationResponse? ExtendedLocation;
        /// <summary>
        /// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// null
        /// </summary>
        public readonly Outputs.PostgresInstancePropertiesResponse Properties;
        /// <summary>
        /// Resource sku.
        /// </summary>
        public readonly Outputs.PostgresInstanceSkuResponse? Sku;
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
        private GetPostgresInstanceResult(
            string azureApiVersion,

            Outputs.ExtendedLocationResponse? extendedLocation,

            string id,

            string location,

            string name,

            Outputs.PostgresInstancePropertiesResponse properties,

            Outputs.PostgresInstanceSkuResponse? sku,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            ExtendedLocation = extendedLocation;
            Id = id;
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
