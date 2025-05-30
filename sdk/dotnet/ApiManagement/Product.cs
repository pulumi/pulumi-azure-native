// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement
{
    /// <summary>
    /// Product details.
    /// 
    /// Uses Azure REST API version 2022-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-08-01.
    /// 
    /// Other available API versions: 2021-04-01-preview, 2021-08-01, 2021-12-01-preview, 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:apimanagement:Product")]
    public partial class Product : global::Pulumi.CustomResource
    {
        /// <summary>
        /// whether subscription approval is required. If false, new subscriptions will be approved automatically enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually approve the subscription before the developer can any of the product’s APIs. Can be present only if subscriptionRequired property is present and has a value of false.
        /// </summary>
        [Output("approvalRequired")]
        public Output<bool?> ApprovalRequired { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Product description. May include HTML formatting tags.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Product name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is notPublished.
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// Whether a product subscription is required for accessing APIs included in this product. If true, the product is referred to as "protected" and a valid subscription key is required for a request to an API included in the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be true.
        /// </summary>
        [Output("subscriptionRequired")]
        public Output<bool?> SubscriptionRequired { get; private set; } = null!;

        /// <summary>
        /// Whether the number of subscriptions a user can have to this product at the same time. Set to null or omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has a value of false.
        /// </summary>
        [Output("subscriptionsLimit")]
        public Output<int?> SubscriptionsLimit { get; private set; } = null!;

        /// <summary>
        /// Product terms of use. Developers trying to subscribe to the product will be presented and required to accept these terms before they can complete the subscription process.
        /// </summary>
        [Output("terms")]
        public Output<string?> Terms { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Product resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Product(string name, ProductArgs args, CustomResourceOptions? options = null)
            : base("azure-native:apimanagement:Product", name, args ?? new ProductArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Product(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:apimanagement:Product", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20160707:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20161010:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20170301:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20180101:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20180601preview:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20190101:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20191201:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20191201preview:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20200601preview:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20201201:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20210101preview:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20210401preview:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20210801:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20211201preview:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20220401preview:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20220801:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20220901preview:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20230301preview:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20230501preview:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20230901preview:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20240501:Product" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20240601preview:Product" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Product resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Product Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Product(name, id, options);
        }
    }

    public sealed class ProductArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// whether subscription approval is required. If false, new subscriptions will be approved automatically enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually approve the subscription before the developer can any of the product’s APIs. Can be present only if subscriptionRequired property is present and has a value of false.
        /// </summary>
        [Input("approvalRequired")]
        public Input<bool>? ApprovalRequired { get; set; }

        /// <summary>
        /// Product description. May include HTML formatting tags.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Product name.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Product identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("productId")]
        public Input<string>? ProductId { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is notPublished.
        /// </summary>
        [Input("state")]
        public Input<Pulumi.AzureNative.ApiManagement.ProductState>? State { get; set; }

        /// <summary>
        /// Whether a product subscription is required for accessing APIs included in this product. If true, the product is referred to as "protected" and a valid subscription key is required for a request to an API included in the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be true.
        /// </summary>
        [Input("subscriptionRequired")]
        public Input<bool>? SubscriptionRequired { get; set; }

        /// <summary>
        /// Whether the number of subscriptions a user can have to this product at the same time. Set to null or omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has a value of false.
        /// </summary>
        [Input("subscriptionsLimit")]
        public Input<int>? SubscriptionsLimit { get; set; }

        /// <summary>
        /// Product terms of use. Developers trying to subscribe to the product will be presented and required to accept these terms before they can complete the subscription process.
        /// </summary>
        [Input("terms")]
        public Input<string>? Terms { get; set; }

        public ProductArgs()
        {
        }
        public static new ProductArgs Empty => new ProductArgs();
    }
}
