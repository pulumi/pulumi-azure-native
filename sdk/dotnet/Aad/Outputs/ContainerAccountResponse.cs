// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Aad.Outputs
{

    /// <summary>
    /// Container Account Description
    /// </summary>
    [OutputType]
    public sealed class ContainerAccountResponse
    {
        /// <summary>
        /// The account name
        /// </summary>
        public readonly string? AccountName;
        /// <summary>
        /// The account password
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// The account spn
        /// </summary>
        public readonly string? Spn;

        [OutputConstructor]
        private ContainerAccountResponse(
            string? accountName,

            string? password,

            string? spn)
        {
            AccountName = accountName;
            Password = password;
            Spn = spn;
        }
    }
}
