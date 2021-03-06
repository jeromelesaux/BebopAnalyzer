package main

//go:generate swagger generate spec -o swagger.json
import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jeromelesaux/bebopanalyzer/configuration"
	"github.com/jeromelesaux/bebopanalyzer/fsmanager"
	"github.com/jeromelesaux/bebopanalyzer/model"
	"github.com/jeromelesaux/bebopanalyzer/routes"
	"net/http"
	"os"
)

var Conf configuration.AppConfiguration

// https://goswagger.io/generate/spec.html
func main() {

	if len(os.Args) == 1 {
		fmt.Println("Usage : app configuration-file.json bebop-fly.json (optionnal for web server)")
		return
	}

	Conf := configuration.LoadConfiguration(os.Args[1])

	if len(os.Args) == 3 {
		bebopJsonFile := os.Args[2]
		fmt.Println("Parsing " + bebopJsonFile)
		pud := model.Load(bebopJsonFile)
		project := fsmanager.Project{BaseDir: Conf.BasepathStorage, Name: pud.SerialNumber, Data: pud, Date: pud.Date}
		project.PerformAnalyse(pud)
		fmt.Println("job done and generated new file at " + project.GeneratedData)
	} else {
		fmt.Println("Starting server web at port " + Conf.HttpPort)
		// gestion des routes http
		routes.Conf = Conf
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/import/", routes.AnalyseFly)
		router.HandleFunc("/get/{serialNumber}/{flyDate}/{type}", routes.GetFile)
		router.HandleFunc("/list", routes.GetListTree)
		router.HandleFunc("/chart/{serialNumber}/{flyDate}", routes.GetChart)
		router.HandleFunc("/displayFly/{serialNumber}/{flyDate}", routes.GetMaps)
		router.HandleFunc("/rebuild", routes.RebuildData)
		router.PathPrefix("/").Handler(http.FileServer(http.Dir("./resources")))
		fmt.Println(http.ListenAndServe(Conf.HttpPort, router))
	}

}
