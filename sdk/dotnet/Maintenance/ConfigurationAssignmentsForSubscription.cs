// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Maintenance
{
    /// <summary>
    /// Configuration Assignment
    /// 
    /// Uses Azure REST API version 2023-10-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
    /// 
    /// Other available API versions: 2023-04-01, 2023-09-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native maintenance [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:maintenance:ConfigurationAssignmentsForSubscription")]
    public partial class ConfigurationAssignmentsForSubscription : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Properties of the configuration assignment
        /// </summary>
        [Output("filter")]
        public Output<Outputs.ConfigurationAssignmentFilterPropertiesResponse?> Filter { get; private set; } = null!;

        /// <summary>
        /// Location of the resource
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The maintenance configuration Id
        /// </summary>
        [Output("maintenanceConfigurationId")]
        public Output<string?> MaintenanceConfigurationId { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The unique resourceId
        /// </summary>
        [Output("resourceId")]
        public Output<string?> ResourceId { get; private set; } = null!;

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
        /// Create a ConfigurationAssignmentsForSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigurationAssignmentsForSubscription(string name, ConfigurationAssignmentsForSubscriptionArgs? args = null, CustomResourceOptions? options = null)
            : base("azure-native:maintenance:ConfigurationAssignmentsForSubscription", name, args ?? new ConfigurationAssignmentsForSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigurationAssignmentsForSubscription(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:maintenance:ConfigurationAssignmentsForSubscription", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:maintenance/v20230401:ConfigurationAssignmentsForSubscription" },
                    new global::Pulumi.Alias { Type = "azure-native:maintenance/v20230901preview:ConfigurationAssignmentsForSubscription" },
                    new global::Pulumi.Alias { Type = "azure-native:maintenance/v20231001preview:ConfigurationAssignmentsForSubscription" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ConfigurationAssignmentsForSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigurationAssignmentsForSubscription Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConfigurationAssignmentsForSubscription(name, id, options);
        }
    }

    public sealed class ConfigurationAssignmentsForSubscriptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the ConfigurationAssignment
        /// </summary>
        [Input("configurationAssignmentName")]
        public Input<string>? ConfigurationAssignmentName { get; set; }

        /// <summary>
        /// Properties of the configuration assignment
        /// </summary>
        [Input("filter")]
        public Input<Inputs.ConfigurationAssignmentFilterPropertiesArgs>? Filter { get; set; }

        /// <summary>
        /// Location of the resource
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The maintenance configuration Id
        /// </summary>
        [Input("maintenanceConfigurationId")]
        public Input<string>? MaintenanceConfigurationId { get; set; }

        /// <summary>
        /// The unique resourceId
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        public ConfigurationAssignmentsForSubscriptionArgs()
        {
        }
        public static new ConfigurationAssignmentsForSubscriptionArgs Empty => new ConfigurationAssignmentsForSubscriptionArgs();
    }
}
