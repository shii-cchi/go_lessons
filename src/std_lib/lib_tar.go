package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func main() {
	archFile, err := os.Create("my_archive.tar")
	if err != nil {
		fmt.Println(err)
	}

	tw := tar.NewWriter(archFile)

	hdr := &tar.Header{
		Name: "my file.txt",
		Mode: 0600,
		Size: int64(len("ffff")),
	}

	if err := tw.WriteHeader(hdr); err != nil {
		fmt.Println(err)
	}

	if _, err := tw.Write([]byte("1234")); err != nil {
		fmt.Println(err)
	}

	if err := tw.Close(); err != nil {
		fmt.Println(err)
	}
	archFile.Seek(0, 0)
	tr := tar.NewReader(archFile)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			fmt.Println(err)
		}
		fmt.Println()
	}

}
