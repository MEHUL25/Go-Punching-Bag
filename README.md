# Go-Algorithms


type EntityDetail struct {
	EntityID           int64  `json:"entity_id"`
	EntityName         string `json:"entity_name"`
}

type UserDetail struct {
	UserID           int64  `json:"user_id"`
	UserName         string `json:"user_name"`
	OrganizationID   int64  `json:"org_id"`
	OrganizationName string `json:"org_name,omitempty"`
}

type ActivityDetail struct {
	User UserDetail
    Entity EntityDetail 
	Status    string `json:"status,omitempty"`
	Activity  string `json:"activity,omitempty"`
	Event     string `json:"event,omitempty"`
	Timestamp time.Time
}


