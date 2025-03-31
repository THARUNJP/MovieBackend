package control

import (
	"MovieBack/config"
	"MovieBack/internal/types"
	"fmt"
)

func InsertRefreshToken(refToken string, userID string) {

	existingUser, _ := config.GetRecords("SELECT user_id FROM refreshTokenList WHERE user_id=$1", types.Slice{userID})

	if len(existingUser) > 0 {
		fmt.Println(existingUser, refToken, userID)
		// update func

	} else {
		// insert function
	}
	fmt.Print("ref token saved successfully")
}
