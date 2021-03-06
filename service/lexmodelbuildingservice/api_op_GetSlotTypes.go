// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns slot type information as follows:
//
// * If you specify the nameContains
// field, returns the $LATEST version of all slot types that contain the specified
// string.
//
// * If you don't specify the nameContains field, returns information
// about the $LATEST version of all slot types.
//
// The operation requires permission
// for the lex:GetSlotTypes action.
func (c *Client) GetSlotTypes(ctx context.Context, params *GetSlotTypesInput, optFns ...func(*Options)) (*GetSlotTypesOutput, error) {
	if params == nil {
		params = &GetSlotTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSlotTypes", params, optFns, addOperationGetSlotTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSlotTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSlotTypesInput struct {

	// The maximum number of slot types to return in the response. The default is 10.
	MaxResults *int32

	// Substring to match in slot type names. A slot type will be returned if any part
	// of its name matches the substring. For example, "xyz" matches both "xyzabc" and
	// "abcxyz."
	NameContains *string

	// A pagination token that fetches the next page of slot types. If the response to
	// this API call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch next page of slot types, specify the pagination token in the
	// next request.
	NextToken *string
}

type GetSlotTypesOutput struct {

	// If the response is truncated, it includes a pagination token that you can
	// specify in your next request to fetch the next page of slot types.
	NextToken *string

	// An array of objects, one for each slot type, that provides information such as
	// the name of the slot type, the version, and a description.
	SlotTypes []*types.SlotTypeMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSlotTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSlotTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSlotTypes{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSlotTypes(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetSlotTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetSlotTypes",
	}
}
