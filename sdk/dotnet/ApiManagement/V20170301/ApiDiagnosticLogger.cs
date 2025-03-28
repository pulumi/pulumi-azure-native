// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.V20170301
{
    /// <summary>
    /// Logger details.
    /// </summary>
    [AzureNativeResourceType("azure-native:apimanagement/v20170301:ApiDiagnosticLogger")]
    public partial class ApiDiagnosticLogger : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name and SendRule connection string of the event hub for azureEventHub logger.
        /// Instrumentation key for applicationInsights logger.
        /// </summary>
        [Output("credentials")]
        public Output<ImmutableDictionary<string, string>> Credentials { get; private set; } = null!;

        /// <summary>
        /// Logger description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether records are buffered in the logger before publishing. Default is assumed to be true.
        /// </summary>
        [Output("isBuffered")]
        public Output<bool?> IsBuffered { get; private set; } = null!;

        /// <summary>
        /// Logger type.
        /// </summary>
        [Output("loggerType")]
        public Output<string> LoggerType { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Sampling settings for an ApplicationInsights logger.
        /// </summary>
        [Output("sampling")]
        public Output<Outputs.LoggerSamplingContractResponse?> Sampling { get; private set; } = null!;

        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ApiDiagnosticLogger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiDiagnosticLogger(string name, ApiDiagnosticLoggerArgs args, CustomResourceOptions? options = null)
            : base("azure-native:apimanagement/v20170301:ApiDiagnosticLogger", name, args ?? new ApiDiagnosticLoggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiDiagnosticLogger(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:apimanagement/v20170301:ApiDiagnosticLogger", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:apimanagement/v20180101:ApiDiagnosticLogger" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApiDiagnosticLogger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiDiagnosticLogger Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ApiDiagnosticLogger(name, id, options);
        }
    }

    public sealed class ApiDiagnosticLoggerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// API identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// Diagnostic identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("diagnosticId", required: true)]
        public Input<string> DiagnosticId { get; set; } = null!;

        /// <summary>
        /// Logger identifier. Must be unique in the API Management service instance.
        /// </summary>
        [Input("loggerid")]
        public Input<string>? Loggerid { get; set; }

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

        public ApiDiagnosticLoggerArgs()
        {
        }
        public static new ApiDiagnosticLoggerArgs Empty => new ApiDiagnosticLoggerArgs();
    }
}
