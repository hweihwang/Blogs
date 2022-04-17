package fiberblog

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hweihwang/go-blogs/common"
	"github.com/hweihwang/go-blogs/component"
	"github.com/hweihwang/go-blogs/modules/blog/blogbiz"
	"github.com/hweihwang/go-blogs/modules/blog/blogmodel"
	"github.com/hweihwang/go-blogs/modules/blog/blogstorage"
)

func UpdateBlog(appCtx component.AppContext) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var blog blogmodel.BlogUpdateRequest
		
		id := c.Params("id")
		
		uintid, err := common.StringToUint(id)
		
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		
		if err = c.BodyParser(&blog); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		store := blogstorage.NewSQLStore(appCtx.GetMainDBConnection())

		biz := blogbiz.NewUpdateBlogBiz(store)

		if err = biz.UpdateBlog(c.Context(), uintid, &blog); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(common.SimpleSuccessResponse(&blog))
	}
}
