// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DigitalTwins
{
    public static class GetDigitalTwin
    {
        /// <summary>
        /// Get DigitalTwinsInstances resource.
        /// 
        /// Uses Azure REST API version 2023-01-31.
        /// </summary>
        public static Task<GetDigitalTwinResult> InvokeAsync(GetDigitalTwinArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDigitalTwinResult>("azure-native:digitaltwins:getDigitalTwin", args ?? new GetDigitalTwinArgs(), options.WithDefaults());

        /// <summary>
        /// Get DigitalTwinsInstances resource.
        /// 
        /// Uses Azure REST API version 2023-01-31.
        /// </summary>
        public static Output<GetDigitalTwinResult> Invoke(GetDigitalTwinInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDigitalTwinResult>("azure-native:digitaltwins:getDigitalTwin", args ?? new GetDigitalTwinInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get DigitalTwinsInstances resource.
        /// 
        /// Uses Azure REST API version 2023-01-31.
        /// </summary>
        public static Output<GetDigitalTwinResult> Invoke(GetDigitalTwinInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDigitalTwinResult>("azure-native:digitaltwins:getDigitalTwin", args ?? new GetDigitalTwinInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDigitalTwinArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group that contains the DigitalTwinsInstance.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the DigitalTwinsInstance.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetDigitalTwinArgs()
        {
        }
        public static new GetDigitalTwinArgs Empty => new GetDigitalTwinArgs();
    }

    public sealed class GetDigitalTwinInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group that contains the DigitalTwinsInstance.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the DigitalTwinsInstance.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public GetDigitalTwinInvokeArgs()
        {
        }
        public static new GetDigitalTwinInvokeArgs Empty => new GetDigitalTwinInvokeArgs();
    }


    [OutputType]
    public sealed class GetDigitalTwinResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Time when DigitalTwinsInstance was created.
        /// </summary>
        public readonly string CreatedTime;
        /// <summary>
        /// Api endpoint to work with DigitalTwinsInstance.
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// The resource identifier.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The managed identity for the DigitalTwinsInstance.
        /// </summary>
        public readonly Outputs.DigitalTwinsIdentityResponse? Identity;
        /// <summary>
        /// Time when DigitalTwinsInstance was updated.
        /// </summary>
        public readonly string LastUpdatedTime;
        /// <summary>
        /// The resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The private endpoint connections.
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionResponse> PrivateEndpointConnections;
        /// <summary>
        /// The provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Public network access for the DigitalTwinsInstance.
        /// </summary>
        public readonly string? PublicNetworkAccess;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the DigitalTwinsInstance.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// The resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDigitalTwinResult(
            string azureApiVersion,

            string createdTime,

            string hostName,

            string id,

            Outputs.DigitalTwinsIdentityResponse? identity,

            string lastUpdatedTime,

            string location,

            string name,

            ImmutableArray<Outputs.PrivateEndpointConnectionResponse> privateEndpointConnections,

            string provisioningState,

            string? publicNetworkAccess,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            CreatedTime = createdTime;
            HostName = hostName;
            Id = id;
            Identity = identity;
            LastUpdatedTime = lastUpdatedTime;
            Location = location;
            Name = name;
            PrivateEndpointConnections = privateEndpointConnections;
            ProvisioningState = provisioningState;
            PublicNetworkAccess = publicNetworkAccess;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
