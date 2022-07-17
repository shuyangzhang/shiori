package service

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/lonelyevil/khl"
	"github.com/shuyangzhang/shiori/constant"
	"github.com/shuyangzhang/shiori/dal"
	"github.com/shuyangzhang/shiori/dal/model"
	"github.com/shuyangzhang/shiori/tools/message"
)

func Arena(ctx *khl.KmarkdownMessageContext, parameters ...string) {
	var result string

	if len(parameters) == 5 {
		var defenderParty []int

		for _, characterName := range parameters {
			characterId, ok := constant.CharactersReverseMap[characterName]
			if ok {
				defenderParty = append(defenderParty, characterId)
			} else {
				panic(fmt.Sprintf("%v is not recognized, please check your input", characterName))
			}
		}

		sort.Slice(defenderParty, func(i, j int) bool {
			return defenderParty[i] < defenderParty[j]
		})

		defenderPartyObj, err := queryCharactersParty(defenderParty)
		if err != nil {
			result = fmt.Sprintf("defender party: %v is not exist", parameters)
		} else {
			solutionObj, err := queryArenaSolutionsByDefenderId(defenderPartyObj.ID)
			if err != nil {
				result = fmt.Sprintf("solution is not exist, you can /submit a solution for defender id: %v", defenderPartyObj.ID)
			} else {
				result = fmt.Sprintf("solution found\nsolution id: %v, solution: ", solutionObj.ID)
				party := queryCharactersPartyById(*solutionObj.AttackerID)
				for _, characterId := range party {
					result += fmt.Sprint(characterIdToName(characterId), " ")
				}
			}
		}

		// result = fmt.Sprintf("%v", defenderParty)
	} else {
		result = "defence party must be 5 characters"
	}

	message.Reply(ctx, result)
}

func queryCharactersParty(party []int) (*model.CharactersParties, error) {
	partyString, _ := json.Marshal(party)
	return dal.Query.CharactersParties.Where(dal.Query.CharactersParties.Characters.Eq(string(partyString))).First()
}

func queryArenaSolutionsByDefenderId(defenderId int32) (*model.ArenaSolutions, error) {
	return dal.Query.ArenaSolutions.Where(
		dal.Query.ArenaSolutions.DefenderID.Eq(defenderId),
		dal.Query.ArenaSolutions.IsDeleted.Is(false),
		dal.Query.ArenaSolutions.IsOutdated.Is(false),
	).First()
}

func queryCharactersPartyById(partyId int32) (party []int) {
	charactersPartyObj, err := dal.Query.CharactersParties.Where(dal.Query.CharactersParties.ID.Eq(partyId)).First()
	if err != nil {
		panic(fmt.Sprintf("character party id: %v is not exist", partyId))
	}
	partyString := charactersPartyObj.Characters
	err = json.Unmarshal([]byte(partyString), &party)
	if err != nil {
		panic(fmt.Sprintf("can not unmarshal party string: %v, from party id: %v", partyString, partyId))
	}
	return
}

func characterIdToName(characterId int) (characterName string) {
	characterAliases := constant.Characters[characterId]
	if len(characterAliases) > 1 {
		characterName = characterAliases[1]
	} else {
		characterName = characterAliases[0]
	}
	return
}
