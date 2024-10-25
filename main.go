package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type WeightedItem struct {
	Value  string
	Weight int
}

type Items struct {
	Items []WeightedItem `yaml:"items"`
}

func weightedRandomChoice(items []WeightedItem, n int) []WeightedItem {
	// Sort items by weight in ascending order
	sort.Slice(items, func(i, j int) bool {
		return items[i].Weight < items[j].Weight
	})

	// Calculate cumulative weights
	cumWeights := make([]int, len(items))
	cumWeights[0] = items[0].Weight
	for i := 1; i < len(items); i++ {
		cumWeights[i] = cumWeights[i-1] + items[i].Weight
	}

	// Generate random numbers
	rand.Seed(time.Now().UnixNano())
	randomNums := make([]int, n)
	for i := 0; i < n; i++ {
		randomNums[i] = rand.Intn(cumWeights[len(cumWeights)-1])
	}

	// Extract values within brackets for alphabetization
	values := make([]string, len(items))
	for i, item := range items {
		values[i] = extractValueInBrackets(item.Value)
	}

	// Find corresponding items based on random numbers
	selectedItems := make([]WeightedItem, n)
	for i, randomNum := range randomNums {
		// Find the original item index based on the sorted value
		originalIndex := findItemIndexByValue(items, values[i])
		if originalIndex == -1 {
			// Handle potential error (item not found)
			continue
		}

		// Use the random number to select the corresponding item
		index := sort.Search(len(cumWeights), func(j int) bool {
			return cumWeights[j] > randomNum
		})

		selectedItems[i] = items[(originalIndex+index-1)%len(items)]
	}

	// Filter out duplicates based on Value
	uniqueItems := make([]WeightedItem, 0, len(selectedItems))
	seen := make(map[string]struct{}, len(selectedItems))
	for _, item := range selectedItems {
		value := extractValueInBrackets(item.Value)
		if _, ok := seen[value]; !ok {
			uniqueItems = append(uniqueItems, item)
			seen[value] = struct{}{}
		}
	}

	// Alphabetize uniqueItems based on characters in brackets
	sort.Slice(uniqueItems, func(i, j int) bool {
		return extractValueInBrackets(uniqueItems[i].Value) < extractValueInBrackets(uniqueItems[j].Value)
	})

	return uniqueItems
}

// Helper function to extract the value within brackets
func extractValueInBrackets(value string) string {
	start := strings.Index(value, "{")
	end := strings.LastIndex(value, "}")
	if start == -1 || end == -1 {
		return value // Return original value if no brackets found
	}
	return value[start+1 : end]
}

// Helper function to find the original item index based on the sorted value
func findItemIndexByValue(items []WeightedItem, value string) int {
	for i, item := range items {
		if extractValueInBrackets(item.Value) == value {
			return i
		}
	}
	return -1 // Item not found
}

func readYamlConfig(filename string) ([]WeightedItem, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading YAML file: %w", err)
	}

	type YamlConfig struct {
		Items []WeightedItem
	}

	var config YamlConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("Error parsing YAML data: %w", err)
	}

	return config.Items, nil
}

func readTextFile(filename string) ([]WeightedItem, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading text file: %w", err)
	}

	strings.ReplaceAll(string(data), "\n\n", "\n")
	lines := strings.Split(string(data), "\n")
	items := make([]WeightedItem, 0)
	for _, line := range lines {
		if line != "\r" && line != "" {
			lineNoCharageReturn := strings.ReplaceAll(line, "\r", "")
			value := lineNoCharageReturn
			items = append(items, WeightedItem{Value: value, Weight: 100})
		}
	}

	return items, nil
}

func writeYamlFile(filename string, items []WeightedItem) error {
	itemsToSave := Items{Items: items}
	data, err := yaml.Marshal(itemsToSave)
	if err != nil {
		return fmt.Errorf("Error marshalling YAML data: %w", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("Error writing YAML file: %w", err)
	}

	return nil
}

func main() {
	fmt.Println("Flags are (yaml <filename.yaml>, num <number>) or (toyaml <filename.txt>)")

	// Define flags
	var yamlFilename string
	var numItems int
	var textToYaml string

	flag.StringVar(&yamlFilename, "yaml", "", "YAML file containing weighted items")
	flag.IntVar(&numItems, "num", 5, "Number of items to generate")
	flag.StringVar(&textToYaml, "toyaml", "", "Text file containing the different action")
	flag.Parse()

	// Check if `doSomethingDifferent` is set
	if textToYaml != "" {
		// Read the text file and perform the action
		selectedItems, err := readTextFile(textToYaml)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Save the selected items to a YAML file
		err = writeYamlFile("generated_output_items_from_textfile.yaml", selectedItems)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	if yamlFilename != "" {
		// Read YAML configuration file and handle errors
		items, err := readYamlConfig(yamlFilename)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Use the items from the configuration file
		selectedItems := weightedRandomChoice(items, numItems)
		for _, item := range selectedItems {
			fmt.Println(item.Value)
		}
	}
}
