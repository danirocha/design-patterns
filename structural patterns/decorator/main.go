package main

import (
	"fmt"
	"regexp"
)

const ENV_ENCRYPTION_ALGORYTHM = "hgfads"
const ENV_COMPRESSION_ALGORYTHM = "20%"

type DataSource interface {
	WriteData(string)
	ReadData() string
}

type FileDataSource struct {
	Name    string
	Content string
}

func (f FileDataSource) WriteData(s string) {
	f.Content = s
}

func (f FileDataSource) ReadData() string {
	return f.Content
}

type EncryptionDecorator struct {
	DataSource
	EncryptionAlgorithm string
}

func (ed EncryptionDecorator) WriteData(s string) {
	if s == "" {
		s = ed.DataSource.ReadData()
	}
	fmt.Println("Encrypting data...")
	ed.DataSource.WriteData(s + ed.EncryptionAlgorithm)
}

func (ed EncryptionDecorator) ReadData() string {
	data := ed.DataSource.ReadData()
	fmt.Println("Decrypting data...")
	return regexp.MustCompile(ed.EncryptionAlgorithm).ReplaceAllString(data, "")
}

type CompressionDecorator struct {
	DataSource
	CompressionAlgorithm string
}

func (cd CompressionDecorator) WriteData(s string) {
	if s == "" {
		s = cd.DataSource.ReadData()
	}
	fmt.Println("Compressing data...")
	cd.DataSource.WriteData(s + cd.CompressionAlgorithm)
}

func (cd CompressionDecorator) ReadData() string {
	data := cd.DataSource.ReadData()
	fmt.Println("Decompressing data...")
	return regexp.MustCompile(cd.CompressionAlgorithm).ReplaceAllString(data, "")
}

func main() {
	var source DataSource
	source = FileDataSource{"someFile.txt", "Heya!"}
	fmt.Println(source.ReadData())

	source = EncryptionDecorator{source, ENV_ENCRYPTION_ALGORYTHM}
	source.WriteData("")
	fmt.Println(source.ReadData())

	source = CompressionDecorator{source, ENV_COMPRESSION_ALGORYTHM}
	source.WriteData("")
	fmt.Println(source.ReadData())
}
