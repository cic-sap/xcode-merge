package pkg

import (
	"errors"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
)

var ErrorFileName = errors.New("project file must be xcodeproj file or project.pbxproj ")

func StartMerge(project string) error {

	if !strings.HasSuffix(project, ".xcodeproj") {
		return ErrorFileName
	}

	projectPath, err := filepath.Abs(project)
	if err != nil {
		log.Error().Err(err).Msgf("get abspath error:%s", project)
		return err
	}
	f := filepath.Join(projectPath, "project.xcworkspace/xcshareddata/swiftpm/Package.resolved")
	if _, err = os.Stat(f); err == nil {
		log.Info().Msgf("check Package.resolved: %s", f)
		err = DoMergePackage(f)
		if err != nil {
			log.Error().Err(err).Msgf("DoMergePackage error:%s", project)
			return err
		}
	}

	f = filepath.Join(projectPath, "project.pbxproj")

	err = DoMergePBXPROJ(f)
	if err != nil {
		log.Error().Err(err).Msgf("DoMergePBXPROJ error:%s", project)
		return err
	}
	return nil
}

// merge "project.pbxproj"

func DoMergePBXPROJ(projectPath string) error {

	result, err := ParseGitConflicts(projectPath)
	if err != nil {
		log.Error().Err(err).Msgf("ParseGitConflicts error:%s", projectPath)
		return err
	}
	if !result.IsConflicts {
		log.Info().Msgf("%s has no conflict", projectPath)
		return nil
	}

	return DoMerge(projectPath, result)

}

func DoMerge(filename string, result *ConflictsResult) error {
	basePath := filepath.Base(filepath.Base(filename))
	println("base", basePath)
	return nil
}
