// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.V20180202
{
    /// <summary>
    /// Azure Migrate Project.
    /// </summary>
    [AzureNativeResourceType("azure-native:migrate/v20180202:Project")]
    public partial class Project : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Time when this project was created. Date-Time represented in ISO-8601 format.
        /// </summary>
        [Output("createdTimestamp")]
        public Output<string> CreatedTimestamp { get; private set; } = null!;

        /// <summary>
        /// ARM ID of the Service Map workspace created by user.
        /// </summary>
        [Output("customerWorkspaceId")]
        public Output<string?> CustomerWorkspaceId { get; private set; } = null!;

        /// <summary>
        /// Location of the Service Map workspace created by user.
        /// </summary>
        [Output("customerWorkspaceLocation")]
        public Output<string?> CustomerWorkspaceLocation { get; private set; } = null!;

        /// <summary>
        /// Reports whether project is under discovery.
        /// </summary>
        [Output("discoveryStatus")]
        public Output<string> DiscoveryStatus { get; private set; } = null!;

        /// <summary>
        /// For optimistic concurrency control.
        /// </summary>
        [Output("eTag")]
        public Output<string?> ETag { get; private set; } = null!;

        /// <summary>
        /// Time when last assessment was created. Date-Time represented in ISO-8601 format. This value will be null until assessment is created.
        /// </summary>
        [Output("lastAssessmentTimestamp")]
        public Output<string> LastAssessmentTimestamp { get; private set; } = null!;

        /// <summary>
        /// Session id of the last discovery.
        /// </summary>
        [Output("lastDiscoverySessionId")]
        public Output<string> LastDiscoverySessionId { get; private set; } = null!;

        /// <summary>
        /// Time when this project was created. Date-Time represented in ISO-8601 format. This value will be null until discovery is complete.
        /// </summary>
        [Output("lastDiscoveryTimestamp")]
        public Output<string> LastDiscoveryTimestamp { get; private set; } = null!;

        /// <summary>
        /// Azure location in which project is created.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Name of the project.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Number of assessments created in the project.
        /// </summary>
        [Output("numberOfAssessments")]
        public Output<int> NumberOfAssessments { get; private set; } = null!;

        /// <summary>
        /// Number of groups created in the project.
        /// </summary>
        [Output("numberOfGroups")]
        public Output<int> NumberOfGroups { get; private set; } = null!;

        /// <summary>
        /// Number of machines in the project.
        /// </summary>
        [Output("numberOfMachines")]
        public Output<int> NumberOfMachines { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the project.
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Tags provided by Azure Tagging service.
        /// </summary>
        [Output("tags")]
        public Output<object?> Tags { get; private set; } = null!;

        /// <summary>
        /// Type of the object = [Microsoft.Migrate/projects].
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Time when this project was last updated. Date-Time represented in ISO-8601 format.
        /// </summary>
        [Output("updatedTimestamp")]
        public Output<string> UpdatedTimestamp { get; private set; } = null!;


        /// <summary>
        /// Create a Project resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Project(string name, ProjectArgs args, CustomResourceOptions? options = null)
            : base("azure-native:migrate/v20180202:Project", name, args ?? new ProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Project(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:migrate/v20180202:Project", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20171111preview:Project" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Project resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Project Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Project(name, id, options);
        }
    }

    public sealed class ProjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARM ID of the Service Map workspace created by user.
        /// </summary>
        [Input("customerWorkspaceId")]
        public Input<string>? CustomerWorkspaceId { get; set; }

        /// <summary>
        /// Location of the Service Map workspace created by user.
        /// </summary>
        [Input("customerWorkspaceLocation")]
        public Input<string>? CustomerWorkspaceLocation { get; set; }

        /// <summary>
        /// For optimistic concurrency control.
        /// </summary>
        [Input("eTag")]
        public Input<string>? ETag { get; set; }

        /// <summary>
        /// Azure location in which project is created.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the Azure Migrate project.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// Provisioning state of the project.
        /// </summary>
        [Input("provisioningState")]
        public InputUnion<string, Pulumi.AzureNative.Migrate.V20180202.ProvisioningState>? ProvisioningState { get; set; }

        /// <summary>
        /// Name of the Azure Resource Group that project is part of.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Tags provided by Azure Tagging service.
        /// </summary>
        [Input("tags")]
        public Input<object>? Tags { get; set; }

        public ProjectArgs()
        {
        }
        public static new ProjectArgs Empty => new ProjectArgs();
    }
}
