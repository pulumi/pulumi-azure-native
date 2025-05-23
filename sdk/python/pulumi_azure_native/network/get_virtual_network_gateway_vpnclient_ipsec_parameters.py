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

__all__ = [
    'GetVirtualNetworkGatewayVpnclientIpsecParametersResult',
    'AwaitableGetVirtualNetworkGatewayVpnclientIpsecParametersResult',
    'get_virtual_network_gateway_vpnclient_ipsec_parameters',
    'get_virtual_network_gateway_vpnclient_ipsec_parameters_output',
]

@pulumi.output_type
class GetVirtualNetworkGatewayVpnclientIpsecParametersResult:
    """
    An IPSec parameters for a virtual network gateway P2S connection.
    """
    def __init__(__self__, dh_group=None, ike_encryption=None, ike_integrity=None, ipsec_encryption=None, ipsec_integrity=None, pfs_group=None, sa_data_size_kilobytes=None, sa_life_time_seconds=None):
        if dh_group and not isinstance(dh_group, str):
            raise TypeError("Expected argument 'dh_group' to be a str")
        pulumi.set(__self__, "dh_group", dh_group)
        if ike_encryption and not isinstance(ike_encryption, str):
            raise TypeError("Expected argument 'ike_encryption' to be a str")
        pulumi.set(__self__, "ike_encryption", ike_encryption)
        if ike_integrity and not isinstance(ike_integrity, str):
            raise TypeError("Expected argument 'ike_integrity' to be a str")
        pulumi.set(__self__, "ike_integrity", ike_integrity)
        if ipsec_encryption and not isinstance(ipsec_encryption, str):
            raise TypeError("Expected argument 'ipsec_encryption' to be a str")
        pulumi.set(__self__, "ipsec_encryption", ipsec_encryption)
        if ipsec_integrity and not isinstance(ipsec_integrity, str):
            raise TypeError("Expected argument 'ipsec_integrity' to be a str")
        pulumi.set(__self__, "ipsec_integrity", ipsec_integrity)
        if pfs_group and not isinstance(pfs_group, str):
            raise TypeError("Expected argument 'pfs_group' to be a str")
        pulumi.set(__self__, "pfs_group", pfs_group)
        if sa_data_size_kilobytes and not isinstance(sa_data_size_kilobytes, int):
            raise TypeError("Expected argument 'sa_data_size_kilobytes' to be a int")
        pulumi.set(__self__, "sa_data_size_kilobytes", sa_data_size_kilobytes)
        if sa_life_time_seconds and not isinstance(sa_life_time_seconds, int):
            raise TypeError("Expected argument 'sa_life_time_seconds' to be a int")
        pulumi.set(__self__, "sa_life_time_seconds", sa_life_time_seconds)

    @property
    @pulumi.getter(name="dhGroup")
    def dh_group(self) -> builtins.str:
        """
        The DH Group used in IKE Phase 1 for initial SA.
        """
        return pulumi.get(self, "dh_group")

    @property
    @pulumi.getter(name="ikeEncryption")
    def ike_encryption(self) -> builtins.str:
        """
        The IKE encryption algorithm (IKE phase 2).
        """
        return pulumi.get(self, "ike_encryption")

    @property
    @pulumi.getter(name="ikeIntegrity")
    def ike_integrity(self) -> builtins.str:
        """
        The IKE integrity algorithm (IKE phase 2).
        """
        return pulumi.get(self, "ike_integrity")

    @property
    @pulumi.getter(name="ipsecEncryption")
    def ipsec_encryption(self) -> builtins.str:
        """
        The IPSec encryption algorithm (IKE phase 1).
        """
        return pulumi.get(self, "ipsec_encryption")

    @property
    @pulumi.getter(name="ipsecIntegrity")
    def ipsec_integrity(self) -> builtins.str:
        """
        The IPSec integrity algorithm (IKE phase 1).
        """
        return pulumi.get(self, "ipsec_integrity")

    @property
    @pulumi.getter(name="pfsGroup")
    def pfs_group(self) -> builtins.str:
        """
        The Pfs Group used in IKE Phase 2 for new child SA.
        """
        return pulumi.get(self, "pfs_group")

    @property
    @pulumi.getter(name="saDataSizeKilobytes")
    def sa_data_size_kilobytes(self) -> builtins.int:
        """
        The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for P2S client..
        """
        return pulumi.get(self, "sa_data_size_kilobytes")

    @property
    @pulumi.getter(name="saLifeTimeSeconds")
    def sa_life_time_seconds(self) -> builtins.int:
        """
        The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for P2S client.
        """
        return pulumi.get(self, "sa_life_time_seconds")


class AwaitableGetVirtualNetworkGatewayVpnclientIpsecParametersResult(GetVirtualNetworkGatewayVpnclientIpsecParametersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualNetworkGatewayVpnclientIpsecParametersResult(
            dh_group=self.dh_group,
            ike_encryption=self.ike_encryption,
            ike_integrity=self.ike_integrity,
            ipsec_encryption=self.ipsec_encryption,
            ipsec_integrity=self.ipsec_integrity,
            pfs_group=self.pfs_group,
            sa_data_size_kilobytes=self.sa_data_size_kilobytes,
            sa_life_time_seconds=self.sa_life_time_seconds)


def get_virtual_network_gateway_vpnclient_ipsec_parameters(resource_group_name: Optional[builtins.str] = None,
                                                           virtual_network_gateway_name: Optional[builtins.str] = None,
                                                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVirtualNetworkGatewayVpnclientIpsecParametersResult:
    """
    The Get VpnclientIpsecParameters operation retrieves information about the vpnclient ipsec policy for P2S client of virtual network gateway in the specified resource group through Network resource provider.

    Uses Azure REST API version 2024-05-01.

    Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group.
    :param builtins.str virtual_network_gateway_name: The virtual network gateway name.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['virtualNetworkGatewayName'] = virtual_network_gateway_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:network:getVirtualNetworkGatewayVpnclientIpsecParameters', __args__, opts=opts, typ=GetVirtualNetworkGatewayVpnclientIpsecParametersResult).value

    return AwaitableGetVirtualNetworkGatewayVpnclientIpsecParametersResult(
        dh_group=pulumi.get(__ret__, 'dh_group'),
        ike_encryption=pulumi.get(__ret__, 'ike_encryption'),
        ike_integrity=pulumi.get(__ret__, 'ike_integrity'),
        ipsec_encryption=pulumi.get(__ret__, 'ipsec_encryption'),
        ipsec_integrity=pulumi.get(__ret__, 'ipsec_integrity'),
        pfs_group=pulumi.get(__ret__, 'pfs_group'),
        sa_data_size_kilobytes=pulumi.get(__ret__, 'sa_data_size_kilobytes'),
        sa_life_time_seconds=pulumi.get(__ret__, 'sa_life_time_seconds'))
def get_virtual_network_gateway_vpnclient_ipsec_parameters_output(resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                                                  virtual_network_gateway_name: Optional[pulumi.Input[builtins.str]] = None,
                                                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVirtualNetworkGatewayVpnclientIpsecParametersResult]:
    """
    The Get VpnclientIpsecParameters operation retrieves information about the vpnclient ipsec policy for P2S client of virtual network gateway in the specified resource group through Network resource provider.

    Uses Azure REST API version 2024-05-01.

    Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str resource_group_name: The name of the resource group.
    :param builtins.str virtual_network_gateway_name: The virtual network gateway name.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['virtualNetworkGatewayName'] = virtual_network_gateway_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:network:getVirtualNetworkGatewayVpnclientIpsecParameters', __args__, opts=opts, typ=GetVirtualNetworkGatewayVpnclientIpsecParametersResult)
    return __ret__.apply(lambda __response__: GetVirtualNetworkGatewayVpnclientIpsecParametersResult(
        dh_group=pulumi.get(__response__, 'dh_group'),
        ike_encryption=pulumi.get(__response__, 'ike_encryption'),
        ike_integrity=pulumi.get(__response__, 'ike_integrity'),
        ipsec_encryption=pulumi.get(__response__, 'ipsec_encryption'),
        ipsec_integrity=pulumi.get(__response__, 'ipsec_integrity'),
        pfs_group=pulumi.get(__response__, 'pfs_group'),
        sa_data_size_kilobytes=pulumi.get(__response__, 'sa_data_size_kilobytes'),
        sa_life_time_seconds=pulumi.get(__response__, 'sa_life_time_seconds')))
