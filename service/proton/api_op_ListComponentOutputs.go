// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get a list of component Infrastructure as Code (IaC) outputs. For more
// information about components, see Proton components
// (https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html) in the
// Proton User Guide.
func (c *Client) ListComponentOutputs(ctx context.Context, params *ListComponentOutputsInput, optFns ...func(*Options)) (*ListComponentOutputsOutput, error) {
	if params == nil {
		params = &ListComponentOutputsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComponentOutputs", params, optFns, c.addOperationListComponentOutputsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComponentOutputsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComponentOutputsInput struct {

	// The name of the component whose outputs you want.
	//
	// This member is required.
	ComponentName *string

	// A token that indicates the location of the next output in the array of outputs,
	// after the list of outputs that was previously requested.
	NextToken *string

	noSmithyDocumentSerde
}

type ListComponentOutputsOutput struct {

	// An array of component Infrastructure as Code (IaC) outputs.
	//
	// This member is required.
	Outputs []types.Output

	// A token that indicates the location of the next output in the array of outputs,
	// after the list of outputs that was previously requested.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListComponentOutputsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListComponentOutputs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListComponentOutputs{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpListComponentOutputsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListComponentOutputs(options.Region), middleware.Before); err != nil {
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

// ListComponentOutputsAPIClient is a client that implements the
// ListComponentOutputs operation.
type ListComponentOutputsAPIClient interface {
	ListComponentOutputs(context.Context, *ListComponentOutputsInput, ...func(*Options)) (*ListComponentOutputsOutput, error)
}

var _ ListComponentOutputsAPIClient = (*Client)(nil)

// ListComponentOutputsPaginatorOptions is the paginator options for
// ListComponentOutputs
type ListComponentOutputsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListComponentOutputsPaginator is a paginator for ListComponentOutputs
type ListComponentOutputsPaginator struct {
	options   ListComponentOutputsPaginatorOptions
	client    ListComponentOutputsAPIClient
	params    *ListComponentOutputsInput
	nextToken *string
	firstPage bool
}

// NewListComponentOutputsPaginator returns a new ListComponentOutputsPaginator
func NewListComponentOutputsPaginator(client ListComponentOutputsAPIClient, params *ListComponentOutputsInput, optFns ...func(*ListComponentOutputsPaginatorOptions)) *ListComponentOutputsPaginator {
	if params == nil {
		params = &ListComponentOutputsInput{}
	}

	options := ListComponentOutputsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListComponentOutputsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListComponentOutputsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListComponentOutputs page.
func (p *ListComponentOutputsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListComponentOutputsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListComponentOutputs(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListComponentOutputs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "ListComponentOutputs",
	}
}
