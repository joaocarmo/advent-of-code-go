package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const ROOT_NAME = "root"
const ROOT_MARKER = "/"
const PARENT_MARKER = ".."
const DIR_SEPARATOR = "/"
const COMMAND_MARKER = "$ "
const COMMAND_ARGS_SEPARATOR = " "
const FOLDER_MARKER = "dir"
const FILE_SIZE_NAME_SEPARATOR = " "
const FOLDER_MARKER_NAME_SEPARATOR = " "
const FOLDER_SIZE_THRESHOLD = int64(100000)

// File represents a file in the file system.
type File struct {
    Name string
	Size int64
}

// Folder represents a folder in the file system.
type Folder struct {
    Name    string             `json:"name"`
    Files   []File             `json:"files"`
    Folders map[string]*Folder `json:"folders"`
	Parent  *Folder            `json:"-"`
}

// newFolder creates a new folder and returns a pointer to it.
func newFolder(name string, parent *Folder) *Folder {
    return &Folder{name, []File{}, make(map[string]*Folder), parent}
}

// addFile adds a file to the current folder.
func (f *Folder) addFile(input string) {
	// split the file size and name
	sizeName := strings.Split(input, FILE_SIZE_NAME_SEPARATOR)
	fileSizeInt, _ := strconv.ParseInt(sizeName[0], 10, 64)
	fileName := sizeName[1]

	// add the file to the folder
	f.Files = append(f.Files, File{fileName, fileSizeInt})
}

// addFolder adds a subfolder to the current folder.
func (f *Folder) addFolder(input string) {
	// split the folder marker and name
	markerName := strings.Split(input, FOLDER_MARKER_NAME_SEPARATOR)
	folderName := markerName[1]

	// add the folder to the folder
	f.Folders[folderName] = newFolder(folderName, f)
}

// getParent returns the parent folder of the current folder.
func (f *Folder) getParent() *Folder {
	return f.Parent
}

// getRoot returns the root folder of the current folder.
func (f *Folder) getRoot() *Folder {
	if f.Parent == nil {
		return f
	}

	return f.Parent.getRoot()
}

// getSize returns the size of the current folder.
func (f *Folder) getSize() int64 {
	size := int64(0)

	for _, file := range f.Files {
		size += file.Size
	}

	for _, folder := range f.Folders {
		size += folder.getSize()
	}

	return size
}

// cd changes the current folder to the specified folder.
func (f *Folder) cd(args string, output []string) *Folder {
	switch args {
	case ROOT_MARKER:
		return f.getRoot()
	case PARENT_MARKER:
		return f.getParent()
	default:
		return f.Folders[args]
	}
	return nil
}

// ls lists the files and folders in the current folder.
func (f *Folder) ls(args string, output []string) {
	for _, line := range output {
		if isCommand(line) {
			break
		} else if isFolder(line) {
			f.addFolder(line)
		} else {
			f.addFile(line)
		}
	}
}

// execCommand executes a command in the current folder.
func (f *Folder) execCommand(input string, output []string) *Folder {
	// remove the command marker
	command := strings.TrimPrefix(input, COMMAND_MARKER)

	// split the command and arguments
	commandAndArgs := strings.Split(command, COMMAND_ARGS_SEPARATOR)
	commandName := commandAndArgs[0]
	commandArgs := ""

	if len(commandAndArgs) > 1 {
		commandArgs = strings.Join(commandAndArgs[1:], COMMAND_ARGS_SEPARATOR)
	}

	// execute the command
	switch commandName {
	case "cd":
		return f.cd(commandArgs, output)
	case "ls":
		f.ls(commandArgs, output)
	}
	return nil
}

// String returns a string representation of the current folder as JSON.
func (f *Folder) String() string {
    json, err := json.MarshalIndent(f, "", "  ")

    if err != nil {
        log.Fatalf(err.Error())
    }

	return string(json)
}

// isFolder checks if the input is a folder marker.
func isFolder(input string) bool {
	MARKER_LEN := len(FOLDER_MARKER)

	if len(input) < MARKER_LEN {
		return false
	}

	return input[0:MARKER_LEN] == FOLDER_MARKER
}

// isCommand checks if the input is a command marker.
func isCommand(input string) bool {
	MARKER_LEN := len(COMMAND_MARKER)

	if len(input) < MARKER_LEN {
		return false
	}

	return input[0:MARKER_LEN] == COMMAND_MARKER
}

// getFileSystem returns the file system represented by the input.
func getFileSystem(txtlines []string) *Folder {
	root := newFolder(ROOT_NAME, nil)

	currentFolder := root
	for i, line := range txtlines {
		if isCommand(line) {
			newFolder := currentFolder.execCommand(line, txtlines[i+1:])

			if newFolder != nil {
				currentFolder = newFolder
			}
		}
	}

	return root
}

// getTotalSizeFoldersToDelete returns the total size of the folders to delete.
func getTotalSizeFoldersToDelete(folders []*Folder) int64 {
	totalSize := int64(0)

	for _, folder := range folders {
		totalSize += folder.getSize()
	}

	return totalSize
}

// findFoldersToDelete returns the folders to delete.
func findFoldersToDelete(folder *Folder, threshold int64) []*Folder {
	foldersToDelete := []*Folder{}

	for _, subfolder := range folder.Folders {
		if subfolder.getSize() <= threshold {
			foldersToDelete = append(foldersToDelete, subfolder)
		}

		foldersToDelete = append(foldersToDelete, findFoldersToDelete(subfolder, threshold)...)
	}

	return foldersToDelete
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// process the file
	filesystem := getFileSystem(txtlines)

	// part 1
	folderDeleteCandidates := findFoldersToDelete(filesystem, FOLDER_SIZE_THRESHOLD)
	totalSizeDeleteCandidates := getTotalSizeFoldersToDelete(folderDeleteCandidates)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		totalSizeDeleteCandidates,
	)
}
