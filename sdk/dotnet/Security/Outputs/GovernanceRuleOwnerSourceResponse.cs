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
    /// Describe the owner source of governance rule
    /// </summary>
    [OutputType]
    public sealed class GovernanceRuleOwnerSourceResponse
    {
        /// <summary>
        /// The owner type for the governance rule owner source
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The source value e.g. tag key like owner name or email address
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private GovernanceRuleOwnerSourceResponse(
            string? type,

            string? value)
        {
            Type = type;
            Value = value;
        }
    }
}
