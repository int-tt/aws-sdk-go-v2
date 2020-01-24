// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDBProxiesInput struct {
	_ struct{} `type:"structure"`

	// The name of the DB proxy.
	DBProxyName *string `type:"string"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `min:"20" type:"integer"`
}

// String returns the string representation
func (s DescribeDBProxiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBProxiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBProxiesInput"}
	if s.MaxRecords != nil && *s.MaxRecords < 20 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxRecords", 20))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDBProxiesOutput struct {
	_ struct{} `type:"structure"`

	// A return value representing an arbitrary number of DBProxy data structures.
	DBProxies []DBProxy `type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBProxiesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBProxies = "DescribeDBProxies"

// DescribeDBProxiesRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
//
// This is prerelease documentation for the RDS Database Proxy feature in preview
// release. It is subject to change.
//
// Returns information about DB proxies.
//
//    // Example sending a request using DescribeDBProxiesRequest.
//    req := client.DescribeDBProxiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBProxies
func (c *Client) DescribeDBProxiesRequest(input *DescribeDBProxiesInput) DescribeDBProxiesRequest {
	op := &aws.Operation{
		Name:       opDescribeDBProxies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeDBProxiesInput{}
	}

	req := c.newRequest(op, input, &DescribeDBProxiesOutput{})
	return DescribeDBProxiesRequest{Request: req, Input: input, Copy: c.DescribeDBProxiesRequest}
}

// DescribeDBProxiesRequest is the request type for the
// DescribeDBProxies API operation.
type DescribeDBProxiesRequest struct {
	*aws.Request
	Input *DescribeDBProxiesInput
	Copy  func(*DescribeDBProxiesInput) DescribeDBProxiesRequest
}

// Send marshals and sends the DescribeDBProxies API request.
func (r DescribeDBProxiesRequest) Send(ctx context.Context) (*DescribeDBProxiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBProxiesResponse{
		DescribeDBProxiesOutput: r.Request.Data.(*DescribeDBProxiesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDBProxiesRequestPaginator returns a paginator for DescribeDBProxies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDBProxiesRequest(input)
//   p := rds.NewDescribeDBProxiesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDBProxiesPaginator(req DescribeDBProxiesRequest) DescribeDBProxiesPaginator {
	return DescribeDBProxiesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDBProxiesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeDBProxiesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDBProxiesPaginator struct {
	aws.Pager
}

func (p *DescribeDBProxiesPaginator) CurrentPage() *DescribeDBProxiesOutput {
	return p.Pager.CurrentPage().(*DescribeDBProxiesOutput)
}

// DescribeDBProxiesResponse is the response type for the
// DescribeDBProxies API operation.
type DescribeDBProxiesResponse struct {
	*DescribeDBProxiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBProxies request.
func (r *DescribeDBProxiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}