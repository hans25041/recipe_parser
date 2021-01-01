package recipe_parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Recipe struct {
	Title           string
	Description     string
	IngredientList  []string
	InstructionList []string
}

type HostNameNotMinimalistBakerError string

func (h HostNameNotMinimalistBakerError) Error() string {
	return fmt.Sprintf("Host name was not minimalistbaker.com: %q", string(h))
}

type RecipeNotFoundError string

func (r RecipeNotFoundError) Error() string {
	return fmt.Sprintf("No recipe coudl be found at URL: %q", string(r))
}

func GetRecipe(doc *goquery.Document) (Recipe, error) {
	title, err := GetRecipeTitle(doc)
	if err != nil {
		return Recipe{}, err
	}
	description, err := GetRecipeDescription(doc)
	if err != nil {
		return Recipe{}, err
	}
	ingredientList, err := GetRecipeIngredientList(doc)
	if err != nil {
		return Recipe{}, err
	}
	instructionList, err := GetRecipeInstructionList(doc)
	if err != nil {
		return Recipe{}, err
	}
	recipe := Recipe{title, description, ingredientList, instructionList}
	return recipe, nil
}

func GetRecipeInstructionList(doc *goquery.Document) ([]string, error) {
	return getTextSliceOfAllElementsInSelector(doc, "li.wprm-recipe-instruction")
}

func GetRecipeIngredientList(doc *goquery.Document) ([]string, error) {
	return getTextSliceOfAllElementsInSelector(doc, "li.wprm-recipe-ingredient")
}

func getTextSliceOfAllElementsInSelector(doc *goquery.Document, selector string) ([]string, error) {
	selection, err := getAllSelectionOfSelector(doc, selector)
	if err != nil {
		return nil, err
	}
	var texts []string
	selection.Each(func(_ int, s *goquery.Selection) {
		texts = append(texts, s.Text())
	})
	return texts, nil
}

func getAllSelectionOfSelector(doc *goquery.Document, selector string) (*goquery.Selection, error) {
	selection := doc.Find(selector)
	return selection, nil
}

func GetRecipeDescription(doc *goquery.Document) (string, error) {
	return getTextOfSelector(doc, "div.wprm-recipe-summary")
}

func GetRecipeTitle(doc *goquery.Document) (string, error) {
	return getTextOfSelector(doc, "h2.wprm-recipe-name")
}

func getTextOfSelector(doc *goquery.Document, selector string) (string, error) {
	text := doc.Find(selector).First().Text()
	return text, nil
}

func GetRecipePage(rawUrl string) (*goquery.Document, error) {
	_, err := ValidUrl(rawUrl)
	if err != nil {
		return nil, err
	}

	res, err := sendRequest(rawUrl)
	if err != nil {
		return nil, err
	}

	doc, err := getDocument(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func getDocument(body io.ReadCloser) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func sendRequest(rawUrl string) (*http.Response, error) {
	res, err := http.Get(rawUrl)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == 404 {
		return nil, RecipeNotFoundError(rawUrl)
	}
	return res, nil
}

func ValidUrl(rawUrl string) (bool, error) {
	recipeUrl, _ := url.Parse(rawUrl)
	hostname := getHostName(recipeUrl)
	if !isMinimalistBaker(hostname) {
		return false, HostNameNotMinimalistBakerError(hostname)
	}
	return true, nil
}

func isMinimalistBaker(hostname string) bool {
	return hostname == "minimalistbaker.com"
}

func getHostName(recipeUrl *url.URL) string {
	hostnameParts := strings.Split(recipeUrl.Hostname(), ".")
	return strings.Join(hostnameParts[len(hostnameParts)-2:], ".")
}
