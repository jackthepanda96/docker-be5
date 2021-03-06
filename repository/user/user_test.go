package user

import (
	"project/playground/be5/docker-play/configs"
	"project/playground/be5/docker-play/entities"
	"project/playground/be5/docker-play/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	repo := New(db)

	t.Run("Create User", func(t *testing.T) {
		mockUser := entities.User{Nama: "Ilham", HP: "08123456999"}
		res, err := repo.Insert(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Nama, res.Nama)
		assert.Equal(t, 1, int(res.ID))
	})
}
