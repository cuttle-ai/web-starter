package web_server

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/cuttle-ai/web-starter/project"
	"github.com/spf13/cobra"
	"github.com/tcnksm/go-input"
)

/* This file contains the generate command specs */

func init() {
	WebServerCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the boiler plate code for the web-server",
	Long:  `All software has versions. This is web-starter's`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
		 * We will initiate the input
		 * Then we will the user for project details
		 * Then will generate the project
		 * Then we will install it
		 */
		//initializing the UI
		ui := &input.UI{
			Writer: os.Stdout,
			Reader: os.Stdin,
		}

		//prompting the user for project details
		pr, err := prompts(ui)
		if err != nil {
			//Error while promoting the user for inputs of project generation
			fmt.Println(err)
			os.Exit(1)
		}

		//generating the project
		pr.InitSources()
		err = pr.Setup()
		if err != nil {
			//Error while genertaing the project
			fmt.Println(err)
			os.Exit(1)
		}

		//installing the project
		c := exec.Command("go", "install", pr.Package)
		fmt.Println("Installing", pr.Name)
		err = c.Run()
		if err != nil {
			//Error while genertaing the project
			fmt.Println("Project is generated. But couldn't install it", err)
			os.Exit(1)
		}

		os.Exit(0)
	},
}

func prompts(ui *input.UI) (*project.Project, error) {
	/*
	 * First we will ask for the name of the project
	 * Then for project description
	 * Then for author
	 * Then for project destination
	 * Then for project package name
	 * then for license
	 */
	name, err := ui.Ask("Project name", &input.Options{
		Default:  "Web Server",
		Required: true,
	})
	if err != nil {
		return nil, err
	}
	description, err := ui.Ask("Project description", &input.Options{
		Default:  "Backend server",
		Required: true,
	})
	if err != nil {
		return nil, err
	}
	author, err := promptAuthor(ui)
	if err != nil {
		return nil, err
	}
	dst, err := ui.Ask("Project destination", &input.Options{
		Default:  project.GoPath() + "github.com" + project.Separator + strings.Split(author.Email, "@")[0] + project.Separator + "web-server",
		Required: true,
	})
	if err != nil {
		return nil, err
	}
	pkgName, err := ui.Ask("Package name", &input.Options{
		Default:  "github.com/" + strings.Split(author.Email, "@")[0] + "/web-server",
		Required: true,
	})
	if err != nil {
		return nil, err
	}
	lic, err := promptLicense(ui)
	if err != nil {
		return nil, err
	}

	return &project.Project{
		Name:        name,
		Description: description,
		Author:      author,
		Destination: dst,
		Package:     pkgName,
		License:     lic,
	}, nil
}

func promptAuthor(ui *input.UI) (project.Author, error) {
	/*
	 * First we will ask for the author name
	 * Then for author email
	 */
	name, err := ui.Ask("Author name", &input.Options{
		Default:  "cuttle.ai",
		Required: true,
	})
	if err != nil {
		return project.Author{}, err
	}
	email, err := ui.Ask("Author email", &input.Options{
		Default:  "hi@cuttle.ai",
		Required: true,
	})
	if err != nil {
		return project.Author{}, err
	}
	return project.Author{Name: name, Email: email}, nil
}

func promptLicense(ui *input.UI) (project.License, error) {
	/*
	 * First we will ask for the license type
	 * Then for copyright year
	 * Then for organisation
	 */
	licType, err := ui.Select("Type of license", []string{
		string(project.AGPL3),
		string(project.BSD2),
		string(project.BSD3),
		string(project.CLOSED),
		string(project.GPL2),
		string(project.MIT),
		string(project.UNLICENSED),
	}, &input.Options{
		Default:  string(project.MIT),
		Required: true,
	})
	if err != nil {
		return project.License{}, err
	}
	year, err := ui.Ask("Copyright year", &input.Options{
		Default:  strconv.Itoa(time.Now().Year()),
		Required: true,
	})
	if err != nil {
		return project.License{}, err
	}

	org, err := ui.Ask("Organisation", &input.Options{
		Default:  "Cuttle.ai",
		Required: true,
	})
	if err != nil {
		return project.License{}, err
	}
	return project.License{Type: project.LicenseType(licType), Year: year, Organisation: org}, nil
}
