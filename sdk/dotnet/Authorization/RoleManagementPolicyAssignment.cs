// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Authorization
{
    /// <summary>
    /// Role management policy
    /// 
    /// Uses Azure REST API version 2024-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2020-10-01.
    /// 
    /// Other available API versions: 2020-10-01, 2020-10-01-preview, 2024-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native authorization [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:authorization:RoleManagementPolicyAssignment")]
    public partial class RoleManagementPolicyAssignment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// The readonly computed rule applied to the policy.
        /// </summary>
        [Output("effectiveRules")]
        public Output<ImmutableArray<object>> EffectiveRules { get; private set; } = null!;

        /// <summary>
        /// The role management policy name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Additional properties of scope, role definition and policy
        /// </summary>
        [Output("policyAssignmentProperties")]
        public Output<Outputs.PolicyAssignmentPropertiesResponse> PolicyAssignmentProperties { get; private set; } = null!;

        /// <summary>
        /// The policy id role management policy assignment.
        /// </summary>
        [Output("policyId")]
        public Output<string?> PolicyId { get; private set; } = null!;

        /// <summary>
        /// The role definition of management policy assignment.
        /// </summary>
        [Output("roleDefinitionId")]
        public Output<string?> RoleDefinitionId { get; private set; } = null!;

        /// <summary>
        /// The role management policy scope.
        /// </summary>
        [Output("scope")]
        public Output<string?> Scope { get; private set; } = null!;

        /// <summary>
        /// The role management policy type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a RoleManagementPolicyAssignment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RoleManagementPolicyAssignment(string name, RoleManagementPolicyAssignmentArgs args, CustomResourceOptions? options = null)
            : base("azure-native:authorization:RoleManagementPolicyAssignment", name, args ?? new RoleManagementPolicyAssignmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RoleManagementPolicyAssignment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:authorization:RoleManagementPolicyAssignment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:authorization/v20201001:RoleManagementPolicyAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:authorization/v20201001preview:RoleManagementPolicyAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:authorization/v20240201preview:RoleManagementPolicyAssignment" },
                    new global::Pulumi.Alias { Type = "azure-native:authorization/v20240901preview:RoleManagementPolicyAssignment" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RoleManagementPolicyAssignment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RoleManagementPolicyAssignment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RoleManagementPolicyAssignment(name, id, options);
        }
    }

    public sealed class RoleManagementPolicyAssignmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The policy id role management policy assignment.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// The role definition of management policy assignment.
        /// </summary>
        [Input("roleDefinitionId")]
        public Input<string>? RoleDefinitionId { get; set; }

        /// <summary>
        /// The name of format {guid_guid} the role management policy assignment to upsert.
        /// </summary>
        [Input("roleManagementPolicyAssignmentName")]
        public Input<string>? RoleManagementPolicyAssignmentName { get; set; }

        /// <summary>
        /// The role management policy scope.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public RoleManagementPolicyAssignmentArgs()
        {
        }
        public static new RoleManagementPolicyAssignmentArgs Empty => new RoleManagementPolicyAssignmentArgs();
    }
}
