// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Management.Outputs
{

    /// <summary>
    /// The ID of the parent management group.
    /// </summary>
    [OutputType]
    public sealed class DescendantParentGroupInfoResponse
    {
        /// <summary>
        /// The fully qualified ID for the parent management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private DescendantParentGroupInfoResponse(string? id)
        {
            Id = id;
        }
    }
}
