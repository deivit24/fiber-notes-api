package users

import (
	"github.com/deivit24/fiber-notes-api/pkg/common/models"

	"github.com/gofiber/fiber/v2"
)

type CreateUserRequestBody struct {
	Username  string `json:"username"`
	Email string  `json:"email"`
	Password string  `json:"password"`
}
func (h handler) CreateUser(c *fiber.Ctx) error {
	body := CreateUserRequestBody{}

	// parse body, attach to AddProductRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var user models.User

	user.Username = body.Username
	user.Email = body.Email
	user.Password = body.Password
	user.Prepare()
	var err = user.Validate("")
	// insert new db entry
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "There was an error creating user")
	}

	// insert new db entry
	if result := h.DB.Create(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&user)
}

// func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 	}
// 	user := models.User{}
// 	err = json.Unmarshal(body, &user)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	user.Prepare()
// 	err = user.Validate("")
// 	if err != nil {
// 		responses.ERROR(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	userCreated, err := user.SaveUser(server.DB)

// 	if err != nil {

// 		formattedError := formaterror.FormatError(err.Error())

// 		responses.ERROR(w, http.StatusInternalServerError, formattedError)
// 		return
// 	}
// 	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
// 	responses.JSON(w, http.StatusCreated, userCreated)
// }