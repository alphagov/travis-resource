package model

type DefaultSource struct {
	Repository     string   `json:"repository"`
	Branch         string   `json:"branch"`
	GithubToken    string   `json:"github-token"`
	TravisToken    string   `json:"travis-token"`
	CheckAllBuilds bool     `json:"check-all-builds"`
}
type OutParams struct {
	Build      interface{}  `json:"build"`
	Branch     string       `json:"branch"`
	Repository string       `json:"repository"`
}
type Version struct {
	BuildNumber int `json:"build"`
}
type VersionString struct {
	BuildNumber string `json:"build"`
}
type Metadata struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}
type CheckRequest struct {
	Source  DefaultSource  `json:"source"`
	Version Version        `json:"version"`
}
type InRequest struct {
	Source  DefaultSource  `json:"source"`
	Version VersionString  `json:"version"`
}
type OutRequest struct {
	InRequest
	OutParams OutParams `json:"params"`
}
type CheckResponse []Version
type InResponse struct {
	Metadata []Metadata   `json:"metadata"`
	Version  Version      `json:"version"`
}

type OutResponse struct {
	InResponse
}