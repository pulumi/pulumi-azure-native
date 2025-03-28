// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Insights
{
    public static class GetGuestDiagnosticsSettingsAssociation
    {
        /// <summary>
        /// Gets guest diagnostics association settings.
        /// 
        /// Uses Azure REST API version 2018-06-01-preview.
        /// </summary>
        public static Task<GetGuestDiagnosticsSettingsAssociationResult> InvokeAsync(GetGuestDiagnosticsSettingsAssociationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGuestDiagnosticsSettingsAssociationResult>("azure-native:insights:getGuestDiagnosticsSettingsAssociation", args ?? new GetGuestDiagnosticsSettingsAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// Gets guest diagnostics association settings.
        /// 
        /// Uses Azure REST API version 2018-06-01-preview.
        /// </summary>
        public static Output<GetGuestDiagnosticsSettingsAssociationResult> Invoke(GetGuestDiagnosticsSettingsAssociationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGuestDiagnosticsSettingsAssociationResult>("azure-native:insights:getGuestDiagnosticsSettingsAssociation", args ?? new GetGuestDiagnosticsSettingsAssociationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets guest diagnostics association settings.
        /// 
        /// Uses Azure REST API version 2018-06-01-preview.
        /// </summary>
        public static Output<GetGuestDiagnosticsSettingsAssociationResult> Invoke(GetGuestDiagnosticsSettingsAssociationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetGuestDiagnosticsSettingsAssociationResult>("azure-native:insights:getGuestDiagnosticsSettingsAssociation", args ?? new GetGuestDiagnosticsSettingsAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGuestDiagnosticsSettingsAssociationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the diagnostic settings association.
        /// </summary>
        [Input("associationName", required: true)]
        public string AssociationName { get; set; } = null!;

        /// <summary>
        /// The fully qualified ID of the resource, including the resource name and resource type.
        /// </summary>
        [Input("resourceUri", required: true)]
        public string ResourceUri { get; set; } = null!;

        public GetGuestDiagnosticsSettingsAssociationArgs()
        {
        }
        public static new GetGuestDiagnosticsSettingsAssociationArgs Empty => new GetGuestDiagnosticsSettingsAssociationArgs();
    }

    public sealed class GetGuestDiagnosticsSettingsAssociationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the diagnostic settings association.
        /// </summary>
        [Input("associationName", required: true)]
        public Input<string> AssociationName { get; set; } = null!;

        /// <summary>
        /// The fully qualified ID of the resource, including the resource name and resource type.
        /// </summary>
        [Input("resourceUri", required: true)]
        public Input<string> ResourceUri { get; set; } = null!;

        public GetGuestDiagnosticsSettingsAssociationInvokeArgs()
        {
        }
        public static new GetGuestDiagnosticsSettingsAssociationInvokeArgs Empty => new GetGuestDiagnosticsSettingsAssociationInvokeArgs();
    }


    [OutputType]
    public sealed class GetGuestDiagnosticsSettingsAssociationResult
    {
        /// <summary>
        /// The guest diagnostic settings name.
        /// </summary>
        public readonly string GuestDiagnosticSettingsName;
        /// <summary>
        /// Azure resource Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Azure resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Azure resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetGuestDiagnosticsSettingsAssociationResult(
            string guestDiagnosticSettingsName,

            string id,

            string location,

            string name,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            GuestDiagnosticSettingsName = guestDiagnosticSettingsName;
            Id = id;
            Location = location;
            Name = name;
            Tags = tags;
            Type = type;
        }
    }
}
