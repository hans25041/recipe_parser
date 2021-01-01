package recipe_parser

import (
	"strings"
	"testing"
)

func assertStringsEqual(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected: %q Got: %q", expected, actual)
	}
}

func assertStringSlicesEqual(t testing.TB, actual, expected []string) {
	t.Helper()
	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("Expected %q to be in %v", v, actual)
		}
	}

	for i, v := range actual {
		if expected[i] != v {
			t.Errorf("Expected %q not to be in %v", v, actual)
		}
	}
}

func assertTrue(t testing.TB, actual bool) {
	t.Helper()
	if !actual {
		t.Errorf("Expected: true, got: %v", actual)
	}

}

func assertStringContains(t testing.TB, s, substr string) {
	t.Helper()
	if !strings.Contains(s, substr) {
		t.Errorf("Expected string: %.10q... to contain %.10q", s, substr)
	}
}

func assertStringSliceContains(t testing.TB, slice []string, str string) {
	t.Helper()
	for _, a := range slice {
		if a == str {
			return
		}
	}
	t.Errorf("Expected slice: %v to contain string: %q", slice, str)
}

func assertNotMinimalistBakerError(t testing.TB, actual error, expected HostNameNotMinimalistBakerError) {
	t.Helper()
	if actual != expected {
		t.Errorf("Got %q but expected %q", actual, expected)
	}
}

func assertRecipeNotFoundError(t testing.TB, actual error, expected RecipeNotFoundError) {
	t.Helper()
	if actual != expected {
		t.Errorf("Got %q but expected %q", actual, expected)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Unexpected error returned: %v", err)
	}
}

var invalidRecipeUrl = "https://www.epicurious.com/some/recipe/"
var epicuriousHostName = "epicurious.com"
var validRecipeUrl = "https://minimalistbaker.com/easy-vegan-fried-rice/"
var easyVeganFriedRiceUrl = validRecipeUrl
var howToMakeChiaPuddingUrl = "https://minimalistbaker.com/how-to-make-chia-pudding/"
var doesNotExistUrl = "https://minimalistbaker.com/not-here/"

func TestMinimalistBakerGetRecipe(t *testing.T) {
	t.Run("Test get recipe for Easy Vegan Fried Rice.", func(t *testing.T) {
		page, err := GetRecipePage(easyVeganFriedRiceUrl)
		assertNoError(t, err)
		title, err := GetRecipeTitle(page)
		assertNoError(t, err)
		description, err := GetRecipeDescription(page)
		assertNoError(t, err)
		ingredientList, err := GetRecipeIngredientList(page)
		assertNoError(t, err)
		instructionList, err := GetRecipeInstructionList(page)
		assertNoError(t, err)

		recipe, err := GetRecipe(page)
		assertNoError(t, err)
		assertStringsEqual(t, recipe.Title, title)
		assertStringsEqual(t, recipe.Description, description)
		assertStringSlicesEqual(t, recipe.IngredientList, ingredientList)
		assertStringSlicesEqual(t, recipe.InstructionList, instructionList)
	})

	t.Run("Test get recipe for How to Make Chia Pudding.", func(t *testing.T) {
		page, err := GetRecipePage(howToMakeChiaPuddingUrl)
		assertNoError(t, err)
		title, err := GetRecipeTitle(page)
		assertNoError(t, err)
		description, err := GetRecipeDescription(page)
		assertNoError(t, err)
		ingredientList, err := GetRecipeIngredientList(page)
		assertNoError(t, err)
		instructionList, err := GetRecipeInstructionList(page)
		assertNoError(t, err)

		recipe, err := GetRecipe(page)
		assertNoError(t, err)
		assertStringsEqual(t, recipe.Title, title)
		assertStringsEqual(t, recipe.Description, description)
		assertStringSlicesEqual(t, recipe.IngredientList, ingredientList)
		assertStringSlicesEqual(t, recipe.InstructionList, instructionList)
	})
}

func TestMinimalistBakerGetRecipeInstructionList(t *testing.T) {
	t.Run("Test get instruction list for Easy Vegan Fried Rice.", func(t *testing.T) {
		page, err := GetRecipePage(easyVeganFriedRiceUrl)
		assertNoError(t, err)
		instructionList, err := GetRecipeInstructionList(page)
		assertNoError(t, err)
		assertStringSliceContains(t, instructionList, "Preheat oven to 400 degrees F (204 C) and line a baking sheet with parchment paper (or lightly grease with non-stick spray).")
		assertStringSliceContains(t, instructionList, "Serve immediately with extra chili garlic sauce or sriracha for heat (optional). Crushed salted, roasted peanuts or cashews make a lovely additional garnish. Leftovers keep well in the refrigerator for 3-4 days, though best when fresh. Reheat in a skillet over medium heat or in the microwave.")
	})

	t.Run("Test get instruction list for How to Make Chia Pudding.", func(t *testing.T) {
		page, err := GetRecipePage(howToMakeChiaPuddingUrl)
		assertNoError(t, err)
		instructionList, err := GetRecipeInstructionList(page)
		assertNoError(t, err)
		assertStringSliceContains(t, instructionList, "To a mixing bowl add dairy free milk, chia seeds, maple syrup (to taste), and vanilla. Whisk to combine.")
		assertStringSliceContains(t, instructionList, "Enjoy as is, or layer with compote or fresh fruit! Will keep covered in the refrigerator up to 5 days.")
	})
}

func TestMinimalistBakerGetRecipeIngredientList(t *testing.T) {
	t.Run("Test get ingredient list for Easy Vegan Fried Rice.", func(t *testing.T) {
		page, err := GetRecipePage(easyVeganFriedRiceUrl)
		assertNoError(t, err)
		ingredientList, err := GetRecipeIngredientList(page)
		assertNoError(t, err)
		assertStringSliceContains(t, ingredientList, "1 cup extra-firm tofu* (8 ounces yields ~1 cup)")
		assertStringSliceContains(t, ingredientList, "1 cup long- or short-grain brown rice* (rinsed thoroughly in a fine mesh strainer)")
		assertStringSliceContains(t, ingredientList, "4 cloves garlic (minced)")
		assertStringSliceContains(t, ingredientList, "1 cup chopped green onion")
		assertStringSliceContains(t, ingredientList, "1/2 cup peas")
		assertStringSliceContains(t, ingredientList, "1/2 cup carrots (finely diced)")
		assertStringSliceContains(t, ingredientList, "3 Tbsp tamari or soy sauce (plus more for veggies + to taste)")
		assertStringSliceContains(t, ingredientList, "1 Tbsp peanut butter")
		assertStringSliceContains(t, ingredientList, "2-3 Tbsp organic brown sugar, muscovado sugar, or maple syrup")
		assertStringSliceContains(t, ingredientList, "1 clove garlic (minced)")
		assertStringSliceContains(t, ingredientList, "1-2 tsp chili garlic sauce  (more or less depending on preferred spice)")
		assertStringSliceContains(t, ingredientList, "1 tsp toasted sesame oil (optional // or sub peanut or avocado oil)")
	})

	t.Run("Test get ingredient list for How to Make Chia Pudding.", func(t *testing.T) {
		page, err := GetRecipePage(howToMakeChiaPuddingUrl)
		assertNoError(t, err)
		ingredientList, err := GetRecipeIngredientList(page)
		assertNoError(t, err)
		assertStringSliceContains(t, ingredientList, "1 1/2 cups dairy-free milk  (we used DIY coconut – use creamier milks for creamier, thicker pudding, such as full fat coconut and cashew)")
		assertStringSliceContains(t, ingredientList, "1/2 cup chia seeds")
		assertStringSliceContains(t, ingredientList, "1-2 Tbsp maple syrup  (more or less to taste)")
		assertStringSliceContains(t, ingredientList, "1 tsp vanilla extract")
		assertStringSliceContains(t, ingredientList, "Compote ")
		assertStringSliceContains(t, ingredientList, "Mint")
		assertStringSliceContains(t, ingredientList, "Fresh Fruit")
	})
}

func TestMinimalistBakerGetRecipeDescription(t *testing.T) {
	t.Run("Test get description for Easy Vegan Fried Rice.", func(t *testing.T) {
		page, err := GetRecipePage(easyVeganFriedRiceUrl)
		assertNoError(t, err)
		description, err := GetRecipeDescription(page)
		assertNoError(t, err)
		assertStringContains(t, description, "Easy, 10-ingredient vegan fried rice that’s loaded with vegetables")
	})

	t.Run("Test get description for How to Make Chia Pudding.", func(t *testing.T) {
		page, err := GetRecipePage(howToMakeChiaPuddingUrl)
		assertNoError(t, err)
		description, err := GetRecipeDescription(page)
		assertNoError(t, err)
		assertStringContains(t, description, "Creamy, thick chia pudding that’s easy to make, nutritious, and so")
	})
}

func TestMinimalistBakerGetRecipeTitle(t *testing.T) {
	t.Run("Test get title for Easy Vegan Fried Rice.", func(t *testing.T) {
		page, err := GetRecipePage(easyVeganFriedRiceUrl)
		assertNoError(t, err)
		title, err := GetRecipeTitle(page)
		assertNoError(t, err)
		expectedTitle := "Easy Vegan Fried Rice"
		assertStringsEqual(t, title, expectedTitle)
	})

	t.Run("Test get title for How to Make Chia Pudding.", func(t *testing.T) {
		page, err := GetRecipePage(howToMakeChiaPuddingUrl)
		assertNoError(t, err)
		title, err := GetRecipeTitle(page)
		assertNoError(t, err)
		expectedTitle := "How to Make Chia Pudding"
		assertStringsEqual(t, title, expectedTitle)
	})
}

func TestMinimalistBakerGetRecipePage(t *testing.T) {
	t.Run("Test parse a URL with the wrong hostname.", func(t *testing.T) {
		_, err := GetRecipePage(invalidRecipeUrl)
		expectedErr := HostNameNotMinimalistBakerError(epicuriousHostName)
		assertNotMinimalistBakerError(t, err, expectedErr)
	})

	t.Run("Test get HTML of recipe that does not exist.", func(t *testing.T) {
		_, err := GetRecipePage(doesNotExistUrl)
		expectedErr := RecipeNotFoundError(doesNotExistUrl)
		assertRecipeNotFoundError(t, err, expectedErr)

	})

	t.Run("Test get HTML for Easy Vegan Fried Rice recipe.", func(t *testing.T) {
		doc, err := GetRecipePage(validRecipeUrl)
		assertNoError(t, err)
		pageString := doc.Text()
		assertStringContains(t, pageString, "Easy Vegan Fried Rice")
	})
}

func TestMinimalistBakerUrlParser(t *testing.T) {
	t.Run("Test parse a URL with the wrong hostname.", func(t *testing.T) {
		_, err := ValidUrl("https://www.epicurious.com/some/recipe/")
		expectedErr := HostNameNotMinimalistBakerError(epicuriousHostName)
		assertNotMinimalistBakerError(t, err, expectedErr)
	})

	t.Run("Test parse a URL with the correct hostname.", func(t *testing.T) {
		valid, err := ValidUrl(validRecipeUrl)
		assertNoError(t, err)
		assertTrue(t, valid)
	})
}
