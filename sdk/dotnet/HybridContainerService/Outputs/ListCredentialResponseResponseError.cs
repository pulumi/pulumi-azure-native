// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridContainerService.Outputs
{

    [OutputType]
    public sealed class ListCredentialResponseResponseError
    {
        public readonly string? Code;
        public readonly string? Message;

        [OutputConstructor]
        private ListCredentialResponseResponseError(
            string? code,

            string? message)
        {
            Code = code;
            Message = message;
        }
    }
}
