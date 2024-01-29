// test_user.go
package unit

import (
    "testing"
    "github.com/Kami0rn/Carne/entity"
    . "github.com/onsi/gomega"
    "github.com/asaskevich/govalidator"
)

func TestOk(t *testing.T) {
    g := NewGomegaWithT(t)

    t.Run(`test_ok`, func(t *testing.T) {

        create := entity.User{
            FirstName:   "Ok",
            LastName:    "Ok",
            PhoneNumber: 123441234,
            Email:       "t@g.com",
            Password:    "OK",
        }
        ok, err := govalidator.ValidateStruct(create)

        g.Expect(ok).To(BeTrue())
        g.Expect(err).To(BeNil())
    })
}
