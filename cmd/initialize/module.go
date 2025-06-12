package initialize

import (
	"fmt"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// InitializeModules initializes all modules and registers routes
func InitializeModules(e *echo.Echo, db *gorm.DB) {
	// Add other modules initialization here
}

func InitializeRoutes(e *echo.Echo) {
	routes := make([][]string, 0)

	for _, route := range e.Routes() {
		routes = append(routes, []string{
			route.Method,
			route.Path,
			route.Name,
		})
	}

	// Sort routes by path
	sort.Slice(routes, func(i, j int) bool {
		return routes[i][1] < routes[j][1]
	})

	// Find the maximum width for each column
	methodWidth, pathWidth, nameWidth := 6, 4, 4 // Minimum widths based on headers
	for _, route := range routes {
		if len(route[0]) > methodWidth {
			methodWidth = len(route[0])
		}
		if len(route[1]) > pathWidth {
			pathWidth = len(route[1])
		}
		if len(route[2]) > nameWidth {
			nameWidth = len(route[2])
		}
	}

	// Add some padding
	methodWidth += 2
	pathWidth += 2
	nameWidth += 2

	// Calculate total width
	totalWidth := methodWidth + pathWidth + nameWidth + 4 // 4 for the separators

	// Print table header
	fmt.Println(strings.Repeat("=", totalWidth))
	fmt.Println("REGISTERED ROUTES:")
	fmt.Println(strings.Repeat("=", totalWidth))

	// Print column headers
	fmt.Printf("| %-*s | %-*s | %-*s |\n", methodWidth-2, "METHOD", pathWidth-2, "PATH", nameWidth-2, "NAME")
	fmt.Println(strings.Repeat("-", totalWidth))

	// Print rows
	for _, route := range routes {
		fmt.Printf("| %-*s | %-*s | %-*s |\n", methodWidth-2, route[0], pathWidth-2, route[1], nameWidth-2, route[2])
	}

	fmt.Println(strings.Repeat("=", totalWidth))
}
