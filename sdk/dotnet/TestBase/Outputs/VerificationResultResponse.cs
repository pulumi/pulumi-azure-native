// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.TestBase.Outputs
{

    /// <summary>
    /// The detailed result of a validation or rule checking.
    /// </summary>
    [OutputType]
    public sealed class VerificationResultResponse
    {
        /// <summary>
        /// Message for clarification.
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// Indicates if the validation or rule checking is passed.
        /// </summary>
        public readonly string Result;
        /// <summary>
        /// The name of the verification rule.
        /// </summary>
        public readonly string? VerificationName;

        [OutputConstructor]
        private VerificationResultResponse(
            string? message,

            string result,

            string? verificationName)
        {
            Message = message;
            Result = result;
            VerificationName = verificationName;
        }
    }
}
