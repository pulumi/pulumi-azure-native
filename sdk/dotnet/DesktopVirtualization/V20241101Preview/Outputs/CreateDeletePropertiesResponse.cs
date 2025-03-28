// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DesktopVirtualization.V20241101Preview.Outputs
{

    /// <summary>
    /// The properties that control how Scaling will manage the size of the hostpool by creating and deleting hosts.
    /// </summary>
    [OutputType]
    public sealed class CreateDeletePropertiesResponse
    {
        /// <summary>
        /// Maximum number of session hosts that may be created by the Scaling Service. This requires the assigned hostpool to have a session host config property.
        /// </summary>
        public readonly int? RampDownMaximumHostPoolSize;
        /// <summary>
        /// Minimum number of session hosts that will be be created by the Scaling Service. Scaling will not delete any hosts when this limit is met. This requires the assigned hostpool to have a session host config property.
        /// </summary>
        public readonly int? RampDownMinimumHostPoolSize;
        /// <summary>
        /// Maximum number of session hosts that may be created by the Scaling Service. This requires the assigned hostpool to have a session host config property.
        /// </summary>
        public readonly int? RampUpMaximumHostPoolSize;
        /// <summary>
        /// Minimum number of session hosts that will be be created by the Scaling Service. Scaling will not delete any hosts when this limit is met. This requires the assigned hostpool to have a session host config property.
        /// </summary>
        public readonly int? RampUpMinimumHostPoolSize;

        [OutputConstructor]
        private CreateDeletePropertiesResponse(
            int? rampDownMaximumHostPoolSize,

            int? rampDownMinimumHostPoolSize,

            int? rampUpMaximumHostPoolSize,

            int? rampUpMinimumHostPoolSize)
        {
            RampDownMaximumHostPoolSize = rampDownMaximumHostPoolSize;
            RampDownMinimumHostPoolSize = rampDownMinimumHostPoolSize;
            RampUpMaximumHostPoolSize = rampUpMaximumHostPoolSize;
            RampUpMinimumHostPoolSize = rampUpMinimumHostPoolSize;
        }
    }
}
