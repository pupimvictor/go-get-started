package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetItemsHandler(t *testing.T) {

	testCases := []struct {
		name               string
		expectedStatusCode int
		initApp            bool
		items              []Item
	}{
		{
			"list has items should pass",
			200,
			true,
			[]Item{{"apple", 2}},
		},
		{
			"empty list should error",
			400,
			true,
			nil,
		},
		{
			"List not initialized should error",
			500,
			false,
			nil,
		},
	}

	for _, tc := range testCases {

		// Set up
		if tc.initApp {
			NewShoppingList("test")

			if tc.items != nil {
				myShoppingList.items = tc.items
			}
		}

		// Test Run
		t.Run(tc.name, func(tt *testing.T) {
			// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
			// pass 'nil' as the third parameter.
			req, err := http.NewRequest("GET", "/item/get", nil)
			if err != nil {
				tt.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetItemsHandler)

			// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			handler.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if status := rr.Code; status != tc.expectedStatusCode {
				tt.Errorf("handler returned wrong status code: got %v want %v",
					status, tc.expectedStatusCode)
			}

		})

		// Tear Down
		myShoppingList = nil
	}
}
