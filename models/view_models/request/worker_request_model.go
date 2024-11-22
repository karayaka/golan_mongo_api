package request

type WorkerRequestModel struct {
	Id         string `json:"id" bson:"id"`
	UserId     int32  `json:"userId" bson:"user_id"`
	Name       string `json:"name" bson:"name"`
	Surname    string `json:"surname" bson:"surname"`
	Email      string `json:"email" bson:"email"`
	Department string `json:"department" bson:"department"`
}
