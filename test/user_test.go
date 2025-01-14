package unit

import (
	"fmt"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/practice-se-lab/entity"
)

func TestStudentID(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`student_id is required`, func(t *testing.T){
		user := entity.User{
			StudentID: "", // ผิดตรงนี้
			FirstName: "John",
			LastName: "Doe",
			Email: "JohnDoe@gmail.com",
			Phone: "0123456789",
			GenderID: 1,
			// Gender: entity.Gender{
			// 	Name: "Male",
			// },
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("StudentID is required"))

	})

}

func TestFirstName(t *testing.T){

	g := NewGomegaWithT(t)

	t.Run(`first_name is required`, func(t *testing.T){
		user := entity.User{
			StudentID: "B1234567",
			FirstName: "", // ผิดตรงนี้
			LastName: "Doe",
			Email: "JohnDoe@gmail.com",
			Phone: "0123456789",
			GenderID: 1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("FirstName is required"))
	})

}

func TestLastName(t *testing.T){

	g := NewGomegaWithT(t)

	t.Run(`last_name is required`, func(t *testing.T){
		user := entity.User{
			StudentID: "M1234567",
			FirstName: "John",
			LastName: "", // <--------------- ผิดตรงนี้
			Email: "JohnDoe@gmail.com",
			Phone: "0741852963",
			GenderID: 1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("LastName: non zero value required"))

	})
}

func TestStudentIDInvalidPattern(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`student_id pattern is not true`, func(t *testing.T){
		user := entity.User{
			StudentID: "K1234567", // <---------------- ผิดตรงนี้
			FirstName: "John",
			LastName: "Doe",
			Email: "JD@gmail.com",
			Phone: "0963852741",
			GenderID: 1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(fmt.Sprintf("StudentID: %s does not validate as matches(^[BMD]\\d{7}$)", user.StudentID)))
	})
}

func TestEmailInvalid(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`email is invalid`, func(t *testing.T){
		user := entity.User{
			StudentID: "B7415263",
			FirstName: "John",
			LastName: "Doe",
			Email: "JDgmail.com", // <------------- ผิดตรงนี้
			Phone: "0852741963",
			GenderID: 1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Email: %s does not validate as email", user.Email)))
	})
}

func TestCheckDigitPhone(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`phone_number check 10 digit`, func(t *testing.T){
		user := entity.User{
			StudentID: "D5382641",
			FirstName: "M",
			LastName: "L",
			Email: "m@gmail.com",
			Phone: "07418529636", // <--------------- ผิดตรงนี้
			GenderID: 1,
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal(fmt.Sprintf("Phone: %s does not validate as matches(^[0]\\d{9}$)", user.Phone)))
	})
}