// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureDataTransfer.Outputs
{

    /// <summary>
    /// Properties of connection
    /// </summary>
    [OutputType]
    public sealed class ConnectionPropertiesResponse
    {
        /// <summary>
        /// Approver of this connection request
        /// </summary>
        public readonly string Approver;
        /// <summary>
        /// The timestamp that this connection request was submitted at
        /// </summary>
        public readonly string DateSubmitted;
        /// <summary>
        /// Direction of data movement
        /// </summary>
        public readonly string? Direction;
        /// <summary>
        /// The flow types being requested for this connection
        /// </summary>
        public readonly ImmutableArray<string> FlowTypes;
        /// <summary>
        /// Justification for the connection request
        /// </summary>
        public readonly string? Justification;
        /// <summary>
        /// Link status of the current connection
        /// </summary>
        public readonly string LinkStatus;
        /// <summary>
        /// Resource ID of the linked connection
        /// </summary>
        public readonly string LinkedConnectionId;
        /// <summary>
        /// PIN to link requests together
        /// </summary>
        public readonly string? Pin;
        /// <summary>
        /// Pipeline to use to transfer data
        /// </summary>
        public readonly string Pipeline;
        /// <summary>
        /// The policies for this connection
        /// </summary>
        public readonly ImmutableArray<string> Policies;
        /// <summary>
        /// The primary contact for this connection request
        /// </summary>
        public readonly string? PrimaryContact;
        /// <summary>
        /// Provisioning state of the connection
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Subscription ID to link cloud subscriptions together
        /// </summary>
        public readonly string? RemoteSubscriptionId;
        /// <summary>
        /// Requirement ID of the connection
        /// </summary>
        public readonly string? RequirementId;
        /// <summary>
        /// The schema URIs for this connection
        /// </summary>
        public readonly ImmutableArray<string> SchemaUris;
        /// <summary>
        /// The schemas for this connection
        /// </summary>
        public readonly ImmutableArray<Outputs.SchemaResponse> Schemas;
        /// <summary>
        /// The secondary contacts for this connection request
        /// </summary>
        public readonly ImmutableArray<string> SecondaryContacts;
        /// <summary>
        /// Status of the connection
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Reason for status
        /// </summary>
        public readonly string StatusReason;

        [OutputConstructor]
        private ConnectionPropertiesResponse(
            string approver,

            string dateSubmitted,

            string? direction,

            ImmutableArray<string> flowTypes,

            string? justification,

            string linkStatus,

            string linkedConnectionId,

            string? pin,

            string pipeline,

            ImmutableArray<string> policies,

            string? primaryContact,

            string provisioningState,

            string? remoteSubscriptionId,

            string? requirementId,

            ImmutableArray<string> schemaUris,

            ImmutableArray<Outputs.SchemaResponse> schemas,

            ImmutableArray<string> secondaryContacts,

            string status,

            string statusReason)
        {
            Approver = approver;
            DateSubmitted = dateSubmitted;
            Direction = direction;
            FlowTypes = flowTypes;
            Justification = justification;
            LinkStatus = linkStatus;
            LinkedConnectionId = linkedConnectionId;
            Pin = pin;
            Pipeline = pipeline;
            Policies = policies;
            PrimaryContact = primaryContact;
            ProvisioningState = provisioningState;
            RemoteSubscriptionId = remoteSubscriptionId;
            RequirementId = requirementId;
            SchemaUris = schemaUris;
            Schemas = schemas;
            SecondaryContacts = secondaryContacts;
            Status = status;
            StatusReason = statusReason;
        }
    }
}
