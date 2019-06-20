package seed

import (
	"github.com/icrowley/fake"
	"schedule-management/model"
)

func (s Seeder) UserSeed() {
	for i := 0; i < 20; i++ {
		u := model.User{}
		u.Username = fake.UserName()
		u.Phone = fake.Phone()
		u.Email = fake.EmailAddress()
		s.DB.Save(&u)
	}
}