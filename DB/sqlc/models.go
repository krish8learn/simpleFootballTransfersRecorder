// Code generated by sqlc. DO NOT EDIT.

package DB

import (
	"time"
)

type Footballclub struct {
	FcID      int32  `json:"fc_id"`
	ClubName  string `json:"club_name"`
	CountryFc string `json:"country_fc"`
	// can be positive or negative
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type Player struct {
	PID        int32  `json:"p_id"`
	PlayerName string `json:"player_name"`
	Position   string `json:"position"`
	CountryPl  string `json:"country_pl"`
	// must be positive
	Value          int64     `json:"value"`
	FootballclubID int64     `json:"footballclub_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type Transfer struct {
	TID             int32 `json:"t_id"`
	Season          int64 `json:"season"`
	PlayerID        int32 `json:"player_id"`
	SourceClub      int32 `json:"source_club"`
	DestinationClub int32 `json:"destination_club"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
