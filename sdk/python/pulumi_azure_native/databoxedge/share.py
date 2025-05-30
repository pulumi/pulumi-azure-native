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

__all__ = ['ShareArgs', 'Share']

@pulumi.input_type
class ShareArgs:
    def __init__(__self__, *,
                 access_protocol: pulumi.Input[Union[builtins.str, 'ShareAccessProtocol']],
                 device_name: pulumi.Input[builtins.str],
                 monitoring_status: pulumi.Input[Union[builtins.str, 'MonitoringStatus']],
                 resource_group_name: pulumi.Input[builtins.str],
                 share_status: pulumi.Input[Union[builtins.str, 'ShareStatus']],
                 azure_container_info: Optional[pulumi.Input['AzureContainerInfoArgs']] = None,
                 client_access_rights: Optional[pulumi.Input[Sequence[pulumi.Input['ClientAccessRightArgs']]]] = None,
                 data_policy: Optional[pulumi.Input[Union[builtins.str, 'DataPolicy']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 refresh_details: Optional[pulumi.Input['RefreshDetailsArgs']] = None,
                 user_access_rights: Optional[pulumi.Input[Sequence[pulumi.Input['UserAccessRightArgs']]]] = None):
        """
        The set of arguments for constructing a Share resource.
        :param pulumi.Input[Union[builtins.str, 'ShareAccessProtocol']] access_protocol: Access protocol to be used by the share.
        :param pulumi.Input[builtins.str] device_name: The device name.
        :param pulumi.Input[Union[builtins.str, 'MonitoringStatus']] monitoring_status: Current monitoring status of the share.
        :param pulumi.Input[builtins.str] resource_group_name: The resource group name.
        :param pulumi.Input[Union[builtins.str, 'ShareStatus']] share_status: Current status of the share.
        :param pulumi.Input['AzureContainerInfoArgs'] azure_container_info: Azure container mapping for the share.
        :param pulumi.Input[Sequence[pulumi.Input['ClientAccessRightArgs']]] client_access_rights: List of IP addresses and corresponding access rights on the share(required for NFS protocol).
        :param pulumi.Input[Union[builtins.str, 'DataPolicy']] data_policy: Data policy of the share.
        :param pulumi.Input[builtins.str] description: Description for the share.
        :param pulumi.Input[builtins.str] name: The share name.
        :param pulumi.Input['RefreshDetailsArgs'] refresh_details: Details of the refresh job on this share.
        :param pulumi.Input[Sequence[pulumi.Input['UserAccessRightArgs']]] user_access_rights: Mapping of users and corresponding access rights on the share (required for SMB protocol).
        """
        pulumi.set(__self__, "access_protocol", access_protocol)
        pulumi.set(__self__, "device_name", device_name)
        pulumi.set(__self__, "monitoring_status", monitoring_status)
        pulumi.set(__self__, "resource_group_name", resource_group_name)
        pulumi.set(__self__, "share_status", share_status)
        if azure_container_info is not None:
            pulumi.set(__self__, "azure_container_info", azure_container_info)
        if client_access_rights is not None:
            pulumi.set(__self__, "client_access_rights", client_access_rights)
        if data_policy is not None:
            pulumi.set(__self__, "data_policy", data_policy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if refresh_details is not None:
            pulumi.set(__self__, "refresh_details", refresh_details)
        if user_access_rights is not None:
            pulumi.set(__self__, "user_access_rights", user_access_rights)

    @property
    @pulumi.getter(name="accessProtocol")
    def access_protocol(self) -> pulumi.Input[Union[builtins.str, 'ShareAccessProtocol']]:
        """
        Access protocol to be used by the share.
        """
        return pulumi.get(self, "access_protocol")

    @access_protocol.setter
    def access_protocol(self, value: pulumi.Input[Union[builtins.str, 'ShareAccessProtocol']]):
        pulumi.set(self, "access_protocol", value)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> pulumi.Input[builtins.str]:
        """
        The device name.
        """
        return pulumi.get(self, "device_name")

    @device_name.setter
    def device_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "device_name", value)

    @property
    @pulumi.getter(name="monitoringStatus")
    def monitoring_status(self) -> pulumi.Input[Union[builtins.str, 'MonitoringStatus']]:
        """
        Current monitoring status of the share.
        """
        return pulumi.get(self, "monitoring_status")

    @monitoring_status.setter
    def monitoring_status(self, value: pulumi.Input[Union[builtins.str, 'MonitoringStatus']]):
        pulumi.set(self, "monitoring_status", value)

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Input[builtins.str]:
        """
        The resource group name.
        """
        return pulumi.get(self, "resource_group_name")

    @resource_group_name.setter
    def resource_group_name(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_group_name", value)

    @property
    @pulumi.getter(name="shareStatus")
    def share_status(self) -> pulumi.Input[Union[builtins.str, 'ShareStatus']]:
        """
        Current status of the share.
        """
        return pulumi.get(self, "share_status")

    @share_status.setter
    def share_status(self, value: pulumi.Input[Union[builtins.str, 'ShareStatus']]):
        pulumi.set(self, "share_status", value)

    @property
    @pulumi.getter(name="azureContainerInfo")
    def azure_container_info(self) -> Optional[pulumi.Input['AzureContainerInfoArgs']]:
        """
        Azure container mapping for the share.
        """
        return pulumi.get(self, "azure_container_info")

    @azure_container_info.setter
    def azure_container_info(self, value: Optional[pulumi.Input['AzureContainerInfoArgs']]):
        pulumi.set(self, "azure_container_info", value)

    @property
    @pulumi.getter(name="clientAccessRights")
    def client_access_rights(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClientAccessRightArgs']]]]:
        """
        List of IP addresses and corresponding access rights on the share(required for NFS protocol).
        """
        return pulumi.get(self, "client_access_rights")

    @client_access_rights.setter
    def client_access_rights(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClientAccessRightArgs']]]]):
        pulumi.set(self, "client_access_rights", value)

    @property
    @pulumi.getter(name="dataPolicy")
    def data_policy(self) -> Optional[pulumi.Input[Union[builtins.str, 'DataPolicy']]]:
        """
        Data policy of the share.
        """
        return pulumi.get(self, "data_policy")

    @data_policy.setter
    def data_policy(self, value: Optional[pulumi.Input[Union[builtins.str, 'DataPolicy']]]):
        pulumi.set(self, "data_policy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description for the share.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The share name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="refreshDetails")
    def refresh_details(self) -> Optional[pulumi.Input['RefreshDetailsArgs']]:
        """
        Details of the refresh job on this share.
        """
        return pulumi.get(self, "refresh_details")

    @refresh_details.setter
    def refresh_details(self, value: Optional[pulumi.Input['RefreshDetailsArgs']]):
        pulumi.set(self, "refresh_details", value)

    @property
    @pulumi.getter(name="userAccessRights")
    def user_access_rights(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserAccessRightArgs']]]]:
        """
        Mapping of users and corresponding access rights on the share (required for SMB protocol).
        """
        return pulumi.get(self, "user_access_rights")

    @user_access_rights.setter
    def user_access_rights(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserAccessRightArgs']]]]):
        pulumi.set(self, "user_access_rights", value)


@pulumi.type_token("azure-native:databoxedge:Share")
class Share(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_protocol: Optional[pulumi.Input[Union[builtins.str, 'ShareAccessProtocol']]] = None,
                 azure_container_info: Optional[pulumi.Input[Union['AzureContainerInfoArgs', 'AzureContainerInfoArgsDict']]] = None,
                 client_access_rights: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ClientAccessRightArgs', 'ClientAccessRightArgsDict']]]]] = None,
                 data_policy: Optional[pulumi.Input[Union[builtins.str, 'DataPolicy']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 device_name: Optional[pulumi.Input[builtins.str]] = None,
                 monitoring_status: Optional[pulumi.Input[Union[builtins.str, 'MonitoringStatus']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 refresh_details: Optional[pulumi.Input[Union['RefreshDetailsArgs', 'RefreshDetailsArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 share_status: Optional[pulumi.Input[Union[builtins.str, 'ShareStatus']]] = None,
                 user_access_rights: Optional[pulumi.Input[Sequence[pulumi.Input[Union['UserAccessRightArgs', 'UserAccessRightArgsDict']]]]] = None,
                 __props__=None):
        """
        Represents a share on the  Data Box Edge/Gateway device.

        Uses Azure REST API version 2023-07-01. In version 2.x of the Azure Native provider, it used API version 2022-03-01.

        Other available API versions: 2022-03-01, 2022-04-01-preview, 2022-12-01-preview, 2023-01-01-preview, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databoxedge [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union[builtins.str, 'ShareAccessProtocol']] access_protocol: Access protocol to be used by the share.
        :param pulumi.Input[Union['AzureContainerInfoArgs', 'AzureContainerInfoArgsDict']] azure_container_info: Azure container mapping for the share.
        :param pulumi.Input[Sequence[pulumi.Input[Union['ClientAccessRightArgs', 'ClientAccessRightArgsDict']]]] client_access_rights: List of IP addresses and corresponding access rights on the share(required for NFS protocol).
        :param pulumi.Input[Union[builtins.str, 'DataPolicy']] data_policy: Data policy of the share.
        :param pulumi.Input[builtins.str] description: Description for the share.
        :param pulumi.Input[builtins.str] device_name: The device name.
        :param pulumi.Input[Union[builtins.str, 'MonitoringStatus']] monitoring_status: Current monitoring status of the share.
        :param pulumi.Input[builtins.str] name: The share name.
        :param pulumi.Input[Union['RefreshDetailsArgs', 'RefreshDetailsArgsDict']] refresh_details: Details of the refresh job on this share.
        :param pulumi.Input[builtins.str] resource_group_name: The resource group name.
        :param pulumi.Input[Union[builtins.str, 'ShareStatus']] share_status: Current status of the share.
        :param pulumi.Input[Sequence[pulumi.Input[Union['UserAccessRightArgs', 'UserAccessRightArgsDict']]]] user_access_rights: Mapping of users and corresponding access rights on the share (required for SMB protocol).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ShareArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a share on the  Data Box Edge/Gateway device.

        Uses Azure REST API version 2023-07-01. In version 2.x of the Azure Native provider, it used API version 2022-03-01.

        Other available API versions: 2022-03-01, 2022-04-01-preview, 2022-12-01-preview, 2023-01-01-preview, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databoxedge [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.

        :param str resource_name: The name of the resource.
        :param ShareArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ShareArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_protocol: Optional[pulumi.Input[Union[builtins.str, 'ShareAccessProtocol']]] = None,
                 azure_container_info: Optional[pulumi.Input[Union['AzureContainerInfoArgs', 'AzureContainerInfoArgsDict']]] = None,
                 client_access_rights: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ClientAccessRightArgs', 'ClientAccessRightArgsDict']]]]] = None,
                 data_policy: Optional[pulumi.Input[Union[builtins.str, 'DataPolicy']]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 device_name: Optional[pulumi.Input[builtins.str]] = None,
                 monitoring_status: Optional[pulumi.Input[Union[builtins.str, 'MonitoringStatus']]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 refresh_details: Optional[pulumi.Input[Union['RefreshDetailsArgs', 'RefreshDetailsArgsDict']]] = None,
                 resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                 share_status: Optional[pulumi.Input[Union[builtins.str, 'ShareStatus']]] = None,
                 user_access_rights: Optional[pulumi.Input[Sequence[pulumi.Input[Union['UserAccessRightArgs', 'UserAccessRightArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ShareArgs.__new__(ShareArgs)

            if access_protocol is None and not opts.urn:
                raise TypeError("Missing required property 'access_protocol'")
            __props__.__dict__["access_protocol"] = access_protocol
            __props__.__dict__["azure_container_info"] = azure_container_info
            __props__.__dict__["client_access_rights"] = client_access_rights
            __props__.__dict__["data_policy"] = data_policy
            __props__.__dict__["description"] = description
            if device_name is None and not opts.urn:
                raise TypeError("Missing required property 'device_name'")
            __props__.__dict__["device_name"] = device_name
            if monitoring_status is None and not opts.urn:
                raise TypeError("Missing required property 'monitoring_status'")
            __props__.__dict__["monitoring_status"] = monitoring_status
            __props__.__dict__["name"] = name
            __props__.__dict__["refresh_details"] = refresh_details
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__.__dict__["resource_group_name"] = resource_group_name
            if share_status is None and not opts.urn:
                raise TypeError("Missing required property 'share_status'")
            __props__.__dict__["share_status"] = share_status
            __props__.__dict__["user_access_rights"] = user_access_rights
            __props__.__dict__["azure_api_version"] = None
            __props__.__dict__["share_mappings"] = None
            __props__.__dict__["system_data"] = None
            __props__.__dict__["type"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-native:databoxedge/v20190301:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20190701:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20190801:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20200501preview:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20200901:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20200901preview:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20201201:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20210201:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20210201preview:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20210601:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20210601preview:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20220301:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20220401preview:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20221201preview:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20230101preview:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20230701:Share"), pulumi.Alias(type_="azure-native:databoxedge/v20231201:Share")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Share, __self__).__init__(
            'azure-native:databoxedge:Share',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Share':
        """
        Get an existing Share resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ShareArgs.__new__(ShareArgs)

        __props__.__dict__["access_protocol"] = None
        __props__.__dict__["azure_api_version"] = None
        __props__.__dict__["azure_container_info"] = None
        __props__.__dict__["client_access_rights"] = None
        __props__.__dict__["data_policy"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["monitoring_status"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["refresh_details"] = None
        __props__.__dict__["share_mappings"] = None
        __props__.__dict__["share_status"] = None
        __props__.__dict__["system_data"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["user_access_rights"] = None
        return Share(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessProtocol")
    def access_protocol(self) -> pulumi.Output[builtins.str]:
        """
        Access protocol to be used by the share.
        """
        return pulumi.get(self, "access_protocol")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> pulumi.Output[builtins.str]:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="azureContainerInfo")
    def azure_container_info(self) -> pulumi.Output[Optional['outputs.AzureContainerInfoResponse']]:
        """
        Azure container mapping for the share.
        """
        return pulumi.get(self, "azure_container_info")

    @property
    @pulumi.getter(name="clientAccessRights")
    def client_access_rights(self) -> pulumi.Output[Optional[Sequence['outputs.ClientAccessRightResponse']]]:
        """
        List of IP addresses and corresponding access rights on the share(required for NFS protocol).
        """
        return pulumi.get(self, "client_access_rights")

    @property
    @pulumi.getter(name="dataPolicy")
    def data_policy(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Data policy of the share.
        """
        return pulumi.get(self, "data_policy")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Description for the share.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="monitoringStatus")
    def monitoring_status(self) -> pulumi.Output[builtins.str]:
        """
        Current monitoring status of the share.
        """
        return pulumi.get(self, "monitoring_status")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The object name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="refreshDetails")
    def refresh_details(self) -> pulumi.Output[Optional['outputs.RefreshDetailsResponse']]:
        """
        Details of the refresh job on this share.
        """
        return pulumi.get(self, "refresh_details")

    @property
    @pulumi.getter(name="shareMappings")
    def share_mappings(self) -> pulumi.Output[Sequence['outputs.MountPointMapResponse']]:
        """
        Share mount point to the role.
        """
        return pulumi.get(self, "share_mappings")

    @property
    @pulumi.getter(name="shareStatus")
    def share_status(self) -> pulumi.Output[builtins.str]:
        """
        Current status of the share.
        """
        return pulumi.get(self, "share_status")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Metadata pertaining to creation and last modification of Share
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[builtins.str]:
        """
        The hierarchical type of the object.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userAccessRights")
    def user_access_rights(self) -> pulumi.Output[Optional[Sequence['outputs.UserAccessRightResponse']]]:
        """
        Mapping of users and corresponding access rights on the share (required for SMB protocol).
        """
        return pulumi.get(self, "user_access_rights")

