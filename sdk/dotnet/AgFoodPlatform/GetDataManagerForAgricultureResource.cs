// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AgFoodPlatform
{
    public static class GetDataManagerForAgricultureResource
    {
        /// <summary>
        /// Get DataManagerForAgriculture resource.
        /// 
        /// Uses Azure REST API version 2023-06-01-preview.
        /// </summary>
        public static Task<GetDataManagerForAgricultureResourceResult> InvokeAsync(GetDataManagerForAgricultureResourceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDataManagerForAgricultureResourceResult>("azure-native:agfoodplatform:getDataManagerForAgricultureResource", args ?? new GetDataManagerForAgricultureResourceArgs(), options.WithDefaults());

        /// <summary>
        /// Get DataManagerForAgriculture resource.
        /// 
        /// Uses Azure REST API version 2023-06-01-preview.
        /// </summary>
        public static Output<GetDataManagerForAgricultureResourceResult> Invoke(GetDataManagerForAgricultureResourceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataManagerForAgricultureResourceResult>("azure-native:agfoodplatform:getDataManagerForAgricultureResource", args ?? new GetDataManagerForAgricultureResourceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get DataManagerForAgriculture resource.
        /// 
        /// Uses Azure REST API version 2023-06-01-preview.
        /// </summary>
        public static Output<GetDataManagerForAgricultureResourceResult> Invoke(GetDataManagerForAgricultureResourceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDataManagerForAgricultureResourceResult>("azure-native:agfoodplatform:getDataManagerForAgricultureResource", args ?? new GetDataManagerForAgricultureResourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDataManagerForAgricultureResourceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// DataManagerForAgriculture resource name.
        /// </summary>
        [Input("dataManagerForAgricultureResourceName", required: true)]
        public string DataManagerForAgricultureResourceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDataManagerForAgricultureResourceArgs()
        {
        }
        public static new GetDataManagerForAgricultureResourceArgs Empty => new GetDataManagerForAgricultureResourceArgs();
    }

    public sealed class GetDataManagerForAgricultureResourceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// DataManagerForAgriculture resource name.
        /// </summary>
        [Input("dataManagerForAgricultureResourceName", required: true)]
        public Input<string> DataManagerForAgricultureResourceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetDataManagerForAgricultureResourceInvokeArgs()
        {
        }
        public static new GetDataManagerForAgricultureResourceInvokeArgs Empty => new GetDataManagerForAgricultureResourceInvokeArgs();
    }


    [OutputType]
    public sealed class GetDataManagerForAgricultureResourceResult
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
        /// Identity for the resource.
        /// </summary>
        public readonly Outputs.IdentityResponse? Identity;
        /// <summary>
        /// Uri of the Data Manager For Agriculture instance.
        /// </summary>
        public readonly string InstanceUri;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Private endpoints.
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionResponse> PrivateEndpointConnections;
        /// <summary>
        /// Data Manager For Agriculture instance provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Property to allow or block public traffic for an Azure Data Manager For Agriculture resource.
        /// </summary>
        public readonly string? PublicNetworkAccess;
        /// <summary>
        /// Sensor integration request model.
        /// </summary>
        public readonly Outputs.SensorIntegrationResponse? SensorIntegration;
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
        private GetDataManagerForAgricultureResourceResult(
            string azureApiVersion,

            string id,

            Outputs.IdentityResponse? identity,

            string instanceUri,

            string location,

            string name,

            ImmutableArray<Outputs.PrivateEndpointConnectionResponse> privateEndpointConnections,

            string provisioningState,

            string? publicNetworkAccess,

            Outputs.SensorIntegrationResponse? sensorIntegration,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Identity = identity;
            InstanceUri = instanceUri;
            Location = location;
            Name = name;
            PrivateEndpointConnections = privateEndpointConnections;
            ProvisioningState = provisioningState;
            PublicNetworkAccess = publicNetworkAccess;
            SensorIntegration = sensorIntegration;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
