package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

var responsibleBot string

type bot struct {
	botID      string
	microChips []int
	low        string
	high       string
}

func (b bot) getLowestMC() int {
	var lowest int
	for _, v := range b.microChips {
		if lowest == 0 {
			lowest = v
		}
		if v < lowest {
			lowest = v
		}
	}
	return lowest
}

func (b bot) getHighestMC() int {
	var highest int
	for _, v := range b.microChips {
		if highest == 0 {
			highest = v
		}
		if v > highest {
			highest = v
		}
	}
	return highest
}

func check(s string, err error) {
	if err != nil {
		log.Println(s, err)
	}
}

func deleteChip(n int, microChips []int) []int {
	for i, v := range microChips {
		if n == v {
			microChips = append(microChips[:i], microChips[1+i:]...)
		}
	}
	return microChips
}

//parsedInstructions [0]value [1]microChip [2]goes [3]to [4]bot [5]botNumber
func assignInputToBot(parsedInstructions []string, botArmy map[string]bot) map[string]bot {

	botID := parsedInstructions[4] + parsedInstructions[5]
	microChip, err := strconv.Atoi(parsedInstructions[1])
	check("trying to convert microChip string to int:", err)

	if _, ok := botArmy[botID]; ok {
		existingBot := botArmy[botID]
		existingBot.botID = botID
		existingBot.microChips = append(existingBot.microChips, microChip)

		botArmy[botID] = existingBot
	} else {

		var newBot bot
		newBot.botID = botID
		newBot.microChips = append(newBot.microChips, microChip)

		botArmy[botID] = newBot
	}
	return botArmy
}

//parsedInstructions [0]bot [1]botNumber [2]gives [3]low [4]to [5]bot
//[6]botNumber [7]and [8]high [9]to [10]bot [11]botNumber
func assignCommandToBot(parsedInstructions []string, botArmy map[string]bot) map[string]bot {
	botID := parsedInstructions[0] + parsedInstructions[1]
	low := parsedInstructions[5] + parsedInstructions[6]
	high := parsedInstructions[10] + parsedInstructions[11]

	if _, ok := botArmy[botID]; ok {
		existingBot := botArmy[botID]
		existingBot.low = low
		existingBot.high = high
		botArmy[botID] = existingBot
	} else {
		var newBot bot
		newBot.botID = botID
		newBot.low = low
		newBot.high = high
		botArmy[botID] = newBot
	}
	return botArmy
}

func recruitBot(instruction string, botArmy map[string]bot) error {
	parsedInstructions := strings.Split(instruction, " ")
	switch parsedInstructions[0] {
	case "value":
		assignInputToBot(parsedInstructions, botArmy)
		return nil

	case "bot":
		assignCommandToBot(parsedInstructions, botArmy)
		return nil

	default:
		return errors.New("cannot read instructions: " + instruction)
	}
}

func buildBotArmy(location string) (map[string]bot, error) {
	botArmy := make(map[string]bot)

	//Read instructions from the input file
	file, err := os.Open(location)
	if err != nil {
		return botArmy, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		instruction := scanner.Text()           //get line with instruction
		err := recruitBot(instruction, botArmy) //recruit a bot with the instruction
		check("while creating bot army", err)

	}

	return botArmy, nil
}

func giveMicroChips(b bot, botArmy map[string]bot) map[string]bot {
	if _, ok := botArmy[b.high]; ok {
		highBot := botArmy[b.high]
		highBot.microChips = append(highBot.microChips, b.getHighestMC())
		botArmy[b.high] = highBot
	} else {
		var newBot bot
		newBot.microChips = append(newBot.microChips, b.getHighestMC())
		botArmy[b.high] = newBot
	}
	if _, ok := botArmy[b.low]; ok {
		lowBot := botArmy[b.low]
		lowBot.microChips = append(lowBot.microChips, b.getLowestMC())
		botArmy[b.low] = lowBot
	} else {
		var newBot bot
		newBot.microChips = append(newBot.microChips, b.getLowestMC())
		botArmy[b.low] = newBot
	}

	b.microChips = []int{}
	botArmy[b.botID] = b
	return botArmy

}

func compareMicroChips(a, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)
	if reflect.DeepEqual(a, b) {
		return true
	}
	return false
}

func runBotSequence(botArmy map[string]bot, microChips []int) bool {
	lastSequence := true

	for _, value := range botArmy {
		if compareMicroChips(microChips, value.microChips) {
			responsibleBot = value.botID
		}
		if len(value.microChips) == 2 {
			giveMicroChips(value, botArmy)
			lastSequence = false
		}
	}
	return lastSequence
}

func main() {
	botArmy, err := buildBotArmy("../input.txt")
	check("drats, could not create my botArmy", err)
	var i int
	for !runBotSequence(botArmy, []int{61, 17}) {
		i++
		fmt.Println("Sequence:", i)
	}
	fmt.Println("Bot's finished")
	fmt.Println("The bot that is responsible for comparing value-61 microchips with value-17 microchips:")
	fmt.Println(responsibleBot)
}
