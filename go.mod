module github.com/camptocamp/terraform-provider-pass

replace github.com/stretchr/testify => github.com/stretchrcom/testify v1.6.1

require (
	github.com/agext/levenshtein v1.2.2 // indirect
	github.com/aws/aws-sdk-go v1.40.27
	github.com/blang/semver v3.5.1+incompatible
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/gopasspw/gopass v1.11.0
	github.com/hashicorp/hcl/v2 v2.6.0
	github.com/kr/pretty v0.2.0
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.3.2
	github.com/pkg/errors v0.9.1
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	github.com/zclconf/go-cty v1.5.1 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

go 1.15
