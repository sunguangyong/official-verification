package convert

import (
	"net/http"

	"google.golang.org/grpc/metadata"
)

func HeaderToMetadata(header http.Header) metadata.MD {
	md := metadata.MD{}
	for key, values := range header {
		for _, value := range values {
			md.Append(key, value)
		}
	}
	return md
}
