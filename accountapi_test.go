package accountapi

import (
	"errors"
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	// expected := `{
	// 	"data": {
	// 		"attributes": {
	// 			"account_classification": "Personal",
	// 			"account_matching_opt_out": false,
	// 			"account_number": "41426819",
	// 			"alternative_names": null,
	// 			"bank_id": "400300",
	// 			"bank_id_code": "GBDSC",
	// 			"base_currency": "GBP",
	// 			"bic": "NWBKGB22",
	// 			"country": "GB",
	// 			"status": "confirmed",
	// 			"name": ["Abrar"],
	// 			"switched": false,
	// 		},
	// 		"created_on": "2022-03-10T17:47:13.244Z",
	// 		"id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
	// 		"modified_on": "2022-03-10T17:47:13.244Z",
	// 		"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	// 		"type": "accounts",
	// 		"version": 0,
	// 	},
	// 	"links": {
	// 		"self": "/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
	// 	},
	// }`
	cases := []struct {
		in   string
		want string
		err  error
	}{
		{"", "empty id", errors.New("empty id")},
		{"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", "NWBKGB22", nil},
	}

	for _, c := range cases {
		got, err := Fetch(c.in)
		//s := fmt.Sprintf("%#v", got) stringify
		if err != nil {
			gotErr := fmt.Sprintf("%#v", err)
			wantErr := fmt.Sprintf("%#v", c.err)
			if gotErr != wantErr {
				t.Errorf("got %q, want %q", err, c.err)
			}
		} else {
			if got.Data.Attributes.Bic != c.want {
				t.Errorf("fetch(%q) == %q, want %q", c.in, got.Data.Attributes.Bic, c.want)
			}
		}

	}
}

// func TestCreate(t *testing.T) {
// 	expected := AccountData{
// 		  ID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
// 		  OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
// 		  Attributes: {
// 			AccountClassification: "Personal",
// 			AccountMatchingOptOut: false,
// 			AccountNumber: "41426819 ",
// 			Name: ["Abrar"],
// 			*Country: "GB",
// 			BaseCurrency: "GBP",
// 			BankID: "400300",
// 			BankIDCode: "GBDSC",
// 			Bic: "NWBKGB22",
// 			Switched: false,
// 			Status: "confirmed",
// 		  },
// 	  }	  
// }
