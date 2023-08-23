package repositories_test

import (
	"testing"
	"time"

	"FC-3.0/Encoder-new/application/repositories"
	"FC-3.0/Encoder-new/domain"
	"FC-3.0/Encoder-new/framework/database"
	"github.com/docker/distribution/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.Generate().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}
