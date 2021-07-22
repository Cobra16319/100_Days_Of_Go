package nhlApi

import "encoding/json"

type Team struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Link  string `json:"link"`
		Venue struct {
			Name     string `json:"name"`
			Link     string `json:"link"`
			City     string `json:"city"`
			TimeZone struct {
				ID     string `json:"id"`
				Offset int    `json:"offset"`
				Tz     string `json:"tz"`
			} `json:"timeZone"`
		} `json:"venue,omitempty"`
		Abbreviation    string `json:"abbreviation"`
		TeamName        string `json:"teamName"`
		LocationName    string `json:"locationName"`
		FirstYearOfPlay string `json:"firstYearOfPlay"`
		Division        struct {
			ID           int    `json:"id"`
			Name         string `json:"name"`
			NameShort    string `json:"nameShort"`
			Link         string `json:"link"`
			Abbreviation string `json:"abbreviation"`
		} `json:"division"`
		Conference struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"conference"`
		Franchise struct {
			FranchiseID int    `json:"franchiseId"`
			TeamName    string `json:"teamName"`
			Link        string `json:"link"`
		} `json:"franchise"`
		ShortName       string `json:"shortName,omitempty"`
		OfficialSiteURL string `json:"officialSiteUrl"`
		FranchiseID     int    `json:"franchiseId"`
		Active          bool   `json:"active"`

}

type nhlTeamsRespone struct { 
     Teams []Team `json:"teams"`
} 

func GetAllTeams() ([]Team, error) { 
    res, err := http.Get(fmt.Sprintf(format: "%s/teams", baseURL))
    if err ≠ nil {
		return nil, err
	}
	defer res.Body.Close()

	var response nhlTeamsRespone
	err = json.NewDecoder(res.Body).Decode(&response)

	return response.Teams, err
}














