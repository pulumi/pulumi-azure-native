// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.Outputs
{

    /// <summary>
    /// Certificate resource specific properties
    /// </summary>
    [OutputType]
    public sealed class ManagedCertificateResponseProperties
    {
        /// <summary>
        /// Selected type of domain control validation for managed certificates.
        /// </summary>
        public readonly string? DomainControlValidation;
        /// <summary>
        /// Any error occurred during the certificate provision.
        /// </summary>
        public readonly string Error;
        /// <summary>
        /// Provisioning state of the certificate.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Subject name of the certificate.
        /// </summary>
        public readonly string? SubjectName;
        /// <summary>
        /// A TXT token used for DNS TXT domain control validation when issuing this type of managed certificates.
        /// </summary>
        public readonly string ValidationToken;

        [OutputConstructor]
        private ManagedCertificateResponseProperties(
            string? domainControlValidation,

            string error,

            string provisioningState,

            string? subjectName,

            string validationToken)
        {
            DomainControlValidation = domainControlValidation;
            Error = error;
            ProvisioningState = provisioningState;
            SubjectName = subjectName;
            ValidationToken = validationToken;
        }
    }
}
