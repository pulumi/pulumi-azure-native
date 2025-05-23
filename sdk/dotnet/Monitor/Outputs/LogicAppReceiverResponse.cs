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
    /// A logic app receiver.
    /// </summary>
    [OutputType]
    public sealed class LogicAppReceiverResponse
    {
        /// <summary>
        /// The callback url where http request sent to.
        /// </summary>
        public readonly string CallbackUrl;
        /// <summary>
        /// The principal id of the managed identity. The value can be "None", "SystemAssigned" 
        /// </summary>
        public readonly string? ManagedIdentity;
        /// <summary>
        /// The name of the logic app receiver. Names must be unique across all receivers within an action group.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The azure resource id of the logic app receiver.
        /// </summary>
        public readonly string ResourceId;
        /// <summary>
        /// Indicates whether to use common alert schema.
        /// </summary>
        public readonly bool? UseCommonAlertSchema;

        [OutputConstructor]
        private LogicAppReceiverResponse(
            string callbackUrl,

            string? managedIdentity,

            string name,

            string resourceId,

            bool? useCommonAlertSchema)
        {
            CallbackUrl = callbackUrl;
            ManagedIdentity = managedIdentity;
            Name = name;
            ResourceId = resourceId;
            UseCommonAlertSchema = useCommonAlertSchema;
        }
    }
}
