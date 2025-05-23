// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataProtection.Outputs
{

    /// <summary>
    /// Policy Info in backupInstance
    /// </summary>
    [OutputType]
    public sealed class PolicyInfoResponse
    {
        public readonly string PolicyId;
        /// <summary>
        /// Policy parameters for the backup instance
        /// </summary>
        public readonly Outputs.PolicyParametersResponse? PolicyParameters;
        public readonly string PolicyVersion;

        [OutputConstructor]
        private PolicyInfoResponse(
            string policyId,

            Outputs.PolicyParametersResponse? policyParameters,

            string policyVersion)
        {
            PolicyId = policyId;
            PolicyParameters = policyParameters;
            PolicyVersion = policyVersion;
        }
    }
}
