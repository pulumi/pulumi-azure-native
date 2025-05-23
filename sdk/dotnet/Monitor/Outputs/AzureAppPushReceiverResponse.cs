// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Monitor.Outputs
{

    /// <summary>
    /// The Azure mobile App push notification receiver.
    /// </summary>
    [OutputType]
    public sealed class AzureAppPushReceiverResponse
    {
        /// <summary>
        /// The email address registered for the Azure mobile app.
        /// </summary>
        public readonly string EmailAddress;
        /// <summary>
        /// The name of the Azure mobile app push receiver. Names must be unique across all receivers within a tenant action group.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private AzureAppPushReceiverResponse(
            string emailAddress,

            string name)
        {
            EmailAddress = emailAddress;
            Name = name;
        }
    }
}
