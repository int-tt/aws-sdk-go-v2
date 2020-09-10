// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Determines the dominant language of the input text. For a list of languages that
// Amazon Comprehend can detect, see Amazon Comprehend Supported Languages
// (https://docs.aws.amazon.com/comprehend/latest/dg/how-languages.html).
func (c *Client) DetectDominantLanguage(ctx context.Context, params *DetectDominantLanguageInput, optFns ...func(*Options)) (*DetectDominantLanguageOutput, error) {
	stack := middleware.NewStack("DetectDominantLanguage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDetectDominantLanguageMiddlewares(stack)
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
	addOpDetectDominantLanguageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetectDominantLanguage(options.Region), middleware.Before)

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
			OperationName: "DetectDominantLanguage",
			Err:           err,
		}
	}
	out := result.(*DetectDominantLanguageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectDominantLanguageInput struct {
	// A UTF-8 text string. Each string should contain at least 20 characters and must
	// contain fewer that 5,000 bytes of UTF-8 encoded characters.
	Text *string
}

type DetectDominantLanguageOutput struct {
	// The languages that Amazon Comprehend detected in the input text. For each
	// language, the response returns the RFC 5646 language code and the level of
	// confidence that Amazon Comprehend has in the accuracy of its inference. For more
	// information about RFC 5646, see Tags for Identifying Languages
	// (https://tools.ietf.org/html/rfc5646) on the IETF Tools web site.
	Languages []*types.DominantLanguage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDetectDominantLanguageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDetectDominantLanguage{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectDominantLanguage{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetectDominantLanguage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DetectDominantLanguage",
	}
}