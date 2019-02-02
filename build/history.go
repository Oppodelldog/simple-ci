package build

import (
	"fmt"
	"github.com/Oppodelldog/simpleci/config"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
)

func GetAllBuilds() map[string][]int {
	res := map[string][]int{}
	repoFolders, err := ioutil.ReadDir(config.BuildDir)
	if err != nil {
		logrus.Errorf("error reading build dir %s: %v", config.BuildDir, err)
		return nil
	}

	for _, repoFolder := range repoFolders {
		if !repoFolder.IsDir() {
			continue
		}

		res[repoFolder.Name()] = GetBuild(repoFolder.Name())
	}
	return res
}

func GetBuild(repoDir string) []int {
	logFiles, err := ioutil.ReadDir(getBuildLogPath(repoDir))
	if err != nil {
		logrus.Errorf("error reading build dir for repositoryURL '%s': %v", err)
		return nil
	}

	var logNumbers []int

	for _, file := range logFiles {
		stringNumber := file.Name()[:strings.LastIndex(file.Name(), ".")]
		number, err := strconv.Atoi(stringNumber)
		if err != nil {
			logrus.Error(err)
		}
		logNumbers = append(logNumbers, number)
	}

	return logNumbers
}

func GetBuildLog(repoDir string, logId string) (string, error) {
	filePath := path.Join(getBuildLogPath(repoDir), fmt.Sprintf("%s.txt", logId))

	b, err := ioutil.ReadFile(filePath)

	return string(b), err
}