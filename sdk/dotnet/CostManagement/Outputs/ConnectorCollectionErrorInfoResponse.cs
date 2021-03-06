// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement.Outputs
{

    [OutputType]
    public sealed class ConnectorCollectionErrorInfoResponse
    {
        /// <summary>
        /// Short error message
        /// </summary>
        public readonly string ErrorCode;
        /// <summary>
        /// Detailed error message
        /// </summary>
        public readonly string ErrorMessage;
        /// <summary>
        /// Time the error started occurring (Last time error occurred in lastRun)
        /// </summary>
        public readonly string ErrorStartTime;

        [OutputConstructor]
        private ConnectorCollectionErrorInfoResponse(
            string errorCode,

            string errorMessage,

            string errorStartTime)
        {
            ErrorCode = errorCode;
            ErrorMessage = errorMessage;
            ErrorStartTime = errorStartTime;
        }
    }
}
