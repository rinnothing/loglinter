package loglinter

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/rinnothing/loglinter/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglinter", New)
}

type Settings struct {
	SensitiveData []string `json:"sensitive_data"`
}

type PluginLoglinter struct {
	settings Settings
}

func New(settings any) (register.LinterPlugin, error) {
	s, err := register.DecodeSettings[Settings](settings)
	if err != nil {
		return nil, err
	}

	return &PluginLoglinter{settings: s}, nil
}

func (f *PluginLoglinter) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		analyzer.New(f.settings.SensitiveData),
	}, nil
}

func (f *PluginLoglinter) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
