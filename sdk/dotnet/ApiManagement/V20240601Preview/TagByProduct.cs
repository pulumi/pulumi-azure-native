// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.V20240601Preview
{
    /// <summary>
    /// Tag Contract details.
    /// </summary>
    [AzureNativeResourceType("azure-native:apimanagement/v20240601preview:TagByProduct")]
    public partial class TagByProduct : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Tag name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a TagByProduct resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TagByProduct(string name, TagByProductArgs args, CustomResourceOptions? options = null)
            : base("azure-native:apimanagement/v20240601preview:TagByProduct", name, args ?? new TagByProductArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TagByProduct(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:apimanagement/v20240601preview:TagByProduct", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20170301:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20180101:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20180601preview:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20190101:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20191201:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20191201preview:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20200601preview:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20201201:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20210101preview:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20210401preview:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20210801:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20211201preview:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20220401preview:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20220801:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20220901preview:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20230301preview:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20230501preview:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20230901preview:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20240501:TagByProduct" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement:TagByProduct" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TagByProduct resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TagByProduct Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TagByProduct(name, id, options);
        }
    }

    public sealed class TagByProductArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Product identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("productId", required: true)]
        public Input<string> ProductId { get; set; } = null!;

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
        /// Tag identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("tagId")]
        public Input<string>? TagId { get; set; }

        public TagByProductArgs()
        {
        }
        public static new TagByProductArgs Empty => new TagByProductArgs();
    }
}
