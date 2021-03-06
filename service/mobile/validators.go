// Code generated by smithy-go-codegen DO NOT EDIT.

package mobile

import (
	"context"
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpDeleteProject struct {
}

func (*validateOpDeleteProject) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteProject) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteProjectInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteProjectInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeBundle struct {
}

func (*validateOpDescribeBundle) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeBundle) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeBundleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeBundleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeProject struct {
}

func (*validateOpDescribeProject) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeProject) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeProjectInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeProjectInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpExportBundle struct {
}

func (*validateOpExportBundle) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpExportBundle) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ExportBundleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpExportBundleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpExportProject struct {
}

func (*validateOpExportProject) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpExportProject) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ExportProjectInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpExportProjectInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateProject struct {
}

func (*validateOpUpdateProject) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateProject) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateProjectInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateProjectInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpDeleteProjectValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteProject{}, middleware.After)
}

func addOpDescribeBundleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeBundle{}, middleware.After)
}

func addOpDescribeProjectValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeProject{}, middleware.After)
}

func addOpExportBundleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpExportBundle{}, middleware.After)
}

func addOpExportProjectValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpExportProject{}, middleware.After)
}

func addOpUpdateProjectValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateProject{}, middleware.After)
}

func validateOpDeleteProjectInput(v *DeleteProjectInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteProjectInput"}
	if v.ProjectId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeBundleInput(v *DescribeBundleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeBundleInput"}
	if v.BundleId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BundleId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeProjectInput(v *DescribeProjectInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeProjectInput"}
	if v.ProjectId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpExportBundleInput(v *ExportBundleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExportBundleInput"}
	if v.BundleId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BundleId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpExportProjectInput(v *ExportProjectInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExportProjectInput"}
	if v.ProjectId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateProjectInput(v *UpdateProjectInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateProjectInput"}
	if v.ProjectId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProjectId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
