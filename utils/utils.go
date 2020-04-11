package utils

import (
	"errors"
	"github.com/digitalocean/godo"
	"github.com/manifoldco/promptui"
	"strconv"
	"strings"
)

// SelectItem is used for custom selects
type SelectItem struct {
	Name  string
	Value string
}

// CreateCustomSelectPrompt will create a customised select prompt to be used to ask the user a multi-select question
func CreateCustomSelectPrompt(title string, completeList []SelectItem) promptui.Select {
	selectList := completeList

	templates := &promptui.SelectTemplates{
		Label: "{{ . }}?",

		Active:   "\U0001F449{{ .Name | cyan }} ({{ .Value | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .Value | red }})",
		Selected: "\U0001F449{{ .Name | red | cyan }}",
	}

	searcher := func(input string, index int) bool {
		item := selectList[index]
		name := strings.Replace(strings.ToLower(item.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     title,
		Items:     selectList,
		Templates: templates,
		Size:      8,
		Searcher:  searcher,
	}

	return prompt
}

// ValidateAreYouSure will check whether they entered yes
func ValidateAreYouSure(input string) error {
	if input == "y" || input == "n" {
		return nil
	}
	return errors.New("Answer must be y/n")
}

// ValidateDropletName will check whether they entered yes
func ValidateDropletName(input string) error {
	if len(input) <= 0 {
		return errors.New("Must have a name")
	}
	if strings.Contains(input, " ") {
		return errors.New("Must not contain a space")
	}
	return nil
}

// ParseRegionListresults will return a list of DigitalOcean regions as SelectItems to be used for promptui
func ParseRegionListresults(list []godo.Region) []SelectItem {
	selectList := []SelectItem{}

	for _, element := range list {
		listItem := SelectItem{Name: element.Name, Value: element.Slug}
		selectList = append(selectList, listItem)
	}

	return selectList
}

// ParseImageListResults will return a list of DigitalOcean images as SelectItems to be used for promptui
func ParseImageListResults(list []godo.Image) []SelectItem {
	selectList := []SelectItem{}

	for _, element := range list {
		listItem := SelectItem{Name: element.Name, Value: element.Slug}
		selectList = append(selectList, listItem)
	}

	return selectList
}

// ParseSizeListResults will return a list of DigitalOcean sizes as SelectItems to be used for promptui
func ParseSizeListResults(list []godo.Size) []SelectItem {
	selectList := []SelectItem{}

	for _, element := range list {
		listItem := SelectItem{Name: element.Slug, Value: element.Slug}
		selectList = append(selectList, listItem)
	}

	return selectList
}

// ParseSSHKeyListResults will return a list of DigitalOcean ssh keys as SelectItems to be used for promptui
func ParseSSHKeyListResults(list []godo.Key) []SelectItem {
	selectList := []SelectItem{}

	for _, element := range list {
		id := strconv.Itoa(element.ID)
		listItem := SelectItem{Name: element.Name, Value: id}
		selectList = append(selectList, listItem)
	}

	return selectList
}

// AskForProvider will ask the user which provider they would like to use
func AskForProvider() (string, error) {
	supportedProviders := []SelectItem{}
	digitalOcean := SelectItem{Name: "DigitalOcean", Value: "DO"}
	supportedProviders = append(supportedProviders, digitalOcean)

	providerPrompt := CreateCustomSelectPrompt("Select Provider", supportedProviders)

	providerIndex, _, providerPromptError := providerPrompt.Run()

	if providerPromptError != nil {
		return "", providerPromptError
	}

	selectedProvider := supportedProviders[providerIndex]

	return selectedProvider.Value, nil
}
