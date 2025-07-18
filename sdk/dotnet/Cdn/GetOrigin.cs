// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn
{
    public static class GetOrigin
    {
        /// <summary>
        /// Gets an existing origin within an endpoint.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// 
        /// Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetOriginResult> InvokeAsync(GetOriginArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOriginResult>("azure-native:cdn:getOrigin", args ?? new GetOriginArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an existing origin within an endpoint.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// 
        /// Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetOriginResult> Invoke(GetOriginInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOriginResult>("azure-native:cdn:getOrigin", args ?? new GetOriginInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an existing origin within an endpoint.
        /// 
        /// Uses Azure REST API version 2024-09-01.
        /// 
        /// Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetOriginResult> Invoke(GetOriginInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOriginResult>("azure-native:cdn:getOrigin", args ?? new GetOriginInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOriginArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the endpoint under the profile which is unique globally.
        /// </summary>
        [Input("endpointName", required: true)]
        public string EndpointName { get; set; } = null!;

        /// <summary>
        /// Name of the origin which is unique within the endpoint.
        /// </summary>
        [Input("originName", required: true)]
        public string OriginName { get; set; } = null!;

        /// <summary>
        /// Name of the CDN profile which is unique within the resource group.
        /// </summary>
        [Input("profileName", required: true)]
        public string ProfileName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetOriginArgs()
        {
        }
        public static new GetOriginArgs Empty => new GetOriginArgs();
    }

    public sealed class GetOriginInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the endpoint under the profile which is unique globally.
        /// </summary>
        [Input("endpointName", required: true)]
        public Input<string> EndpointName { get; set; } = null!;

        /// <summary>
        /// Name of the origin which is unique within the endpoint.
        /// </summary>
        [Input("originName", required: true)]
        public Input<string> OriginName { get; set; } = null!;

        /// <summary>
        /// Name of the CDN profile which is unique within the resource group.
        /// </summary>
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetOriginInvokeArgs()
        {
        }
        public static new GetOriginInvokeArgs Empty => new GetOriginInvokeArgs();
    }


    [OutputType]
    public sealed class GetOriginResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        /// <summary>
        /// Origin is enabled for load balancing or not
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// The value of the HTTP port. Must be between 1 and 65535.
        /// </summary>
        public readonly int? HttpPort;
        /// <summary>
        /// The value of the HTTPS port. Must be between 1 and 65535.
        /// </summary>
        public readonly int? HttpsPort;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
        /// </summary>
        public readonly string? OriginHostHeader;
        /// <summary>
        /// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// The approval status for the connection to the Private Link
        /// </summary>
        public readonly string PrivateEndpointStatus;
        /// <summary>
        /// The Alias of the Private Link resource. Populating this optional field indicates that this origin is 'Private'
        /// </summary>
        public readonly string? PrivateLinkAlias;
        /// <summary>
        /// A custom message to be included in the approval request to connect to the Private Link.
        /// </summary>
        public readonly string? PrivateLinkApprovalMessage;
        /// <summary>
        /// The location of the Private Link resource. Required only if 'privateLinkResourceId' is populated
        /// </summary>
        public readonly string? PrivateLinkLocation;
        /// <summary>
        /// The Resource Id of the Private Link resource. Populating this optional field indicates that this backend is 'Private'
        /// </summary>
        public readonly string? PrivateLinkResourceId;
        /// <summary>
        /// Provisioning status of the origin.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource status of the origin.
        /// </summary>
        public readonly string ResourceState;
        /// <summary>
        /// Read only system data
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
        /// </summary>
        public readonly int? Weight;

        [OutputConstructor]
        private GetOriginResult(
            string azureApiVersion,

            bool? enabled,

            string hostName,

            int? httpPort,

            int? httpsPort,

            string id,

            string name,

            string? originHostHeader,

            int? priority,

            string privateEndpointStatus,

            string? privateLinkAlias,

            string? privateLinkApprovalMessage,

            string? privateLinkLocation,

            string? privateLinkResourceId,

            string provisioningState,

            string resourceState,

            Outputs.SystemDataResponse systemData,

            string type,

            int? weight)
        {
            AzureApiVersion = azureApiVersion;
            Enabled = enabled;
            HostName = hostName;
            HttpPort = httpPort;
            HttpsPort = httpsPort;
            Id = id;
            Name = name;
            OriginHostHeader = originHostHeader;
            Priority = priority;
            PrivateEndpointStatus = privateEndpointStatus;
            PrivateLinkAlias = privateLinkAlias;
            PrivateLinkApprovalMessage = privateLinkApprovalMessage;
            PrivateLinkLocation = privateLinkLocation;
            PrivateLinkResourceId = privateLinkResourceId;
            ProvisioningState = provisioningState;
            ResourceState = resourceState;
            SystemData = systemData;
            Type = type;
            Weight = weight;
        }
    }
}
