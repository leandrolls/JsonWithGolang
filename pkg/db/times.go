package db

import "github.com/leandrolls/JsonWithGo/pkg/models"

var times = []models.Time{models.Time{Nome: "São Paulo Futebol Clube", Divisao: "Serie A", Grandeza: "ENORME", ID: "sao-paulo"},
	models.Time{Nome: "Vasco da Gama", Divisao: "Serie B", Grandeza: "Grande",ID: "vasco"},
	models.Time{Nome: "Parmera", Divisao: "Serie B", Grandeza: "Minúsculo",ID: "palmeiras"}}

func GetTeam(time string) *models.Time{
	for _, v := range times {
		if v.ID == time {
			return &v
		}
	}
	return nil

}
func GetTeams() []models.Time{
	return times
}
