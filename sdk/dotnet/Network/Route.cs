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
    /// Route resource.
    /// 
    /// Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
    /// 
    /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:network:Route")]
    public partial class Route : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The destination CIDR to which the route applies.
        /// </summary>
        [Output("addressPrefix")]
        public Output<string?> AddressPrefix { get; private set; } = null!;

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
        /// A value indicating whether this route overrides overlapping BGP routes regardless of LPM.
        /// </summary>
        [Output("hasBgpOverride")]
        public Output<bool> HasBgpOverride { get; private set; } = null!;

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
        /// </summary>
        [Output("nextHopIpAddress")]
        public Output<string?> NextHopIpAddress { get; private set; } = null!;

        /// <summary>
        /// The type of Azure hop the packet should be sent to.
        /// </summary>
        [Output("nextHopType")]
        public Output<string> NextHopType { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the route resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Route resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Route(string name, RouteArgs args, CustomResourceOptions? options = null)
            : base("azure-native:network:Route", name, args ?? new RouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Route(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:network:Route", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:network/v20150501preview:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20150615:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20160330:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20160601:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20160901:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20161201:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20170301:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20170601:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20170801:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20170901:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20171001:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20171101:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180101:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180201:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180401:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180601:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180701:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180801:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20181001:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20181101:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20181201:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190201:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190401:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190601:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190701:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190801:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190901:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191101:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191201:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200301:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200401:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200501:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200601:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200701:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200801:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20201101:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210201:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210301:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210501:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210801:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220101:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220501:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220701:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220901:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20221101:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230201:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230401:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230501:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230601:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230901:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20231101:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240101:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240301:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240501:Route" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240701:Route" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Route resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Route Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Route(name, id, options);
        }
    }

    public sealed class RouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination CIDR to which the route applies.
        /// </summary>
        [Input("addressPrefix")]
        public Input<string>? AddressPrefix { get; set; }

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
        /// The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
        /// </summary>
        [Input("nextHopIpAddress")]
        public Input<string>? NextHopIpAddress { get; set; }

        /// <summary>
        /// The type of Azure hop the packet should be sent to.
        /// </summary>
        [Input("nextHopType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Network.RouteNextHopType> NextHopType { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the route.
        /// </summary>
        [Input("routeName")]
        public Input<string>? RouteName { get; set; }

        /// <summary>
        /// The name of the route table.
        /// </summary>
        [Input("routeTableName", required: true)]
        public Input<string> RouteTableName { get; set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public RouteArgs()
        {
        }
        public static new RouteArgs Empty => new RouteArgs();
    }
}
