package handler

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mercan/RandomUserGenerator/docs"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type User struct {
	Gender string `json:"gender" fake:"{gender}"`
	Name   struct {
		FirstName string `json:"firstName" fake:"{firstname}"`
		LastName  string `json:"lastName" fake:"{lastname}"`
		FullName  string `json:"fullName"`
	} `json:"name"`
	Login struct {
		Id           int    `json:"id" fake:"{number:1000000,99999999}"`
		Uuid         string `json:"uuid" fake:"{uuid}"`
		Username     string `json:"username" fake:"{username}"`
		Email        string `json:"email" fake:"{email}"`
		Password     string `json:"password" fake:"{password}"`
		PasswordHash struct {
			Md5    string `json:"md5"`
			Sha1   string `json:"sha1"`
			Sha256 string `json:"sha256"`
			Sha512 string `json:"sha512"`
		} `json:"password_hash"`
		PhoneNumber string `json:"phoneNumber" fake:"{phoneformatted}"`
		Role        string `json:"role" fake:"{randomstring:[guest, user, administrator]}"`
		UserAgent   struct {
			Chrome  string `json:"chrome" fake:"{chromeuseragent}"`
			Firefox string `json:"firefox" fake:"{firefoxuseragent}"`
			Opera   string `json:"opera" fake:"{operauseragent}"`
			Safari  string `json:"safari" fake:"{safariuseragent}"`
		} `json:"user_agent"`
	} `json:"login"`
	BirthDate      string `json:"birthDate" fake:"{daterange:1950-01-01,2001-01-01}"`
	RegisteredDate string `json:"registeredDate" fake:"{daterange:2001-01-01,2019-01-01}"`
	MacAddress     string `json:"macAddress"`
	IpAddress      struct {
		V4 string `json:"v4" fake:"{ipv4address}"`
		V6 string `json:"v6" fake:"{ipv6address}"`
	} `json:"ip_address"`
	Location struct {
		Country     string `json:"country" fake:"{country}"`
		City        string `json:"city" fake:"{city}"`
		State       string `json:"state" fake:"{state}"`
		Street      string `json:"street" fake:"{street}"`
		ZipCode     int    `json:"zipCode" fake:"{zip}"`
		Coordinates struct {
			Latitude  float64 `json:"latitude" fake:"{latitude}"`
			Longitude float64 `json:"longitude" fake:"{longitude}"`
		} `json:"coordinates"`
		TimeZone string `json:"timeZone" fake:"{timezoneregion}"`
	} `json:"location"`
	CreditCard struct {
		Type   string `json:"type" fake:"{creditcardtype}"`
		Number string `json:"number" fake:"{creditcardnumber}"`
		Exp    string `json:"exp" fake:"{creditcardexp}"`
		Cvv    string `json:"cvv" fake:"{creditcardcvv}"`
	} `json:"credit_card"`
	Bitcoin struct {
		Address    string `json:"address" fake:"{bitcoinaddress}"`
		PrivateKey string `json:"privateKey" fake:"{bitcoinprivatekey}"`
	} `json:"bitcoin"`
	Company struct {
		Name   string `json:"name" fake:"{company}"`
		Suffix string `json:"suffix" fake:"{companysuffix}"`
		Job    struct {
			Title      string `json:"title" fake:"{jobtitle}"`
			Level      string `json:"level" fake:"{joblevel}"`
			Descriptor string `json:"descriptor" fake:"{jobdescriptor}"`
		} `json:"job"`
	} `json:"company"`
	Size      string    `json:"size" fake:"{randomstring:[small, medium, large]}"`
	Rating    float32   `json:"rating" fake:"{number:0,5}"`
	CreatedAt time.Time `json:"createdAt"`
}

// @Tags Generate
// @Produce json
// @Success 200 {object} User
// @Failure 500 {object} Response
// @Router /v1/user [get]
func GenerateUser(c *fiber.Ctx) error {
	var user User
	if err := gofakeit.Struct(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&Response{
			Status:  fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// Name.FullName
	user.Name.FullName = user.Name.FirstName + " " + user.Name.LastName

	// MD5 Password Hash
	md5hash := md5.New()
	md5hash.Write([]byte(user.Login.Password))
	user.Login.PasswordHash.Md5 = hex.EncodeToString(md5hash.Sum(nil))

	// SHA1 Password Hash
	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(user.Login.Password))
	user.Login.PasswordHash.Sha1 = hex.EncodeToString(sha1Hash.Sum(nil))

	// SHA256 Password Hash
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(user.Login.Password))
	user.Login.PasswordHash.Sha256 = hex.EncodeToString(sha256Hash.Sum(nil))

	// SHA256 Password Hash
	sha512Hash := sha512.New()
	sha512Hash.Write([]byte(user.Login.Password))
	user.Login.PasswordHash.Sha512 = hex.EncodeToString(sha512Hash.Sum(nil))

	// Mac Address
	user.MacAddress = gofakeit.MacAddress()

	// CreatedAt
	user.CreatedAt = time.Now()

	return c.Status(fiber.StatusOK).JSON(&Response{
		Status: fiber.StatusOK,
		Data:   &user,
	})
}
