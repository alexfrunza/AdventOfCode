package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Directory struct {
	name        string
	size        int
	files       []*File
	directories []*Directory
	parent      *Directory
}

type File struct {
	size   int
	name   string
	parent *Directory
}

func subDirectory(currentDirectory *Directory, dirName string) *Directory {
	for i := 0; i < len(currentDirectory.directories); i++ {
		if currentDirectory.directories[i].name == dirName {
			return currentDirectory.directories[i]
		}
	}

	return nil
}

func fileIsInDirectory(currentDirectory *Directory, fileName string) bool {
	for i := 0; i < len(currentDirectory.files); i++ {
		if currentDirectory.files[i].name == fileName {
			return true
		}
	}

	return false
}

func goBackToRoot(currentDirectory **Directory) {
	for (*currentDirectory).name != "/" {
		doCommandCd(currentDirectory, []string{"", "", ".."})
	}
}

func doCommandCd(currentDirectory **Directory, prompt []string) {
	if prompt[2] == ".." {
		*currentDirectory = (*currentDirectory).parent
		return
	}

	var newDir Directory
	newDir.name = filepath.Join((*currentDirectory).name, prompt[2])
	newDir.parent = *currentDirectory
	if dir := subDirectory((*currentDirectory), newDir.name); dir == nil {
		(*currentDirectory).directories = append((*currentDirectory).directories, &newDir)
		(*currentDirectory) = &newDir
	} else {
		(*currentDirectory) = dir
	}

}

func addFile(currentDirectory *Directory, prompt []string) {
	if prompt[0] == "dir" {
		var newDir Directory
		newDir.name = filepath.Join(currentDirectory.name, prompt[1])
		newDir.parent = currentDirectory

		if subDirectory(currentDirectory, newDir.name) == nil {
			currentDirectory.directories = append(currentDirectory.directories, &newDir)
		}
		return
	}

	var file File
	file.name = filepath.Join(currentDirectory.name, prompt[1])

	size, err := strconv.Atoi(prompt[0])

	if err != nil {
		log.Fatalln("The file uses unknown commands", err)
	}

	file.size = size
	file.parent = currentDirectory

	if !fileIsInDirectory(currentDirectory, file.name) {
		currentDirectory.files = append(currentDirectory.files, &file)
	}
}

func calculateDirSizes(directory *Directory) int {
	for i := 0; i < len(directory.files); i++ {
		directory.size += directory.files[i].size
	}

	for i := 0; i < len(directory.directories); i++ {
		directory.size += calculateDirSizes(directory.directories[i])
	}

	return directory.size
}

func drawTree(directory *Directory, nestedLevel int) {
	for i := 0; i < nestedLevel; i++ {
		fmt.Print("\t")
	}
	fmt.Printf("Type: directory\tname: %s\tsize: %d\n", directory.name, directory.size)
	for i := 0; i < len(directory.directories); i++ {
		drawTree(directory.directories[i], nestedLevel+1)
	}
	for i := 0; i < len(directory.files); i++ {
		for i := 0; i <= nestedLevel; i++ {
			fmt.Print("\t")
		}
		fmt.Printf("Type: file\tname: %s\tsize: %d\n", directory.files[i].name, directory.files[i].size)
	}
}

func calculateSum(d *Directory) int {
	var sum int
	if d.size <= 100000 {
		sum += d.size
	}
	for i := 0; i < len(d.directories); i++ {
		sum += calculateSum(d.directories[i])
	}
	return sum
}

var totalSizeOfRemovedDir int
var totalUsedSpace int

func calculateSizeOfRemovedDir(d *Directory) {
	if totalUsedSpace-d.size <= 40000000 {
		if d.size < totalSizeOfRemovedDir {
			totalSizeOfRemovedDir = d.size
		}
	}
	for i := 0; i < len(d.directories); i++ {
		calculateSizeOfRemovedDir(d.directories[i])
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var dir Directory
	var currentDirectory *Directory = &dir

	scanner.Scan()
	prompt := strings.Split(scanner.Text(), " ")
	if prompt[0] == "$" && prompt[1] == "cd" && prompt[2] == "/" {
		currentDirectory.name = "/"
	}

	for scanner.Scan() {
		prompt := strings.Split(scanner.Text(), " ")
		if prompt[0] == "$" && prompt[1] == "cd" {
			doCommandCd(&currentDirectory, prompt)
		} else if prompt[0] == "$" && prompt[1] == "ls" {
		} else {
			addFile(currentDirectory, prompt)
		}
	}

	goBackToRoot(&currentDirectory)
	calculateDirSizes(currentDirectory)
	totalSizeOfRemovedDir = currentDirectory.size
	totalUsedSpace = currentDirectory.size
	calculateSizeOfRemovedDir(currentDirectory)
	// drawTree(currentDirectory, 0)

	fmt.Println("The total size of directory that must be removed is: ", totalSizeOfRemovedDir)
}
