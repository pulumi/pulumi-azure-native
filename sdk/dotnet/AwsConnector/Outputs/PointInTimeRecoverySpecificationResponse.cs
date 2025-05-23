// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Outputs
{

    /// <summary>
    /// Definition of PointInTimeRecoverySpecification
    /// </summary>
    [OutputType]
    public sealed class PointInTimeRecoverySpecificationResponse
    {
        /// <summary>
        /// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table.
        /// </summary>
        public readonly bool? PointInTimeRecoveryEnabled;

        [OutputConstructor]
        private PointInTimeRecoverySpecificationResponse(bool? pointInTimeRecoveryEnabled)
        {
            PointInTimeRecoveryEnabled = pointInTimeRecoveryEnabled;
        }
    }
}
