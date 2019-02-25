package role

import (
	"log"
	"strconv"

	"github.com/terraform-providers/terraform-provider-vsphere/vsphere/internal/helper/provider"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
	"golang.org/x/net/context"
)

// PermissionsList Reference List for Permissions
var PermissionsList = []string{
	"Alarm.Acknowledge",
	"Alarm.Create",
	"Alarm.Delete",
	"Alarm.DisableActions",
	"Alarm.Edit",
	"Alarm.SetStatus",
	"Authorization.ModifyPermissions",
	"Authorization.ModifyPrivileges",
	"Authorization.ModifyRoles",
	"Authorization.ReassignRolePermissions",
	"AutoDeploy.Host.AssociateMachine",
	"AutoDeploy.Profile.Create",
	"AutoDeploy.Profile.Edit",
	"AutoDeploy.Rule.Create",
	"AutoDeploy.Rule.Delete",
	"AutoDeploy.Rule.Edit",
	"AutoDeploy.RuleSet.Activate",
	"AutoDeploy.RuleSet.Edit",
	"Certificate.Manage",
	"ContentLibrary.AddLibraryItem",
	"ContentLibrary.CreateLocalLibrary",
	"ContentLibrary.CreateSubscribedLibrary",
	"ContentLibrary.DeleteLibraryItem",
	"ContentLibrary.DeleteLocalLibrary",
	"ContentLibrary.DeleteSubscribedLibrary",
	"ContentLibrary.DownloadSession",
	"ContentLibrary.EvictLibraryItem",
	"ContentLibrary.EvictSubscribedLibrary",
	"ContentLibrary.GetConfiguration",
	"ContentLibrary.ImportStorage",
	"ContentLibrary.ProbeSubscription",
	"ContentLibrary.ReadStorage",
	"ContentLibrary.SyncLibrary",
	"ContentLibrary.SyncLibraryItem",
	"ContentLibrary.TypeIntrospection",
	"ContentLibrary.UpdateConfiguration",
	"ContentLibrary.UpdateLibrary",
	"ContentLibrary.UpdateLibraryItem",
	"ContentLibrary.UpdateLocalLibrary",
	"ContentLibrary.UpdateSession",
	"ContentLibrary.UpdateSubscribedLibrary",
	"Cryptographer.Access",
	"Cryptographer.AddDisk",
	"Cryptographer.Clone",
	"Cryptographer.Decrypt",
	"Cryptographer.Encrypt",
	"Cryptographer.EncryptNew",
	"Cryptographer.ManageEncryptionPolicy",
	"Cryptographer.ManageKeyServers",
	"Cryptographer.ManageKeys",
	"Cryptographer.Migrate",
	"Cryptographer.Recrypt",
	"Cryptographer.RegisterHost",
	"Cryptographer.RegisterVM",
	"DVPortgroup.Create",
	"DVPortgroup.Delete",
	"DVPortgroup.Modify",
	"DVPortgroup.PolicyOp",
	"DVPortgroup.ScopeOp",
	"DVSwitch.Create",
	"DVSwitch.Delete",
	"DVSwitch.HostOp",
	"DVSwitch.Modify",
	"DVSwitch.Move",
	"DVSwitch.PolicyOp",
	"DVSwitch.PortConfig",
	"DVSwitch.PortSetting",
	"DVSwitch.ResourceManagement",
	"DVSwitch.Vspan",
	"Datacenter.Create",
	"Datacenter.Delete",
	"Datacenter.IpPoolConfig",
	"Datacenter.IpPoolQueryAllocations",
	"Datacenter.IpPoolReleaseIp",
	"Datacenter.Move",
	"Datacenter.Reconfigure",
	"Datacenter.Rename",
	"Datastore.AllocateSpace",
	"Datastore.Browse",
	"Datastore.Config",
	"Datastore.Delete",
	"Datastore.DeleteFile",
	"Datastore.FileManagement",
	"Datastore.Move",
	"Datastore.Rename",
	"Datastore.UpdateVirtualMachineFiles",
	"Datastore.UpdateVirtualMachineMetadata",
	"EAM.Config",
	"EAM.Modify",
	"EAM.View",
	"Extension.Register",
	"Extension.Unregister",
	"Extension.Update",
	"ExternalStatsProvider.Register",
	"ExternalStatsProvider.Unregister",
	"ExternalStatsProvider.Update",
	"Folder.Create",
	"Folder.Delete",
	"Folder.Move",
	"Folder.Rename",
	"Global.CancelTask",
	"Global.CapacityPlanning",
	"Global.Diagnostics",
	"Global.DisableMethods",
	"Global.EnableMethods",
	"Global.GlobalTag",
	"Global.Health",
	"Global.Licenses",
	"Global.LogEvent",
	"Global.ManageCustomFields",
	"Global.Proxy",
	"Global.ScriptAction",
	"Global.ServiceManagers",
	"Global.SetCustomField",
	"Global.Settings",
	"Global.SystemTag",
	"Global.VCServer",
	"HealthUpdateProvider.Register",
	"HealthUpdateProvider.Unregister",
	"HealthUpdateProvider.Update",
	"Host.Cim.CimInteraction",
	"Host.Config.AdvancedConfig",
	"Host.Config.AuthenticationStore",
	"Host.Config.AutoStart",
	"Host.Config.Connection",
	"Host.Config.DateTime",
	"Host.Config.Firmware",
	"Host.Config.HyperThreading",
	"Host.Config.Image",
	"Host.Config.Maintenance",
	"Host.Config.Memory",
	"Host.Config.NetService",
	"Host.Config.Network",
	"Host.Config.Patch",
	"Host.Config.PciPassthru",
	"Host.Config.Power",
	"Host.Config.Quarantine",
	"Host.Config.Resources",
	"Host.Config.Settings",
	"Host.Config.Snmp",
	"Host.Config.Storage",
	"Host.Config.SystemManagement",
	"Host.Hbr.HbrManagement",
	"Host.Inventory.AddHostToCluster",
	"Host.Inventory.AddStandaloneHost",
	"Host.Inventory.CreateCluster",
	"Host.Inventory.DeleteCluster",
	"Host.Inventory.EditCluster",
	"Host.Inventory.MoveCluster",
	"Host.Inventory.MoveHost",
	"Host.Inventory.RemoveHostFromCluster",
	"Host.Inventory.RenameCluster",
	"Host.Local.CreateVM",
	"Host.Local.DeleteVM",
	"Host.Local.InstallAgent",
	"Host.Local.ManageUserGroups",
	"Host.Local.ReconfigVM",
	"InventoryService.Tagging.AttachTag",
	"InventoryService.Tagging.CreateCategory",
	"InventoryService.Tagging.CreateTag",
	"InventoryService.Tagging.DeleteCategory",
	"InventoryService.Tagging.DeleteTag",
	"InventoryService.Tagging.EditCategory",
	"InventoryService.Tagging.EditTag",
	"InventoryService.Tagging.ModifyUsedByForCategory",
	"InventoryService.Tagging.ModifyUsedByForTag",
	"Network.Assign",
	"Network.Config",
	"Network.Delete",
	"Network.Move",
	"Performance.ModifyIntervals",
	"Profile.Clear",
	"Profile.Create",
	"Profile.Delete",
	"Profile.Edit",
	"Profile.Export",
	"Profile.View",
	"Resource.ApplyRecommendation",
	"Resource.AssignVAppToPool",
	"Resource.AssignVMToPool",
	"Resource.ColdMigrate",
	"Resource.CreatePool",
	"Resource.DeletePool",
	"Resource.EditPool",
	"Resource.HotMigrate",
	"Resource.MovePool",
	"Resource.QueryVMotion",
	"Resource.RenamePool",
	"ScheduledTask.Create",
	"ScheduledTask.Delete",
	"ScheduledTask.Edit",
	"ScheduledTask.Run",
	"Sessions.GlobalMessage",
	"Sessions.ImpersonateUser",
	"Sessions.TerminateSession",
	"Sessions.ValidateSession",
	"StoragePod.Config",
	"StorageProfile.Update",
	"StorageProfile.View",
	"StorageViews.ConfigureService",
	"StorageViews.View",
	"System.Anonymous",
	"System.Read",
	"System.View",
	"Task.Create",
	"Task.Update",
	"TransferService.Manage",
	"TransferService.Monitor",
	"VApp.ApplicationConfig",
	"VApp.AssignResourcePool",
	"VApp.AssignVApp",
	"VApp.AssignVM",
	"VApp.Clone",
	"VApp.Create",
	"VApp.Delete",
	"VApp.Export",
	"VApp.ExtractOvfEnvironment",
	"VApp.Import",
	"VApp.InstanceConfig",
	"VApp.ManagedByConfig",
	"VApp.Move",
	"VApp.PowerOff",
	"VApp.PowerOn",
	"VApp.Rename",
	"VApp.ResourceConfig",
	"VApp.Suspend",
	"VApp.Unregister",
	"VRMPolicy.Query",
	"VRMPolicy.Update",
	"VcIntegrity.Baseline.com.vmware.vcIntegrity.AssignBaselines",
	"VcIntegrity.Baseline.com.vmware.vcIntegrity.ManageBaselines",
	"VcIntegrity.FileUpload.com.vmware.vcIntegrity.ImportFile",
	"VcIntegrity.General.com.vmware.vcIntegrity.Configure",
	"VcIntegrity.Updates.com.vmware.vcIntegrity.Remediate",
	"VcIntegrity.Updates.com.vmware.vcIntegrity.Scan",
	"VcIntegrity.Updates.com.vmware.vcIntegrity.Stage",
	"VcIntegrity.Updates.com.vmware.vcIntegrity.ViewStatus",
	"VirtualMachine.Config.AddExistingDisk",
	"VirtualMachine.Config.AddNewDisk",
	"VirtualMachine.Config.AddRemoveDevice",
	"VirtualMachine.Config.AdvancedConfig",
	"VirtualMachine.Config.Annotation",
	"VirtualMachine.Config.CPUCount",
	"VirtualMachine.Config.ChangeTracking",
	"VirtualMachine.Config.DiskExtend",
	"VirtualMachine.Config.DiskLease",
	"VirtualMachine.Config.EditDevice",
	"VirtualMachine.Config.HostUSBDevice",
	"VirtualMachine.Config.ManagedBy",
	"VirtualMachine.Config.Memory",
	"VirtualMachine.Config.MksControl",
	"VirtualMachine.Config.QueryFTCompatibility",
	"VirtualMachine.Config.QueryUnownedFiles",
	"VirtualMachine.Config.RawDevice",
	"VirtualMachine.Config.ReloadFromPath",
	"VirtualMachine.Config.RemoveDisk",
	"VirtualMachine.Config.Rename",
	"VirtualMachine.Config.ResetGuestInfo",
	"VirtualMachine.Config.Resource",
	"VirtualMachine.Config.Settings",
	"VirtualMachine.Config.SwapPlacement",
	"VirtualMachine.Config.ToggleForkParent",
	"VirtualMachine.Config.Unlock",
	"VirtualMachine.Config.UpgradeVirtualHardware",
	"VirtualMachine.GuestOperations.Execute",
	"VirtualMachine.GuestOperations.Modify",
	"VirtualMachine.GuestOperations.ModifyAliases",
	"VirtualMachine.GuestOperations.Query",
	"VirtualMachine.GuestOperations.QueryAliases",
	"VirtualMachine.Hbr.ConfigureReplication",
	"VirtualMachine.Hbr.MonitorReplication",
	"VirtualMachine.Hbr.ReplicaManagement",
	"VirtualMachine.Interact.AnswerQuestion",
	"VirtualMachine.Interact.Backup",
	"VirtualMachine.Interact.ConsoleInteract",
	"VirtualMachine.Interact.CreateScreenshot",
	"VirtualMachine.Interact.CreateSecondary",
	"VirtualMachine.Interact.DefragmentAllDisks",
	"VirtualMachine.Interact.DeviceConnection",
	"VirtualMachine.Interact.DisableSecondary",
	"VirtualMachine.Interact.DnD",
	"VirtualMachine.Interact.EnableSecondary",
	"VirtualMachine.Interact.GuestControl",
	"VirtualMachine.Interact.MakePrimary",
	"VirtualMachine.Interact.Pause",
	"VirtualMachine.Interact.PowerOff",
	"VirtualMachine.Interact.PowerOn",
	"VirtualMachine.Interact.PutUsbScanCodes",
	"VirtualMachine.Interact.Record",
	"VirtualMachine.Interact.Replay",
	"VirtualMachine.Interact.Reset",
	"VirtualMachine.Interact.SESparseMaintenance",
	"VirtualMachine.Interact.SetCDMedia",
	"VirtualMachine.Interact.SetFloppyMedia",
	"VirtualMachine.Interact.Suspend",
	"VirtualMachine.Interact.TerminateFaultTolerantVM",
	"VirtualMachine.Interact.ToolsInstall",
	"VirtualMachine.Interact.TurnOffFaultTolerance",
	"VirtualMachine.Inventory.Create",
	"VirtualMachine.Inventory.CreateFromExisting",
	"VirtualMachine.Inventory.Delete",
	"VirtualMachine.Inventory.Move",
	"VirtualMachine.Inventory.Register",
	"VirtualMachine.Inventory.Unregister",
	"VirtualMachine.Namespace.Event",
	"VirtualMachine.Namespace.EventNotify",
	"VirtualMachine.Namespace.Management",
	"VirtualMachine.Namespace.ModifyContent",
	"VirtualMachine.Namespace.Query",
	"VirtualMachine.Namespace.ReadContent",
	"VirtualMachine.Provisioning.Clone",
	"VirtualMachine.Provisioning.CloneTemplate",
	"VirtualMachine.Provisioning.CreateTemplateFromVM",
	"VirtualMachine.Provisioning.Customize",
	"VirtualMachine.Provisioning.DeployTemplate",
	"VirtualMachine.Provisioning.DiskRandomAccess",
	"VirtualMachine.Provisioning.DiskRandomRead",
	"VirtualMachine.Provisioning.FileRandomAccess",
	"VirtualMachine.Provisioning.GetVmFiles",
	"VirtualMachine.Provisioning.MarkAsTemplate",
	"VirtualMachine.Provisioning.MarkAsVM",
	"VirtualMachine.Provisioning.ModifyCustSpecs",
	"VirtualMachine.Provisioning.PromoteDisks",
	"VirtualMachine.Provisioning.PutVmFiles",
	"VirtualMachine.Provisioning.ReadCustSpecs",
	"VirtualMachine.State.CreateSnapshot",
	"VirtualMachine.State.RemoveSnapshot",
	"VirtualMachine.State.RenameSnapshot",
	"VirtualMachine.State.RevertToSnapshot",
	"vService.CreateDependency",
	"vService.DestroyDependency",
	"vService.ReconfigureDependency",
	"vService.UpdateDependency",
}

// ByID check if a role exist using id, and return that role
func ByID(client *govmomi.Client, id string) (*types.AuthorizationRole, error) {
	log.Printf("[DEBUG] Locating role with ID %q", id)
	m := object.NewAuthorizationManager(client.Client)
	ctx, cancel := context.WithTimeout(context.Background(), provider.DefaultAPITimeout)
	defer cancel()
	roles, err := m.RoleList(ctx)
	if err != nil {
		log.Printf("Role Listing error: %s", err)
		return nil, err
	}
	nid, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	role := roles.ById(int32(nid))
	log.Printf("[DEBUG] Successfully located role with ID %q", id)
	return role, nil
}

// ByName check if a role exist using name, and return that role
func ByName(client *govmomi.Client, name string) (*types.AuthorizationRole, error) {
	log.Printf("[DEBUG] Locating role with name %q", name)
	m := object.NewAuthorizationManager(client.Client)
	ctx, cancel := context.WithTimeout(context.Background(), provider.DefaultAPITimeout)
	defer cancel()
	roles, err := m.RoleList(ctx)
	if err != nil {
		log.Printf("Role Listing error: %s", err)
		return nil, err
	}
	role := roles.ByName(name)
	log.Printf("[DEBUG] Successfully located role with name %q", name)
	return role, nil
}

// Update Role permissions
func Update(client *govmomi.Client, roleID int32, name string, perms []string) error {
	log.Printf("[DEBUG] Updating role with ID %q", roleID)
	m := object.NewAuthorizationManager(client.Client)
	ctx, cancel := context.WithTimeout(context.Background(), provider.DefaultAPITimeout)
	defer cancel()

	if err := m.UpdateRole(ctx, roleID, name, perms); err != nil {
		return err
	}
	log.Printf("[DEBUG] Successfully updated role with ID %q", roleID)
	return nil
}

// Create Role
func Create(client *govmomi.Client, name string, perms []string) (int32, error) {
	log.Printf("[DEBUG] Creating role with name %q", name)
	m := object.NewAuthorizationManager(client.Client)
	ctx, cancel := context.WithTimeout(context.Background(), provider.DefaultAPITimeout)
	defer cancel()
	roleID, err := m.AddRole(ctx, name, perms)
	if err != nil {
		return 0, err
	}
	log.Printf("[DEBUG] Successfully created role with name %q", name)
	return roleID, nil
}

// Remove Role
func Remove(client *govmomi.Client, roleID int32) error {
	log.Printf("[DEBUG] Removing role with ID %q", roleID)
	m := object.NewAuthorizationManager(client.Client)
	ctx, cancel := context.WithTimeout(context.Background(), provider.DefaultAPITimeout)
	defer cancel()
	if err := m.RemoveRole(ctx, roleID, false); err != nil {
		return err
	}
	log.Printf("[DEBUG] Successfully removed role with ID %q", roleID)
	return nil
}
