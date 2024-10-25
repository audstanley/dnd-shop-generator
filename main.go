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
	// Calculate the total weight
	totalWeight := 0
	for _, item := range items {
		totalWeight += item.Weight
	}

	// Normalize weights to probabilities
	probabilities := make([]float64, len(items))
	for i, item := range items {
		probabilities[i] = float64(item.Weight) / float64(totalWeight)
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	selectedItems := make([]WeightedItem, 0, n)
	for i := 0; i < n; i++ {
		// Generate a random number between 0 and 1
		randomNum := rand.Float64()

		cumulativeProbability := 0.0
		for i, item := range items {
			cumulativeProbability += probabilities[i]
			if randomNum <= cumulativeProbability {
				selectedItems = append(selectedItems, item)
				break
			}
		}
	}

	// Extract values within brackets for alphabetization
	values := make([]string, len(items))
	for i, item := range items {
		values[i] = extractValueInBrackets(item.Value)
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

func readYamlConfig(filename string) ([]WeightedItem, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %w", err)
	}

	type YamlConfig struct {
		Items []WeightedItem
	}

	var config YamlConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML data: %w", err)
	}

	return config.Items, nil
}

func readTextFile(filename string) ([]WeightedItem, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading text file: %w", err)
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
		return fmt.Errorf("error marshalling YAML data: %w", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing YAML file: %w", err)
	}

	return nil
}

func main() {
	// Define flags
	var yamlFilename string
	var numItems int
	var textToYaml string
	var silent bool

	flag.StringVar(&yamlFilename, "yaml", "", "YAML file containing weighted items")
	flag.IntVar(&numItems, "num", 5, "Number of items to generate")
	flag.StringVar(&textToYaml, "toyaml", "", "Text file containing the different action")
	flag.BoolVar(&silent, "silent", false, "Silence the flag helper")
	flag.Parse()

	if !silent {
		fmt.Println("Flags are (yaml <filename.yaml>, num <number>) or (toyaml <filename.txt>)")
	}

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
