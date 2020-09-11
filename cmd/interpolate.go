package cmd

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var primaryPrefix string
var alternativePrefix string
var filePath string

// interpolateCmd represents the interpolate command
var interpolateCmd = &cobra.Command{
	Use: "interpolate",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("no file specified")
		}
		_, err := os.Stat(args[0])
		if os.IsNotExist(err) {
			return errors.New("file " + args[0] + " does not exists")
		}

		return nil
	},
	Short: "Interpolate variables in file",
	Long:  "Interpolate the environment variables inside {{}} in file and substitutes them with the corresponding value",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		file, err := ioutil.ReadFile(filePath)
		checkError(err)

		interpolatedFile := interpolate(file)

		err = ioutil.WriteFile(filePath, interpolatedFile, 0444)
		checkError(err)
	},
}

func init() {
	rootCmd.AddCommand(interpolateCmd)

	interpolateCmd.Flags().StringVarP(&primaryPrefix, "prefix", "p", "", "primary prefix to add when looking for envs")
	interpolateCmd.Flags().StringVarP(&alternativePrefix, "alternative-prefix", "a", "", "prefix to use when the primary prefix env does not exists")
}

type env_var struct {
	name  string
	value string
}

func interpolate(file []byte) []byte {
	envs := make(map[string]env_var)
	envs = getVariablesToInterpolate(file, envs)

	// exit if there are no variables to interpolate
	if len(envs) == 0 {
		os.Exit(0)
	}

	err := checkEnvs(envs)
	checkError(err)

	return interpolateVariables(file, envs)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getVariablesToInterpolate(file_content []byte, envs map[string]env_var) map[string]env_var {
	re := regexp.MustCompile("\\{\\{(.+?)\\}\\}")
	match := re.FindAllStringSubmatch(string(file_content), -1)

	for parsedVar := range match {
		varName := strings.ReplaceAll(match[parsedVar][1], " ", "")
		// keep track of the entire pattern found by the regex
		// using as key the variable name
		if _, exists := envs[varName]; !exists {
			envs[varName] = env_var{name: match[parsedVar][0]}
		}
	}
	return envs
}

func checkEnvs(envs map[string]env_var) error {
	for varName, _ := range envs {
		varPrefixed := primaryPrefix + "_" + varName
		varPrefixedAlternative := alternativePrefix + "_" + varName
		var envValue string

		// get the escaped environment variable
		switch {
		case os.Getenv(varPrefixed) != "":
			envValue = strconv.Quote(os.Getenv(varPrefixed))
		case os.Getenv(varPrefixedAlternative) != "":
			envValue = strconv.Quote(os.Getenv(varPrefixedAlternative))
		default:
			return errors.New("environment variables " + varPrefixed + " and " + varPrefixedAlternative + " do not exist")
		}

		// discard the initial and final double quotes generated by strconv
		envValue = envValue[1 : len(envValue)-1]
		envs[varName] = env_var{name: envs[varName].name, value: envValue}

	}
	return nil
}

func interpolateVariables(file []byte, envs map[string]env_var) []byte {
	fileString := string(file)

	for varName, _ := range envs {
		env := envs[varName]
		fileString = strings.ReplaceAll(fileString, env.name, env.value)
	}

	return []byte(fileString)
}
