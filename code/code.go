package code

import (
	"bufio"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	CodesFile     string
	UsedCodesFile string
}

type Code struct {
	Config *Config
}

// GetCode - Get code to use it with 2FA.
func (c *Code) GetCode() string {
	code := c.readCode()
	c.writeCodeToUsed(code)
	c.removeUsedCode(code)

	return string(code)
}

// readCode - Read the code that we will use to pass the 2FA.
func (c *Code) readCode() []byte {
	f := openFile(c.Config.CodesFile, os.O_RDONLY, 0)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	line := scanner.Bytes()
	return line
}

// writeCodeToUsed - Write the used code to the file that
// contains the used codes.
func (c *Code) writeCodeToUsed(code []byte) {
	f := openFile(c.Config.UsedCodesFile, os.O_WRONLY|os.O_APPEND, 0600)
	defer f.Close()

	f.WriteString(string(code) + "\n")
}

// removeUsedCode - Removes the code that we will use from
// the codes file to prevent use it again.
func (c *Code) removeUsedCode(code []byte) {
	f := openFile(c.Config.CodesFile, os.O_RDWR, 0)
	defer f.Close()

	var keepData []byte
	data, _ := ioutil.ReadAll(f)
	size := len(data)
	length := len(code) + 1 // +1 for new line

	if size < length {
		// in case the size of the file data is less than the code length
		// we need to empty the file so we will make the keepData slice size
		// equal to 0.
		keepData = make([]byte, 0)
	} else {
		// else, set the keepData size to be (data size - code length)
		// this size is the other file data length without the code.
		keepData = make([]byte, size-length)
		// put the file data without the code to the keepData slice.
		copy(keepData, data[length:])
	}

	f.Truncate(0)     // empty the file
	f.Seek(0, 0)      // go to the offset 0 of the file to start the write on it
	f.Write(keepData) // write the data after we removed the used code.
}

// openFile - alternative for os.OpenFile.
func openFile(fileName string, flag int, perm fs.FileMode) *os.File {
	f, err := os.OpenFile(fileName, flag, perm)
	if err != nil {
		log.Printf("Get error while open the file: %v", err)
	}

	return f
}
