// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Security.Outputs
{

    /// <summary>
    /// Configuration payload for PR Annotations.
    /// </summary>
    [OutputType]
    public sealed class ActionableRemediationResponse
    {
        /// <summary>
        /// Repository branch configuration for PR Annotations.
        /// </summary>
        public readonly Outputs.TargetBranchConfigurationResponse? BranchConfiguration;
        /// <summary>
        /// Gets or sets list of categories and severity levels.
        /// </summary>
        public readonly ImmutableArray<Outputs.CategoryConfigurationResponse> CategoryConfigurations;
        /// <summary>
        /// Update Settings.
        /// 
        /// Enabled - Resource should inherit configurations from parent.
        /// Disabled - Resource should not inherit configurations from parent.
        /// </summary>
        public readonly string? InheritFromParentState;
        /// <summary>
        /// ActionableRemediation Setting.
        /// None - the setting was never set.
        /// Enabled - ActionableRemediation is enabled.
        /// Disabled - ActionableRemediation is disabled.
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private ActionableRemediationResponse(
            Outputs.TargetBranchConfigurationResponse? branchConfiguration,

            ImmutableArray<Outputs.CategoryConfigurationResponse> categoryConfigurations,

            string? inheritFromParentState,

            string? state)
        {
            BranchConfiguration = branchConfiguration;
            CategoryConfigurations = categoryConfigurations;
            InheritFromParentState = inheritFromParentState;
            State = state;
        }
    }
}
