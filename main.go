package main

import (
	"fmt"
	"os"

	"github.com/google/logger"
	"go.opentelemetry.io/collector/exporter/debugexporter"
)

func main() {
	f := debugexporter.NewFactory()
	fmt.Println("Created factory:", f)
	logger.Init("DebugExporter", true, false, os.Stderr)
	logger.Info("Debug exporter factory created successfully")
}
