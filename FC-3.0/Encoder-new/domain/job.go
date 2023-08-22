package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/docker/distribution/uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)

}

type Job struct {
	ID               string    `json:"job_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OutuptBucketPath string    `json:"output_bucket_path" valid:"notnull"`
	Status           string    `json:"status" valid:"notnull"`
	Video            *Video    `json:"video" valid:"-"`
	VideoID          string    `json:"-" valid:"-" gorm:"column:video_id;type:uuid;notnull"`
	Error            string    `valid:"-"`
	CreatedAt        time.Time `json:"createdAt" valid:"-"`
	UpdatedAt        time.Time `json:"updatedAt" valid:"-"`
}

func NewJob(output string, status string, video *Video) (*Job, error) {
	job := Job{
		OutuptBucketPath: output,
		Status:           status,
		Video:            video,
	}

	job.prepare()

	err := job.Validate()

	if err != nil {
		return nil, err
	}

	return &job, nil

}

func (job *Job) prepare() {

	job.ID = uuid.Generate().String()
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()

}

func (job *Job) Validate() error {
	_, err := govalidator.ValidateStruct(job)

	if err != nil {
		return err
	}

	return nil
}
