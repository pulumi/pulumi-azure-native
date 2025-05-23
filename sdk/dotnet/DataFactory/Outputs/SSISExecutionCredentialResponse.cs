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
    /// SSIS package execution credential.
    /// </summary>
    [OutputType]
    public sealed class SSISExecutionCredentialResponse
    {
        /// <summary>
        /// Domain for windows authentication. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object Domain;
        /// <summary>
        /// Password for windows authentication.
        /// </summary>
        public readonly Outputs.SecureStringResponse Password;
        /// <summary>
        /// UseName for windows authentication. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly object UserName;

        [OutputConstructor]
        private SSISExecutionCredentialResponse(
            object domain,

            Outputs.SecureStringResponse password,

            object userName)
        {
            Domain = domain;
            Password = password;
            UserName = userName;
        }
    }
}
