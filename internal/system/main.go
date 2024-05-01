package system

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func MoveFile(src, dest string) error {
	// Obtenga la información del archivo de origen.
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Crea la carpeta de destino si no existe.
	if !info.IsDir() {
		err = os.MkdirAll(filepath.Dir(dest), os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Mueve el archivo a la carpeta de destino.
	err = os.Rename(src, dest)
	if err != nil {
		return err
	}

	return nil
}

func MoveDir(src, dest string) error {
	// Obtenga la información del archivo o carpeta de origen.
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Si el archivo o carpeta es un archivo, muévalo directamente.
	if !info.IsDir() {
		return MoveFile(src, dest)
	}

	// Si el archivo o carpeta es una carpeta, muévala recursivamente.
	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := MoveDir(src+"/"+file.Name(), dest+"/"+file.Name())
		if err != nil {
			return err
		}
	}

	return nil
}

func UnzipSource(source, destination string) error {
	// Open the zip file
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	// Get the absolute destination path
	destination, err = filepath.Abs(destination)
	if err != nil {
		return err
	}

	// Iterate over zip files inside the archive and unzip each of them
	for _, f := range reader.File {
		err := UnzipFile(f, destination)
		if err != nil {
			return err
		}
	}

	return nil
}

func UnzipFile(f *zip.File, destination string) error {
	// Check if file paths are not vulnerable to Zip Slip
	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	// Create directory tree
	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// Create a destination file for unzipped content
	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// Unzip the content of a file and copy it to the destination file
	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}
	return nil
}
