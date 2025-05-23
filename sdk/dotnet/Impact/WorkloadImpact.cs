// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Impact
{
    /// <summary>
    /// Workload Impact properties
    /// 
    /// Uses Azure REST API version 2024-05-01-preview. In version 2.x of the Azure Native provider, it used API version 2024-05-01-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:impact:WorkloadImpact")]
    public partial class WorkloadImpact : global::Pulumi.CustomResource
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
        /// The resource-specific properties for this resource.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.WorkloadImpactPropertiesResponse> Properties { get; private set; } = null!;

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
        /// Create a WorkloadImpact resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkloadImpact(string name, WorkloadImpactArgs? args = null, CustomResourceOptions? options = null)
            : base("azure-native:impact:WorkloadImpact", name, args ?? new WorkloadImpactArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkloadImpact(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:impact:WorkloadImpact", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:impact/v20240501preview:WorkloadImpact" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WorkloadImpact resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkloadImpact Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WorkloadImpact(name, id, options);
        }
    }

    public sealed class WorkloadImpactArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource-specific properties for this resource.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.WorkloadImpactPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// workloadImpact resource 
        /// </summary>
        [Input("workloadImpactName")]
        public Input<string>? WorkloadImpactName { get; set; }

        public WorkloadImpactArgs()
        {
        }
        public static new WorkloadImpactArgs Empty => new WorkloadImpactArgs();
    }
}
