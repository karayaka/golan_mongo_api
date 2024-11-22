package entitymodels

import "time"

type Worker struct {
	Id          string    `bson:"_id,omitempty"`
	UserId      int32     `bson:"user_id"`
	Name        string    `bson:"name"`
	Surname     string    `bson:"surname"`
	Email       string    `bson:"email"`
	Department  string    `bson:"department"`
	CreatedDate time.Time `bson:"created_date,omitempty"`
	UpdatedDate time.Time `bson:"updated_date,omitempty"`
	CreatedBy   uint64    `bson:"created_by"`
	UpdatedBy   uint64    `bson:"updated_by"`
}

func (w Worker) ToMap() map[string]interface{} {
	out := make(map[string]interface{})
	out["user_id"] = w.UserId
	out["name"] = w.Name
	out["surname"] = w.Surname
	out["email"] = w.Email
	out["department"] = w.Department
	out["updated_date"] = w.UpdatedDate
	out["created_date"] = w.CreatedDate
	out["created_by"] = w.CreatedBy
	out["updated_by"] = w.UpdatedBy
	return out
}
