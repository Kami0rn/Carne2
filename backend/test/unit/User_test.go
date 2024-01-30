// test_user.go
package unit

import (
    "testing"
    "github.com/Kami0rn/Carne/entity"
    . "github.com/onsi/gomega"
    "github.com/asaskevich/govalidator"
)

func Test_Ok(t *testing.T) {
    g := NewGomegaWithT(t)

    t.Run(`test_ok`, func(t *testing.T) {

        create := entity.User{
            FirstName:   "Ok",
            LastName:    "Ok",
            PhoneNumber: 96341085,
            Email:       "t@g.com",
            Password:    "OK",
        }
        ok, err := govalidator.ValidateStruct(create)

        g.Expect(ok).To(BeTrue())
        g.Expect(err).To(BeNil())
    })
}

func TestFname(t *testing.T){
    g := NewGomegaWithT(t)

    t.Run(`tes_firstname`,func (t *testing.T)  {

        create := entity.User{
            FirstName:   "",
            LastName:    "Ok",
            PhoneNumber: 96341085,
            Email:       "t@g.com",
            Password:    "OK",
        }
        ok,err := govalidator.ValidateStruct(create)

        g.Expect(ok).NotTo(BeTrue())
        g.Expect(err).NotTo(BeNil())

        g.Expect(err.Error()).To(Equal("Please fill FirstName"))
        
    })
}

func TestLname (t *testing.T){
    g := NewGomegaWithT(t)

    t.Run(`test_lastname`, func(t *testing.T){

        create := entity.User {
            FirstName:   "Ok",
            LastName:    "",
            PhoneNumber: 96341085,
            Email:       "t@g.com",
            Password:    "OK",
        }
        ok,err := govalidator.ValidateStruct(create)

        g.Expect(ok).NotTo(BeTrue())
        g.Expect(err).NotTo(BeNil())

        g.Expect(err.Error()).To(Equal("Please fill LastName"))
    })
}


func TestEmailValid(t *testing.T)  {

    g := NewGomegaWithT(t)

    t.Run(`test_email`,func (t *testing.T)  {

        create := entity.User{
            FirstName:   "Ok",
            LastName:    "Ok",
            PhoneNumber: 96341085,
            Email:       "tg",
            Password:    "OK",
        }
        ok,err := govalidator.ValidateStruct(create)

        g.Expect(ok).NotTo(BeTrue())
        g.Expect(err).NotTo(BeNil())

        g.Expect(err.Error()).To(Equal("Email is invalid"))
        
    })
    
}

func TestEmailIsFill (t *testing.T){
    g := NewGomegaWithT(t)

    t.Run(`test_emailfill`,func(t *testing.T)  {

        create := entity.User{
            FirstName:   "Ok",
            LastName:    "Ok",
            PhoneNumber: 96341085,
            Email:       "",
            Password:    "OK",
        }
        ok,err := govalidator.ValidateStruct(create)

        g.Expect(ok).NotTo(BeTrue())
        g.Expect(err).NotTo(BeNil())

        g.Expect(err.Error()).To(Equal("Please fill email"))
        
    })
}


