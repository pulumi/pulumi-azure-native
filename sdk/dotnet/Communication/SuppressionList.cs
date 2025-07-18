// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Communication
{
    /// <summary>
    /// A class representing a SuppressionList resource.
    /// 
    /// Uses Azure REST API version 2023-06-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-06-01-preview.
    /// 
    /// Other available API versions: 2024-09-01-preview, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native communication [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:communication:SuppressionList")]
    public partial class SuppressionList : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The date the resource was created.
        /// </summary>
        [Output("createdTimeStamp")]
        public Output<string> CreatedTimeStamp { get; private set; } = null!;

        /// <summary>
        /// The location where the SuppressionListAddress data is stored at rest. This value is inherited from the parent Domains resource.
        /// </summary>
        [Output("dataLocation")]
        public Output<string> DataLocation { get; private set; } = null!;

        /// <summary>
        /// The date the resource was last updated.
        /// </summary>
        [Output("lastUpdatedTimeStamp")]
        public Output<string> LastUpdatedTimeStamp { get; private set; } = null!;

        /// <summary>
        /// The the name of the suppression list. This value must match one of the valid sender usernames of the sending domain.
        /// </summary>
        [Output("listName")]
        public Output<string?> ListName { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SuppressionList resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SuppressionList(string name, SuppressionListArgs args, CustomResourceOptions? options = null)
            : base("azure-native:communication:SuppressionList", name, args ?? new SuppressionListArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SuppressionList(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:communication:SuppressionList", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:communication/v20230601preview:SuppressionList" },
                    new global::Pulumi.Alias { Type = "azure-native:communication/v20240901preview:SuppressionList" },
                    new global::Pulumi.Alias { Type = "azure-native:communication/v20250501preview:SuppressionList" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SuppressionList resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SuppressionList Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SuppressionList(name, id, options);
        }
    }

    public sealed class SuppressionListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Domains resource.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The name of the EmailService resource.
        /// </summary>
        [Input("emailServiceName", required: true)]
        public Input<string> EmailServiceName { get; set; } = null!;

        /// <summary>
        /// The the name of the suppression list. This value must match one of the valid sender usernames of the sending domain.
        /// </summary>
        [Input("listName")]
        public Input<string>? ListName { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the suppression list.
        /// </summary>
        [Input("suppressionListName")]
        public Input<string>? SuppressionListName { get; set; }

        public SuppressionListArgs()
        {
        }
        public static new SuppressionListArgs Empty => new SuppressionListArgs();
    }
}
