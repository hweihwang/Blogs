package fiberblog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hweihwang/go-blogs/common"
	"github.com/hweihwang/go-blogs/component"
	"github.com/hweihwang/go-blogs/modules/blog/blogbiz"
	"github.com/hweihwang/go-blogs/modules/blog/blogmodel"
	"github.com/hweihwang/go-blogs/modules/blog/blogstorage"
)

func CreateBLog(appCtx component.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var blog blogmodel.BlogCreateRequest

		if err := c.BodyParser(&blog); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		store := blogstorage.NewSQLStore(appCtx.GetMainDBConnection())

		biz := blogbiz.NewCreateBlogBiz(store)

		if err := biz.CreateBlog(c.Context(), &blog); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(common.SimpleSuccessResponse(&blog))
	}
}
