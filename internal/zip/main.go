package zip

import (
	"log"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkcarreno/avdm/internal/system"
)

func UnzipAsTemp(basepath string, location string) string {

	regexToPickFileNameWithExtension := regexp.MustCompile(`[^\\\/]+(=\.[\w]+$)|[^\\\/]+$`)
	PickedFileName := regexToPickFileNameWithExtension.FindString(location)

	PickedFileNameWithoutExtension := strings.TrimSuffix(PickedFileName, filepath.Ext(PickedFileName))

	newUnzipFolderName := "unzip-" + PickedFileNameWithoutExtension

	newFolderPath := basepath + "temp/" + newUnzipFolderName

	err := system.UnzipSource(location, newFolderPath)
	if err != nil {
		log.Fatal(err)
	}

	return newFolderPath
}
