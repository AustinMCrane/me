package main

// JSONResume is a json format for a resume
// schema: https://jsonresume.org/schema/
// generated from https://mholt.github.io/json-to-go/
type JSONResume struct {
	Basics struct {
		Name     string `json:"name"`
		Label    string `json:"label"`
		Image    string `json:"image"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		URL      string `json:"url"`
		Summary  string `json:"summary"`
		Location struct {
			Address     string `json:"address"`
			PostalCode  string `json:"postalCode"`
			City        string `json:"city"`
			CountryCode string `json:"countryCode"`
			Region      string `json:"region"`
		} `json:"location"`
		Profiles []struct {
			Network  string `json:"network"`
			Username string `json:"username"`
			URL      string `json:"url"`
		} `json:"profiles"`
	} `json:"basics"`
	Work []struct {
		Name       string   `json:"name"`
		Position   string   `json:"position"`
		URL        string   `json:"url"`
		StartDate  string   `json:"startDate"`
		EndDate    string   `json:"endDate"`
		Summary    string   `json:"summary"`
		Highlights []string `json:"highlights"`
	} `json:"work"`
	Volunteer []struct {
		Organization string   `json:"organization"`
		Position     string   `json:"position"`
		URL          string   `json:"url"`
		StartDate    string   `json:"startDate"`
		EndDate      string   `json:"endDate"`
		Summary      string   `json:"summary"`
		Highlights   []string `json:"highlights"`
	} `json:"volunteer"`
	Education []struct {
		Institution string   `json:"institution"`
		URL         string   `json:"url"`
		Area        string   `json:"area"`
		StudyType   string   `json:"studyType"`
		StartDate   string   `json:"startDate"`
		EndDate     string   `json:"endDate"`
		Score       string   `json:"score"`
		Courses     []string `json:"courses"`
	} `json:"education"`
	Awards []struct {
		Title   string `json:"title"`
		Date    string `json:"date"`
		Awarder string `json:"awarder"`
		Summary string `json:"summary"`
	} `json:"awards"`
	Certificates []struct {
		Name   string `json:"name"`
		Date   string `json:"date"`
		Issuer string `json:"issuer"`
		URL    string `json:"url"`
	} `json:"certificates"`
	Publications []struct {
		Name        string `json:"name"`
		Publisher   string `json:"publisher"`
		ReleaseDate string `json:"releaseDate"`
		URL         string `json:"url"`
		Summary     string `json:"summary"`
	} `json:"publications"`
	Skills []struct {
		Name     string   `json:"name"`
		Level    string   `json:"level"`
		Keywords []string `json:"keywords"`
	} `json:"skills"`
	Languages []struct {
		Language string `json:"language"`
		Fluency  string `json:"fluency"`
	} `json:"languages"`
	Interests []struct {
		Name     string   `json:"name"`
		Keywords []string `json:"keywords"`
	} `json:"interests"`
	References []struct {
		Name      string `json:"name"`
		Reference string `json:"reference"`
	} `json:"references"`
	Projects []struct {
		Name        string   `json:"name"`
		StartDate   string   `json:"startDate"`
		EndDate     string   `json:"endDate"`
		Description string   `json:"description"`
		Highlights  []string `json:"highlights"`
		URL         string   `json:"url"`
	} `json:"projects"`
}
