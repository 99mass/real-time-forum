package helper

import (
	"database/sql"
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"forum/controller"
	"forum/models"
)

func CheckRegisterFormat(username, email, password, confirmPassword string, db *sql.DB) (bool, models.ErrorAuth) {
	ErrAuth := new(models.ErrorAuth)
	ok := true

	if !ConfirmPasswordsMatch(password, confirmPassword) {
		//fmt.Println("Les mots de passe ne sont pas conformes")
		ok = false
		ErrAuth.PasswordError = "The passwords do not match"
	} else {
		okPassWord, errP := CheckPassword(password)
		if !okPassWord {
			ok = false
			ErrAuth.PasswordError = errP.Error()
		}
	}

	okUserName, errUN := CheckUserName(username)
	if !okUserName {
		ok = false
		ErrAuth.UserNameError = errUN.Error()
	}

	okEmail, errE := CheckEmail(email)
	//fmt.Println("checkemail:",okEmail)
	if !okEmail {
		ok = false
		ErrAuth.EmailError = errE.Error()
	} else {
		//fmt.Println("checking dupli")
		email, errdup := controller.IsDuplicateEmail(db, email)

		if email {
			ok = false
			ErrAuth.EmailError = errdup.Error()
			//return false,models.ErrorAuth{}
		}
		username, err := controller.IsDuplicateUsername(db,username)
		if username {
			ok = false
			ErrAuth.UserNameError = err.Error()
		}
	}
	//fmt.Println(ErrAuth.EmailError,)
	return ok, *ErrAuth
}

// Check if the form it valid and try to connect the user

// Vérification du format de l'email "name@name.ext"
func CheckEmail(email string) (bool, error) {

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, _ := regexp.MatchString(emailRegex, email)
	if !match {
		return false, errors.New("invalid email format")
	}

	return match, nil
}

// Vérification du format du mot de passe.
// Il doit avoir 8 à 25 caractère.
// Il doit contenir au moins un chiffre, une lettre majuscule, une lettre minuscule et un caractère spécial
func CheckPassword(password string) (bool, error) {
	// Cette expression exige au moins 8 caractères avec au moins une lettre majuscule,
	// une lettre minuscule, un chiffre et un caractère spécial.
	// Vérification de la longueur

	if len(password) < 8 || len(password) > 25 {

		return false, errors.New("Invalid password length: minimum 8, maximum 25")
	}

	// Vérification des autres conditions avec des expressions régulières
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	digitRegex := regexp.MustCompile(`\d`)
	specialCharRegex := regexp.MustCompile(`[@$!%*?&_\-]`)

	if !lowercaseRegex.MatchString(password) {

		return false, errors.New("the password must contain at least one lower-case letter")
	}
	if !uppercaseRegex.MatchString(password) {

		return false, errors.New("the password must contain at least one capital letter")
	}
	if !digitRegex.MatchString(password) {

		return false, errors.New("the password must contain at least one digit")
	}
	if !specialCharRegex.MatchString(password) {

		return false, errors.New("the password must contain at least one special character")
	}

	// if !match {
	// 	return false, errors.New("Format mot de passe non valide!")
	// }

	return true, nil

}

// Vérification du format du UserName
// Il doit avoir 5 à 20 Caractères alpha_numérique
func CheckUserName(username string) (bool, error) {
	// Cette expression exige que le pseudo ait entre 5 et 20 caractères alphanumériques.
	usernameRegex := `^[a-zA-Z0-9]{3,20}$`

	match, _ := regexp.MatchString(usernameRegex, username)

	if !match {
		return false, errors.New("invalid username format")
	}

	return match, nil

}

//---------Post------------

func CheckTitle(title string) {

}

//--------Comment--------

func CheckContent(content string) {

}

// Cryptage du mot de passe
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Confirmation du mot de passe lors de l'inscription d'uun nouveau client
func ConfirmPasswordsMatch(password, confirmPassword string) bool {
	return password == confirmPassword
}
