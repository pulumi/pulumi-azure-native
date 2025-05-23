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

__all__ = ['ScriptArgs', 'Script']

@pulumi.input_type
class ScriptArgs:
    def __init__(__self__, *,
                 cluster_name: pulumi.Input[builtins.str],
                 database_name: pulumi.Input[builtins.str],
                 resource_group_name: pulumi.Input[builtins.str],
                 continue_on_errors: Optional[pulumi.Input[builtins.bool]] = None,
                 force_update_tag: Optional[pulumi.Input[builtins.str]] = None,
                 principal_permissions_action: Optional[pulumi.Input[Union[builtins.str, 'PrincipalPermissionsAction']]] = None,
                 script_content: Optional[pulumi.Input[builtins.str]] = None,
                 script_level: Optional[pulumi.Input[Union[builtins.str, 'ScriptLevel']]] = None,
                 script_name: Optional[pulumi.Input[builtins.str]] = None,
                 script_url: Optional[pulumi.Input[builtins.str]] = None,
                 script_url_sas_token: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Script resource.
        :param pulumi.Input[builtins.str] cluster_name: The name of the Kusto cluster.
        :param pulumi.Input[builtins.str] database_name: The name of the database in the Kusto cluster.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.bool] continue_on_errors: Flag that indicates whether to continue if one of the command fails.
        :param pulumi.Input[builtins.str] force_update_tag: A unique string. If changed the script will be applied again.
        :param pulumi.Input[Union[builtins.str, 'PrincipalPermissionsAction']] principal_permissions_action: Indicates if the permissions for the script caller are kept following completion of the script.
        :param pulumi.Input[builtins.str] script_content: The script content. This property should be used when the script is provide inline and not through file in a SA. Must not be used together with scriptUrl and scriptUrlSasToken properties.
        :param pulumi.Input[Union[builtins.str, 'ScriptLevel']] script_level: Differentiates between the type of script commands included - Database or Cluster. The default is Database.
        :param pulumi.Input[builtins.str] script_name: The name of the Kusto database script.
        :param pulumi.Input[builtins.str] script_url: The url to the KQL script blob file. Must not be used together with scriptContent property
        :param pulumi.Input[builtins.str] script_url_sas_token: The SaS token that provide read access to the file which contain the script. Must be provided when using scriptUrl property.
        """
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        if continue_on_errors is None:
            continue_on_errors = False
        if continue_on_errors is not None:
            pulumi.set(__self__, "continue_on_errors", continue_on_errors)
        if force_update_tag is not None:
            pulumi.set(__self__, "force_update_tag", force_update_tag)
        if principal_permissions_action is not None:
            pulumi.set(__self__, "principal_permissions_action", principal_permissions_action)
        if script_content is not None:
            pulumi.set(__self__, "script_content", script_content)
        if script_level is not None:
            pulumi.set(__self__, "script_level", script_level)
        if script_name is not None:
            pulumi.set(__self__, "script_name", script_name)
        if script_url is not None:
            pulumi.set(__self__, "script_url", script_url)
        if script_url_sas_token is not None:
            pulumi.set(__self__, "script_url_sas_token", script_url_sas_token)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the Kusto cluster.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[builtins.str]:
        """
        The name of the database in the Kusto cluster.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "database_name", value)

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
    @pulumi.getter(name="continueOnErrors")
    def continue_on_errors(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Flag that indicates whether to continue if one of the command fails.
        """
        return pulumi.get(self, "continue_on_errors")

    @continue_on_errors.setter
    def continue_on_errors(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "continue_on_errors", value)

    @property
    @pulumi.getter(name="forceUpdateTag")
    def force_update_tag(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A unique string. If changed the script will be applied again.
        """
        return pulumi.get(self, "force_update_tag")

    @force_update_tag.setter
    def force_update_tag(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "force_update_tag", value)

    @property
    @pulumi.getter(name="principalPermissionsAction")
    def principal_permissions_action(self) -> Optional[pulumi.Input[Union[builtins.str, 'PrincipalPermissionsAction']]]:
        """
        Indicates if the permissions for the script caller are kept following completion of the script.
        """
        return pulumi.get(self, "principal_permissions_action")

    @principal_permissions_action.setter
    def principal_permissions_action(self, value: Optional[pulumi.Input[Union[builtins.str, 'PrincipalPermissionsAction']]]):
        pulumi.set(self, "principal_permissions_action", value)

    @property
    @pulumi.getter(name="scriptContent")
    def script_content(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The script content. This property should be used when the script is provide inline and not through file in a SA. Must not be used together with scriptUrl and scriptUrlSasToken properties.
        """
        return pulumi.get(self, "script_content")

    @script_content.setter
    def script_content(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "script_content", value)

    @property
    @pulumi.getter(name="scriptLevel")
    def script_level(self) -> Optional[pulumi.Input[Union[builtins.str, 'ScriptLevel']]]:
        """
        Differentiates between the type of script commands included - Database or Cluster. The default is Database.
        """
        return pulumi.get(self, "script_level")

    @script_level.setter
    def script_level(self, value: Optional[pulumi.Input[Union[builtins.str, 'ScriptLevel']]]):
        pulumi.set(self, "script_level", value)

    @property
    @pulumi.getter(name="scriptName")
    def script_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the Kusto database script.
        """
        return pulumi.get(self, "script_name")

    @script_name.setter
    def script_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "script_name", value)

    @property
    @pulumi.getter(name="scriptUrl")
    def script_url(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The url to the KQL script blob file. Must not be used together with scriptContent property
        """
        return pulumi.get(self, "script_url")

    @script_url.setter
    def script_url(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "script_url", value)

    @property
    @pulumi.getter(name="scriptUrlSasToken")
    def script_url_sas_token(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The SaS token that provide read access to the file which contain the script. Must be provided when using scriptUrl property.
        """
        return pulumi.get(self, "script_url_sas_token")

    @script_url_sas_token.setter
    def script_url_sas_token(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "script_url_sas_token", value)


@pulumi.type_token("azure-native:kusto:Script")
class Script(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                 continue_on_errors: Optional[pulumi.Input[builtins.bool]] = None,
                 database_name: Optional[pulumi.Input[builtins.str]] = None,
                 force_update_tag: Optional[pulumi.Input[builtins.str]] = None,
                 principal_permissions_action: Optional[pulumi.Input[Union[builtins.str, 'PrincipalPermissionsAction']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 script_content: Optional[pulumi.Input[builtins.str]] = None,
                 script_level: Optional[pulumi.Input[Union[builtins.str, 'ScriptLevel']]] = None,
                 script_name: Optional[pulumi.Input[builtins.str]] = None,
                 script_url: Optional[pulumi.Input[builtins.str]] = None,
                 script_url_sas_token: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Class representing a database script.

        Uses Azure REST API version 2024-04-13. In version 2.x of the Azure Native provider, it used API version 2022-12-29.

        Other available API versions: 2021-01-01, 2021-08-27, 2022-02-01, 2022-07-07, 2022-11-11, 2022-12-29, 2023-05-02, 2023-08-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native kusto [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] cluster_name: The name of the Kusto cluster.
        :param pulumi.Input[builtins.bool] continue_on_errors: Flag that indicates whether to continue if one of the command fails.
        :param pulumi.Input[builtins.str] database_name: The name of the database in the Kusto cluster.
        :param pulumi.Input[builtins.str] force_update_tag: A unique string. If changed the script will be applied again.
        :param pulumi.Input[Union[builtins.str, 'PrincipalPermissionsAction']] principal_permissions_action: Indicates if the permissions for the script caller are kept following completion of the script.
        :param pulumi.Input[builtins.str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[builtins.str] script_content: The script content. This property should be used when the script is provide inline and not through file in a SA. Must not be used together with scriptUrl and scriptUrlSasToken properties.
        :param pulumi.Input[Union[builtins.str, 'ScriptLevel']] script_level: Differentiates between the type of script commands included - Database or Cluster. The default is Database.
        :param pulumi.Input[builtins.str] script_name: The name of the Kusto database script.
        :param pulumi.Input[builtins.str] script_url: The url to the KQL script blob file. Must not be used together with scriptContent property
        :param pulumi.Input[builtins.str] script_url_sas_token: The SaS token that provide read access to the file which contain the script. Must be provided when using scriptUrl property.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ScriptArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Class representing a database script.

        Uses Azure REST API version 2024-04-13. In version 2.x of the Azure Native provider, it used API version 2022-12-29.

        Other available API versions: 2021-01-01, 2021-08-27, 2022-02-01, 2022-07-07, 2022-11-11, 2022-12-29, 2023-05-02, 2023-08-15. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native kusto [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param ScriptArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ScriptArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[builtins.str]] = None,
                 continue_on_errors: Optional[pulumi.Input[builtins.bool]] = None,
                 database_name: Optional[pulumi.Input[builtins.str]] = None,
                 force_update_tag: Optional[pulumi.Input[builtins.str]] = None,
                 principal_permissions_action: Optional[pulumi.Input[Union[builtins.str, 'PrincipalPermissionsAction']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 script_content: Optional[pulumi.Input[builtins.str]] = None,
                 script_level: Optional[pulumi.Input[Union[builtins.str, 'ScriptLevel']]] = None,
                 script_name: Optional[pulumi.Input[builtins.str]] = None,
                 script_url: Optional[pulumi.Input[builtins.str]] = None,
                 script_url_sas_token: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ScriptArgs.__new__(ScriptArgs)

            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            if continue_on_errors is None:
                continue_on_errors = False
            __props__.__dict__["continue_on_errors"] = continue_on_errors
            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            __props__.__dict__["force_update_tag"] = force_update_tag
            __props__.__dict__["principal_permissions_action"] = principal_permissions_action
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            __props__.__dict__["script_content"] = script_content
            __props__.__dict__["script_level"] = script_level
            __props__.__dict__["script_name"] = script_name
            __props__.__dict__["script_url"] = script_url
            __props__.__dict__["script_url_sas_token"] = script_url_sas_token
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["provisioning_state"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:kusto/v20210101:Script"), pulumi.Alias(type_="azure-native:kusto/v20210827:Script"), pulumi.Alias(type_="azure-native:kusto/v20220201:Script"), pulumi.Alias(type_="azure-native:kusto/v20220707:Script"), pulumi.Alias(type_="azure-native:kusto/v20221111:Script"), pulumi.Alias(type_="azure-native:kusto/v20221229:Script"), pulumi.Alias(type_="azure-native:kusto/v20230502:Script"), pulumi.Alias(type_="azure-native:kusto/v20230815:Script"), pulumi.Alias(type_="azure-native:kusto/v20240413:Script")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Script, __self__).__init__(
            'azure-native:kusto:Script',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Script':
        """
        Get an existing Script resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ScriptArgs.__new__(ScriptArgs)

        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["continue_on_errors"] = None
        __props__.__dict__["force_update_tag"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["principal_permissions_action"] = None
        __props__.__dict__["provisioning_state"] = None
        __props__.__dict__["script_level"] = None
        __props__.__dict__["script_url"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        return Script(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="continueOnErrors")
    def continue_on_errors(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Flag that indicates whether to continue if one of the command fails.
        """
        return pulumi.get(self, "continue_on_errors")

    @property
    @pulumi.getter(name="forceUpdateTag")
    def force_update_tag(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        A unique string. If changed the script will be applied again.
        """
        return pulumi.get(self, "force_update_tag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="principalPermissionsAction")
    def principal_permissions_action(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Indicates if the permissions for the script caller are kept following completion of the script.
        """
        return pulumi.get(self, "principal_permissions_action")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[builtins.str]:
        """
        The provisioned state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="scriptLevel")
    def script_level(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Differentiates between the type of script commands included - Database or Cluster. The default is Database.
        """
        return pulumi.get(self, "script_level")

    @property
    @pulumi.getter(name="scriptUrl")
    def script_url(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The url to the KQL script blob file. Must not be used together with scriptContent property
        """
        return pulumi.get(self, "script_url")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Metadata pertaining to creation and last modification of the resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

