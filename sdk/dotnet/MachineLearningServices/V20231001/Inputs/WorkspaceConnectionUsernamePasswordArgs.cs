// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20231001.Inputs
{

    public sealed class WorkspaceConnectionUsernamePasswordArgs : global::Pulumi.ResourceArgs
    {
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("username")]
        public Input<string>? Username { get; set; }

        public WorkspaceConnectionUsernamePasswordArgs()
        {
        }
        public static new WorkspaceConnectionUsernamePasswordArgs Empty => new WorkspaceConnectionUsernamePasswordArgs();
    }
}
