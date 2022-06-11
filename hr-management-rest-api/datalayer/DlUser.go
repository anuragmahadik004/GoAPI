package datalayer

import (
	"fmt"
	"log"

	"github.com/Knetic/go-namedParameterQuery"
	encryptdecrypt "github.com/anuragmahadik004/hr_api/encrypt_decrypt"
	"github.com/anuragmahadik004/hr_api/models"
	"github.com/google/uuid"
)

type DLUser struct {
}

func (DLUser) SaveUser(User models.User) bool {

	if len(User.UserId) == 0 {
		User.UserId = uuid.New().String()
	}

	User.Password, _ = encryptdecrypt.Encrypt(User.Password)

	//defining query parameters
	insQuery := `INSERT INTO [LoginMaster]
	([UserId],[UserName],[Password],[Email])
	VALUES(:UserId, :UserName, :Password, :Email)`

	updQuery := `UPDATE [LoginMaster]
		SET   [UserName] = :UserName 
		,[Password] = :Password
		,[Email] = :Email
		WHERE [UserId] = :UserId `

	sb := namedParameterQuery.NewNamedParameterQuery(
		`IF NOT EXISTS (SELECT * FROM LoginMaster WHERE [UserId] = :UserId)` +
			` BEGIN ` +
			insQuery +
			` END ` +
			` ELSE ` +
			` BEGIN ` +
			updQuery +
			` END `,
	)

	//query params
	sb.SetValue("UserId", User.UserId)
	sb.SetValue("UserName", User.UserName)
	sb.SetValue("Password", User.Password)
	sb.SetValue("Email", User.Email)

	conn, err := GetDbCOnnection()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	result, err := conn.Exec(sb.GetParsedQuery(), sb.GetParsedParameters()...)

	if err != nil {
		fmt.Println(err)
	}

	count, err := result.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		return true
	}

	return false
}

func (DLUser) GetUser(UserName string) models.User {

	var User models.User

	conn, err := GetDbCOnnection()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	query := namedParameterQuery.NewNamedParameterQuery(" SELECT * FROM LoginMaster WHERE UserName = :UserName ")

	query.SetValue("UserName", UserName)

	err = conn.QueryRow(query.GetParsedQuery(), query.GetParsedParameters()...).Scan(&User.UserId,
		&User.UserName,
		&User.Password,
		&User.Email)

	if err != nil {
		fmt.Println(err)
	}

	User.Password, _ = encryptdecrypt.Decrypt(User.Password)

	return User
}

func (DLUser) GetUserByUserId(UserId string) models.User {

	var User models.User

	conn, err := GetDbCOnnection()

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	query := namedParameterQuery.NewNamedParameterQuery(" SELECT * FROM LoginMaster WHERE UserId = :UserId ")

	query.SetValue("UserId", UserId)

	err = conn.QueryRow(query.GetParsedQuery(), query.GetParsedParameters()...).Scan(
		&User.UserId,
		&User.UserName,
		&User.Password,
		&User.Email,
	)

	if err != nil {
		fmt.Println(err)
	}

	User.Password, _ = encryptdecrypt.Decrypt(User.Password)

	return User
}
