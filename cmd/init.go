/*
Copyright © 2024 Pedro Carreño <pkcarrenodev@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mingrammer/cfmt"
	"github.com/pkcarreno/avdm/internal/cmdlinetoolsutils"
	"github.com/pkcarreno/avdm/internal/dir"
	"github.com/pkcarreno/avdm/internal/system"
	"github.com/pkcarreno/avdm/internal/zip"
	"github.com/pkcarreno/avdm/pkg/prompt"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup android environment",
	Long:  `Take the Android command line tools and set up a base Android development environment so you can start coding quickly without having to deal with other configurations or download the entire Android Studio IDE (which helps you save space if you don't intend to use the IDE).`,
	Run: func(cmd *cobra.Command, args []string) {
		HOME_DIR, homeDirErr := os.UserHomeDir()
		if homeDirErr != nil {
			log.Fatal(homeDirErr)
			return
		}

		// TODO: Add flag to choose the android installation path

		AVDM_BASE_PATH := HOME_DIR + "/.avdm/"
		TEMP_PATH := AVDM_BASE_PATH + "temp/"
		ANDROID_BASE_PATH := AVDM_BASE_PATH + "android/"

		cfmt.Infoln("Checking for Android environment")
		// TODO: check if android enviroment exists and if not create it
		existAndroidEnviroment := false
		if existAndroidEnviroment {
			fmt.Println(cfmt.Serrorf("Android environment already exists"))
			return
		}

		proceedWithInstalation := prompt.YesNoPrompt("Do you want to proceed with the installation?", true)
		if !proceedWithInstalation {
			fmt.Println(cfmt.Serrorf("User refuses to proceed"))
			return
		}

		cfmt.Infoln("Checking latest version of Android command line tools")
		CLT_LASTEST_VERSION, err := cmdlinetoolsutils.GetLatestVersion()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Done")

		cfmt.Infoln("Downloading Android command line tools")
		CLT_ZIP_PATH, err := cmdlinetoolsutils.DownloadPackage(CLT_LASTEST_VERSION, AVDM_BASE_PATH)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Done")

		cfmt.Infoln("Unzipping Android command line tools")
		CLT_RAW_PATH := zip.UnzipAsTemp(AVDM_BASE_PATH, CLT_ZIP_PATH)
		fmt.Println("Done")

		cfmt.Infoln("Moving Android command line tools to Android base path")
		CLT_PATH := CLT_RAW_PATH + "/cmdline-tools"
		moveUnzipFolderErr := system.MoveDir(CLT_PATH, ANDROID_BASE_PATH)
		if moveUnzipFolderErr != nil {
			fmt.Println(moveUnzipFolderErr)
			return
		}
		fmt.Println("Done")

		cfmt.Infoln("Removing temporary files")
		dir.RemoveDir(TEMP_PATH)
		fmt.Println("Done")

		cfmt.Infoln("Set Android tools executable permissions")
		ANDROID_BIN_LOCATION := ANDROID_BASE_PATH + "/bin"
		dir.SetExecutablePermissions(ANDROID_BIN_LOCATION)
		fmt.Println("Done")

		cfmt.Successln("Environment setup successfully!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
