// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of predictors created using the CreatePredictor () operation. For
// each predictor, this operation returns a summary of its properties, including
// its Amazon Resource Name (ARN). You can retrieve the complete set of properties
// by using the ARN with the DescribePredictor () operation. You can filter the
// list using an array of Filter () objects.
func (c *Client) ListPredictors(ctx context.Context, params *ListPredictorsInput, optFns ...func(*Options)) (*ListPredictorsOutput, error) {
	stack := middleware.NewStack("ListPredictors", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListPredictorsMiddlewares(stack)
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
	addOpListPredictorsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPredictors(options.Region), middleware.Before)

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
			OperationName: "ListPredictors",
			Err:           err,
		}
	}
	out := result.(*ListPredictorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPredictorsInput struct {
	// The number of items to return in the response.
	MaxResults *int32
	// If the result of the previous request was truncated, the response includes a
	// NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string
	// An array of filters. For each filter, you provide a condition and a match
	// statement. The condition is either IS or IS_NOT, which specifies whether to
	// include or exclude the predictors that match the statement from the list,
	// respectively. The match statement consists of a key and a value. Filter
	// properties
	//
	//     * Condition - The condition to apply. Valid values are IS and
	// IS_NOT. To include the predictors that match the statement, specify IS. To
	// exclude matching predictors, specify IS_NOT.
	//
	//     * Key - The name of the
	// parameter to filter on. Valid values are DatasetGroupArn and Status.
	//
	//     *
	// Value - The value to match.
	//
	// For example, to list all predictors whose status is
	// ACTIVE, you would specify: "Filters": [ { "Condition": "IS", "Key": "Status",
	// "Value": "ACTIVE" } ]
	Filters []*types.Filter
}

type ListPredictorsOutput struct {
	// If the response is truncated, Amazon Forecast returns this token. To retrieve
	// the next set of results, use the token in the next request.
	NextToken *string
	// An array of objects that summarize each predictor's properties.
	Predictors []*types.PredictorSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListPredictorsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListPredictors{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPredictors{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPredictors(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "ListPredictors",
	}
}