package cmd

import (
	"os"
	"path/filepath"

	"github.com/blackfly19/easypdf/src/easypdf"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:    "convert",
	Short:  "Converts single or multiple markdown files to pdf",
	Long:   "Converts single or multiple markdown files to pdf",
	PreRun: prerun,
	RunE:   run,
}

func prerun(cmd *cobra.Command, args []string) {

	if cmd.Flags().Changed("ddl") || cmd.Flags().Changed("dtl") || cmd.Flags().Changed("tht") ||
		cmd.Flags().Changed("tli") || cmd.Flags().Changed("ttss") {
		cmd.MarkFlagRequired("toc")
	}
}

func run(cmd *cobra.Command, args []string) error {

	fileConvert, err := parseValues(cmd)
	if err != nil {
		return err
	}

	if fileConvert.WatchMode {
		change := make(chan bool)

		if fileConvert.Directory != "" {
			go easypdf.DirWatcher(fileConvert.Directory, change)
		} else {
			go easypdf.FileWatcher(fileConvert.Files, change)
		}

		for {
			if fileConvert.Directory != "" {
				fileConvert.Files, err = getFilesFromDir(fileConvert.Directory)
				if err != nil {
					return err
				}
			}

			err = fileConvert.ConvertFileToPDF()
			if err != nil {
				return err
			}

			<-change
		}
	}

	err = fileConvert.ConvertFileToPDF()
	if err != nil {
		return err
	}

	return nil
}

func parseValues(cmd *cobra.Command) (*easypdf.MdToPdf, error) {

	var mdFiles []string
	var directory string
	var err error

	if cmd.Flags().Changed("files") {
		mdFiles, err = cmd.Flags().GetStringArray("files")
		if err != nil {
			return nil, err
		}
	} else {
		directory, err = cmd.Flags().GetString("dir")
		if err != nil {
			return nil, err
		}

		mdFiles, err = getFilesFromDir(directory)
		if err != nil {
			return nil, err
		}

	}

	outputFilename, err := cmd.Flags().GetString("output")
	if err != nil {
		return nil, err
	}

	cssFilename, err := cmd.Flags().GetString("css")
	if err != nil {
		return nil, err
	}

	enableToc, err := cmd.Flags().GetBool("toc")
	if err != nil {
		return nil, err
	}

	disableDottedLines, err := cmd.Flags().GetBool("ddl")
	if err != nil {
		return nil, err
	}

	disableTocLinks, err := cmd.Flags().GetBool("dtl")
	if err != nil {
		return nil, err
	}

	tocHeaderText, err := cmd.Flags().GetString("tht")
	if err != nil {
		return nil, err
	}

	tocLevelIndentation, err := cmd.Flags().GetUint("tli")
	if err != nil {
		return nil, err
	}

	tocTextSizeShrink, err := cmd.Flags().GetFloat64("ttss")
	if err != nil {
		return nil, err
	}

	watchMode, err := cmd.Flags().GetBool("watch")
	if err != nil {
		return nil, err
	}

	coverPage, err := cmd.Flags().GetString("cover")
	if err != nil {
		return nil, err
	}

	return &easypdf.MdToPdf{
		Files:             mdFiles,
		Directory:         directory,
		OutputFilename:    outputFilename,
		CssFilename:       cssFilename,
		CoverPageFileName: coverPage,
		WatchMode:         watchMode,
		Toc: easypdf.Toc{
			Include:             enableToc,
			DisableDottedLines:  disableDottedLines,
			DisableTocLinks:     disableTocLinks,
			TocHeaderText:       tocHeaderText,
			TocLevelIndentation: tocLevelIndentation,
			TocTextSizeShrink:   tocTextSizeShrink,
		},
	}, nil
}

func getFilesFromDir(directory string) ([]string, error) {

	var mdFiles []string

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			mdFiles = append(mdFiles, filepath.Join(wd, directory, file.Name()))
		}
	}

	return mdFiles, nil
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringArrayP("files", "f", make([]string, 0), "Single file or multiple files")
	convertCmd.Flags().StringP("dir", "d", "", "Directory name containing markdown files")
	convertCmd.Flags().StringP("output", "o", uuid.New().String(), "Output file name")
	convertCmd.Flags().String("css", "", "css file name")
	convertCmd.Flags().BoolP("toc", "t", false, "Add table of contents")
	convertCmd.Flags().Bool("ddl", false, "Disable dotted lines in toc")
	convertCmd.Flags().Bool("dtl", false, "Disable toc links")
	convertCmd.Flags().String("tht", "Table of Contents", "Toc header text")
	convertCmd.Flags().Uint("tli", 1, "Toc level indentation")
	convertCmd.Flags().Float64("ttss", 0.8, "Font scaling for each level of heading")
	convertCmd.Flags().BoolP("watch", "w", false, "Enable watch mode to see changes in real time")
	convertCmd.Flags().StringP("cover", "c", "", "Cover page of the document")

	convertCmd.MarkFlagsMutuallyExclusive("files", "dir")
}
