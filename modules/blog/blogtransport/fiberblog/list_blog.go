package fiberblog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hweihwang/go-blogs/common"
	"github.com/hweihwang/go-blogs/component"
	"github.com/hweihwang/go-blogs/modules/blog/blogbiz"
	"github.com/hweihwang/go-blogs/modules/blog/blogmodel"
	"github.com/hweihwang/go-blogs/modules/blog/blogstorage"
)

func ListBlog(appCtx component.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var filter blogmodel.Filter

		if err := c.BodyParser(&filter); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		var paging common.Paging

		if err := c.BodyParser(&paging); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		
		paging.FulFill()

		store := blogstorage.NewSQLStore(appCtx.GetMainDBConnection())

		biz := blogbiz.NewListBlogBiz(store)

		blogs, err := biz.ListBlog(c.Context(), &filter, &paging)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(common.NewSuccessResponse(blogs, paging, filter))
	}
}
