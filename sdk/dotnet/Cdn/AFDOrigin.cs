// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Cdn
{
    /// <summary>
    /// Azure Front Door origin is the source of the content being delivered via Azure Front Door. When the edge nodes represented by an endpoint do not have the requested content cached, they attempt to fetch it from one or more of the configured origins.
    /// 
    /// Uses Azure REST API version 2024-09-01. In version 2.x of the Azure Native provider, it used API version 2023-05-01.
    /// 
    /// Other available API versions: 2023-05-01, 2023-07-01-preview, 2024-02-01, 2024-05-01-preview, 2024-06-01-preview, 2025-01-01-preview, 2025-04-15, 2025-06-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cdn [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:cdn:AFDOrigin")]
    public partial class AFDOrigin : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Resource reference to the Azure origin resource.
        /// </summary>
        [Output("azureOrigin")]
        public Output<Outputs.ResourceReferenceResponse?> AzureOrigin { get; private set; } = null!;

        [Output("deploymentStatus")]
        public Output<string> DeploymentStatus { get; private set; } = null!;

        /// <summary>
        /// Whether to enable health probes to be made against backends defined under backendPools. Health probes can only be disabled if there is a single enabled backend in single enabled backend pool.
        /// </summary>
        [Output("enabledState")]
        public Output<string?> EnabledState { get; private set; } = null!;

        /// <summary>
        /// Whether to enable certificate name check at origin level
        /// </summary>
        [Output("enforceCertificateNameCheck")]
        public Output<bool?> EnforceCertificateNameCheck { get; private set; } = null!;

        /// <summary>
        /// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
        /// </summary>
        [Output("hostName")]
        public Output<string> HostName { get; private set; } = null!;

        /// <summary>
        /// The value of the HTTP port. Must be between 1 and 65535.
        /// </summary>
        [Output("httpPort")]
        public Output<int?> HttpPort { get; private set; } = null!;

        /// <summary>
        /// The value of the HTTPS port. Must be between 1 and 65535.
        /// </summary>
        [Output("httpsPort")]
        public Output<int?> HttpsPort { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the origin group which contains this origin.
        /// </summary>
        [Output("originGroupName")]
        public Output<string> OriginGroupName { get; private set; } = null!;

        /// <summary>
        /// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure Front Door origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
        /// </summary>
        [Output("originHostHeader")]
        public Output<string?> OriginHostHeader { get; private set; } = null!;

        /// <summary>
        /// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// Provisioning status
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The properties of the private link resource for private origin.
        /// </summary>
        [Output("sharedPrivateLinkResource")]
        public Output<Outputs.SharedPrivateLinkResourcePropertiesResponse?> SharedPrivateLinkResource { get; private set; } = null!;

        /// <summary>
        /// Read only system data
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
        /// </summary>
        [Output("weight")]
        public Output<int?> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a AFDOrigin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AFDOrigin(string name, AFDOriginArgs args, CustomResourceOptions? options = null)
            : base("azure-native:cdn:AFDOrigin", name, args ?? new AFDOriginArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AFDOrigin(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:cdn:AFDOrigin", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:cdn/v20200901:AFDOrigin" },
                    new global::Pulumi.Alias { Type = "azure-native:cdn/v20210601:AFDOrigin" },
                    new global::Pulumi.Alias { Type = "azure-native:cdn/v20220501preview:AFDOrigin" },
                    new global::Pulumi.Alias { Type = "azure-native:cdn/v20221101preview:AFDOrigin" },
                    new global::Pulumi.Alias { Type = "azure-native:cdn/v20230501:AFDOrigin" },
                    new global::Pulumi.Alias { Type = "azure-native:cdn/v20230701preview:AFDOrigin" },
                    new global::Pulumi.Alias { Type = "azure-native:cdn/v20240201:AFDOrigin" },
                    new global::Pulumi.Alias { Type = "azure-native:cdn/v20240501preview:AFDOrigin" },
                    new global::Pulumi.Alias { Type = "azure-native:cdn/v20240601preview:AFDOrigin" },
                    new global::Pulumi.Alias { Type = "azure-native:cdn/v20240901:AFDOrigin" },
                    new global::Pulumi.Alias { Type = "azure-native:cdn/v20250101preview:AFDOrigin" },
                    new global::Pulumi.Alias { Type = "azure-native:cdn/v20250415:AFDOrigin" },
                    new global::Pulumi.Alias { Type = "azure-native:cdn/v20250601:AFDOrigin" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AFDOrigin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AFDOrigin Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AFDOrigin(name, id, options);
        }
    }

    public sealed class AFDOriginArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource reference to the Azure origin resource.
        /// </summary>
        [Input("azureOrigin")]
        public Input<Inputs.ResourceReferenceArgs>? AzureOrigin { get; set; }

        /// <summary>
        /// Whether to enable health probes to be made against backends defined under backendPools. Health probes can only be disabled if there is a single enabled backend in single enabled backend pool.
        /// </summary>
        [Input("enabledState")]
        public InputUnion<string, Pulumi.AzureNative.Cdn.EnabledState>? EnabledState { get; set; }

        /// <summary>
        /// Whether to enable certificate name check at origin level
        /// </summary>
        [Input("enforceCertificateNameCheck")]
        public Input<bool>? EnforceCertificateNameCheck { get; set; }

        /// <summary>
        /// The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
        /// </summary>
        [Input("hostName", required: true)]
        public Input<string> HostName { get; set; } = null!;

        /// <summary>
        /// The value of the HTTP port. Must be between 1 and 65535.
        /// </summary>
        [Input("httpPort")]
        public Input<int>? HttpPort { get; set; }

        /// <summary>
        /// The value of the HTTPS port. Must be between 1 and 65535.
        /// </summary>
        [Input("httpsPort")]
        public Input<int>? HttpsPort { get; set; }

        /// <summary>
        /// Name of the origin group which is unique within the profile.
        /// </summary>
        [Input("originGroupName", required: true)]
        public Input<string> OriginGroupName { get; set; } = null!;

        /// <summary>
        /// The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure Front Door origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
        /// </summary>
        [Input("originHostHeader")]
        public Input<string>? OriginHostHeader { get; set; }

        /// <summary>
        /// Name of the origin that is unique within the profile.
        /// </summary>
        [Input("originName")]
        public Input<string>? OriginName { get; set; }

        /// <summary>
        /// Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Name of the Azure Front Door Standard or Azure Front Door Premium profile which is unique within the resource group.
        /// </summary>
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The properties of the private link resource for private origin.
        /// </summary>
        [Input("sharedPrivateLinkResource")]
        public Input<Inputs.SharedPrivateLinkResourcePropertiesArgs>? SharedPrivateLinkResource { get; set; }

        /// <summary>
        /// Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public AFDOriginArgs()
        {
            EnforceCertificateNameCheck = true;
            HttpPort = 80;
            HttpsPort = 443;
        }
        public static new AFDOriginArgs Empty => new AFDOriginArgs();
    }
}
