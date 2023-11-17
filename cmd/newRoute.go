// package cmd

// import (
// 	"github.com/spf13/cobra"
// 	// Other imports...
// )

// var newRouteCmd = &cobra.Command{
// 	// Command definition...
// }

// func init() {
// 	routeCmd.AddCommand(newRouteCmd)
// 	// Set up flags and other configurations for newRouteCmd
// }

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const routeTemplate = `package routes

import (
	"github.com/gofiber/fiber/v2"
	"{{.ModuleName}}/controllers"
)

func {{.RouteName}}Routes(api fiber.Router) {
	{{.RouteVar}} := api.Group("/{{.URLPath}}")
	{{.RouteVar}}.{{.HTTPMethod}}("/", controllers.{{.FunctionName}})
}
`

const controllerTemplate = `package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func {{.FunctionName}}(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "{{.FunctionName}} called"})
}
`

type RouteInfo struct {
	RouteName    string
	RouteVar     string
	URLPath      string
	HTTPMethod   string
	FunctionName string
	ModuleName   string
}

var newRouteCmd = &cobra.Command{
	Use:     "route [route name]",
	Aliases: []string{"n"},
	Short:   "Create a new route",
	Long:    `Creates a new route file with a specified route group and an initial route.`,
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

		createRouteFile(routeName, routeInfo)
		createControllerFile(routeName, routeInfo)
	},
}

func init() {
	routeCmd.AddCommand(newRouteCmd)
}

func createRouteFile(filename string, routeInfo RouteInfo) {
	file, err := os.Create("routes/" + filename + ".routes.go")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	template, err := template.New("route").Parse(routeTemplate)
	if err != nil {
		fmt.Println("Unable to parse template:", err)
		os.Exit(1)
	}

	err = template.Execute(file, routeInfo)
	if err != nil {
		fmt.Println("Unable to execute template:", err)
		os.Exit(1)
	}

	fmt.Println("Route file created successfully:", filename)
}

func getModuleName() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module ") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				return parts[1], nil
			}
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("module directive not found in go.mod")
}

func createControllerFile(filename string, routeInfo RouteInfo) {
	file, err := os.Create("controllers/" + filename + ".controller.go")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	template, err := template.New("controller").Parse(controllerTemplate)
	if err != nil {
		fmt.Println("Unable to parse template:", err)
		os.Exit(1)
	}

	err = template.Execute(file, routeInfo)
	if err != nil {
		fmt.Println("Unable to execute template:", err)
		os.Exit(1)
	}

	fmt.Println("Controller file created successfully:", filename)
}
