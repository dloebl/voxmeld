package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <input.wem>\n", os.Args[0])
	}
	wemPath := os.Args[1]
	bnkPath := "BNKs/Play_" + filepath.Base(wemPath) + ".bnk"

	bnk, err := os.ReadFile(bnkPath)
	if err != nil {
		panic(err)
	}

	pattern := []byte{0x01, 0x00, 0x14, 0x00}
	newCodec := []byte{0x01, 0x00, 0x04, 0x00}
	// Find the pattern in the file
	pos := bytes.Index(bnk, pattern)
	if pos == -1 {
		pos = bytes.Index(bnk, newCodec)
		if pos == -1 {
			panic("Pattern not found")
		}
	}
	// Read the values that we need
	codec := bnk[pos : pos+4]
	dummy := bnk[pos+4]
	id := binary.LittleEndian.Uint32(bnk[pos+5 : pos+9])
	fileSize := binary.LittleEndian.Uint32(bnk[pos+9 : pos+13])

	fmt.Printf("Codec:      %02X %02X %02X %02X\n", codec[0], codec[1], codec[2], codec[3])
	fmt.Printf("Dummy:      %02X\n", dummy)
	fmt.Printf("ID:         %d\n", id)
	fmt.Printf("File Size:  %d bytes\n", fileSize)

	// Get size of the .wem file
	wemInfo, err := os.Stat(wemPath)
	if err != nil {
		log.Fatalf("Failed to read .wem file: %v\n", err)
	}
	wemSize := uint32(wemInfo.Size())
	// Update the codex to VORBIS
	copy(bnk[pos:pos+4], newCodec)
	// Update file size (4 bytes after dummy byte and ID)
	fileSizeOffset := pos + 9
	if fileSizeOffset+4 > len(bnk) {
		log.Fatalf("Not enough data to update file size in .bnk")
	}
	binary.LittleEndian.PutUint32(bnk[fileSizeOffset:fileSizeOffset+4], wemSize)
	// write the modified .bnk file to out_bnks/
	outBnkPath := "out_bnks/" + filepath.Base(bnkPath)
	err = os.WriteFile(outBnkPath, bnk, 0644)
	if err != nil {
		log.Fatalf("Failed to write modified .bnk file: %v\n", err)
	}
	fmt.Printf("Modified .bnk file written to: %s\n", outBnkPath)
	// write the .wem file to out_wems/
	wem, err := os.ReadFile(wemPath)
	if err != nil {
		log.Fatalf("Failed to read .wem file: %v\n", err)
	}
	outWemPath := "out_wems/" + strconv.Itoa(int(id)) + ".wem"
	err = os.WriteFile(outWemPath, wem, 0644)
	if err != nil {
		log.Fatalf("Failed to write .wem file: %v\n", err)
	}
	fmt.Printf("Extracted .wem file written to: %s\n", outWemPath)
}
