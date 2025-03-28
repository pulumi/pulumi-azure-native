// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HealthcareApis.V20250301Preview
{
    public static class GetIotConnector
    {
        /// <summary>
        /// Gets the properties of the specified IoT Connector.
        /// </summary>
        public static Task<GetIotConnectorResult> InvokeAsync(GetIotConnectorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIotConnectorResult>("azure-native:healthcareapis/v20250301preview:getIotConnector", args ?? new GetIotConnectorArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the properties of the specified IoT Connector.
        /// </summary>
        public static Output<GetIotConnectorResult> Invoke(GetIotConnectorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIotConnectorResult>("azure-native:healthcareapis/v20250301preview:getIotConnector", args ?? new GetIotConnectorInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the properties of the specified IoT Connector.
        /// </summary>
        public static Output<GetIotConnectorResult> Invoke(GetIotConnectorInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIotConnectorResult>("azure-native:healthcareapis/v20250301preview:getIotConnector", args ?? new GetIotConnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIotConnectorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of IoT Connector resource.
        /// </summary>
        [Input("iotConnectorName", required: true)]
        public string IotConnectorName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the service instance.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of workspace resource.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetIotConnectorArgs()
        {
        }
        public static new GetIotConnectorArgs Empty => new GetIotConnectorArgs();
    }

    public sealed class GetIotConnectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of IoT Connector resource.
        /// </summary>
        [Input("iotConnectorName", required: true)]
        public Input<string> IotConnectorName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the service instance.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of workspace resource.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public GetIotConnectorInvokeArgs()
        {
        }
        public static new GetIotConnectorInvokeArgs Empty => new GetIotConnectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetIotConnectorResult
    {
        /// <summary>
        /// Device Mappings.
        /// </summary>
        public readonly Outputs.IotMappingPropertiesResponse? DeviceMapping;
        /// <summary>
        /// An etag associated with the resource, used for optimistic concurrency when editing it.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// The resource identifier.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Setting indicating whether the service has a managed identity associated with it.
        /// </summary>
        public readonly Outputs.ServiceManagedIdentityResponseIdentity? Identity;
        /// <summary>
        /// Source configuration.
        /// </summary>
        public readonly Outputs.IotEventHubIngestionEndpointConfigurationResponse? IngestionEndpointConfiguration;
        /// <summary>
        /// The resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetIotConnectorResult(
            Outputs.IotMappingPropertiesResponse? deviceMapping,

            string? etag,

            string id,

            Outputs.ServiceManagedIdentityResponseIdentity? identity,

            Outputs.IotEventHubIngestionEndpointConfigurationResponse? ingestionEndpointConfiguration,

            string? location,

            string name,

            string provisioningState,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            DeviceMapping = deviceMapping;
            Etag = etag;
            Id = id;
            Identity = identity;
            IngestionEndpointConfiguration = ingestionEndpointConfiguration;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
