# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['AksAssessmentOperationArgs', 'AksAssessmentOperation']

@pulumi.input_type
class AksAssessmentOperationArgs:
    def __init__(__self__, *,
                 project_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 settings: pulumi.Input['AKSAssessmentSettingsArgs'],
                 assessment_name: Optional[pulumi.Input[builtins.str]] = None,
                 scope: Optional[pulumi.Input['AssessmentScopeParametersArgs']] = None):
        """
        The set of arguments for constructing a AksAssessmentOperation resource.
        :param pulumi.Input[builtins.str] project_name: Assessment Project Name
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input['AKSAssessmentSettingsArgs'] settings: Gets or sets AKS Assessment Settings.
        :param pulumi.Input[builtins.str] assessment_name: AKS Assessment Name.
        :param pulumi.Input['AssessmentScopeParametersArgs'] scope: Gets or sets scope parameters to identify inventory items for assessment.
        """
        pulumi.set(__self__, "project_name", project_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "settings", settings)
        if assessment_name is not None:
            pulumi.set(__self__, "assessment_name", assessment_name)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Input[builtins.str]:
        """
        Assessment Project Name
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the resource group. The name is case insensitive.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter
    def settings(self) -> pulumi.Input['AKSAssessmentSettingsArgs']:
        """
        Gets or sets AKS Assessment Settings.
        """
        return pulumi.get(self, "settings")

    @settings.setter
    def settings(self, value: pulumi.Input['AKSAssessmentSettingsArgs']):
        pulumi.set(self, "settings", value)

    @property
    @pulumi.getter(name="assessmentName")
    def assessment_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        AKS Assessment Name.
        """
        return pulumi.get(self, "assessment_name")

    @assessment_name.setter
    def assessment_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "assessment_name", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input['AssessmentScopeParametersArgs']]:
        """
        Gets or sets scope parameters to identify inventory items for assessment.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input['AssessmentScopeParametersArgs']]):
        pulumi.set(self, "scope", value)


@pulumi.type_token("azure-native:migrate:AksAssessmentOperation")
class AksAssessmentOperation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assessment_name: Optional[pulumi.Input[builtins.str]] = None,
                 project_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 scope: Optional[pulumi.Input[Union['AssessmentScopeParametersArgs', 'AssessmentScopeParametersArgsDict']]] = None,
                 settings: Optional[pulumi.Input[Union['AKSAssessmentSettingsArgs', 'AKSAssessmentSettingsArgsDict']]] = None,
                 __props__=None):
        """
        ARM model of AKS Assessment.

        Uses Azure REST API version 2024-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-04-01-preview.

        Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-09-09-preview, 2024-01-15, 2024-03-03-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] assessment_name: AKS Assessment Name.
        :param pulumi.Input[builtins.str] project_name: Assessment Project Name
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[Union['AssessmentScopeParametersArgs', 'AssessmentScopeParametersArgsDict']] scope: Gets or sets scope parameters to identify inventory items for assessment.
        :param pulumi.Input[Union['AKSAssessmentSettingsArgs', 'AKSAssessmentSettingsArgsDict']] settings: Gets or sets AKS Assessment Settings.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AksAssessmentOperationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ARM model of AKS Assessment.

        Uses Azure REST API version 2024-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-04-01-preview.

        Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-09-09-preview, 2024-01-15, 2024-03-03-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param AksAssessmentOperationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AksAssessmentOperationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assessment_name: Optional[pulumi.Input[builtins.str]] = None,
                 project_name: Optional[pulumi.Input[builtins.str]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 scope: Optional[pulumi.Input[Union['AssessmentScopeParametersArgs', 'AssessmentScopeParametersArgsDict']]] = None,
                 settings: Optional[pulumi.Input[Union['AKSAssessmentSettingsArgs', 'AKSAssessmentSettingsArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AksAssessmentOperationArgs.__new__(AksAssessmentOperationArgs)

            __props__.__dict__["assessment_name"] = assessment_name
            if project_name is None and not opts.urn:
                raise TypeError("Missing required property 'project_name'")
            __props__.__dict__["project_name"] = project_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["scope"] = scope
            if settings is None and not opts.urn:
                raise TypeError("Missing required property 'settings'")
            __props__.__dict__["settings"] = settings
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["details"] = None
            __props__.__dict__["e_tag"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:migrate/v20230401preview:AksAssessmentOperation"), pulumi.Alias(type_="azure-native:migrate/v20230501preview:AksAssessmentOperation"), pulumi.Alias(type_="azure-native:migrate/v20230909preview:AksAssessmentOperation"), pulumi.Alias(type_="azure-native:migrate/v20240101preview:AksAssessmentOperation"), pulumi.Alias(type_="azure-native:migrate/v20240115:AksAssessmentOperation"), pulumi.Alias(type_="azure-native:migrate/v20240303preview:AksAssessmentOperation")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AksAssessmentOperation, __self__).__init__(
            'azure-native:migrate:AksAssessmentOperation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AksAssessmentOperation':
        """
        Get an existing AksAssessmentOperation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AksAssessmentOperationArgs.__new__(AksAssessmentOperationArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["details"] = None
        __props__.__dict__["e_tag"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["scope"] = None
        __props__.__dict__["settings"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return AksAssessmentOperation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def details(self) -> pulumi.Output['outputs.AKSAssessmentDetailsResponse']:
        """
        Gets AKS Assessment Details.
        """
        return pulumi.get(self, "details")

    @property
    @pulumi.getter(name="eTag")
    def e_tag(self) -> pulumi.Output[builtins.str]:
        """
        If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields.
        """
        return pulumi.get(self, "e_tag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        Gets the provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[Optional['outputs.AssessmentScopeParametersResponse']]:
        """
        Gets or sets scope parameters to identify inventory items for assessment.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def settings(self) -> pulumi.Output['outputs.AKSAssessmentSettingsResponse']:
        """
        Gets or sets AKS Assessment Settings.
        """
        return pulumi.get(self, "settings")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

