// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package route53resolver provides the client and types for making API
// requests to Amazon Route 53 Resolver.
//
// When you create a VPC using Amazon VPC, you automatically get DNS resolution
// within the VPC from Route 53 Resolver. By default, Resolver answers DNS queries
// for VPC domain names such as domain names for EC2 instances or Elastic Load
// Balancing load balancers. Resolver performs recursive lookups against public
// name servers for all other domain names.
//
// You can also configure DNS resolution between your VPC and your network over
// a Direct Connect or VPN connection:
//
// Forward DNS queries from resolvers on your network to Route 53 Resolver
//
// DNS resolvers on your network can forward DNS queries to Resolver in a specified
// VPC. This allows your DNS resolvers to easily resolve domain names for Amazon
// Web Services resources such as EC2 instances or records in a Route 53 private
// hosted zone. For more information, see How DNS Resolvers on Your Network
// Forward DNS Queries to Route 53 Resolver (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resolver.html#resolver-overview-forward-network-to-vpc)
// in the Amazon Route 53 Developer Guide.
//
// Conditionally forward queries from a VPC to resolvers on your network
//
// You can configure Resolver to forward queries that it receives from EC2 instances
// in your VPCs to DNS resolvers on your network. To forward selected queries,
// you create Resolver rules that specify the domain names for the DNS queries
// that you want to forward (such as example.com), and the IP addresses of the
// DNS resolvers on your network that you want to forward the queries to. If
// a query matches multiple rules (example.com, acme.example.com), Resolver
// chooses the rule with the most specific match (acme.example.com) and forwards
// the query to the IP addresses that you specified in that rule. For more information,
// see How Route 53 Resolver Forwards DNS Queries from Your VPCs to Your Network
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resolver.html#resolver-overview-forward-vpc-to-network)
// in the Amazon Route 53 Developer Guide.
//
// Like Amazon VPC, Resolver is Regional. In each Region where you have VPCs,
// you can choose whether to forward queries from your VPCs to your network
// (outbound queries), from your network to your VPCs (inbound queries), or
// both.
//
// See https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01 for more information on this service.
//
// See route53resolver package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/route53resolver/
//
// Using the Client
//
// To contact Amazon Route 53 Resolver with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon Route 53 Resolver client Route53Resolver for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/route53resolver/#New
package route53resolver
