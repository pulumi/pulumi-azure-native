// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Authorization.Inputs
{

    /// <summary>
    /// The role management policy expiration rule.
    /// </summary>
    public sealed class RoleManagementPolicyExpirationRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("exceptionMembers")]
        private InputList<Inputs.UserSetArgs>? _exceptionMembers;

        /// <summary>
        /// The members not restricted by expiration rule.
        /// </summary>
        public InputList<Inputs.UserSetArgs> ExceptionMembers
        {
            get => _exceptionMembers ?? (_exceptionMembers = new InputList<Inputs.UserSetArgs>());
            set => _exceptionMembers = value;
        }

        /// <summary>
        /// The id of the rule.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The value indicating whether expiration is required.
        /// </summary>
        [Input("isExpirationRequired")]
        public Input<bool>? IsExpirationRequired { get; set; }

        /// <summary>
        /// The maximum duration of expiration in timespan.
        /// </summary>
        [Input("maximumDuration")]
        public Input<string>? MaximumDuration { get; set; }

        /// <summary>
        /// The type of rule
        /// Expected value is 'RoleManagementPolicyExpirationRule'.
        /// </summary>
        [Input("ruleType", required: true)]
        public Input<string> RuleType { get; set; } = null!;

        /// <summary>
        /// The target of the current rule.
        /// </summary>
        [Input("target")]
        public Input<Inputs.RoleManagementPolicyRuleTargetArgs>? Target { get; set; }

        public RoleManagementPolicyExpirationRuleArgs()
        {
        }
        public static new RoleManagementPolicyExpirationRuleArgs Empty => new RoleManagementPolicyExpirationRuleArgs();
    }
}
