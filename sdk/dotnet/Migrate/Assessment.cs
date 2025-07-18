// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate
{
    /// <summary>
    /// An assessment created for a group in the Migration project.
    /// 
    /// Uses Azure REST API version 2019-10-01. In version 2.x of the Azure Native provider, it used API version 2019-10-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:migrate:Assessment")]
    public partial class Assessment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// For optimistic concurrency control.
        /// </summary>
        [Output("eTag")]
        public Output<string?> ETag { get; private set; } = null!;

        /// <summary>
        /// Unique name of an assessment.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the assessment.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.AssessmentPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Type of the object = [Microsoft.Migrate/assessmentProjects/groups/assessments].
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Assessment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Assessment(string name, AssessmentArgs args, CustomResourceOptions? options = null)
            : base("azure-native:migrate:Assessment", name, args ?? new AssessmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Assessment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:migrate:Assessment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20180202:Assessment" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20191001:Assessment" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230315:Assessment" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230315:AssessmentsOperation" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230401preview:Assessment" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230401preview:AssessmentsOperation" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230501preview:Assessment" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230501preview:AssessmentsOperation" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230909preview:Assessment" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20230909preview:AssessmentsOperation" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20240101preview:Assessment" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20240101preview:AssessmentsOperation" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate/v20240115:Assessment" },
                    new global::Pulumi.Alias { Type = "azure-native:migrate:AssessmentsOperation" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Assessment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Assessment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Assessment(name, id, options);
        }
    }

    public sealed class AssessmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique name of an assessment within a project.
        /// </summary>
        [Input("assessmentName")]
        public Input<string>? AssessmentName { get; set; }

        /// <summary>
        /// For optimistic concurrency control.
        /// </summary>
        [Input("eTag")]
        public Input<string>? ETag { get; set; }

        /// <summary>
        /// Unique name of a group within a project.
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Migrate project.
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        /// <summary>
        /// Properties of the assessment.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.AssessmentPropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Resource Group that project is part of.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public AssessmentArgs()
        {
        }
        public static new AssessmentArgs Empty => new AssessmentArgs();
    }
}
