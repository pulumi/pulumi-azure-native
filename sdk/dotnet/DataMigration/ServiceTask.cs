// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataMigration
{
    /// <summary>
    /// A task resource
    /// 
    /// Uses Azure REST API version 2023-07-15-preview. In version 2.x of the Azure Native provider, it used API version 2021-06-30.
    /// 
    /// Other available API versions: 2021-06-30, 2021-10-30-preview, 2022-01-30-preview, 2022-03-30-preview, 2025-03-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datamigration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:datamigration:ServiceTask")]
    public partial class ServiceTask : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// HTTP strong entity tag value. This is ignored if submitted.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Custom task properties
        /// </summary>
        [Output("properties")]
        public Output<object> Properties { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceTask resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceTask(string name, ServiceTaskArgs args, CustomResourceOptions? options = null)
            : base("azure-native:datamigration:ServiceTask", name, args ?? new ServiceTaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceTask(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:datamigration:ServiceTask", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20180715preview:ServiceTask" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20210630:ServiceTask" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20211030preview:ServiceTask" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20220130preview:ServiceTask" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20220330preview:ServiceTask" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20230715preview:ServiceTask" },
                    new global::Pulumi.Alias { Type = "azure-native:datamigration/v20250315preview:ServiceTask" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceTask resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceTask Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServiceTask(name, id, options);
        }
    }

    public sealed class ServiceTaskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the resource group
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// Custom task properties
        /// </summary>
        [Input("properties")]
        public object? Properties { get; set; }

        /// <summary>
        /// Name of the service
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Name of the Task
        /// </summary>
        [Input("taskName")]
        public Input<string>? TaskName { get; set; }

        public ServiceTaskArgs()
        {
        }
        public static new ServiceTaskArgs Empty => new ServiceTaskArgs();
    }
}
