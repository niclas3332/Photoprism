package config

import (
	"github.com/photoprism/photoprism/internal/thumb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_ConvertSize(t *testing.T) {
	c := NewTestConfig()
	assert.Equal(t, int(720), c.ConvertSize())
	c.params.ConvertSize = 31000
	assert.Equal(t, int(30000), c.ConvertSize())
	c.params.ConvertSize = 800
	assert.Equal(t, int(800), c.ConvertSize())
}

func TestConfig_JpegQuality(t *testing.T) {
	c := NewTestConfig()
	assert.Equal(t, int(25), c.JpegQuality())
	c.params.JpegQuality = 110
	assert.Equal(t, int(100), c.JpegQuality())
	c.params.JpegQuality = 98
	assert.Equal(t, int(98), c.JpegQuality())
}

func TestConfig_ThumbFilter(t *testing.T) {
	c := NewTestConfig()
	assert.Equal(t, thumb.ResampleFilter("cubic"), c.ThumbFilter())
	c.params.ThumbFilter = "blackman"
	assert.Equal(t, thumb.ResampleFilter("blackman"), c.ThumbFilter())
	c.params.ThumbFilter = "lanczos"
	assert.Equal(t, thumb.ResampleFilter("lanczos"), c.ThumbFilter())
	c.params.ThumbFilter = "linear"
	assert.Equal(t, thumb.ResampleFilter("linear"), c.ThumbFilter())
}

func TestConfig_ThumbSizeUncached(t *testing.T) {
	c := NewTestConfig()
	assert.False(t, c.ThumbUncached())
}

func TestConfig_ThumbSize(t *testing.T) {
	c := NewTestConfig()
	assert.Equal(t, int(720), c.ThumbSize())
	c.params.ThumbSize = 7681
	assert.Equal(t, int(7680), c.ThumbSize())
}

func TestConfig_ThumbSizeUncached2(t *testing.T) {
	c := NewTestConfig()
	assert.Equal(t, int(720), c.ThumbSizeUncached())
	c.params.ThumbSizeUncached = 7681
	assert.Equal(t, int(7680), c.ThumbSizeUncached())
	c.params.ThumbSizeUncached = 800
	c.params.ThumbSize = 900
	assert.Equal(t, int(900), c.ThumbSize())
}