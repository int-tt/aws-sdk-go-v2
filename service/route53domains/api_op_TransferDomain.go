// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Transfers a domain from another registrar to Amazon Route 53. When the transfer
// is complete, the domain is registered either with Amazon Registrar (for .com,
// .net, and .org domains) or with our registrar associate, Gandi (for all other
// TLDs). For more information about transferring domains, see the following
// topics:
//
//     * For transfer requirements, a detailed procedure, and information
// about viewing the status of a domain that you're transferring to Route 53, see
// Transferring Registration for a Domain to Amazon Route 53
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-transfer-to-route-53.html)
// in the Amazon Route 53 Developer Guide.
//
//     * For information about how to
// transfer a domain from one AWS account to another, see
// TransferDomainToAnotherAwsAccount
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html).
//
//
// * For information about how to transfer a domain to another domain registrar,
// see Transferring a Domain from Amazon Route 53 to Another Registrar
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-transfer-from-route-53.html)
// in the Amazon Route 53 Developer Guide.
//
// If the registrar for your domain is
// also the DNS service provider for the domain, we highly recommend that you
// transfer your DNS service to Route 53 or to another DNS service provider before
// you transfer your registration. Some registrars provide free DNS service when
// you purchase a domain registration. When you transfer the registration, the
// previous registrar will not renew your domain registration and could end your
// DNS service at any time.  <important> <p>If the registrar for your domain is
// also the DNS service provider for the domain and you don't transfer DNS service
// to another provider, your website, email, and the web applications associated
// with the domain might become unavailable.</p> </important> <p>If the transfer is
// successful, this method returns an operation ID that you can use to track the
// progress and completion of the action. If the transfer doesn't complete
// successfully, the domain registrant will be notified by email.</p>
func (c *Client) TransferDomain(ctx context.Context, params *TransferDomainInput, optFns ...func(*Options)) (*TransferDomainOutput, error) {
	stack := middleware.NewStack("TransferDomain", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpTransferDomainMiddlewares(stack)
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
	addOpTransferDomainValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTransferDomain(options.Region), middleware.Before)

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
			OperationName: "TransferDomain",
			Err:           err,
		}
	}
	out := result.(*TransferDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The TransferDomain request includes the following elements.
type TransferDomainInput struct {
	// The number of years that you want to register the domain for. Domains are
	// registered for a minimum of one year. The maximum period depends on the
	// top-level domain. Default: 1
	DurationInYears *int32
	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the technical contact. Default: true
	PrivacyProtectTechContact *bool
	// Indicates whether the domain will be automatically renewed (true) or not
	// (false). Autorenewal only takes effect after the account is charged. Default:
	// true
	AutoRenew *bool
	// Contains details for the host and glue IP addresses.
	Nameservers []*types.Nameserver
	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the registrant contact (domain
	// owner). Default: true
	PrivacyProtectRegistrantContact *bool
	// The authorization code for the domain. You get this value from the current
	// registrar.
	AuthCode *string
	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either for
	// Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the admin contact. Default: true
	PrivacyProtectAdminContact *bool
	// Reserved for future use.
	IdnLangCode *string
	// The name of the domain that you want to transfer to Route 53. The top-level
	// domain (TLD), such as .com, must be a TLD that Route 53 supports. For a list of
	// supported TLDs, see Domains that You Can Register with Amazon Route 53
	// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html)
	// in the Amazon Route 53 Developer Guide. The domain name can contain only the
	// following characters:
	//
	//     * Letters a through z. Domain names are not case
	// sensitive.
	//
	//     * Numbers 0 through 9.
	//
	//     * Hyphen (-). You can't specify a
	// hyphen at the beginning or end of a label.
	//
	//     * Period (.) to separate the
	// labels in the name, such as the . in example.com.
	DomainName *string
	// Provides detailed contact information.
	AdminContact *types.ContactDetail
	// Provides detailed contact information.
	RegistrantContact *types.ContactDetail
	// Provides detailed contact information.
	TechContact *types.ContactDetail
}

// The TransferDomain response includes the following element.
type TransferDomainOutput struct {
	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html).
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpTransferDomainMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpTransferDomain{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpTransferDomain{}, middleware.After)
}

func newServiceMetadataMiddleware_opTransferDomain(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "TransferDomain",
	}
}