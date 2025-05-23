// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AppPlatform.Outputs
{

    /// <summary>
    /// The error code compose of code and message.
    /// </summary>
    [OutputType]
    public sealed class ErrorResponse
    {
        /// <summary>
        /// The code of error.
        /// </summary>
        public readonly string? Code;
        /// <summary>
        /// The message of error.
        /// </summary>
        public readonly string? Message;

        [OutputConstructor]
        private ErrorResponse(
            string? code,

            string? message)
        {
            Code = code;
            Message = message;
        }
    }
}
