package bd

import (
	"context"
	"time"
	"github.com/Rodripaz45/twitter/models"
)
/* Borra el follow de un usuario a otro */
func BorroRelacion(t models.Relacion)(bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}