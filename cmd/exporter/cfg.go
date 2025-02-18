package exporter

type ExportOptions struct {
	Dockerfile     string // docker filename
	OutputFilename string // output filename
	ExportPath     string // container export path
}
