// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web
{
    /// <summary>
    /// Premier add-on.
    /// 
    /// Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
    /// 
    /// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:web:WebAppPremierAddOnSlot")]
    public partial class WebAppPremierAddOnSlot : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Kind of resource. If the resource is an app, you can refer to https://github.com/Azure/app-service-linux-docs/blob/master/Things_You_Should_Know/kind_property.md#app-service-resource-kind-reference for details supported values for kind.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource Location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Premier add on Marketplace offer.
        /// </summary>
        [Output("marketplaceOffer")]
        public Output<string?> MarketplaceOffer { get; private set; } = null!;

        /// <summary>
        /// Premier add on Marketplace publisher.
        /// </summary>
        [Output("marketplacePublisher")]
        public Output<string?> MarketplacePublisher { get; private set; } = null!;

        /// <summary>
        /// Resource Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Premier add on Product.
        /// </summary>
        [Output("product")]
        public Output<string?> Product { get; private set; } = null!;

        /// <summary>
        /// Premier add on SKU.
        /// </summary>
        [Output("sku")]
        public Output<string?> Sku { get; private set; } = null!;

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
        /// Premier add on Vendor.
        /// </summary>
        [Output("vendor")]
        public Output<string?> Vendor { get; private set; } = null!;


        /// <summary>
        /// Create a WebAppPremierAddOnSlot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WebAppPremierAddOnSlot(string name, WebAppPremierAddOnSlotArgs args, CustomResourceOptions? options = null)
            : base("azure-native:web:WebAppPremierAddOnSlot", name, args ?? new WebAppPremierAddOnSlotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WebAppPremierAddOnSlot(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:web:WebAppPremierAddOnSlot", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:web/v20150801:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20160801:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20180201:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20181101:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20190801:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20200601:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20200901:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20201001:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20201201:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20210101:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20210115:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20210201:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20210301:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20220301:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20220901:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20230101:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20231201:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20240401:WebAppPremierAddOnSlot" },
                    new global::Pulumi.Alias { Type = "azure-native:web/v20241101:WebAppPremierAddOnSlot" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WebAppPremierAddOnSlot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WebAppPremierAddOnSlot Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WebAppPremierAddOnSlot(name, id, options);
        }
    }

    public sealed class WebAppPremierAddOnSlotArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Kind of resource. If the resource is an app, you can refer to https://github.com/Azure/app-service-linux-docs/blob/master/Things_You_Should_Know/kind_property.md#app-service-resource-kind-reference for details supported values for kind.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Resource Location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Premier add on Marketplace offer.
        /// </summary>
        [Input("marketplaceOffer")]
        public Input<string>? MarketplaceOffer { get; set; }

        /// <summary>
        /// Premier add on Marketplace publisher.
        /// </summary>
        [Input("marketplacePublisher")]
        public Input<string>? MarketplacePublisher { get; set; }

        /// <summary>
        /// Name of the app.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Add-on name.
        /// </summary>
        [Input("premierAddOnName")]
        public Input<string>? PremierAddOnName { get; set; }

        /// <summary>
        /// Premier add on Product.
        /// </summary>
        [Input("product")]
        public Input<string>? Product { get; set; }

        /// <summary>
        /// Name of the resource group to which the resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Premier add on SKU.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

        /// <summary>
        /// Name of the deployment slot. If a slot is not specified, the API will update the named add-on for the production slot.
        /// </summary>
        [Input("slot", required: true)]
        public Input<string> Slot { get; set; } = null!;

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

        /// <summary>
        /// Premier add on Vendor.
        /// </summary>
        [Input("vendor")]
        public Input<string>? Vendor { get; set; }

        public WebAppPremierAddOnSlotArgs()
        {
        }
        public static new WebAppPremierAddOnSlotArgs Empty => new WebAppPremierAddOnSlotArgs();
    }
}
