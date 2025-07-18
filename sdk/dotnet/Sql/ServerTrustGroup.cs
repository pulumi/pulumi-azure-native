// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql
{
    /// <summary>
    /// A server trust group.
    /// 
    /// Uses Azure REST API version 2023-08-01. In version 2.x of the Azure Native provider, it used API version 2021-11-01.
    /// 
    /// Other available API versions: 2020-02-02-preview, 2020-08-01-preview, 2020-11-01-preview, 2021-02-01-preview, 2021-05-01-preview, 2021-08-01-preview, 2021-11-01, 2021-11-01-preview, 2022-02-01-preview, 2022-05-01-preview, 2022-08-01-preview, 2022-11-01-preview, 2023-02-01-preview, 2023-05-01-preview, 2023-08-01-preview, 2024-05-01-preview, 2024-11-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native sql [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:sql:ServerTrustGroup")]
    public partial class ServerTrustGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Group members information for the server trust group.
        /// </summary>
        [Output("groupMembers")]
        public Output<ImmutableArray<Outputs.ServerInfoResponse>> GroupMembers { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Trust scope of the server trust group.
        /// </summary>
        [Output("trustScopes")]
        public Output<ImmutableArray<string>> TrustScopes { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ServerTrustGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerTrustGroup(string name, ServerTrustGroupArgs args, CustomResourceOptions? options = null)
            : base("azure-native:sql:ServerTrustGroup", name, args ?? new ServerTrustGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerTrustGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:sql:ServerTrustGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20200202preview:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20200801preview:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20201101preview:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20210201preview:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20210501preview:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20210801preview:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20211101:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20211101preview:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220201preview:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220501preview:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20220801preview:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20221101preview:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230201preview:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230501preview:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230801:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20230801preview:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20240501preview:ServerTrustGroup" },
                    new global::Pulumi.Alias { Type = "azure-native:sql/v20241101preview:ServerTrustGroup" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServerTrustGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerTrustGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServerTrustGroup(name, id, options);
        }
    }

    public sealed class ServerTrustGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("groupMembers", required: true)]
        private InputList<Inputs.ServerInfoArgs>? _groupMembers;

        /// <summary>
        /// Group members information for the server trust group.
        /// </summary>
        public InputList<Inputs.ServerInfoArgs> GroupMembers
        {
            get => _groupMembers ?? (_groupMembers = new InputList<Inputs.ServerInfoArgs>());
            set => _groupMembers = value;
        }

        /// <summary>
        /// The name of the region where the resource is located.
        /// </summary>
        [Input("locationName", required: true)]
        public Input<string> LocationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server trust group.
        /// </summary>
        [Input("serverTrustGroupName")]
        public Input<string>? ServerTrustGroupName { get; set; }

        [Input("trustScopes", required: true)]
        private InputList<Union<string, Pulumi.AzureNative.Sql.TrustScope>>? _trustScopes;

        /// <summary>
        /// Trust scope of the server trust group.
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.Sql.TrustScope>> TrustScopes
        {
            get => _trustScopes ?? (_trustScopes = new InputList<Union<string, Pulumi.AzureNative.Sql.TrustScope>>());
            set => _trustScopes = value;
        }

        public ServerTrustGroupArgs()
        {
        }
        public static new ServerTrustGroupArgs Empty => new ServerTrustGroupArgs();
    }
}
