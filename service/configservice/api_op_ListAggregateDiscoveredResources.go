// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Accepts a resource type and returns a list of resource identifiers that are
// aggregated for a specific resource type across accounts and regions. A resource
// identifier includes the resource type, ID, (if available) the custom resource
// name, source account, and source region. You can narrow the results to include
// only resources that have specific resource IDs, or a resource name, or source
// account ID, or source region. For example, if the input consists of accountID
// 12345678910 and the region is us-east-1 for resource type AWS::EC2::Instance
// then the API returns all the EC2 instance identifiers of accountID 12345678910
// and region us-east-1.
func (c *Client) ListAggregateDiscoveredResources(ctx context.Context, params *ListAggregateDiscoveredResourcesInput, optFns ...func(*Options)) (*ListAggregateDiscoveredResourcesOutput, error) {
	if params == nil {
		params = &ListAggregateDiscoveredResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAggregateDiscoveredResources", params, optFns, c.addOperationListAggregateDiscoveredResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAggregateDiscoveredResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAggregateDiscoveredResourcesInput struct {

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// The type of resources that you want Config to list in the response.
	//
	// This member is required.
	ResourceType types.ResourceType

	// Filters the results based on the ResourceFilters object.
	Filters *types.ResourceFilters

	// The maximum number of resource identifiers returned on each page. You cannot
	// specify a number greater than 100. If you specify 0, Config uses the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAggregateDiscoveredResourcesOutput struct {

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// Returns a list of ResourceIdentifiers objects.
	ResourceIdentifiers []types.AggregateResourceIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAggregateDiscoveredResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAggregateDiscoveredResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAggregateDiscoveredResources{}, middleware.After)
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
	if err = addOpListAggregateDiscoveredResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAggregateDiscoveredResources(options.Region), middleware.Before); err != nil {
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

// ListAggregateDiscoveredResourcesAPIClient is a client that implements the
// ListAggregateDiscoveredResources operation.
type ListAggregateDiscoveredResourcesAPIClient interface {
	ListAggregateDiscoveredResources(context.Context, *ListAggregateDiscoveredResourcesInput, ...func(*Options)) (*ListAggregateDiscoveredResourcesOutput, error)
}

var _ ListAggregateDiscoveredResourcesAPIClient = (*Client)(nil)

// ListAggregateDiscoveredResourcesPaginatorOptions is the paginator options for
// ListAggregateDiscoveredResources
type ListAggregateDiscoveredResourcesPaginatorOptions struct {
	// The maximum number of resource identifiers returned on each page. You cannot
	// specify a number greater than 100. If you specify 0, Config uses the default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAggregateDiscoveredResourcesPaginator is a paginator for
// ListAggregateDiscoveredResources
type ListAggregateDiscoveredResourcesPaginator struct {
	options   ListAggregateDiscoveredResourcesPaginatorOptions
	client    ListAggregateDiscoveredResourcesAPIClient
	params    *ListAggregateDiscoveredResourcesInput
	nextToken *string
	firstPage bool
}

// NewListAggregateDiscoveredResourcesPaginator returns a new
// ListAggregateDiscoveredResourcesPaginator
func NewListAggregateDiscoveredResourcesPaginator(client ListAggregateDiscoveredResourcesAPIClient, params *ListAggregateDiscoveredResourcesInput, optFns ...func(*ListAggregateDiscoveredResourcesPaginatorOptions)) *ListAggregateDiscoveredResourcesPaginator {
	if params == nil {
		params = &ListAggregateDiscoveredResourcesInput{}
	}

	options := ListAggregateDiscoveredResourcesPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAggregateDiscoveredResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAggregateDiscoveredResourcesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListAggregateDiscoveredResources page.
func (p *ListAggregateDiscoveredResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAggregateDiscoveredResourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.ListAggregateDiscoveredResources(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListAggregateDiscoveredResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "ListAggregateDiscoveredResources",
	}
}
