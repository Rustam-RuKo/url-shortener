package routes 

import (
	"github.com/akhil/awesomeproject/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"


)

func ResolveURL(c *fiber.Ctx) error{

	url:= c.Params("url")
	r:= database.CreateClient(0)
	defer r.Close()


	value,err := r.Get(database.Ctx, url).Result
	if err == redis.Nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":"short not found in the database"
		})
	} 
	
	else if err != nil {	

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"cannot connect toi DB"})
			
	}
	
	rInr := database.CreateClient(1)
	defer rInr.Close

	_ = rInr.Incr(database.Ctx , "counter")

	return c.Redirect(value, 301)

}