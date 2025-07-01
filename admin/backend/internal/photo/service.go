package photo

import (
	"errors"
	"strings"
)

func TransformGoogleDriveLink(originalLink string) (string, error) {
	if !strings.Contains(originalLink, "drive.google.com/file/d/") {
		return "", errors.New("invalid Google Drive link format")
	}

	parts := strings.Split(originalLink, "/")
	if len(parts) < 5 {
		return "", errors.New("failed to extract file ID from link")
	}

	fileID := parts[4]
	if fileID == "" {
		return "", errors.New("file ID is empty")
	}

	directLink := "https://drive.google.com/uc?id=" + fileID
	return directLink, nil
}
