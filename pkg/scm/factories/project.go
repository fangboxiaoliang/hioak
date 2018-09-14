package factories

import (
	"fmt"
	"github.com/hidevopsio/hiboot/pkg/log"
	"errors"
	"github.com/hidevopsio/hioak/pkg/scm"
	"github.com/hidevopsio/hioak/pkg/scm/gitlab"
)

func (s *ScmFactory) NewProject(provider int) (scm.ProjectInterface, error)  {
	log.Debug("scm.NewSession()")
	switch provider {
	case GithubScmType:
		return nil, errors.New(fmt.Sprintf("SCM of type %d not implemented\n", provider))
	case GitlabScmType:
		return new(gitlab.Project), nil
	default:
		return nil, errors.New(fmt.Sprintf("SCM of type %d not recognized\n", provider))
	}
	return nil, nil
}