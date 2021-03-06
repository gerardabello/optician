package store

import (
	"errors"
	"image"

	"github.com/theopticians/optician-api/core/structs"
)

var NotFoundError = errors.New("Not found in DB")

type Store interface {
	Close()
	GetResults() ([]structs.Result, error)
	GetResultsByBatch(string) ([]structs.Result, error)
	GetResult(string) (structs.Result, error)
	GetLastResult(projectID, branch, target, browser string) (structs.Result, error)
	StoreResult(structs.Result) error

	GetBatchs() ([]structs.BatchInfo, error)

	GetMask(string) (structs.Mask, error)
	StoreMask(masks structs.Mask) (string, error)

	GetImage(string) (image.Image, error)
	StoreImage(image.Image) (string, error)

	GetBaseImageID(projectID, branch, target, browser string) (string, error)
	SetBaseImageID(baseImageID, projectID, branch, target, browser string) error

	GetBaseMaskID(projectID, branch, target, browser string) (string, error)
	SetBaseMaskID(baseImageID, projectID, branch, target, browser string) error
}
