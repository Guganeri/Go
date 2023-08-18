package domain_test

import (
	"testing"
	"time"

	"FC-3.0/Encoder-new/domain"
	"github.com/docker/distribution/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewJOb(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.Generate().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path", "Converted", video)
	require.NotNil(t, job)
	require.Nil(t, err)
}
