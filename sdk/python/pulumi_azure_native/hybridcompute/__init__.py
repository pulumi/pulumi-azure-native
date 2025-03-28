# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .gateway import *
from .get_gateway import *
from .get_license import *
from .get_license_profile import *
from .get_machine import *
from .get_machine_extension import *
from .get_machine_run_command import *
from .get_private_endpoint_connection import *
from .get_private_link_scope import *
from .get_private_link_scoped_resource import *
from .license import *
from .license_profile import *
from .machine import *
from .machine_extension import *
from .machine_run_command import *
from .private_endpoint_connection import *
from .private_link_scope import *
from .private_link_scoped_resource import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_azure_native.hybridcompute.v20200802 as __v20200802
    v20200802 = __v20200802
    import pulumi_azure_native.hybridcompute.v20200815preview as __v20200815preview
    v20200815preview = __v20200815preview
    import pulumi_azure_native.hybridcompute.v20220510preview as __v20220510preview
    v20220510preview = __v20220510preview
    import pulumi_azure_native.hybridcompute.v20221227 as __v20221227
    v20221227 = __v20221227
    import pulumi_azure_native.hybridcompute.v20230620preview as __v20230620preview
    v20230620preview = __v20230620preview
    import pulumi_azure_native.hybridcompute.v20231003preview as __v20231003preview
    v20231003preview = __v20231003preview
    import pulumi_azure_native.hybridcompute.v20240331preview as __v20240331preview
    v20240331preview = __v20240331preview
    import pulumi_azure_native.hybridcompute.v20240520preview as __v20240520preview
    v20240520preview = __v20240520preview
    import pulumi_azure_native.hybridcompute.v20240710 as __v20240710
    v20240710 = __v20240710
    import pulumi_azure_native.hybridcompute.v20240731preview as __v20240731preview
    v20240731preview = __v20240731preview
    import pulumi_azure_native.hybridcompute.v20240910preview as __v20240910preview
    v20240910preview = __v20240910preview
    import pulumi_azure_native.hybridcompute.v20241110preview as __v20241110preview
    v20241110preview = __v20241110preview
    import pulumi_azure_native.hybridcompute.v20250113 as __v20250113
    v20250113 = __v20250113
else:
    v20200802 = _utilities.lazy_import('pulumi_azure_native.hybridcompute.v20200802')
    v20200815preview = _utilities.lazy_import('pulumi_azure_native.hybridcompute.v20200815preview')
    v20220510preview = _utilities.lazy_import('pulumi_azure_native.hybridcompute.v20220510preview')
    v20221227 = _utilities.lazy_import('pulumi_azure_native.hybridcompute.v20221227')
    v20230620preview = _utilities.lazy_import('pulumi_azure_native.hybridcompute.v20230620preview')
    v20231003preview = _utilities.lazy_import('pulumi_azure_native.hybridcompute.v20231003preview')
    v20240331preview = _utilities.lazy_import('pulumi_azure_native.hybridcompute.v20240331preview')
    v20240520preview = _utilities.lazy_import('pulumi_azure_native.hybridcompute.v20240520preview')
    v20240710 = _utilities.lazy_import('pulumi_azure_native.hybridcompute.v20240710')
    v20240731preview = _utilities.lazy_import('pulumi_azure_native.hybridcompute.v20240731preview')
    v20240910preview = _utilities.lazy_import('pulumi_azure_native.hybridcompute.v20240910preview')
    v20241110preview = _utilities.lazy_import('pulumi_azure_native.hybridcompute.v20241110preview')
    v20250113 = _utilities.lazy_import('pulumi_azure_native.hybridcompute.v20250113')

