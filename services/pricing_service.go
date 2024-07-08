package services

import (
	"HackOn/models"
	"HackOn/repositories"
	"fmt"
	"strconv"
	"strings"
)

var gmvLevels = map[string]int{
	"<1 L":    0,
	"1L-5L":   1,
	"5L-10L":  2,
	"10L-20L": 3,
	"20L-50L": 4,
	"50L-75L": 5,
	"75L-1Cr": 6,
	"1Cr-5Cr": 7,
	"5Cr+":    8,
}

var levelWeightage = map[int]int{
	0: 0,
	1: 1,
	2: 1,
	3: 2,
	4: 2,
	5: 3,
	6: 3,
	7: 4,
	8: 4,
	9: 5,
}

func GetGMVForLevel(level int) string {
	for str, lvl := range gmvLevels {
		if lvl == level {
			return str
		}
	}
	return "" // Handle case where level is not found
}

func CalculatePricing(merchantID string, round string, volume string, servicesNegotiated string, competition string) (models.Parameter, error) {
	var parameter models.Parameter

	merchantGivenLevel := gmvLevels[volume]
	merchantActualLevel := 2
	levelDiff := merchantGivenLevel - merchantActualLevel
	fmt.Println(merchantGivenLevel, merchantActualLevel, levelDiff)

	if levelDiff > 0 {
		levelDiffAccToWeightage, ok := levelWeightage[levelDiff]
		if ok {
			volume = GetGMVForLevel(merchantGivenLevel - levelDiffAccToWeightage)
		}
	}

	fmt.Println("new volume is", volume)

	if round == "1" {
		fmt.Println("calculating for round 1")
		param, err := repositories.GetParamerterByRoundCompAndVolume(round, volume, competition)
		if err != nil {
			return parameter, err
		}
		parameter = *param
	} else {
		fmt.Println("calculating for round 2/ other")
		param, err := repositories.GetParamerterByRoundCompAndVolume(round, volume, competition)
		prevRound, err := strconv.Atoi(round)
		prevRound--
		prevParam, err := repositories.GetParamerterByRoundCompAndVolume(strconv.Itoa(prevRound), volume, competition)

		if err != nil {
			return parameter, err
		}
		fmt.Println(prevParam)
		fmt.Println(param)

		parameter = *prevParam
		parameter.Round = "Round " + round
		fmt.Println("current parameter:")
		fmt.Println(parameter)

		servicesNegotiated = "upi and credit"
		// if the service is negotiated, then pick param.<service_name> else pick prevParam.<service_name>
		servicesNegotiated = strings.ToLower(servicesNegotiated)
		if strings.Contains(servicesNegotiated, "upi") {
			fmt.Println("1")
			parameter.UPI = param.UPI
		}
		if strings.Contains(servicesNegotiated, "banking") {
			fmt.Println("2")
			parameter.NB = param.NB
		}
		if strings.Contains(servicesNegotiated, "credit") {
			fmt.Println("3")
			parameter.CC = param.CC
		}
		if strings.Contains(servicesNegotiated, "debit") {
			fmt.Println("4")
			parameter.DC = param.DC
		}
		if strings.Contains(servicesNegotiated, "others") {
			fmt.Println("5")
			parameter.Others = param.Others
		}
		fmt.Println(parameter)
	}

	return parameter, nil
}
