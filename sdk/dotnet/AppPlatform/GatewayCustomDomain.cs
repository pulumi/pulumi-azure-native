// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform
{
    /// <summary>
    /// Custom domain of the Spring Cloud Gateway
    /// 
    /// Uses Azure REST API version 2024-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-05-01-preview.
    /// 
    /// Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:appplatform:GatewayCustomDomain")]
    public partial class GatewayCustomDomain : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties of custom domain for Spring Cloud Gateway
        /// </summary>
        [Output("properties")]
        public Output<Outputs.GatewayCustomDomainPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a GatewayCustomDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewayCustomDomain(string name, GatewayCustomDomainArgs args, CustomResourceOptions? options = null)
            : base("azure-native:appplatform:GatewayCustomDomain", name, args ?? new GatewayCustomDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewayCustomDomain(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:appplatform:GatewayCustomDomain", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:appplatform/v20220101preview:GatewayCustomDomain" },
                    new global::Pulumi.Alias { Type = "azure-native:appplatform/v20220301preview:GatewayCustomDomain" },
                    new global::Pulumi.Alias { Type = "azure-native:appplatform/v20220501preview:GatewayCustomDomain" },
                    new global::Pulumi.Alias { Type = "azure-native:appplatform/v20220901preview:GatewayCustomDomain" },
                    new global::Pulumi.Alias { Type = "azure-native:appplatform/v20221101preview:GatewayCustomDomain" },
                    new global::Pulumi.Alias { Type = "azure-native:appplatform/v20221201:GatewayCustomDomain" },
                    new global::Pulumi.Alias { Type = "azure-native:appplatform/v20230101preview:GatewayCustomDomain" },
                    new global::Pulumi.Alias { Type = "azure-native:appplatform/v20230301preview:GatewayCustomDomain" },
                    new global::Pulumi.Alias { Type = "azure-native:appplatform/v20230501preview:GatewayCustomDomain" },
                    new global::Pulumi.Alias { Type = "azure-native:appplatform/v20230701preview:GatewayCustomDomain" },
                    new global::Pulumi.Alias { Type = "azure-native:appplatform/v20230901preview:GatewayCustomDomain" },
                    new global::Pulumi.Alias { Type = "azure-native:appplatform/v20231101preview:GatewayCustomDomain" },
                    new global::Pulumi.Alias { Type = "azure-native:appplatform/v20231201:GatewayCustomDomain" },
                    new global::Pulumi.Alias { Type = "azure-native:appplatform/v20240101preview:GatewayCustomDomain" },
                    new global::Pulumi.Alias { Type = "azure-native:appplatform/v20240501preview:GatewayCustomDomain" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GatewayCustomDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewayCustomDomain Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GatewayCustomDomain(name, id, options);
        }
    }

    public sealed class GatewayCustomDomainArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Spring Cloud Gateway custom domain.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// The name of Spring Cloud Gateway.
        /// </summary>
        [Input("gatewayName", required: true)]
        public Input<string> GatewayName { get; set; } = null!;

        /// <summary>
        /// The properties of custom domain for Spring Cloud Gateway
        /// </summary>
        [Input("properties")]
        public Input<Inputs.GatewayCustomDomainPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Service resource.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GatewayCustomDomainArgs()
        {
        }
        public static new GatewayCustomDomainArgs Empty => new GatewayCustomDomainArgs();
    }
}
