// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Outputs
{

    /// <summary>
    /// Skip error file.
    /// </summary>
    [OutputType]
    public sealed class SkipErrorFileResponse
    {
        /// <summary>
        /// Skip if source/sink file changed by other concurrent write. Default is false. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public readonly object? DataInconsistency;
        /// <summary>
        /// Skip if file is deleted by other client during copy. Default is true. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public readonly object? FileMissing;

        [OutputConstructor]
        private SkipErrorFileResponse(
            object? dataInconsistency,

            object? fileMissing)
        {
            DataInconsistency = dataInconsistency;
            FileMissing = fileMissing;
        }
    }
}
