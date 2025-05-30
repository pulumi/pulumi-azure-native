// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// Account key datastore credentials configuration.
    /// </summary>
    public sealed class AccountKeyDatastoreCredentialsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enum to determine the datastore credentials type.
        /// Expected value is 'AccountKey'.
        /// </summary>
        [Input("credentialsType", required: true)]
        public Input<string> CredentialsType { get; set; } = null!;

        /// <summary>
        /// [Required] Storage account secrets.
        /// </summary>
        [Input("secrets", required: true)]
        public Input<Inputs.AccountKeyDatastoreSecretsArgs> Secrets { get; set; } = null!;

        public AccountKeyDatastoreCredentialsArgs()
        {
        }
        public static new AccountKeyDatastoreCredentialsArgs Empty => new AccountKeyDatastoreCredentialsArgs();
    }
}
