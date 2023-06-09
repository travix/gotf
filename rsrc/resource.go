package rsrc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Resource required methods on datasource block.
//
// T is the protobuf message generated by protoc-gen-go.
type Resource[T any] interface {
	// Create receive a context, a CreateRequest, a CreateResponse and a data object filled with the data from the resource.
	// It should return the same or update data object updated or diag.Diagnostics if error occurs during the creation.
	Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse, data T) (T, diag.Diagnostics)
	// Read receive a context, a ReadRequest, a ReadResponse and a data object filled with the data from the resource.
	// It should return the same or update data object updated or diag.Diagnostics if error occurs during the read.
	Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse, data T) (T, diag.Diagnostics)
	// Update receives a context, a UpdateRequest, a UpdateResponse and a data object filled with the data from the resource.
	// It should return the same or update data object updated or diag.Diagnostics if error occurs during the update.
	Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse, data T) (T, diag.Diagnostics)
	// Delete receive a context, a DeleteRequest, a DeleteResponse and a data object filled with the data from the resource.
	// It should return the same or update data object updated or diag.Diagnostics if error occurs during the deletion.
	Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse, data T) diag.Diagnostics
}

// CanMetadata optional interface to inject metadata method on resource block.
type CanMetadata interface {
	Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse)
}

// CanSchema optional interface to inject schema method on resource block.
type CanSchema interface {
	Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse)
}

// CanConfigure optional interface to inject configure method on resource block.
type CanConfigure interface {
	Configure(context.Context, resource.ConfigureRequest, *resource.ConfigureResponse)
}

// CanConfigValidators optional interface to inject config validators method on resource block.
type CanConfigValidators interface {
	ConfigValidators(context.Context) []resource.ConfigValidator
}

// CanImportState optional interface to inject import state method on resource block.
type CanImportState interface {
	ImportState(context.Context, resource.ImportStateRequest, *resource.ImportStateResponse)
}

// CanModifyPlan optional interface to inject modify plan method on resource block.
type CanModifyPlan interface {
	ModifyPlan(context.Context, resource.ModifyPlanRequest, *resource.ModifyPlanResponse)
}

// CanUpgradeState optional interface to inject upgrade state method on resource block.
type CanUpgradeState interface {
	UpgradeState(context.Context) map[int64]resource.StateUpgrader
}

// CanValidateConfig optional interface to inject validate config method on resource block.
type CanValidateConfig interface {
	ValidateConfig(context.Context, resource.ValidateConfigRequest, *resource.ValidateConfigResponse)
}
