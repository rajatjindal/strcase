package strcase

import (
	"github.com/ettle/strcase"
)

// TODO: maybe change this to sync.Map
var overrides = map[string]bool{}
var caser *strcase.Caser
var ocaser *strcase.Caser

func init() {
	caser = newCaser(overrides)
}

func newCaser(overrides map[string]bool) *strcase.Caser {
	var splitFn = strcase.NewSplitFn(
		[]rune{'*', '.', ',', '-', '_'},
		strcase.SplitCase,
		strcase.SplitAcronym,
		strcase.PreserveNumberFormatting,
		strcase.SplitBeforeNumber,
		strcase.SplitAfterNumber,
	)

	return strcase.NewCaser(false, overrides, splitFn)
}

// ToPascal returns words in PascalCase (capitalized words concatenated together).
func ToPascal(inp string) string {
	if ocaser != nil {
		return ocaser.ToPascal(inp)
	}

	return caser.ToPascal(inp)
}

// ToCamel returns words in camelCase (capitalized words concatenated together, with first word lower case).
func ToCamel(inp string) string {
	if ocaser != nil {
		return ocaser.ToCamel(inp)
	}

	return caser.ToCamel(inp)
}

// ToKebab returns words in kebab-case (lower case words with dashes).
func ToKebab(inp string) string {
	return caser.ToKebab(inp)
}

// ToScreamingSnake returns words in SNAKE_CASE (upper case words with underscores).
func ToScreamingSnake(inp string) string {
	return caser.ToSNAKE(inp)
}

// ConfigureAcronym configures the acronym override
// once ConfigureAcronym is called, the ToCamel and ToPascal
// functions will make use of these overrides
func ConfigureAcronym(acronym string) {
	overrides[acronym] = true
	ocaser = newCaser(overrides)
}

