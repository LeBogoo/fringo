package main

// proxy for github api that caches themes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var repoDir string = "themes-repo"
var summaryCache []ThemeSummary

var version = AppVersion{
	Version: os.Getenv("VERSION"),
	Commit:  os.Getenv("GIT_COMMIT"),
}

func getVersionHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, version)
}

func getRepoHandler(c echo.Context) error {
	return c.String(http.StatusOK, os.Getenv("GH_REPO"))
}

func cloneRepo() error {
	ghRepo := os.Getenv("GH_REPO")
	if ghRepo == "" {
		return fmt.Errorf("GH_REPO is not set")
	}

	url := fmt.Sprintf("https://github.com/%s.git", ghRepo)

	var err error
	err = os.Mkdir(repoDir, 0755)
	if err != nil {
		fmt.Println("Failed to create repo dir:", err)
		return err
	}

	fmt.Println("Cloning repository from", url)
	cmd := exec.Command("git", "clone", url, repoDir)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to clone repo:", err)
		return err
	}

	fmt.Println("Repository cloned successfully")
	return nil
}

func fetchAndPullRepo() error {
	fmt.Println("Fetching updates for repository")
	cmd := exec.Command("git", "-C", repoDir, "fetch", "--all")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Failed to fetch themes-repo:", err)
		return err
	}

	fmt.Println("Pulling latest changes for repository")
	cmd = exec.Command("git", "-C", repoDir, "pull")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to pull repo:", err)
		return err
	}

	fmt.Println("Repository updated successfully")
	return nil
}

func getUserThemesHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, summaryCache)
}

func updateSummaryCache() {
	fmt.Println("Updating summary cache")
	err := fetchAndPullRepo()
	if err != nil {
		fmt.Println("Failed to update repository:", err)
		return
	}

	themes, err := readThemes()
	if err != nil {
		fmt.Println("Failed to read themes:", err)
		return
	}

	summaryCache = []ThemeSummary{}

	for filename, theme := range themes {
		summaryCache = append(summaryCache, ThemeSummary{
			ID:       theme.ID,
			Name:     theme.Name,
			Author:   theme.Author,
			FileName: filename,
			Hash:     theme.GetHash(),
		})
	}

	// sort by id
	sort.Slice(summaryCache, func(i, j int) bool {
		return summaryCache[i].ID < summaryCache[j].ID
	})

	fmt.Println("Theme summary cache updated successfully")
}

func readThemes() (map[string]Theme, error) {
	themes := make(map[string]Theme)

	fmt.Println("Reading themes from repository")
	files, err := os.ReadDir(repoDir + "/themes")
	if err != nil {
		fmt.Println("Failed to read themes dir:", err)
		return themes, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		f, err := os.Open(repoDir + "/themes/" + file.Name())
		if err != nil {
			fmt.Println("Failed to open theme file:", err)
			return themes, err
		}

		theme := Theme{}
		err = json.NewDecoder(f).Decode(&theme)
		if err != nil {
			fmt.Println("Failed to decode theme file:", err)
			return themes, err
		}

		themes[file.Name()] = theme
	}

	fmt.Println("Themes read successfully")
	return themes, nil
}

func main() {
	// clone repo if it doesn't exist
	if _, err := os.Stat(repoDir); os.IsNotExist(err) {
		fmt.Println("Repository directory does not exist, cloning repository")
		err := cloneRepo()
		if err != nil {
			fmt.Println("Failed to clone repo:", err)
			return
		}
	}

	go func() {
		for {
			updateSummaryCache()
			time.Sleep(5 * time.Minute)
		}
	}()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/user-themes", getUserThemesHandler)
	e.GET("/version", getVersionHandler)
	e.GET("/repo", getRepoHandler)

	// Serve static files from the public and themes-repo directories
	e.Static("/", "public")
	e.Static("/user-themes", repoDir+"/themes")

	fmt.Println("Server started at :8080")
	e.Start(":8080")
}
