module github.com/GoogleCloudPlatform/proto-gen-parser

replace github.com/GoogleCloudPlatform/proto-gen-parser => ./cmd

replace github.com/GoogleCloudPlatform/proto-gen-parser/pkg/api => ./pkg/api

replace github.com/GoogleCloudPlatform/proto-gen-parser/pkg/pb => ./pkg/pb

replace github.com/GoogleCloudPlatform/proto-gen-parser/pkg/reader => ./pkg/reader

go 1.20

require github.com/stretchr/testify v1.8.2

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/api v0.121.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/grpc v1.55.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
