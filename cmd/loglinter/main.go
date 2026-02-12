package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/rinnothing/loglinter/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
