package gogdrive

import "fmt"

func (dc *driveClient) List() (string, error) {
	r, err := dc.ds.Files.List().PageSize(10).Fields("files").Do()

	fmt.Println("Responses:")
	for _, f := range r.Files {
		fmt.Printf("File Name: %s\nFile type: %s", f.Name, f.MimeType)
		fmt.Printf("\n\n")
	}

	if err != nil {
		return "", err
	}

	return "", nil

}
