package middleware

// func TokenValidation(c *fiber.Ctx) error {
// 	// Get the Authorization header
// 	authHeader := c.Get("Authorization")

// 	// Example: "Bearer abc.def.ghi"
// 	if authHeader == "" {
// 		return c.Status(401).JSON(fiber.Map{
// 			"error": "Authorization header missing",
// 		})
// 	} else {
// 		fmt.Println("authHeader", authHeader)
// 		return c.Status(200).JSON("sucess")
// 	}

// }
