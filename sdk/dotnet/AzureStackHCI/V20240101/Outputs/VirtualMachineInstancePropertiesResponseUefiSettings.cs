// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureStackHCI.V20240101.Outputs
{

    [OutputType]
    public sealed class VirtualMachineInstancePropertiesResponseUefiSettings
    {
        /// <summary>
        /// Specifies whether secure boot should be enabled on the virtual machine instance.
        /// </summary>
        public readonly bool? SecureBootEnabled;

        [OutputConstructor]
        private VirtualMachineInstancePropertiesResponseUefiSettings(bool? secureBootEnabled)
        {
            SecureBootEnabled = secureBootEnabled;
        }
    }
}
