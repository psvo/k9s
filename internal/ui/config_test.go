package ui_test

import (
	"path/filepath"
	"testing"

	"github.com/derailed/k9s/internal/config"
	"github.com/derailed/k9s/internal/render"
	"github.com/derailed/k9s/internal/ui"
	"github.com/gdamore/tcell"
	"github.com/stretchr/testify/assert"
)

func TestBenchConfig(t *testing.T) {
	config.K9sHome = "/tmp/blee"
	assert.Equal(t, "/tmp/blee/bench-fred.yml", ui.BenchConfig("fred"))
}

func TestConfiguratorRefreshStyle(t *testing.T) {
	config.K9sStylesFile = filepath.Join("..", "config", "testdata", "black_and_wtf.yml")

	cfg := ui.Configurator{}
	cfg.RefreshStyles("")

	assert.True(t, cfg.HasSkin())
	assert.Equal(t, tcell.ColorGhostWhite, render.StdColor)
	assert.Equal(t, tcell.ColorWhiteSmoke, render.ErrColor)
}
