// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media.Outputs
{

    [OutputType]
    public sealed class KeyDeliveryResponse
    {
        /// <summary>
        /// The access control properties for Key Delivery.
        /// </summary>
        public readonly Outputs.AccessControlResponse? AccessControl;

        [OutputConstructor]
        private KeyDeliveryResponse(Outputs.AccessControlResponse? accessControl)
        {
            AccessControl = accessControl;
        }
    }
}
