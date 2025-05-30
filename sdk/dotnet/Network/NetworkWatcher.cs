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
    /// Network watcher in a resource group.
    /// 
    /// Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
    /// 
    /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:network:NetworkWatcher")]
    public partial class NetworkWatcher : global::Pulumi.CustomResource
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
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the network watcher resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkWatcher resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkWatcher(string name, NetworkWatcherArgs args, CustomResourceOptions? options = null)
            : base("azure-native:network:NetworkWatcher", name, args ?? new NetworkWatcherArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkWatcher(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:network:NetworkWatcher", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:network/v20160901:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20161201:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20170301:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20170601:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20170801:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20170901:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20171001:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20171101:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180101:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180201:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180401:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180601:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180701:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180801:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20181001:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20181101:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20181201:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190201:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190401:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190601:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190701:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190801:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190901:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191101:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191201:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200301:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200401:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200501:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200601:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200701:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200801:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20201101:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210201:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210301:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210501:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210801:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220101:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220501:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220701:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220901:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20221101:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230201:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230401:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230501:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230601:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230901:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20231101:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240101:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240301:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240501:NetworkWatcher" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240701:NetworkWatcher" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NetworkWatcher resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkWatcher Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkWatcher(name, id, options);
        }
    }

    public sealed class NetworkWatcherArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the network watcher.
        /// </summary>
        [Input("networkWatcherName")]
        public Input<string>? NetworkWatcherName { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public NetworkWatcherArgs()
        {
        }
        public static new NetworkWatcherArgs Empty => new NetworkWatcherArgs();
    }
}
