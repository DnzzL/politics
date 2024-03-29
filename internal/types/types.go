package types

import "politics/.gen/model"

type EnrichedStanced struct {
	model.Stance
	model.Party
}
