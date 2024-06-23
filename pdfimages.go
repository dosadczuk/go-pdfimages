// Package pdfimages is a wrapper for Xpdf command line tool `pdfimages`.
//
// What is `pdfimages`?
//
//	Pdfimages saves images from a Portable Document Format (PDF) file as
//	Portable Pixmap (PPM), Portable Graymap (PGM), Portable Bitmap (PBM),
//	or JPEG files.
//
//	Pdfimages reads the PDF file, scans one or more pages, PDF-file, and
//	writes one PPM, PGM, PBM, or JPEG file for each image.
//
//	Note: `pdfimages` extracts the raw image data from the PDF file,
//	without performing any additional transforms. Any rotation, clipping,
//	color inversion, etc. done by the PDF content stream is ignored.
//
// Reference: https://www.xpdfreader.com/pdfimages-man.html
package pdfimages

import (
	"context"
	"os/exec"
	"strconv"
)

// ----------------------------------------------------------------------------
// -- `pdfimages`
// ----------------------------------------------------------------------------

type Command struct {
	path string
	args []string
}

// NewCommand creates new `pdfimages` command.
func NewCommand(opts ...option) (*Command, error) {
	cmd := &Command{path: "pdfimages"}
	for _, opt := range opts {
		opt(cmd)
	}

	var err error

	// assert that executable exists and get absolute path
	cmd.path, err = exec.LookPath(cmd.path)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

// Run executes prepared `pdfimages` command.
func (c *Command) Run(ctx context.Context, inpath, outdir string) error {
	cmd := exec.CommandContext(ctx, c.path, append(c.args, inpath, outdir)...)

	return cmd.Run()
}

// String returns a human-readable description of the command.
func (c *Command) String() string {
	return exec.Command(c.path, append(c.args, "<inpath>", "<outdir>")...).String()
}

// ----------------------------------------------------------------------------
// -- `pdfimages` options
// ----------------------------------------------------------------------------

type option func(*Command)

// Set custom location for `pdfimages` executable.
func WithCustomPath(path string) option {
	return func(c *Command) {
		c.path = path
	}
}

// Read config-file in place of ~/.xpdfrc or the system-wide config file.
func WithCustomConfig(path string) option {
	return func(c *Command) {
		c.args = append(c.args, "-cfg", path)
	}
}

// Specifies the first page to scan.
func WithPageFrom(from uint64) option {
	return func(c *Command) {
		c.args = append(c.args, "-f", strconv.FormatUint(from, 10))
	}
}

// Specifies the last page to scan.
func WithPageTo(to uint64) option {
	return func(c *Command) {
		c.args = append(c.args, "-l", strconv.FormatUint(to, 10))
	}
}

// Specifies the range of pages to convert.
func WithPageRange(from, to uint64) option {
	return func(c *Command) {
		WithPageFrom(from)
		WithPageTo(to)
	}
}

// With this option, images in DCT format are saved as JPEG files. All non-DCT
// images are saved in PBM/PGM/PPM format as usual.
//
// Normally, all images are written as:
//   - PBM (for monochrome images),
//   - PGM (grayscale images),
//   - PPM (for color images).
//
// Note: inline images are always saved in PBM/PGM/PPM format.
func WithSaveDctAsJpeg() option {
	return func(c *Command) {
		c.args = append(c.args, "-j")
	}
}

// Write all images in PDF-native formats.
//
// Most of the formats are not standard image formats, so this option is primarily
// useful as input to a tool that generates PDF files.
//
// Note: inline images are always saved in PBM/PGM/PPM format.
func WithSaveRaw() option {
	return func(c *Command) {
		c.args = append(c.args, "-raw")
	}
}

// Specify the owner password for the PDF file.
//
// Providing this will bypass all security restrictions.
func WithOwnerPassword(password string) option {
	return func(c *Command) {
		c.args = append(c.args, "-opw", password)
	}
}

// Specify the user password for the PDF file.
func WithUserPassword(password string) option {
	return func(c *Command) {
		c.args = append(c.args, "-upw", password)
	}
}
