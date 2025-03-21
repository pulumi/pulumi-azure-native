// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.V20180101
{
    public static class GetDiagnostic
    {
        /// <summary>
        /// Gets the details of the Diagnostic specified by its identifier.
        /// </summary>
        public static Task<GetDiagnosticResult> InvokeAsync(GetDiagnosticArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDiagnosticResult>("azure-native:apimanagement/v20180101:getDiagnostic", args ?? new GetDiagnosticArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the Diagnostic specified by its identifier.
        /// </summary>
        public static Output<GetDiagnosticResult> Invoke(GetDiagnosticInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDiagnosticResult>("azure-native:apimanagement/v20180101:getDiagnostic", args ?? new GetDiagnosticInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of the Diagnostic specified by its identifier.
        /// </summary>
        public static Output<GetDiagnosticResult> Invoke(GetDiagnosticInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDiagnosticResult>("azure-native:apimanagement/v20180101:getDiagnostic", args ?? new GetDiagnosticInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDiagnosticArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Diagnostic identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("diagnosticId", required: true)]
        public string DiagnosticId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetDiagnosticArgs()
        {
        }
        public static new GetDiagnosticArgs Empty => new GetDiagnosticArgs();
    }

    public sealed class GetDiagnosticInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Diagnostic identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("diagnosticId", required: true)]
        public Input<string> DiagnosticId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetDiagnosticInvokeArgs()
        {
        }
        public static new GetDiagnosticInvokeArgs Empty => new GetDiagnosticInvokeArgs();
    }


    [OutputType]
    public sealed class GetDiagnosticResult
    {
        /// <summary>
        /// Indicates whether a diagnostic should receive data or not.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDiagnosticResult(
            bool enabled,

            string id,

            string name,

            string type)
        {
            Enabled = enabled;
            Id = id;
            Name = name;
            Type = type;
        }
    }
}
