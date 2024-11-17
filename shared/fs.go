package shared

import (
	"ac/domain/entities"
	"bufio"
	"embed"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func ReadFileWithFileInfo(filePath string, fs *embed.FS) (*entities.FileInfo, string, error) {
	file, err := fs.Open(filePath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	var jsonText string
	var contentBuilder strings.Builder
	var insideCommentBlock bool

	scanner := bufio.NewScanner(file)
	jsonRegex := regexp.MustCompile(`(?s)\/\*.*?\*\/`)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "/*") {
			insideCommentBlock = true
		}

		if !insideCommentBlock {
			if !strings.HasPrefix(strings.TrimSpace(line), "//") {
				contentBuilder.WriteString(line + "\n")
			}
		} else {
			jsonText += line + "\n"
		}

		if strings.Contains(line, "*/") {
			insideCommentBlock = false
		}
	}

	jsonText = jsonRegex.FindString(jsonText)
	if jsonText == "" {
		return nil, "", fmt.Errorf("no JSON found in comments")
	}

	jsonText = strings.TrimPrefix(jsonText, "/*")
	jsonText = strings.TrimSuffix(jsonText, "*/")

	var fileInfo entities.FileInfo
	err = json.Unmarshal([]byte(jsonText), &fileInfo)
	if err != nil {
		return nil, "", err
	}

	noCommentContent := strings.TrimSpace(contentBuilder.String())

	noCommentContent = noCommentContent + "\n"

	return &fileInfo, noCommentContent, nil
}
