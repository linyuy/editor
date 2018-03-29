package service

import (
    "net/http"
    "github.com/unrolled/render"

    "crypto/md5"
    "regexp"
    "fmt"
    "io"
    "net/url"
    "os"
    "errors"
)

// Image : type of image saved
// type Image struct {
//     Link     string `json:"link"`
//     Filename string `json:"filename"`
//     Time     int64  `json:"created_at"`
// }

// ResText : response of save text
type ResText struct {
    Status      int32 `json:"success"`
    URL         string `json:"url"`
    Message     string `json:"message"`
}

func saveTextError(formatter *render.Render, w http.ResponseWriter, err error)  {
    formatter.JSON(w, http.StatusNotAcceptable, 
        ResText{
            Status: 0,
            URL: "",
            Message: err.Error() })
}

func saveMd(imageBytes []byte, imageName string, imageFile io.Reader) (string, error) {

    md5Checksum := md5.Sum(imageBytes)

    // format file name
    uploadFilePath := fmt.Sprintf("/images/upload/%x-%v.%v", md5Checksum, imageName, ".md")
    uploadedFile, err := os.Create("./assets" + uploadFilePath)
    if err != nil {
        return "", errors.New("Unable to create the file for writing. Check your write access privilege")
    }
    defer uploadedFile.Close()

    _, err = io.Copy(uploadedFile, imageFile)
    if err != nil {
        return "", errors.New("Unable to create the file for writing. Check your write access privilege")
    }

    return uploadFilePath, nil
}

func saveMdHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

        req.ParseMultipartForm(32 << 20)

        imageFile, imageHeader, err := req.FormFile("image")
        if err != nil {
            imageFile, imageHeader, err = req.FormFile("editormd-image-file")
            if err != nil {
                saveImageError(formatter, w, errors.New("Set file: image"))
                return
            }
        }
        defer imageFile.Close()

        // Read 512 bytes to recognize file type
        imageBytes := make([]byte, 512)
        _, err = imageFile.Read(imageBytes)
        if err != nil {
            saveImageError(formatter, w, err)
            return
        }
        
        extensionMatcher := regexp.MustCompile("\\.\\w+$")
        imageName := extensionMatcher.ReplaceAllString(imageHeader.Filename, "")
        imageName = url.PathEscape(imageName)
        
        imageFile.Seek(0, 0)
        url, err := saveImage(imageBytes, imageName, imageFile)
        if err != nil {
            saveImageError(formatter, w, err)
            return
        }

        formatter.JSON(w, http.StatusOK, 
            ResText{
                Status: 1,
                URL: "http://" + req.Host + "/static" + url,
                Message: "" })

        // timestamp := time.Now().UnixNano() / int64(time.Millisecond)
        // image := Image{link, imageHeader.Filename, timestamp}
        // jsonBytes, _ := json.Marshal(image)
        // w.Write(jsonBytes)
	}
}

