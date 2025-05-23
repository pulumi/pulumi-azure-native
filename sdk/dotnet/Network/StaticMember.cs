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
    /// StaticMember Item.
    /// 
    /// Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
    /// 
    /// Other available API versions: 2022-01-01, 2022-02-01-preview, 2022-04-01-preview, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:network:StaticMember")]
    public partial class StaticMember : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the scope assignment resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Resource region.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Resource Id.
        /// </summary>
        [Output("resourceId")]
        public Output<string?> ResourceId { get; private set; } = null!;

        /// <summary>
        /// The system metadata related to this resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a StaticMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StaticMember(string name, StaticMemberArgs args, CustomResourceOptions? options = null)
            : base("azure-native:network:StaticMember", name, args ?? new StaticMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StaticMember(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:network:StaticMember", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210501preview:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220101:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220201preview:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220401preview:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220501:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220701:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220901:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20221101:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230201:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230401:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230501:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230601:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230901:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20231101:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240101:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240301:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240501:StaticMember" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240701:StaticMember" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StaticMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StaticMember Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StaticMember(name, id, options);
        }
    }

    public sealed class StaticMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the network group.
        /// </summary>
        [Input("networkGroupName", required: true)]
        public Input<string> NetworkGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the network manager.
        /// </summary>
        [Input("networkManagerName", required: true)]
        public Input<string> NetworkManagerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Resource Id.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The name of the static member.
        /// </summary>
        [Input("staticMemberName")]
        public Input<string>? StaticMemberName { get; set; }

        public StaticMemberArgs()
        {
        }
        public static new StaticMemberArgs Empty => new StaticMemberArgs();
    }
}
