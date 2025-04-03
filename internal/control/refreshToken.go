package control

import (
	"MovieBack/config"
	"MovieBack/internal/types"
	"fmt"
	"time"
)

func InsertRefreshToken(refToken string, userID string) {

	expiresAt := time.Now().Add(30 * 24 * time.Hour)

	existingUser, _ := config.GetRecords("SELECT user_id FROM refreshTokenList WHERE user_id=$1", types.Slice{userID})

	if len(existingUser) > 0 {
		fmt.Println(existingUser, refToken, userID)

		updateQuery, err := config.ExecuteQuery("UPDATE refreshTokenList SET ref_token = $1,expires_at = $2 WHERE user_id=$3", types.Slice{refToken, expiresAt, userID})
		config.SetKey(userID, refToken, time.Until(expiresAt))
		// update func
		if err != nil {
			fmt.Printf("err in updating: %v\n", err.Error())
		} else {
			fmt.Printf("ref token updated: %v\n", updateQuery)
		}
	} else {
		// fmt.Println(err, refToken, userID)
		// Insert
		InsertQuery, _ := config.ExecuteQuery("INSERT INTO refreshTokenList(user_id,ref_token,expires_at) VALUES($1,$2,$3)", types.Slice{userID, refToken, expiresAt})
		config.SetKey(userID, refToken, time.Until(expiresAt))
		fmt.Print(InsertQuery, "insert")
	}
	// fmt.Println(config.GetKey("*"))
	fmt.Println("ref token saved successfully")

}
