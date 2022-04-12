package fiberblog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hweihwang/go-blogs/common"
	"github.com/hweihwang/go-blogs/component"
	"github.com/hweihwang/go-blogs/modules/blog/blogbiz"
	"github.com/hweihwang/go-blogs/modules/blog/blogstorage"
	"strconv"
)

func GetBlog(ctx component.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		atoi, err := strconv.Atoi(id)
		
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		
		store := blogstorage.NewSQLStore(ctx.GetMainDBConnection())
		
		biz := blogbiz.NewGetBlogBiz(store)
		
		blog, err := biz.GetBlog(c.Context(), int64(atoi))
		
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		
		return c.Status(fiber.StatusOK).JSON(common.SimpleSuccessResponse(blog))
	}
}