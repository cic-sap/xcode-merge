package pkg

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
)

// Package.resolved

type PackageResolvedItem struct {
	Package       string `json:"package"`
	RepositoryURL string `json:"repositoryURL"`
	State         struct {
		Branch   interface{} `json:"branch"`
		Revision string      `json:"revision"`
		Version  string      `json:"version"`
	} `json:"state"`
}

type PackageResolved struct {
	Object struct {
		Pins []PackageResolvedItem `json:"pins"`
	} `json:"object"`
	Version int `json:"version"`
}

func DoMergePackage(packageFile string) error {

	result, err := ParseGitConflicts(packageFile)
	if err != nil {
		return err
	}
	log.Debug().Msgf("file:IsConflicts:%v", result.IsConflicts)
	if !result.IsConflicts {
		_ = MarkingConflictResolved(packageFile)
		return nil
	}

	var data1, data2 PackageResolved

	err = json.Unmarshal([]byte(result.ConflictFileA.Data), &data1)
	if err != nil {
		log.Error().Err(err).Msgf("load PackageResolved error")
		return err
	}

	err = json.Unmarshal([]byte(result.ConflictFileB.Data), &data2)
	if err != nil {
		log.Error().Err(err).Msgf("load PackageResolved error")
		return err
	}

	log.Debug().Interface("a", data1).Interface("b", data2).Msg(string(result.ConflictFileA.Data))
	packages := make(map[string]PackageResolvedItem)
	for _, v := range data1.Object.Pins {
		packages[v.Package] = v
	}
	for _, v := range data2.Object.Pins {
		if _, ok := packages[v.Package]; !ok {
			packages[v.Package] = v
			data1.Object.Pins = append(data1.Object.Pins, v)
		}
	}
	buf, err := json.MarshalIndent(data1, "", "  ")
	if err != nil {
		log.Error().Err(err).Msgf("json MarshalIndent error")
		return err
	}

	log.Info().Msgf("write PackageResolved:%s", packageFile)

	err = ioutil.WriteFile(packageFile, buf, 0644)
	if err != nil {
		log.Error().Err(err).Msgf("write file error:%s", packageFile)
		return err
	}

	// Marking conflict resolved
	err = MarkingConflictResolved(packageFile)
	if err != nil {
		return err
	}
	return nil
}
