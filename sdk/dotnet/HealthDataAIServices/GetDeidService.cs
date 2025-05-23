// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HealthDataAIServices
{
    public static class GetDeidService
    {
        /// <summary>
        /// Get a DeidService
        /// 
        /// Uses Azure REST API version 2024-09-20.
        /// 
        /// Other available API versions: 2024-02-28-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native healthdataaiservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetDeidServiceResult> InvokeAsync(GetDeidServiceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeidServiceResult>("azure-native:healthdataaiservices:getDeidService", args ?? new GetDeidServiceArgs(), options.WithDefaults());

        /// <summary>
        /// Get a DeidService
        /// 
        /// Uses Azure REST API version 2024-09-20.
        /// 
        /// Other available API versions: 2024-02-28-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native healthdataaiservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDeidServiceResult> Invoke(GetDeidServiceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeidServiceResult>("azure-native:healthdataaiservices:getDeidService", args ?? new GetDeidServiceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get a DeidService
        /// 
        /// Uses Azure REST API version 2024-09-20.
        /// 
        /// Other available API versions: 2024-02-28-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native healthdataaiservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDeidServiceResult> Invoke(GetDeidServiceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeidServiceResult>("azure-native:healthdataaiservices:getDeidService", args ?? new GetDeidServiceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeidServiceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the deid service
        /// </summary>
        [Input("deidServiceName", required: true)]
        public string DeidServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDeidServiceArgs()
        {
        }
        public static new GetDeidServiceArgs Empty => new GetDeidServiceArgs();
    }

    public sealed class GetDeidServiceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the deid service
        /// </summary>
        [Input("deidServiceName", required: true)]
        public Input<string> DeidServiceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetDeidServiceInvokeArgs()
        {
        }
        public static new GetDeidServiceInvokeArgs Empty => new GetDeidServiceInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeidServiceResult
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
        public readonly Outputs.DeidServicePropertiesResponse Properties;
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
        private GetDeidServiceResult(
            string azureApiVersion,

            string id,

            Outputs.ManagedServiceIdentityResponse? identity,

            string location,

            string name,

            Outputs.DeidServicePropertiesResponse properties,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
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
