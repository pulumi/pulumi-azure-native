// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement
{
    /// <summary>
    /// Tag Inheritance Setting definition.
    /// 
    /// Uses Azure REST API version 2024-08-01. In version 2.x of the Azure Native provider, it used API version 2022-10-05-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:costmanagement:TagInheritanceSetting")]
    public partial class TagInheritanceSetting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Specifies the kind of settings.
        /// Expected value is 'taginheritance'.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties of the tag inheritance setting.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.TagInheritancePropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a TagInheritanceSetting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TagInheritanceSetting(string name, TagInheritanceSettingArgs args, CustomResourceOptions? options = null)
            : base("azure-native:costmanagement:TagInheritanceSetting", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }

        private TagInheritanceSetting(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:costmanagement:TagInheritanceSetting", name, null, MakeResourceOptions(options, id))
        {
        }

        private static TagInheritanceSettingArgs MakeArgs(TagInheritanceSettingArgs args)
        {
            args ??= new TagInheritanceSettingArgs();
            args.Kind = "taginheritance";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20221001preview:TagInheritanceSetting" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20221005preview:TagInheritanceSetting" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20230801:TagInheritanceSetting" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20230901:TagInheritanceSetting" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20231101:TagInheritanceSetting" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20240801:TagInheritanceSetting" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20241001preview:TagInheritanceSetting" },
                    new global::Pulumi.Alias { Type = "azure-native:costmanagement/v20250301:TagInheritanceSetting" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TagInheritanceSetting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TagInheritanceSetting Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TagInheritanceSetting(name, id, options);
        }
    }

    public sealed class TagInheritanceSettingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the kind of settings.
        /// Expected value is 'taginheritance'.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// The properties of the tag inheritance setting.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.TagInheritancePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The scope associated with this setting. This includes 'subscriptions/{subscriptionId}' for subscription scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope, 'providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billing profile scope.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        /// <summary>
        /// Setting type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public TagInheritanceSettingArgs()
        {
        }
        public static new TagInheritanceSettingArgs Empty => new TagInheritanceSettingArgs();
    }
}
