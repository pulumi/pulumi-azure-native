// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security
{
    /// <summary>
    /// Custom Assessment Automation
    /// 
    /// Uses Azure REST API version 2021-07-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-07-01-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:security:CustomAssessmentAutomation")]
    public partial class CustomAssessmentAutomation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The assessment metadata key used when an assessment is generated for this assessment automation.
        /// </summary>
        [Output("assessmentKey")]
        public Output<string?> AssessmentKey { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// GZip encoded KQL query representing the assessment automation results required.
        /// </summary>
        [Output("compressedQuery")]
        public Output<string?> CompressedQuery { get; private set; } = null!;

        /// <summary>
        /// The description to relate to the assessments generated by this assessment automation.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the assessments generated by this assessment automation.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The remediation description to relate to the assessments generated by this assessment automation.
        /// </summary>
        [Output("remediationDescription")]
        public Output<string?> RemediationDescription { get; private set; } = null!;

        /// <summary>
        /// The severity to relate to the assessments generated by this assessment automation.
        /// </summary>
        [Output("severity")]
        public Output<string?> Severity { get; private set; } = null!;

        /// <summary>
        /// Relevant cloud for the custom assessment automation.
        /// </summary>
        [Output("supportedCloud")]
        public Output<string?> SupportedCloud { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a CustomAssessmentAutomation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomAssessmentAutomation(string name, CustomAssessmentAutomationArgs args, CustomResourceOptions? options = null)
            : base("azure-native:security:CustomAssessmentAutomation", name, args ?? new CustomAssessmentAutomationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomAssessmentAutomation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:security:CustomAssessmentAutomation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:security/v20210701preview:CustomAssessmentAutomation" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CustomAssessmentAutomation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomAssessmentAutomation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CustomAssessmentAutomation(name, id, options);
        }
    }

    public sealed class CustomAssessmentAutomationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Base 64 encoded KQL query representing the assessment automation results required.
        /// </summary>
        [Input("compressedQuery")]
        public Input<string>? CompressedQuery { get; set; }

        /// <summary>
        /// Name of the Custom Assessment Automation.
        /// </summary>
        [Input("customAssessmentAutomationName")]
        public Input<string>? CustomAssessmentAutomationName { get; set; }

        /// <summary>
        /// The description to relate to the assessments generated by this assessment automation.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the assessments generated by this assessment automation.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The remediation description to relate to the assessments generated by this assessment automation.
        /// </summary>
        [Input("remediationDescription")]
        public Input<string>? RemediationDescription { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The severity to relate to the assessments generated by this assessment automation.
        /// </summary>
        [Input("severity")]
        public InputUnion<string, Pulumi.AzureNative.Security.SeverityEnum>? Severity { get; set; }

        /// <summary>
        /// Relevant cloud for the custom assessment automation.
        /// </summary>
        [Input("supportedCloud")]
        public InputUnion<string, Pulumi.AzureNative.Security.SupportedCloudEnum>? SupportedCloud { get; set; }

        public CustomAssessmentAutomationArgs()
        {
        }
        public static new CustomAssessmentAutomationArgs Empty => new CustomAssessmentAutomationArgs();
    }
}
