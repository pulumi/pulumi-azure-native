// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HealthcareApis.Inputs
{

    /// <summary>
    /// The encryption settings for the customer-managed key
    /// </summary>
    public sealed class EncryptionCustomerManagedKeyEncryptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL of the key to use for encryption
        /// </summary>
        [Input("keyEncryptionKeyUrl")]
        public Input<string>? KeyEncryptionKeyUrl { get; set; }

        public EncryptionCustomerManagedKeyEncryptionArgs()
        {
        }
        public static new EncryptionCustomerManagedKeyEncryptionArgs Empty => new EncryptionCustomerManagedKeyEncryptionArgs();
    }
}
