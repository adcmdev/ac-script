package mobile

import (
	"ac/domain/entities"
	"ac/domain/port/in"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type UsecaseService struct{}

func NewUsecaseService() in.IMobileUsecase {
	return &UsecaseService{}
}

func (u *UsecaseService) DeployLocal(releaseType entities.TulReleaseType) error {
	pubspecFile := "pubspec.yaml"
	data, err := os.ReadFile(pubspecFile)
	if err != nil {
		return err
	}

	pubspec := string(data)
	oldVersion, found := getVersionFromYaml(pubspec)
	if !found {
		return fmt.Errorf("version not found in pubspec.yaml")
	}

	fullVersion := updateVersion(oldVersion, releaseType)
	version := strings.Split(fullVersion, "+")[0]
	fmt.Println(version)

	versionRegex := regexp.MustCompile(`version: \d+\.\d+\.\d+\+\d+`)
	pubspec = versionRegex.ReplaceAllString(pubspec, "version: "+fullVersion)

	updatedPubspec, err := u.updateVersionCommentsInPubspec(pubspec, version)
	if err != nil {
		return err
	}
	updatedPubspec = updateMsixVersionInPubspec(updatedPubspec, version)

	if err := u.updateIOSProjectVersion(fullVersion, version); err != nil {
		return err
	}

	err = os.WriteFile(pubspecFile, []byte(updatedPubspec), 0644)
	if err != nil {
		return err
	}

	return nil
}

func getVersionFromYaml(yaml string) (string, bool) {
	versionRegex := regexp.MustCompile(`version: (\d+\.\d+\.\d+\+\d+)`)
	matches := versionRegex.FindStringSubmatch(yaml)

	if len(matches) > 1 {
		return matches[1], true
	}

	return "", false
}

func updateVersion(currentVersion string, releaseType entities.TulReleaseType) string {
	versionParts := strings.Split(currentVersion, "+")
	versionNumbers := strings.Split(versionParts[0], ".")
	major, _ := strconv.Atoi(versionNumbers[0])
	minor, _ := strconv.Atoi(versionNumbers[1])
	bugfix, _ := strconv.Atoi(versionNumbers[2])

	switch releaseType {
	case entities.Bugfix:
		bugfix++
	case entities.Minor:
		minor++
		bugfix = 0
	case entities.Major:
		major++
		minor = 0
		bugfix = 0
	}

	buildNumber := major*1_000_000 + minor*10_000 + bugfix*10

	buildNumberStr := strconv.Itoa(buildNumber)

	newVersion := fmt.Sprintf("%d.%d.%d+%s", major, minor, bugfix, buildNumberStr)
	return newVersion
}

func (u *UsecaseService) updateIOSProjectVersion(fullVersion, version string) error {
	iosProjectFilePath := "ios/Runner.xcodeproj/project.pbxproj"
	data, err := os.ReadFile(iosProjectFilePath)
	if err != nil {
		return fmt.Errorf("failed to read iOS project file: %v", err)
	}

	projectContent := string(data)
	buildVersion := strings.Split(fullVersion, "+")[1]

	currentProjectVersionRegex := regexp.MustCompile(`CURRENT_PROJECT_VERSION = \d+;`)
	projectContent = currentProjectVersionRegex.ReplaceAllString(projectContent, fmt.Sprintf("CURRENT_PROJECT_VERSION = %s;", buildVersion))

	marketingVersionRegex := regexp.MustCompile(`MARKETING_VERSION = \d+\.\d+\.\d+;`)
	projectContent = marketingVersionRegex.ReplaceAllString(projectContent, fmt.Sprintf("MARKETING_VERSION = %s;", version))

	if err := os.WriteFile(iosProjectFilePath, []byte(projectContent), 0644); err != nil {
		return fmt.Errorf("failed to write updated iOS project file: %v", err)
	}

	return nil
}

func (u *UsecaseService) updateVersionCommentsInPubspec(pubspec, version string) (string, error) {
	versionParts := strings.Split(version, ".")
	if len(versionParts) < 3 {
		return "", fmt.Errorf("invalid version format: %s", version)
	}
	major, minor, patch := versionParts[0], versionParts[1], versionParts[2]

	commentPattern := regexp.MustCompile(`# \(\(\d+\*1000000\)\+\(\d+\*10000\)\+\(\d+\*10\)\)`)
	newComment := fmt.Sprintf("# ((%s*1000000)+(%s*10000)+(%s*10))", major, minor, patch)

	updatedPubspec := commentPattern.ReplaceAllString(pubspec, newComment)
	return updatedPubspec, nil
}

func updateMsixVersionInPubspec(pubspec, version string) string {
	msixVersionRegex := regexp.MustCompile(`(?m)^(\s+msix_version:\s*)\d+\.\d+\.\d+\.\d+`)
	return msixVersionRegex.ReplaceAllString(pubspec, "${1}"+version+".0")
}
