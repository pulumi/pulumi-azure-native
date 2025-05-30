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
    /// Governance rule over a given scope
    /// 
    /// Uses Azure REST API version 2022-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-01-01-preview.
    /// </summary>
    [AzureNativeResourceType("azure-native:security:GovernanceRule")]
    public partial class GovernanceRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Description of the governance rule
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Display name of the governance rule
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Excluded scopes, filter out the descendants of the scope (on management scopes)
        /// </summary>
        [Output("excludedScopes")]
        public Output<ImmutableArray<string>> ExcludedScopes { get; private set; } = null!;

        /// <summary>
        /// The email notifications settings for the governance rule, states whether to disable notifications for mangers and owners
        /// </summary>
        [Output("governanceEmailNotification")]
        public Output<Outputs.GovernanceRuleEmailNotificationResponse?> GovernanceEmailNotification { get; private set; } = null!;

        /// <summary>
        /// Defines whether the rule is management scope rule (master connector as a single scope or management scope)
        /// </summary>
        [Output("includeMemberScopes")]
        public Output<bool?> IncludeMemberScopes { get; private set; } = null!;

        /// <summary>
        /// Defines whether the rule is active/inactive
        /// </summary>
        [Output("isDisabled")]
        public Output<bool?> IsDisabled { get; private set; } = null!;

        /// <summary>
        /// Defines whether there is a grace period on the governance rule
        /// </summary>
        [Output("isGracePeriod")]
        public Output<bool?> IsGracePeriod { get; private set; } = null!;

        /// <summary>
        /// The governance rule metadata
        /// </summary>
        [Output("metadata")]
        public Output<Outputs.GovernanceRuleMetadataResponse?> Metadata { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The owner source for the governance rule - e.g. Manually by user@contoso.com - see example
        /// </summary>
        [Output("ownerSource")]
        public Output<Outputs.GovernanceRuleOwnerSourceResponse> OwnerSource { get; private set; } = null!;

        /// <summary>
        /// Governance rule remediation timeframe - this is the time that will affect on the grace-period duration e.g. 7.00:00:00 - means 7 days
        /// </summary>
        [Output("remediationTimeframe")]
        public Output<string?> RemediationTimeframe { get; private set; } = null!;

        /// <summary>
        /// The governance rule priority, priority to the lower number. Rules with the same priority on the same scope will not be allowed
        /// </summary>
        [Output("rulePriority")]
        public Output<int> RulePriority { get; private set; } = null!;

        /// <summary>
        /// The rule type of the governance rule, defines the source of the rule e.g. Integrated
        /// </summary>
        [Output("ruleType")]
        public Output<string> RuleType { get; private set; } = null!;

        /// <summary>
        /// The governance rule source, what the rule affects, e.g. Assessments
        /// </summary>
        [Output("sourceResourceType")]
        public Output<string> SourceResourceType { get; private set; } = null!;

        /// <summary>
        /// The tenantId (GUID)
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a GovernanceRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GovernanceRule(string name, GovernanceRuleArgs args, CustomResourceOptions? options = null)
            : base("azure-native:security:GovernanceRule", name, args ?? new GovernanceRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GovernanceRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:security:GovernanceRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:security/v20220101preview:GovernanceRule" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GovernanceRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GovernanceRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GovernanceRule(name, id, options);
        }
    }

    public sealed class GovernanceRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the governance rule
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Display name of the governance rule
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("excludedScopes")]
        private InputList<string>? _excludedScopes;

        /// <summary>
        /// Excluded scopes, filter out the descendants of the scope (on management scopes)
        /// </summary>
        public InputList<string> ExcludedScopes
        {
            get => _excludedScopes ?? (_excludedScopes = new InputList<string>());
            set => _excludedScopes = value;
        }

        /// <summary>
        /// The email notifications settings for the governance rule, states whether to disable notifications for mangers and owners
        /// </summary>
        [Input("governanceEmailNotification")]
        public Input<Inputs.GovernanceRuleEmailNotificationArgs>? GovernanceEmailNotification { get; set; }

        /// <summary>
        /// Defines whether the rule is management scope rule (master connector as a single scope or management scope)
        /// </summary>
        [Input("includeMemberScopes")]
        public Input<bool>? IncludeMemberScopes { get; set; }

        /// <summary>
        /// Defines whether the rule is active/inactive
        /// </summary>
        [Input("isDisabled")]
        public Input<bool>? IsDisabled { get; set; }

        /// <summary>
        /// Defines whether there is a grace period on the governance rule
        /// </summary>
        [Input("isGracePeriod")]
        public Input<bool>? IsGracePeriod { get; set; }

        /// <summary>
        /// The owner source for the governance rule - e.g. Manually by user@contoso.com - see example
        /// </summary>
        [Input("ownerSource", required: true)]
        public Input<Inputs.GovernanceRuleOwnerSourceArgs> OwnerSource { get; set; } = null!;

        /// <summary>
        /// Governance rule remediation timeframe - this is the time that will affect on the grace-period duration e.g. 7.00:00:00 - means 7 days
        /// </summary>
        [Input("remediationTimeframe")]
        public Input<string>? RemediationTimeframe { get; set; }

        /// <summary>
        /// The governance rule key - unique key for the standard governance rule (GUID)
        /// </summary>
        [Input("ruleId")]
        public Input<string>? RuleId { get; set; }

        /// <summary>
        /// The governance rule priority, priority to the lower number. Rules with the same priority on the same scope will not be allowed
        /// </summary>
        [Input("rulePriority", required: true)]
        public Input<int> RulePriority { get; set; } = null!;

        /// <summary>
        /// The rule type of the governance rule, defines the source of the rule e.g. Integrated
        /// </summary>
        [Input("ruleType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Security.GovernanceRuleType> RuleType { get; set; } = null!;

        /// <summary>
        /// The scope of the Governance rules. Valid scopes are: management group (format: 'providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: 'subscriptions/{subscriptionId}'), or security connector (format: 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        /// <summary>
        /// The governance rule source, what the rule affects, e.g. Assessments
        /// </summary>
        [Input("sourceResourceType", required: true)]
        public InputUnion<string, Pulumi.AzureNative.Security.GovernanceRuleSourceResourceType> SourceResourceType { get; set; } = null!;

        public GovernanceRuleArgs()
        {
        }
        public static new GovernanceRuleArgs Empty => new GovernanceRuleArgs();
    }
}
