package reader

import (
	"github.com/GoogleCloudPlatform/proto-gen-parser/pkg/api"
	"github.com/GoogleCloudPlatform/proto-gen-parser/pkg/logging"
	"strings"
)

// NewAttributeVisitor - Constructor for the AttributeVisitor
func NewAttributeVisitor(debug bool) *AttributeVisitor {
	return &AttributeVisitor{Log: logging.NewLogger(debug, "attribute_visitor")}
}

// AttributeVisitor implementation for attributes
type AttributeVisitor struct {
	Log *logging.Logger
}

// CanVisit - Determines if the line is an attribute, it doesn't end in a brace,
// it's a map, repeated, or can effectively be split
func (av *AttributeVisitor) CanVisit(in *Line) bool {
	return (!strings.HasSuffix(in.Syntax, OpenBrace) || !strings.HasSuffix(in.Syntax, ClosedBrace)) &&
		strings.HasPrefix(in.Syntax, "repeated") ||
		strings.HasPrefix(in.Syntax, "map") || len(SplitSyntax(in.Syntax)) >= 4
}

// HandleRepeated marshals the attribute into a repeated representation, e.g. List.
func HandleRepeated(qualifier string,
	comment string,
	split []string,
) api.Attribute {
	return ProtobufFactory.NewAttribute(qualifier, split[2], comment, true, false, ParseOrdinal(split[4]), split[1])
}

// HandleMap marshals the attribute into a Map type by using multiple types for key and value.
func HandleMap(qualifier string,
	comment string,
	split []string,
) api.Attribute {
	mapValue := Join(Space, split[0], split[1])
	innerTypes := mapValue[strings.Index(mapValue, OpenMap)+len(OpenMap) : strings.Index(mapValue, ClosedMap)]
	splitTypes := strings.Split(innerTypes, Comma)

	return ProtobufFactory.NewAttribute(qualifier, split[2], comment, false, true,
		ParseOrdinal(split[4]), splitTypes...)
}

// HandleDefaultAttribute marshals a standard attribute type.
func HandleDefaultAttribute(qualifier string,
	comment string,
	split []string,
) api.Attribute {
	if len(split) >= 3 {
		return ProtobufFactory.NewAttribute(qualifier, split[1], comment, false, false, ParseOrdinal(split[3]), split[0])
	}
	return nil
}

// Visit is used for marshalling an attribute into a struct.
func (av *AttributeVisitor) Visit(_ Scanner, in *Line, namespace string) interface{} {
	av.Log.Debug("Visiting Attribute")

	split := SplitSyntax(in.Syntax)
	comment := ""
	var out api.Attribute

	if strings.HasPrefix(in.Syntax, PrefixReserved) {
		av.Log.Debug("\t processing reserved attribute")
		comment += in.Comment
	} else if strings.HasPrefix(in.Syntax, PrefixRepeated) {
		out = HandleRepeated(namespace, in.Comment, split)
	} else if strings.HasPrefix(in.Syntax, PrefixMap) {
		out = HandleMap(namespace, in.Comment, split)
	} else {
		out = HandleDefaultAttribute(namespace, in.Comment, split)
	}

	return out
}
