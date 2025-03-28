// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CognitiveServices.V20230501
{
    /// <summary>
    /// The commitment plan association.
    /// </summary>
    [AzureNativeResourceType("azure-native:cognitiveservices/v20230501:CommitmentPlanAssociation")]
    public partial class CommitmentPlanAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure resource id of the account.
        /// </summary>
        [Output("accountId")]
        public Output<string?> AccountId { get; private set; } = null!;

        /// <summary>
        /// Resource Etag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a CommitmentPlanAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CommitmentPlanAssociation(string name, CommitmentPlanAssociationArgs args, CustomResourceOptions? options = null)
            : base("azure-native:cognitiveservices/v20230501:CommitmentPlanAssociation", name, args ?? new CommitmentPlanAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CommitmentPlanAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:cognitiveservices/v20230501:CommitmentPlanAssociation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20221201:CommitmentPlanAssociation" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20231001preview:CommitmentPlanAssociation" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20240401preview:CommitmentPlanAssociation" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20240601preview:CommitmentPlanAssociation" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20241001:CommitmentPlanAssociation" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices/v20250401preview:CommitmentPlanAssociation" },
                    new global::Pulumi.Alias { Type = "azure-native:cognitiveservices:CommitmentPlanAssociation" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CommitmentPlanAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CommitmentPlanAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CommitmentPlanAssociation(name, id, options);
        }
    }

    public sealed class CommitmentPlanAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure resource id of the account.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The name of the commitment plan association with the Cognitive Services Account
        /// </summary>
        [Input("commitmentPlanAssociationName")]
        public Input<string>? CommitmentPlanAssociationName { get; set; }

        /// <summary>
        /// The name of the commitmentPlan associated with the Cognitive Services Account
        /// </summary>
        [Input("commitmentPlanName", required: true)]
        public Input<string> CommitmentPlanName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public CommitmentPlanAssociationArgs()
        {
        }
        public static new CommitmentPlanAssociationArgs Empty => new CommitmentPlanAssociationArgs();
    }
}
