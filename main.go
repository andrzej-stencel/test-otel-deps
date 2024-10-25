package main

import (
	"fmt"

	"go.opentelemetry.io/collector/exporter/debugexporter"
)

func main() {
	f := debugexporter.NewFactory()
	fmt.Println("Created factory:", f)
}
