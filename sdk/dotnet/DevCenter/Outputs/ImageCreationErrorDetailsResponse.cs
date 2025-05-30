// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DevCenter.Outputs
{

    /// <summary>
    /// Image creation error details
    /// </summary>
    [OutputType]
    public sealed class ImageCreationErrorDetailsResponse
    {
        /// <summary>
        /// An identifier for the error.
        /// </summary>
        public readonly string? Code;
        /// <summary>
        /// A message describing the error.
        /// </summary>
        public readonly string? Message;

        [OutputConstructor]
        private ImageCreationErrorDetailsResponse(
            string? code,

            string? message)
        {
            Code = code;
            Message = message;
        }
    }
}
