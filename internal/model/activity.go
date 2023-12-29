package model

var (
	ActivityTypePhysical ActivityType = "physical activity"
	ActivityTypeSocial   ActivityType = "social activity"
	ActivityTypeEvent    ActivityType = "event"
)

type Activity struct {
	base BaseModel
	Type ActivityType `json:"type"`
}

type ActivityType string
