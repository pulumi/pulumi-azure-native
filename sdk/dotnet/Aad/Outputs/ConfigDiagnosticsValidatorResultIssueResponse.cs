// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Aad.Outputs
{

    /// <summary>
    /// Specific issue for a particular config diagnostics validator
    /// </summary>
    [OutputType]
    public sealed class ConfigDiagnosticsValidatorResultIssueResponse
    {
        /// <summary>
        /// List of domain resource property name or values used to compose a rich description.
        /// </summary>
        public readonly ImmutableArray<string> DescriptionParams;
        /// <summary>
        /// Validation issue identifier.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private ConfigDiagnosticsValidatorResultIssueResponse(
            ImmutableArray<string> descriptionParams,

            string? id)
        {
            DescriptionParams = descriptionParams;
            Id = id;
        }
    }
}
