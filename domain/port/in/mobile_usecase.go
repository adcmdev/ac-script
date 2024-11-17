package in

import "ac/domain/entities"

type IMobileUsecase interface {
	DeployLocal(releaseType entities.TulReleaseType) error
}
