// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Devices.V20221115Preview
{
    public static class GetIotHubResource
    {
        /// <summary>
        /// Get the non-security related metadata of an IoT hub.
        /// </summary>
        public static Task<GetIotHubResourceResult> InvokeAsync(GetIotHubResourceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIotHubResourceResult>("azure-native:devices/v20221115preview:getIotHubResource", args ?? new GetIotHubResourceArgs(), options.WithDefaults());

        /// <summary>
        /// Get the non-security related metadata of an IoT hub.
        /// </summary>
        public static Output<GetIotHubResourceResult> Invoke(GetIotHubResourceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIotHubResourceResult>("azure-native:devices/v20221115preview:getIotHubResource", args ?? new GetIotHubResourceInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Get the non-security related metadata of an IoT hub.
        /// </summary>
        public static Output<GetIotHubResourceResult> Invoke(GetIotHubResourceInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIotHubResourceResult>("azure-native:devices/v20221115preview:getIotHubResource", args ?? new GetIotHubResourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIotHubResourceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group that contains the IoT hub.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the IoT hub.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetIotHubResourceArgs()
        {
        }
        public static new GetIotHubResourceArgs Empty => new GetIotHubResourceArgs();
    }

    public sealed class GetIotHubResourceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group that contains the IoT hub.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the IoT hub.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        public GetIotHubResourceInvokeArgs()
        {
        }
        public static new GetIotHubResourceInvokeArgs Empty => new GetIotHubResourceInvokeArgs();
    }


    [OutputType]
    public sealed class GetIotHubResourceResult
    {
        /// <summary>
        /// The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// The resource identifier.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The managed identities for the IotHub.
        /// </summary>
        public readonly Outputs.ArmIdentityResponse? Identity;
        /// <summary>
        /// The resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// IotHub properties
        /// </summary>
        public readonly Outputs.IotHubPropertiesResponse Properties;
        /// <summary>
        /// IotHub SKU info
        /// </summary>
        public readonly Outputs.IotHubSkuInfoResponse Sku;
        /// <summary>
        /// The system meta data relating to this resource.
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
        private GetIotHubResourceResult(
            string? etag,

            string id,

            Outputs.ArmIdentityResponse? identity,

            string location,

            string name,

            Outputs.IotHubPropertiesResponse properties,

            Outputs.IotHubSkuInfoResponse sku,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Etag = etag;
            Id = id;
            Identity = identity;
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
