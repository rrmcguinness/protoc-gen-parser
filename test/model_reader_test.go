package test

import (
	"testing"

	"github.com/GoogleCloudPlatform/proto-gen-parser/pkg/logging"
	"github.com/GoogleCloudPlatform/proto-gen-parser/pkg/reader"
	"github.com/stretchr/testify/assert"
)

func TestMessageReader(t *testing.T) {
	log := logging.NewLogger(true, "model read test")

	pkg, err := reader.ReadPackage("data/api/location/model.proto", false)
	if err != nil {
		log.Errorf("Error reading file: %v", err)
		assert.Fail(t, "failed to read file")
	}

	assert.NotNil(t, pkg)

	assert.Equal(t, "test.location", pkg.Name())

	assert.Equal(t, 3, len(pkg.Imports()))
	assert.Equal(t, 3, len(pkg.Options()))
	assert.Equal(t, 2, len(pkg.Messages()))
	assert.Equal(t, 1, len(pkg.Enums()))
}
