package cmd

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//addcmd permet d'ajouter des projets dans l'appli
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "this command allows you to add a new project and its path",
	Long:  `this command allows you to add a new project and its path`,
	Run: func(cmd *cobra.Command, args []string) {
		//init struct projecto
		var project Project

		if args == nil && args[0] == "" && args[1] == "" {
			fmt.Println("Erreur sur la commande")
		} else if args[0] != "" && args[1] == "." {
			project.name = args[0]
			addProjectActually(&project)
		} else {
			project.name = args[0]
			project.path = args[1]
			addProject(&project)
		}

	},
}

//une struct ayant les champs necessaire pour enregistrer un projet
type Project struct {
	name, path string
}

func addProjectActually(project *Project) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	project.path = dir
	addProject(project)
}

//Create function pour rechercher le fichier json et enregistrer le projet
func addProject(project *Project) {
	//Tout d'abord on retroubve le fichier de configuration des projets
	vp := viper.New()
	vp.SetConfigName("project")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	vp.Set(project.name, project.path)
	vp.WriteConfig()

	vp.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("le projet %s ete ajouter", in.Name)
	})

	vp.WatchConfig()
}

func init() {
	rootCmd.AddCommand(addCmd)

}
