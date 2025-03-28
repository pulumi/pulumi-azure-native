// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20250101.Inputs
{

    /// <summary>
    /// Certificate resource specific properties
    /// </summary>
    public sealed class ManagedCertificatePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Selected type of domain control validation for managed certificates.
        /// </summary>
        [Input("domainControlValidation")]
        public InputUnion<string, Pulumi.AzureNative.App.V20250101.ManagedCertificateDomainControlValidation>? DomainControlValidation { get; set; }

        /// <summary>
        /// Subject name of the certificate.
        /// </summary>
        [Input("subjectName")]
        public Input<string>? SubjectName { get; set; }

        public ManagedCertificatePropertiesArgs()
        {
        }
        public static new ManagedCertificatePropertiesArgs Empty => new ManagedCertificatePropertiesArgs();
    }
}
