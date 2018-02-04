package views

import (
	"github.com/hjkelly/budget-category-service/categories"
	"github.com/hjkelly/budget-category-service/common"
	"net/http"
)

func ListCategories(w http.ResponseWriter, r *http.Request) {
	result, err := categories.List()
	if err != nil {
		sendServerError(w, err, "couldn't list categories using the controller")
		return
	}

	sendDataAsJSON(w, result, 200)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	// Parse category input via JSON.
	catInput := new(UserCategoryInput)
	err := getRequestBody(r, catInput)
	if err != nil {
		sendParseError(w)
		return
	}

	// Per that input struct, is the provided data enough/valid?
	if errors := catInput.ValidationErrors(); errors != nil {
		sendDataAsJSON(w, errors, 422)
		return
	}

	tokenSub, err := common.GetAuth().GetSub(r)
	if err != nil {
		sendServerError(w, err, "couldn't get sub claim from token") // TODO: parse error?
		return
	}

	// If all has gone well, convert it to a category proper and create it.
	result, err := categories.Create(categories.NewCategory(catInput.Name, catInput.Type, tokenSub))
	if err != nil {
		sendServerError(w, err, "couldn't create the category using its controller")
		return
	}

	sendDataAsJSON(w, result, 201)
}
