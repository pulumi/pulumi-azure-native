# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .azure_ad_metric import *
from .diagnostic_setting import *
from .get_azure_ad_metric import *
from .get_diagnostic_setting import *
from .get_private_endpoint_connection import *
from .get_private_link_for_azure_ad import *
from .private_endpoint_connection import *
from .private_link_for_azure_ad import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_azure_native.aadiam.v20170401 as __v20170401
    v20170401 = __v20170401
    import pulumi_azure_native.aadiam.v20170401preview as __v20170401preview
    v20170401preview = __v20170401preview
    import pulumi_azure_native.aadiam.v20200301 as __v20200301
    v20200301 = __v20200301
    import pulumi_azure_native.aadiam.v20200301preview as __v20200301preview
    v20200301preview = __v20200301preview
    import pulumi_azure_native.aadiam.v20200701preview as __v20200701preview
    v20200701preview = __v20200701preview
else:
    v20170401 = _utilities.lazy_import('pulumi_azure_native.aadiam.v20170401')
    v20170401preview = _utilities.lazy_import('pulumi_azure_native.aadiam.v20170401preview')
    v20200301 = _utilities.lazy_import('pulumi_azure_native.aadiam.v20200301')
    v20200301preview = _utilities.lazy_import('pulumi_azure_native.aadiam.v20200301preview')
    v20200701preview = _utilities.lazy_import('pulumi_azure_native.aadiam.v20200701preview')

