package domain_test

import (
	"testing"
	"time"

	"FC-3.0/Encoder-new/domain"
	"github.com/docker/distribution/uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNotAUuid(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "abc"
	video.ResourceID = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.Generate().String()
	video.ResourceID = "a"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}
