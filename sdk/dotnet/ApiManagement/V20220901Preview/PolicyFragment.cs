// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.V20220901Preview
{
    /// <summary>
    /// Policy fragment contract details.
    /// </summary>
    [AzureNativeResourceType("azure-native:apimanagement/v20220901preview:PolicyFragment")]
    public partial class PolicyFragment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Policy fragment description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Format of the policy fragment content.
        /// </summary>
        [Output("format")]
        public Output<string?> Format { get; private set; } = null!;

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
        /// Contents of the policy fragment.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyFragment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyFragment(string name, PolicyFragmentArgs args, CustomResourceOptions? options = null)
            : base("azure-native:apimanagement/v20220901preview:PolicyFragment", name, args ?? new PolicyFragmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyFragment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:apimanagement/v20220901preview:PolicyFragment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20211201preview:PolicyFragment" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20220401preview:PolicyFragment" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20220801:PolicyFragment" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20230301preview:PolicyFragment" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20230501preview:PolicyFragment" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20230901preview:PolicyFragment" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20240501:PolicyFragment" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20240601preview:PolicyFragment" },
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement:PolicyFragment" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PolicyFragment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyFragment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PolicyFragment(name, id, options);
        }
    }

    public sealed class PolicyFragmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy fragment description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Format of the policy fragment content.
        /// </summary>
        [Input("format")]
        public InputUnion<string, Pulumi.AzureNative.ApiManagement.V20220901Preview.PolicyFragmentContentFormat>? Format { get; set; }

        /// <summary>
        /// A resource identifier.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

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
        /// Contents of the policy fragment.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public PolicyFragmentArgs()
        {
            Format = "xml";
        }
        public static new PolicyFragmentArgs Empty => new PolicyFragmentArgs();
    }
}
