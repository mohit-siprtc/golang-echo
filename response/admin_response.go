package response

// "go_rest_mohit/request"

type Response struct {
	// request.Request
	ID     int `json:"id" bson:"_id"`
	Name   string
	Gender string
	Age    float64
}
type IdResponse struct {
	ID int `bson:"seq"`
}
