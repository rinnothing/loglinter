package analyzer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/rinnothing/loglinter/pkg/analyzer"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)

	dir := filepath.Join(wd, "testdata")
	analysistest.Run(t, dir, analyzer.Analyzer, "p")
}
