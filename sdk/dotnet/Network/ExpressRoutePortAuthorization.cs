// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    /// <summary>
    /// ExpressRoutePort Authorization resource definition.
    /// 
    /// Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
    /// 
    /// Other available API versions: 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:network:ExpressRoutePortAuthorization")]
    public partial class ExpressRoutePortAuthorization : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The authorization key.
        /// </summary>
        [Output("authorizationKey")]
        public Output<string> AuthorizationKey { get; private set; } = null!;

        /// <summary>
        /// The authorization use status.
        /// </summary>
        [Output("authorizationUseStatus")]
        public Output<string> AuthorizationUseStatus { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The reference to the ExpressRoute circuit resource using the authorization.
        /// </summary>
        [Output("circuitResourceUri")]
        public Output<string> CircuitResourceUri { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the authorization resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ExpressRoutePortAuthorization resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExpressRoutePortAuthorization(string name, ExpressRoutePortAuthorizationArgs args, CustomResourceOptions? options = null)
            : base("azure-native:network:ExpressRoutePortAuthorization", name, args ?? new ExpressRoutePortAuthorizationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExpressRoutePortAuthorization(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:network:ExpressRoutePortAuthorization", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210801:ExpressRoutePortAuthorization" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220101:ExpressRoutePortAuthorization" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220501:ExpressRoutePortAuthorization" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220701:ExpressRoutePortAuthorization" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220901:ExpressRoutePortAuthorization" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20221101:ExpressRoutePortAuthorization" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230201:ExpressRoutePortAuthorization" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230401:ExpressRoutePortAuthorization" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230501:ExpressRoutePortAuthorization" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230601:ExpressRoutePortAuthorization" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230901:ExpressRoutePortAuthorization" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20231101:ExpressRoutePortAuthorization" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240101:ExpressRoutePortAuthorization" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240301:ExpressRoutePortAuthorization" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240501:ExpressRoutePortAuthorization" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240701:ExpressRoutePortAuthorization" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ExpressRoutePortAuthorization resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExpressRoutePortAuthorization Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ExpressRoutePortAuthorization(name, id, options);
        }
    }

    public sealed class ExpressRoutePortAuthorizationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the authorization.
        /// </summary>
        [Input("authorizationName")]
        public Input<string>? AuthorizationName { get; set; }

        /// <summary>
        /// The name of the express route port.
        /// </summary>
        [Input("expressRoutePortName", required: true)]
        public Input<string> ExpressRoutePortName { get; set; } = null!;

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ExpressRoutePortAuthorizationArgs()
        {
        }
        public static new ExpressRoutePortAuthorizationArgs Empty => new ExpressRoutePortAuthorizationArgs();
    }
}
