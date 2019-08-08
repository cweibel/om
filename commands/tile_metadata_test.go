package commands_test

import (
	"archive/zip"

	"os"

	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pivotal-cf/om/commands"
	"github.com/pivotal-cf/om/commands/fakes"
)

var _ = Describe("TileMetadata", func() {
	Describe("Execute", func() {
		var (
			command *commands.TileMetadata
			stdout  *fakes.Logger

			productFile *os.File
			err         error
		)

		BeforeEach(func() {
			stdout = &fakes.Logger{}

			command = commands.NewTileMetadata(stdout)

			// write fake file
			productFile, err = ioutil.TempFile("", "fake-tile")
			z := zip.NewWriter(productFile)

			// https://github.com/pivotal-cf/om/issues/239
			// writing a "directory" as well, because some tiles seem to
			// have this as a separate file in the zip, which influences the regexp
			// needed to capture the metadata file
			_, err := z.Create("metadata/")
			Expect(err).NotTo(HaveOccurred())

			f, err := z.Create("metadata/fake-tile.yml")
			Expect(err).NotTo(HaveOccurred())

			_, err = f.Write([]byte(`
name: fake-tile
product_version: 1.2.3
`))
			Expect(err).NotTo(HaveOccurred())

			Expect(z.Close()).To(Succeed())
		})

		AfterEach(func() {
			Expect(os.RemoveAll(productFile.Name())).To(Succeed())
		})

		It("shows product name from tile metadata file", func() {
			err = executeCommand(command, []string{
				"-p",
				productFile.Name(),
				"--product-name",
			}, nil)
			Expect(err).NotTo(HaveOccurred())

			content := stdout.PrintlnArgsForCall(0)
			Expect(content).To(ContainElement("fake-tile"))
		})

		It("shows product version from tile metadata file", func() {
			err = executeCommand(command, []string{
				"-p",
				productFile.Name(),
				"--product-version",
			}, nil)
			Expect(err).NotTo(HaveOccurred())

			content := stdout.PrintlnArgsForCall(0)
			Expect(content).To(ContainElement("1.2.3"))
		})

		Context("failure cases", func() {
			Context("when the flags cannot be parsed", func() {
				It("returns an error", func() {
					err = executeCommand(command, []string{"--bad-flag", "some-value"}, nil)
					Expect(err).To(MatchError(MatchRegexp("could not parse tile-metadata flags")))
				})
			})

			Context("when the flags are not specified", func() {
				It("returns an error", func() {
					err = executeCommand(command, []string{"-p", productFile.Name()}, nil)
					Expect(err).To(MatchError(MatchRegexp("you must specify product-name and/or product-version")))
				})
			})

			Context("when the specified product file is not found", func() {
				It("returns an error", func() {
					err = executeCommand(command, []string{"-p", "non-existent-file", "--product-name"}, nil)
					Expect(err).To(MatchError(MatchRegexp("failed to open product file")))
				})
			})

			Context("when the file does not have metadata", func() {
				var (
					badTile *os.File
				)

				BeforeEach(func() {
					badTile, err = ioutil.TempFile("", "bad-tile")
					Expect(err).NotTo(HaveOccurred())
					z := zip.NewWriter(badTile)
					Expect(z.Close()).To(Succeed())
				})

				AfterEach(func() {
					Expect(os.RemoveAll(badTile.Name())).To(Succeed())
				})

				It("returns an error", func() {
					err = executeCommand(command, []string{"-p", badTile.Name(), "--product-name"}, nil)
					Expect(err).To(MatchError(MatchRegexp("failed to find metadata file")))
				})
			})
		})
	})
})
