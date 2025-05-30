// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Outputs
{

    /// <summary>
    /// SAS datastore credentials configuration.
    /// </summary>
    [OutputType]
    public sealed class SasDatastoreCredentialsResponse
    {
        /// <summary>
        /// Enum to determine the datastore credentials type.
        /// Expected value is 'Sas'.
        /// </summary>
        public readonly string CredentialsType;

        [OutputConstructor]
        private SasDatastoreCredentialsResponse(string credentialsType)
        {
            CredentialsType = credentialsType;
        }
    }
}
