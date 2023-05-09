package reader

import (
	"github.com/GoogleCloudPlatform/proto-gen-parser/pkg/api"
	"github.com/GoogleCloudPlatform/proto-gen-parser/pkg/logging"
	"strings"
)

type AnnotationVisitor struct {
	Log *logging.Logger
}

func NewAnnotationVisitor(debug bool) *AnnotationVisitor {
	return &AnnotationVisitor{
		Log: logging.NewLogger(debug, "annotation_visitor"),
	}
}

func (v *AnnotationVisitor) CanVisit(in *Line) bool {
	return in.Token == OpenBracket
}

func splitValue(in string) api.Annotation {
	if strings.Contains(in, "=") && strings.Count(in, "=") == 1 {
		split := strings.Split(in, "=")
		return ProtobufFactory.NewAnnotation(strings.Trim(split[0], Space), strings.Trim(split[1], Space))
	}
	return nil
}

func (v *AnnotationVisitor) Visit(scanner Scanner, in *Line, namespace string) interface{} {

	annotationComment := ""
	annotationLine := ""

	for scanner.Scan() {
		line := scanner.ReadLine()

		if line.Token == ClosedBracket {
			break
		}
		annotationLine += line.Syntax + Space + line.Token
		annotationComment += Space + line.Comment
	}

	v.Log.Debugf("\nAnnotation Line: %s\n", annotationLine)

	out := make([]api.Annotation, 0)
	if strings.Contains(annotationLine, Comma) {
		for _, a := range strings.Split(annotationLine, Comma) {
			v := splitValue(a)
			if v != nil {
				out = append(out, v)
			}
		}
	} else {
		v := splitValue(annotationLine)
		if v != nil {
			out = append(out, v)
		}
	}
	return out //ProtobufFactory.NewAnnotation("", "")
}
