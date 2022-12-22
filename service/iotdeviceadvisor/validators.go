// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdeviceadvisor

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/iotdeviceadvisor/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateSuiteDefinition struct {
}

func (*validateOpCreateSuiteDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSuiteDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateSuiteDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateSuiteDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSuiteDefinition struct {
}

func (*validateOpDeleteSuiteDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSuiteDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSuiteDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSuiteDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSuiteDefinition struct {
}

func (*validateOpGetSuiteDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSuiteDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSuiteDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSuiteDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSuiteRun struct {
}

func (*validateOpGetSuiteRun) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSuiteRun) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSuiteRunInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSuiteRunInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSuiteRunReport struct {
}

func (*validateOpGetSuiteRunReport) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSuiteRunReport) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSuiteRunReportInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSuiteRunReportInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartSuiteRun struct {
}

func (*validateOpStartSuiteRun) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartSuiteRun) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartSuiteRunInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartSuiteRunInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopSuiteRun struct {
}

func (*validateOpStopSuiteRun) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopSuiteRun) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopSuiteRunInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopSuiteRunInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSuiteDefinition struct {
}

func (*validateOpUpdateSuiteDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSuiteDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSuiteDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSuiteDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateSuiteDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSuiteDefinition{}, middleware.After)
}

func addOpDeleteSuiteDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSuiteDefinition{}, middleware.After)
}

func addOpGetSuiteDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSuiteDefinition{}, middleware.After)
}

func addOpGetSuiteRunValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSuiteRun{}, middleware.After)
}

func addOpGetSuiteRunReportValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSuiteRunReport{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpStartSuiteRunValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartSuiteRun{}, middleware.After)
}

func addOpStopSuiteRunValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopSuiteRun{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateSuiteDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSuiteDefinition{}, middleware.After)
}

func validateSuiteDefinitionConfiguration(v *types.SuiteDefinitionConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SuiteDefinitionConfiguration"}
	if v.SuiteDefinitionName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuiteDefinitionName"))
	}
	if v.RootGroup == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RootGroup"))
	}
	if v.DevicePermissionRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DevicePermissionRoleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSuiteRunConfiguration(v *types.SuiteRunConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SuiteRunConfiguration"}
	if v.PrimaryDevice == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PrimaryDevice"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateSuiteDefinitionInput(v *CreateSuiteDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateSuiteDefinitionInput"}
	if v.SuiteDefinitionConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuiteDefinitionConfiguration"))
	} else if v.SuiteDefinitionConfiguration != nil {
		if err := validateSuiteDefinitionConfiguration(v.SuiteDefinitionConfiguration); err != nil {
			invalidParams.AddNested("SuiteDefinitionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSuiteDefinitionInput(v *DeleteSuiteDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSuiteDefinitionInput"}
	if v.SuiteDefinitionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuiteDefinitionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSuiteDefinitionInput(v *GetSuiteDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSuiteDefinitionInput"}
	if v.SuiteDefinitionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuiteDefinitionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSuiteRunInput(v *GetSuiteRunInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSuiteRunInput"}
	if v.SuiteDefinitionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuiteDefinitionId"))
	}
	if v.SuiteRunId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuiteRunId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSuiteRunReportInput(v *GetSuiteRunReportInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSuiteRunReportInput"}
	if v.SuiteDefinitionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuiteDefinitionId"))
	}
	if v.SuiteRunId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuiteRunId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartSuiteRunInput(v *StartSuiteRunInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartSuiteRunInput"}
	if v.SuiteDefinitionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuiteDefinitionId"))
	}
	if v.SuiteRunConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuiteRunConfiguration"))
	} else if v.SuiteRunConfiguration != nil {
		if err := validateSuiteRunConfiguration(v.SuiteRunConfiguration); err != nil {
			invalidParams.AddNested("SuiteRunConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopSuiteRunInput(v *StopSuiteRunInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopSuiteRunInput"}
	if v.SuiteDefinitionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuiteDefinitionId"))
	}
	if v.SuiteRunId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuiteRunId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSuiteDefinitionInput(v *UpdateSuiteDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSuiteDefinitionInput"}
	if v.SuiteDefinitionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuiteDefinitionId"))
	}
	if v.SuiteDefinitionConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuiteDefinitionConfiguration"))
	} else if v.SuiteDefinitionConfiguration != nil {
		if err := validateSuiteDefinitionConfiguration(v.SuiteDefinitionConfiguration); err != nil {
			invalidParams.AddNested("SuiteDefinitionConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
