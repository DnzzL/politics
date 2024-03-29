package types

import "goth/.gen/model"

type EnrichedStanced struct {
	model.Stance
	model.Party
}
