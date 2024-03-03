# Terraform-provider-vmware-fusion

Reference:

  - [Using VMware Fusion REST API](https://docs.vmware.com/en/VMware-Fusion/13/com.vmware.fusion.using.doc/GUID-5F89D1FE-A95D-4C3C-894F-0084027CF66F.html)

Write `~/.terraformrc` for local developing.

```hcl
provider_installation {
  dev_overrides {
    "registry.terraform.io/leryn1122/vmware-fusion" = "${PWD}/terraform-provider-vmware-fusion/target"
  }
}
```