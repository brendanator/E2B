// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.10.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xYbW/bNhD+K8RtH9XIfRlQ+FvTZJuBJQtmoF8KY6Cls81MIlXy5Cww9N+Ho2SZevHs",
	"oGnaL45CkffyPA+PR+0gMXlhNGpyMN2BSzaYS/94rbfKGp2jJv5XZtmfK5h+3sHPFlcwhZ/iw9K4WRff",
	"4kO4rop2UFhToCWF3qpK+ZceC4QpOLJKr6GKwJGk0k9AXeYw/QyXpcpSfhvBr1JlmEIEV0YjLKL+8ioC",
	"i19KZTHllYqnNgYPk83yHhOCalFFYWrz1nE3zucM6GgsEVxba+zQe2JS5L8pusSqgpTRMK0nC/8ugpWx",
	"uSSYgtL09g20tpUmXKNl4zk6J9fHDMGpsBtHeysMXI/e0bjnWhUF0uxqlOgUi1oHhLkbndEMSGvlo/8f",
	"8yKThCEXtybFe3ca+W48ga0mkDFKbvFhjs55pE7m18V1lqImtVJohVkJ6akSrp4vHjYq2TS/ygnaoMAD",
	"lkI6ZxIlyevqKVkxL8cDzhTqM2LlaOq5Q+8RuNr8eWaaySezOBiNDmEuKp6m9MoMXV3hdmnMP+LD3YyN",
	"K8rwMAoRbNHWIMDri8nFhAM3BWpZKJjCWz8UQSFp45GJ+WeNNPTzO8qMNiLZYMJmGU7Jr2YpTOE3ZIQs",
	"usJoV2P8ZjIZGvkLv5ToSDxIJ1yZJOjcqsyg8vnFqLd1zTHOB9D1cWccXfOMcUeJ0bQvykWRqcSvjO9d",
	"LYC6FPPTUwp15fdmmMG8jTp7FIlFlqaQOhSt55Tk2jGfPqdFm16868i0CuAeIMrJfuxt1UJamSOhdf7Q",
	"URwSkwcRaJkz8/3NfRAX2RKjAIi+EBffENgeqk9AKD6cOf+viw5UzQH2AoB5QV+a9PFbYNWkUXWrBMdV",
	"dcjSZZaNItqUE3d0W/+hHAmZZaKdObK354d3X6WQ9ng7r23a12/fnXQPwBMb0yKVVvPODDOrIvjlOTXt",
	"+4WRUOqOJFOOlF4H7gPRt4Oc2l7XXSsffXERcm9AGN2cJHaLdkATb4MOT88vzKAHOKlIFsfrZ/PccXtO",
	"QW7P2yqCdy/B+aVMRYO5WJr0UTSd6MtqzucfiO6Y5sLSEO/anqOqVZgh4VijweNn67GevlfkPGhrTpfk",
	"sAn6mvPr3TCLjlzqVNNOg/ayhPkImLBOCE8jLba4sug2/ogsaazr8u9DHwL/JdR8cROKnCCVoyAjMrXF",
	"YWUpaUBjY/KHYrOB4bvy2cRwLqP+FsHbpwava/HD3UygTguj/P2jtBlMYUNUuGkcG53WLf5FYvLYf1Lo",
	"QbMnurEgHhRtTEnCFZjw1SQVhbFDw4GwXu32N5Dqoucvgq20Si6zsQvVSpYZ38H/bof7YM2ueter4XWq",
	"NRLK55gVF5wPZwGxRyE5CgOPVq++Gxo+qM7695P3k8HSO2OJN25itMaEH58RykUr2j6mN1LLNYv8003Y",
	"OHY3vBtRZdDb8gkpPt0clvm2dYQ/3wPvw7s3S7Fsvjn1vxZ0TOFWZo5bx/8CAAD//yl3BnXNEwAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
