// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Billing
{
    /// <summary>
    /// The properties of the billing role assignment.
    /// 
    /// Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2019-10-01-preview.
    /// 
    /// Other available API versions: 2019-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native billing [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:billing:BillingRoleAssignmentByBillingAccount")]
    public partial class BillingRoleAssignmentByBillingAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties of the billing role assignment.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.BillingRoleAssignmentPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Dictionary of metadata associated with the resource. It may not be populated for all resource types. Maximum key/value length supported of 256 characters. Keys/value should not empty value nor null. Keys can not contain &lt; &gt; % &amp; \ ? /
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a BillingRoleAssignmentByBillingAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BillingRoleAssignmentByBillingAccount(string name, BillingRoleAssignmentByBillingAccountArgs args, CustomResourceOptions? options = null)
            : base("azure-native:billing:BillingRoleAssignmentByBillingAccount", name, args ?? new BillingRoleAssignmentByBillingAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BillingRoleAssignmentByBillingAccount(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:billing:BillingRoleAssignmentByBillingAccount", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:billing/v20191001preview:BillingRoleAssignmentByBillingAccount" },
                    new global::Pulumi.Alias { Type = "azure-native:billing/v20240401:BillingRoleAssignmentByBillingAccount" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BillingRoleAssignmentByBillingAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BillingRoleAssignmentByBillingAccount Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BillingRoleAssignmentByBillingAccount(name, id, options);
        }
    }

    public sealed class BillingRoleAssignmentByBillingAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID that uniquely identifies a billing account.
        /// </summary>
        [Input("billingAccountName", required: true)]
        public Input<string> BillingAccountName { get; set; } = null!;

        /// <summary>
        /// The ID that uniquely identifies a role assignment.
        /// </summary>
        [Input("billingRoleAssignmentName")]
        public Input<string>? BillingRoleAssignmentName { get; set; }

        /// <summary>
        /// The properties of the billing role assignment.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.BillingRoleAssignmentPropertiesArgs>? Properties { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Dictionary of metadata associated with the resource. It may not be populated for all resource types. Maximum key/value length supported of 256 characters. Keys/value should not empty value nor null. Keys can not contain &lt; &gt; % &amp; \ ? /
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public BillingRoleAssignmentByBillingAccountArgs()
        {
        }
        public static new BillingRoleAssignmentByBillingAccountArgs Empty => new BillingRoleAssignmentByBillingAccountArgs();
    }
}
