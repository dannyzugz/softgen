package services

/*

func HandleDownload(w http.ResponseWriter, r *http.Request) {
	// Gets the name of the folder to be compressed from the URL
	projectname := c.Param("projectname")

	// Create a temporary file to store the ZIP file
	zipFile, err := os.CreateTemp("", "")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer os.Remove(zipFile.Name())

	// Create a writer for the ZIP file
	zipWriter := zip.NewWriter(zipFile)

	dir := "/generatedProjects/" + projectname

	// Compress the folder into the ZIP file
	folderPath := filepath.Join(".", dir)
	err = compressFolder(folderPath, zipWriter)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Close the ZIP file writer
	err = zipWriter.Close()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Send ZIP file as HTTP response
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", "attachment; filename="+projectname+".zip")
	c.File(zipFile.Name())

	err = os.RemoveAll(dir)
	if err != nil {
		fmt.Println("Error deleting temporary directory:", err)
		return
	}
} */
