package main

import "cloud.google.com/go/compute/metadata"

func GetProjectID() (string, error) {
	if !metadata.OnGCE() {
		// TODO 環境変数から取得する
		return "default value", nil
	}
	projectID, err := metadata.ProjectID()
	if err != nil {
		return "", err
	}
	return projectID, nil
}
