module github.com/GoogleCloudPlatform/proto-gen-parser

replace github.com/GoogleCloudPlatform/proto-gen-parser => ./cmd

replace github.com/GoogleCloudPlatform/proto-gen-parser/pkg/api => ./pkg/api

replace github.com/GoogleCloudPlatform/proto-gen-parser/pkg/pb => ./pkg/pb

replace github.com/GoogleCloudPlatform/proto-gen-parser/pkg/reader => ./pkg/reader

go 1.20

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
