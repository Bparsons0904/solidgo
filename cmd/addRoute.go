package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var addRouteCmd = &cobra.Command{
	Use:     "add [route name]",
	Aliases: []string{"a"},
	Short:   "Add a route to an existing file",
	Long:    `Adds a new route to an existing route and controller file.`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		routeName := args[0]
		httpMethod, _ := cmd.Flags().GetString("method")
		functionName, _ := cmd.Flags().GetString("function")
		pathName, _ := cmd.Flags().GetString("path")

		caseType := cases.Title(language.English)
		if httpMethod == "" {
			httpMethod = "Get"
		}
		httpMethod = caseType.String(httpMethod)

		titledRouteName := caseType.String(routeName)
		if functionName == "" {
			functionName = "Get" + titledRouteName
		}

		if pathName == "" {
			pathName = strings.ToLower(routeName)
		}

		moduleName, err := getModuleName()
		if err != nil {
			fmt.Println("Error reading module name:", err)
			os.Exit(1)
		}

		routeInfo := RouteInfo{
			RouteName:    titledRouteName,
			RouteVar:     strings.ToLower(routeName) + "Routes",
			URLPath:      pathName,
			HTTPMethod:   httpMethod,
			FunctionName: functionName,
			ModuleName:   moduleName,
		}

		addToRouteFile(routeName, routeInfo)
		addToControllerFile(routeName, routeInfo)
	},
}

func init() {
	routeCmd.AddCommand(addRouteCmd)
}

func addToRouteFile(routeName string, routeInfo RouteInfo) {
	filename := "server/routes/" + strings.ToLower(routeName) + ".routes.go"

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}

	content := string(fileContent)

	insertionIndex := strings.LastIndex(content, "}")
	if insertionIndex == -1 {
		fmt.Println("No closing brace found in the file:", filename)
		return
	}

	newRoute := fmt.Sprintf("\t%s.%s(\"/%s\", controllers.%s)\n",
		routeInfo.RouteVar, routeInfo.HTTPMethod, routeInfo.URLPath, routeInfo.FunctionName)
	updatedContent := content[:insertionIndex] + newRoute + content[insertionIndex:]

	err = os.WriteFile(filename, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", filename, err)
		return
	}

	fmt.Printf("Successfully added new route to %s\n", filename)

	sortedContent := sortRoutes(updatedContent)

	// Write the updated, sorted content back to the file
	err = os.WriteFile(filename, []byte(sortedContent), 0644)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", filename, err)
		return
	}
}

func addToControllerFile(routeName string, routeInfo RouteInfo) {
	filename := "server/controllers/" + strings.ToLower(routeName) + ".controller.go"

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}

	content := string(fileContent)

	newFunction := fmt.Sprintf("\nfunc %s(c *fiber.Ctx) error {\n\treturn c.Status(fiber.StatusOK).JSON(fiber.Map{\"status\": \"success\", \"message\": \"%s called\"})\n}\n",
		routeInfo.FunctionName, routeInfo.FunctionName)

	updatedContent := content + newFunction

	err = os.WriteFile(filename, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", filename, err)
		return
	}

	fmt.Printf("Successfully added new controller function to %s\n", filename)
}

func sortRoutes(fileContent string) string {
	lines := strings.Split(fileContent, "\n")
	var beforeRoutes, routes, afterRoutes []string
	insideRoutesBlock := false
	foundGroupDefinition := false

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}
		if strings.HasSuffix(trimmedLine, "{") {
			insideRoutesBlock = true
			beforeRoutes = append(beforeRoutes, line)
		} else if strings.HasPrefix(trimmedLine, "}") {
			insideRoutesBlock = false
			afterRoutes = append(afterRoutes, line)
		} else if insideRoutesBlock {
			if strings.Contains(trimmedLine, ".Group(") {
				beforeRoutes = append(beforeRoutes, line)
				foundGroupDefinition = true
			} else if foundGroupDefinition && strings.Contains(trimmedLine, "Routes.") {
				routes = append(routes, trimmedLine)
			}
		} else {
			beforeRoutes = append(beforeRoutes, line)
		}
	}

	sort.Slice(routes, func(i, j int) bool {
		return routes[i] < routes[j]
	})

	return strings.Join(beforeRoutes, "\n") + "\n\t" + strings.Join(routes, "\n\t") + "\n" + strings.Join(afterRoutes, "\n")
}
