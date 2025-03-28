// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Insights
{
    public static class GetAutoscaleSetting
    {
        /// <summary>
        /// Gets an autoscale setting
        /// 
        /// Uses Azure REST API version 2022-10-01.
        /// </summary>
        public static Task<GetAutoscaleSettingResult> InvokeAsync(GetAutoscaleSettingArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAutoscaleSettingResult>("azure-native:insights:getAutoscaleSetting", args ?? new GetAutoscaleSettingArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an autoscale setting
        /// 
        /// Uses Azure REST API version 2022-10-01.
        /// </summary>
        public static Output<GetAutoscaleSettingResult> Invoke(GetAutoscaleSettingInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAutoscaleSettingResult>("azure-native:insights:getAutoscaleSetting", args ?? new GetAutoscaleSettingInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an autoscale setting
        /// 
        /// Uses Azure REST API version 2022-10-01.
        /// </summary>
        public static Output<GetAutoscaleSettingResult> Invoke(GetAutoscaleSettingInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAutoscaleSettingResult>("azure-native:insights:getAutoscaleSetting", args ?? new GetAutoscaleSettingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAutoscaleSettingArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The autoscale setting name.
        /// </summary>
        [Input("autoscaleSettingName", required: true)]
        public string AutoscaleSettingName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAutoscaleSettingArgs()
        {
        }
        public static new GetAutoscaleSettingArgs Empty => new GetAutoscaleSettingArgs();
    }

    public sealed class GetAutoscaleSettingInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The autoscale setting name.
        /// </summary>
        [Input("autoscaleSettingName", required: true)]
        public Input<string> AutoscaleSettingName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public GetAutoscaleSettingInvokeArgs()
        {
        }
        public static new GetAutoscaleSettingInvokeArgs Empty => new GetAutoscaleSettingInvokeArgs();
    }


    [OutputType]
    public sealed class GetAutoscaleSettingResult
    {
        /// <summary>
        /// the enabled flag. Specifies whether automatic scaling is enabled for the resource. The default value is 'false'.
        /// </summary>
        public readonly bool? Enabled;
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
        /// the collection of notifications.
        /// </summary>
        public readonly ImmutableArray<Outputs.AutoscaleNotificationResponse> Notifications;
        /// <summary>
        /// the predictive autoscale policy mode.
        /// </summary>
        public readonly Outputs.PredictiveAutoscalePolicyResponse? PredictiveAutoscalePolicy;
        /// <summary>
        /// the collection of automatic scaling profiles that specify different scaling parameters for different time periods. A maximum of 20 profiles can be specified.
        /// </summary>
        public readonly ImmutableArray<Outputs.AutoscaleProfileResponse> Profiles;
        /// <summary>
        /// The system metadata related to the response.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater in length than 128 characters and a value no greater in length than 256 characters.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// the location of the resource that the autoscale setting should be added to.
        /// </summary>
        public readonly string? TargetResourceLocation;
        /// <summary>
        /// the resource identifier of the resource that the autoscale setting should be added to.
        /// </summary>
        public readonly string? TargetResourceUri;
        /// <summary>
        /// Azure resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAutoscaleSettingResult(
            bool? enabled,

            string id,

            string location,

            string name,

            ImmutableArray<Outputs.AutoscaleNotificationResponse> notifications,

            Outputs.PredictiveAutoscalePolicyResponse? predictiveAutoscalePolicy,

            ImmutableArray<Outputs.AutoscaleProfileResponse> profiles,

            Outputs.SystemDataResponse systemData,

            ImmutableDictionary<string, string>? tags,

            string? targetResourceLocation,

            string? targetResourceUri,

            string type)
        {
            Enabled = enabled;
            Id = id;
            Location = location;
            Name = name;
            Notifications = notifications;
            PredictiveAutoscalePolicy = predictiveAutoscalePolicy;
            Profiles = profiles;
            SystemData = systemData;
            Tags = tags;
            TargetResourceLocation = targetResourceLocation;
            TargetResourceUri = targetResourceUri;
            Type = type;
        }
    }
}
