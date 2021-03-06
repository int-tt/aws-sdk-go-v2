// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsRestjson1_serializeOpCancelSigningProfile struct {
}

func (*awsRestjson1_serializeOpCancelSigningProfile) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCancelSigningProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CancelSigningProfileInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/signing-profiles/{profileName}")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "DELETE"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsCancelSigningProfileInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsCancelSigningProfileInput(v *CancelSigningProfileInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ProfileName == nil {
		return &smithy.SerializationError{Err: fmt.Errorf("input member profileName must not be empty")}
	}
	if v.ProfileName != nil {
		if len(*v.ProfileName) == 0 {
			return &smithy.SerializationError{Err: fmt.Errorf("input member profileName must not be empty")}
		}
		if err := encoder.SetURI("profileName").String(*v.ProfileName); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpDescribeSigningJob struct {
}

func (*awsRestjson1_serializeOpDescribeSigningJob) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDescribeSigningJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeSigningJobInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/signing-jobs/{jobId}")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDescribeSigningJobInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDescribeSigningJobInput(v *DescribeSigningJobInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.JobId == nil {
		return &smithy.SerializationError{Err: fmt.Errorf("input member jobId must not be empty")}
	}
	if v.JobId != nil {
		if len(*v.JobId) == 0 {
			return &smithy.SerializationError{Err: fmt.Errorf("input member jobId must not be empty")}
		}
		if err := encoder.SetURI("jobId").String(*v.JobId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetSigningPlatform struct {
}

func (*awsRestjson1_serializeOpGetSigningPlatform) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetSigningPlatform) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetSigningPlatformInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/signing-platforms/{platformId}")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetSigningPlatformInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetSigningPlatformInput(v *GetSigningPlatformInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.PlatformId == nil {
		return &smithy.SerializationError{Err: fmt.Errorf("input member platformId must not be empty")}
	}
	if v.PlatformId != nil {
		if len(*v.PlatformId) == 0 {
			return &smithy.SerializationError{Err: fmt.Errorf("input member platformId must not be empty")}
		}
		if err := encoder.SetURI("platformId").String(*v.PlatformId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetSigningProfile struct {
}

func (*awsRestjson1_serializeOpGetSigningProfile) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetSigningProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetSigningProfileInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/signing-profiles/{profileName}")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetSigningProfileInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetSigningProfileInput(v *GetSigningProfileInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ProfileName == nil {
		return &smithy.SerializationError{Err: fmt.Errorf("input member profileName must not be empty")}
	}
	if v.ProfileName != nil {
		if len(*v.ProfileName) == 0 {
			return &smithy.SerializationError{Err: fmt.Errorf("input member profileName must not be empty")}
		}
		if err := encoder.SetURI("profileName").String(*v.ProfileName); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListSigningJobs struct {
}

func (*awsRestjson1_serializeOpListSigningJobs) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListSigningJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListSigningJobsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/signing-jobs")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListSigningJobsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListSigningJobsInput(v *ListSigningJobsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	if v.PlatformId != nil {
		encoder.SetQuery("platformId").String(*v.PlatformId)
	}

	if v.RequestedBy != nil {
		encoder.SetQuery("requestedBy").String(*v.RequestedBy)
	}

	if len(v.Status) > 0 {
		encoder.SetQuery("status").String(string(v.Status))
	}

	return nil
}

type awsRestjson1_serializeOpListSigningPlatforms struct {
}

func (*awsRestjson1_serializeOpListSigningPlatforms) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListSigningPlatforms) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListSigningPlatformsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/signing-platforms")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListSigningPlatformsInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListSigningPlatformsInput(v *ListSigningPlatformsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Category != nil {
		encoder.SetQuery("category").String(*v.Category)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	if v.Partner != nil {
		encoder.SetQuery("partner").String(*v.Partner)
	}

	if v.Target != nil {
		encoder.SetQuery("target").String(*v.Target)
	}

	return nil
}

type awsRestjson1_serializeOpListSigningProfiles struct {
}

func (*awsRestjson1_serializeOpListSigningProfiles) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListSigningProfiles) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListSigningProfilesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/signing-profiles")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListSigningProfilesInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListSigningProfilesInput(v *ListSigningProfilesInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.IncludeCanceled != nil {
		encoder.SetQuery("includeCanceled").Boolean(*v.IncludeCanceled)
	}

	if v.MaxResults != nil {
		encoder.SetQuery("maxResults").Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	return nil
}

type awsRestjson1_serializeOpListTagsForResource struct {
}

func (*awsRestjson1_serializeOpListTagsForResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListTagsForResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/tags/{resourceArn}")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceArn == nil {
		return &smithy.SerializationError{Err: fmt.Errorf("input member resourceArn must not be empty")}
	}
	if v.ResourceArn != nil {
		if len(*v.ResourceArn) == 0 {
			return &smithy.SerializationError{Err: fmt.Errorf("input member resourceArn must not be empty")}
		}
		if err := encoder.SetURI("resourceArn").String(*v.ResourceArn); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpPutSigningProfile struct {
}

func (*awsRestjson1_serializeOpPutSigningProfile) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpPutSigningProfile) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PutSigningProfileInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/signing-profiles/{profileName}")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "PUT"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsPutSigningProfileInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentPutSigningProfileInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsPutSigningProfileInput(v *PutSigningProfileInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ProfileName == nil {
		return &smithy.SerializationError{Err: fmt.Errorf("input member profileName must not be empty")}
	}
	if v.ProfileName != nil {
		if len(*v.ProfileName) == 0 {
			return &smithy.SerializationError{Err: fmt.Errorf("input member profileName must not be empty")}
		}
		if err := encoder.SetURI("profileName").String(*v.ProfileName); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentPutSigningProfileInput(v *PutSigningProfileInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Overrides != nil {
		ok := object.Key("overrides")
		if err := awsRestjson1_serializeDocumentSigningPlatformOverrides(v.Overrides, ok); err != nil {
			return err
		}
	}

	if v.PlatformId != nil {
		ok := object.Key("platformId")
		ok.String(*v.PlatformId)
	}

	if v.SigningMaterial != nil {
		ok := object.Key("signingMaterial")
		if err := awsRestjson1_serializeDocumentSigningMaterial(v.SigningMaterial, ok); err != nil {
			return err
		}
	}

	if v.SigningParameters != nil {
		ok := object.Key("signingParameters")
		if err := awsRestjson1_serializeDocumentSigningParameters(v.SigningParameters, ok); err != nil {
			return err
		}
	}

	if v.Tags != nil {
		ok := object.Key("tags")
		if err := awsRestjson1_serializeDocumentTagMap(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpStartSigningJob struct {
}

func (*awsRestjson1_serializeOpStartSigningJob) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpStartSigningJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*StartSigningJobInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/signing-jobs")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentStartSigningJobInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsStartSigningJobInput(v *StartSigningJobInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentStartSigningJobInput(v *StartSigningJobInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientRequestToken != nil {
		ok := object.Key("clientRequestToken")
		ok.String(*v.ClientRequestToken)
	}

	if v.Destination != nil {
		ok := object.Key("destination")
		if err := awsRestjson1_serializeDocumentDestination(v.Destination, ok); err != nil {
			return err
		}
	}

	if v.ProfileName != nil {
		ok := object.Key("profileName")
		ok.String(*v.ProfileName)
	}

	if v.Source != nil {
		ok := object.Key("source")
		if err := awsRestjson1_serializeDocumentSource(v.Source, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpTagResource struct {
}

func (*awsRestjson1_serializeOpTagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*TagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/tags/{resourceArn}")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsTagResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentTagResourceInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceArn == nil {
		return &smithy.SerializationError{Err: fmt.Errorf("input member resourceArn must not be empty")}
	}
	if v.ResourceArn != nil {
		if len(*v.ResourceArn) == 0 {
			return &smithy.SerializationError{Err: fmt.Errorf("input member resourceArn must not be empty")}
		}
		if err := encoder.SetURI("resourceArn").String(*v.ResourceArn); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentTagResourceInput(v *TagResourceInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Tags != nil {
		ok := object.Key("tags")
		if err := awsRestjson1_serializeDocumentTagMap(v.Tags, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpUntagResource struct {
}

func (*awsRestjson1_serializeOpUntagResource) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UntagResourceInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/tags/{resourceArn}")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "DELETE"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsUntagResourceInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ResourceArn == nil {
		return &smithy.SerializationError{Err: fmt.Errorf("input member resourceArn must not be empty")}
	}
	if v.ResourceArn != nil {
		if len(*v.ResourceArn) == 0 {
			return &smithy.SerializationError{Err: fmt.Errorf("input member resourceArn must not be empty")}
		}
		if err := encoder.SetURI("resourceArn").String(*v.ResourceArn); err != nil {
			return err
		}
	}

	if v.TagKeys != nil {
		for i := range v.TagKeys {
			if v.TagKeys[i] == nil {
				continue
			}
			encoder.AddQuery("tagKeys").String(*v.TagKeys[i])
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentDestination(v *types.Destination, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.S3 != nil {
		ok := object.Key("s3")
		if err := awsRestjson1_serializeDocumentS3Destination(v.S3, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentS3Destination(v *types.S3Destination, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.BucketName != nil {
		ok := object.Key("bucketName")
		ok.String(*v.BucketName)
	}

	if v.Prefix != nil {
		ok := object.Key("prefix")
		ok.String(*v.Prefix)
	}

	return nil
}

func awsRestjson1_serializeDocumentS3Source(v *types.S3Source, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.BucketName != nil {
		ok := object.Key("bucketName")
		ok.String(*v.BucketName)
	}

	if v.Key != nil {
		ok := object.Key("key")
		ok.String(*v.Key)
	}

	if v.Version != nil {
		ok := object.Key("version")
		ok.String(*v.Version)
	}

	return nil
}

func awsRestjson1_serializeDocumentSigningConfigurationOverrides(v *types.SigningConfigurationOverrides, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.EncryptionAlgorithm) > 0 {
		ok := object.Key("encryptionAlgorithm")
		ok.String(string(v.EncryptionAlgorithm))
	}

	if len(v.HashAlgorithm) > 0 {
		ok := object.Key("hashAlgorithm")
		ok.String(string(v.HashAlgorithm))
	}

	return nil
}

func awsRestjson1_serializeDocumentSigningMaterial(v *types.SigningMaterial, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CertificateArn != nil {
		ok := object.Key("certificateArn")
		ok.String(*v.CertificateArn)
	}

	return nil
}

func awsRestjson1_serializeDocumentSigningParameters(v map[string]*string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		if vv := v[key]; vv == nil {
			om.Null()
			continue
		}
		om.String(*v[key])
	}
	return nil
}

func awsRestjson1_serializeDocumentSigningPlatformOverrides(v *types.SigningPlatformOverrides, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.SigningConfiguration != nil {
		ok := object.Key("signingConfiguration")
		if err := awsRestjson1_serializeDocumentSigningConfigurationOverrides(v.SigningConfiguration, ok); err != nil {
			return err
		}
	}

	if len(v.SigningImageFormat) > 0 {
		ok := object.Key("signingImageFormat")
		ok.String(string(v.SigningImageFormat))
	}

	return nil
}

func awsRestjson1_serializeDocumentSource(v *types.Source, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.S3 != nil {
		ok := object.Key("s3")
		if err := awsRestjson1_serializeDocumentS3Source(v.S3, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentTagMap(v map[string]*string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		if vv := v[key]; vv == nil {
			om.Null()
			continue
		}
		om.String(*v[key])
	}
	return nil
}
