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
    /// Contract details.
    /// 
    /// Uses Azure REST API version 2022-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-08-01.
    /// 
    /// Other available API versions: 2021-04-01-preview, 2021-08-01, 2021-12-01-preview, 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:apimanagement:ApiTagDescription")]
    public partial class ApiTagDescription : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Description of the Tag.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Tag name.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Description of the external resources describing the tag.
        /// </summary>
        [Output("externalDocsDescription")]
        public Output<string?> ExternalDocsDescription { get; private set; } = null!;

        /// <summary>
        /// Absolute URL of external resources describing the tag.
        /// </summary>
        [Output("externalDocsUrl")]
        public Output<string?> ExternalDocsUrl { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Identifier of the tag in the form of /tags/{tagId}
        /// </summary>
        [Output("tagId")]
        public Output<string?> TagId { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ApiTagDescription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiTagDescription(string name, ApiTagDescriptionArgs args, CustomResourceOptions? options = null)
            : base("azure-native:apimanagement:ApiTagDescription", name, args ?? new ApiTagDescriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiTagDescription(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:apimanagement:ApiTagDescription", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20170301:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20180101:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20180601preview:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20180601preview:TagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20190101:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20191201:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20191201preview:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20200601preview:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20201201:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20210101preview:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20210401preview:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20210801:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20211201preview:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20220401preview:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20220801:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20220901preview:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20230301preview:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20230501preview:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20230901preview:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20240501:ApiTagDescription" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20240601preview:ApiTagDescription" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApiTagDescription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiTagDescription Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ApiTagDescription(name, id, options);
        }
    }

    public sealed class ApiTagDescriptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// API revision identifier. Must be unique in the current API Management service instance. Non-current revision has ;rev=n as a suffix where n is the revision number.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// Description of the Tag.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Description of the external resources describing the tag.
        /// </summary>
        [Input("externalDocsDescription")]
        public Input<string>? ExternalDocsDescription { get; set; }

        /// <summary>
        /// Absolute URL of external resources describing the tag.
        /// </summary>
        [Input("externalDocsUrl")]
        public Input<string>? ExternalDocsUrl { get; set; }

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
        /// Tag description identifier. Used when creating tagDescription for API/Tag association. Based on API and Tag names.
        /// </summary>
        [Input("tagDescriptionId")]
        public Input<string>? TagDescriptionId { get; set; }

        public ApiTagDescriptionArgs()
        {
        }
        public static new ApiTagDescriptionArgs Empty => new ApiTagDescriptionArgs();
    }
}
