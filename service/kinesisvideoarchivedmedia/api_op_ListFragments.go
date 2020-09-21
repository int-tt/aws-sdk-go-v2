// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideoarchivedmedia

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of Fragment () objects from the specified stream and timestamp
// range within the archived data. Listing fragments is eventually consistent. This
// means that even if the producer receives an acknowledgment that a fragment is
// persisted, the result might not be returned immediately from a request to
// ListFragments. However, results are typically available in less than one second.
// You must first call the GetDataEndpoint API to get an endpoint. Then send the
// ListFragments requests to this endpoint using the --endpoint-url parameter
// (https://docs.aws.amazon.com/cli/latest/reference/).  <important> <p>If an error
// is thrown after invoking a Kinesis Video Streams archived media API, in addition
// to the HTTP status code and the response body, it includes the following pieces
// of information: </p> <ul> <li> <p> <code>x-amz-ErrorType</code> HTTP header –
// contains a more specific error type in addition to what the HTTP status code
// provides. </p> </li> <li> <p> <code>x-amz-RequestId</code> HTTP header – if you
// want to report an issue to AWS, the support team can better diagnose the problem
// if given the Request Id.</p> </li> </ul> <p>Both the HTTP status code and the
// ErrorType header can be utilized to make programmatic decisions about whether
// errors are retry-able and under what conditions, as well as provide information
// on what actions the client programmer might need to take in order to
// successfully try again.</p> <p>For more information, see the <b>Errors</b>
// section at the bottom of this topic, as well as <a
// href="https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/CommonErrors.html">Common
// Errors</a>. </p> </important>
func (c *Client) ListFragments(ctx context.Context, params *ListFragmentsInput, optFns ...func(*Options)) (*ListFragmentsOutput, error) {
	stack := middleware.NewStack("ListFragments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListFragmentsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListFragmentsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListFragments(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "ListFragments",
			Err:           err,
		}
	}
	out := result.(*ListFragmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFragmentsInput struct {
	// The name of the stream from which to retrieve a fragment list.
	StreamName *string
	// The total number of fragments to return. If the total number of fragments
	// available is more than the value specified in max-results, then a
	// ListFragmentsOutput$NextToken () is provided in the output that you can use to
	// resume pagination.
	MaxResults *int64
	// A token to specify where to start paginating. This is the
	// ListFragmentsOutput$NextToken () from a previously truncated response.
	NextToken *string
	// Describes the timestamp range and timestamp origin for the range of fragments to
	// return.
	FragmentSelector *types.FragmentSelector
}

type ListFragmentsOutput struct {
	// A list of archived Fragment () objects from the stream that meet the selector
	// criteria. Results are in no specific order, even across pages.
	Fragments []*types.Fragment
	// If the returned list is truncated, the operation returns this token to use to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListFragmentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListFragments{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListFragments{}, middleware.After)
}

func newServiceMetadataMiddleware_opListFragments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "ListFragments",
	}
}