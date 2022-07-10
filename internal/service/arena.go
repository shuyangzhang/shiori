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
			result = fmt.Sprintf("defender id: %v", defenderPartyObj.ID)
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
