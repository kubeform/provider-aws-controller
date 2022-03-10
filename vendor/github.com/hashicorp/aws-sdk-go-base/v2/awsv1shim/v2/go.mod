module github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2

require (
	github.com/aws/aws-sdk-go v1.42.52
	github.com/aws/aws-sdk-go-v2 v1.14.0
	github.com/aws/aws-sdk-go-v2/config v1.14.0
	github.com/aws/aws-sdk-go-v2/credentials v1.9.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.15.0
	github.com/google/go-cmp v0.5.7
	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.12
	github.com/hashicorp/go-cleanhttp v0.5.2
)

go 1.16

replace github.com/hashicorp/aws-sdk-go-base/v2 => ../..
