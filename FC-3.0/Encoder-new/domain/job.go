package domain

import "time"

type Job struct {
	ID               string
	OutuptBucketPath string
	Status           string
	Video            *Video
	Error            string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
