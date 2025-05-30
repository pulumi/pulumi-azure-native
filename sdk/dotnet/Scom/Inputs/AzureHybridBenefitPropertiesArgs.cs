// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Scom.Inputs
{

    /// <summary>
    /// The properties to maximize savings by using Azure Hybrid Benefit
    /// </summary>
    public sealed class AzureHybridBenefitPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SCOM license type. Maximize savings by using license you already own
        /// </summary>
        [Input("scomLicenseType")]
        public InputUnion<string, Pulumi.AzureNative.Scom.HybridLicenseType>? ScomLicenseType { get; set; }

        /// <summary>
        /// SQL Server license type. Maximize savings by using Azure Hybrid Benefit for SQL Server with license you already own
        /// </summary>
        [Input("sqlServerLicenseType")]
        public InputUnion<string, Pulumi.AzureNative.Scom.HybridLicenseType>? SqlServerLicenseType { get; set; }

        /// <summary>
        /// Specifies that the image or disk that is being used was licensed on-premises. &lt;br&gt;&lt;br&gt; For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json)
        /// </summary>
        [Input("windowsServerLicenseType")]
        public InputUnion<string, Pulumi.AzureNative.Scom.HybridLicenseType>? WindowsServerLicenseType { get; set; }

        public AzureHybridBenefitPropertiesArgs()
        {
        }
        public static new AzureHybridBenefitPropertiesArgs Empty => new AzureHybridBenefitPropertiesArgs();
    }
}
