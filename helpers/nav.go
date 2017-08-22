package helpers

import "fmt"

type NavItem struct {
	Title   string
	Current bool
	Link    string
}

type CardListItem struct {
	Value string
}

type Card struct {
	Title       string
	Description string
	Href        string
	ListItem    []CardListItem
}

func MakeData(api *OpsManAPI) (string, interface{}) {
	var data struct {
		Title    string
		Template string
		Nav      struct {
			Title    string
			NavItems []NavItem
		}
		Card []Card
	}

	if report, err := api.GetDiagnosticReport(); err == nil {

		// START

		data.Title = report.Type
		/*
			data.Nav.Title = report.Type
			data.Nav.NavItems = append(data.Nav.NavItems, NavItem{
				Title:   "First",
				Current: true,
				Link:    "#foo",
			})
		*/
		// END

		card := Card{
			Title:       "Releases",
			Description: fmt.Sprintf("Count: %d", len(report.Releases)),
			Href:        "releases",
		}

		data.Card = append(data.Card, card)

	}

	return "data", data
}
