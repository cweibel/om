package commands_test

import (
	"errors"

	"github.com/pivotal-cf/om/api"
	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/commands/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeleteProduct", func() {
	var (
		command     *commands.DeleteProduct
		fakeService *fakes.DeleteProductService
	)

	BeforeEach(func() {
		fakeService = &fakes.DeleteProductService{}
		command = commands.NewDeleteProduct(fakeService)
	})

	Describe("Execute", func() {
		It("deletes the specific product", func() {
			err := executeCommand(command, []string{"-p", "some-product-name", "-v", "1.2.3-build.4"})
			Expect(err).NotTo(HaveOccurred())

			Expect(fakeService.DeleteAvailableProductsCallCount()).To(Equal(1))

			input := fakeService.DeleteAvailableProductsArgsForCall(0)
			Expect(input).To(Equal(api.DeleteAvailableProductsInput{
				ProductName:             "some-product-name",
				ProductVersion:          "1.2.3-build.4",
				ShouldDeleteAllProducts: false,
			}))
		})

		Context("failure cases", func() {
			Context("when deleting a product fails", func() {
				It("returns an error", func() {
					fakeService.DeleteAvailableProductsReturns(errors.New("something bad happened"))

					err := executeCommand(command, []string{"-p", "nah", "-v", "nope"})
					Expect(err).To(MatchError("something bad happened"))
				})
			})

			Context("when the --product-name flag is missing", func() {
				It("returns an error", func() {
					err := executeCommand(command, []string{
						"--product-version", "1.2.3",
					})
					Expect(err.Error()).To(MatchRegexp("the required flag.*--product-name"))
				})
			})

			Context("when the --product-version flag is missing", func() {
				It("returns an error", func() {
					err := executeCommand(command, []string{
						"--product-name", "some-product",
					})
					Expect(err.Error()).To(MatchRegexp("the required flag.*--product-version"))
				})
			})
		})
	})
})
