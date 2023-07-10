package main

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrSameFile              = errors.New("input and output files are the same")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	inputFile, err := os.Open(fromPath)
	if err != nil {
		return err
	}
	defer func() {
		err := inputFile.Close()
		if err != nil {
			return
		}
	}()

	inputFileInfo, err := inputFile.Stat()
	if err != nil {
		return err
	}

	if !inputFileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	if inputFileInfo.Size() < offset {
		return ErrOffsetExceedsFileSize
	}

	outFileInfo, err := os.Stat(toPath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if os.SameFile(inputFileInfo, outFileInfo) {
		return ErrSameFile
	}

	outFile, err := os.Create(toPath)
	if err != nil {
		return err
	}

	defer func() {
		err = outFile.Close()
		if err != nil {
			return
		}
	}()

	if limit == 0 {
		limit = inputFileInfo.Size()
	}

	if offset > 0 {
		_, err := inputFile.Seek(offset, 0)
		if err != nil {
			return err
		}
	}

	return process(inputFile, outFile, limit)
}

func process(in io.Reader, out io.Writer, limit int64) error {
	bar := pb.New(int(limit)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)

	defer bar.Finish()

	bar.ShowSpeed = true

	bar.Start()

	reader := bar.NewProxyReader(in)

	_, err := io.CopyN(out, reader, limit)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	return nil
}
