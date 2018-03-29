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

// Response : response of save image
type Response struct {
    Status      int32 `json:"success"`
    URL         string `json:"url"`
    Message     string `json:"message"`
}

func saveImageError(formatter *render.Render, w http.ResponseWriter, err error)  {
    formatter.JSON(w, http.StatusNotAcceptable, 
        Response{
            Status: 0,
            URL: "",
            Message: err.Error() })
}

func getImageExtension(imageBytes []byte) string {
    fileType := http.DetectContentType(imageBytes)

    switch fileType {
    case "image/jpeg", "image/jpg":
        return "jpg"
    case "image/gif":
        return "gif"
    case "image/png":
        return "png"
    case "application/pdf":
        return "pdf"
    default:
        return ""
    }
}

func saveImage(imageBytes []byte, imageName string, imageFile io.Reader) (string, error) {

    extension := getImageExtension(imageBytes)
    if len(extension) == 0 {
        return "", errors.New("Unsupported file types")
    }

    md5Checksum := md5.Sum(imageBytes)

    // format file name
    uploadFilePath := fmt.Sprintf("/images/upload/%x-%v.%v", md5Checksum, imageName, extension)
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

func saveImageHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
        // req.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
        // //注意:如果没有调用ParseForm方法，下面无法获取表单的数据
        // fmt.Println(req.Form) //这些信息是输出到服务器端的打印信息
        // fmt.Println("path", req.URL.Path)
        // fmt.Println("scheme", req.URL.Scheme)
        // fmt.Println(req.Form["url_long"])
        // for k, v := range req.Form {
        //     fmt.Println("key:", k)
        //     fmt.Println("val:", strings.Join(v, ""))
        // }

        // fmt.Println("SAVE")

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
            Response{
                Status: 1,
                URL: "http://" + req.Host + "/static" + url,
                Message: "" })

        // timestamp := time.Now().UnixNano() / int64(time.Millisecond)
        // image := Image{link, imageHeader.Filename, timestamp}
        // jsonBytes, _ := json.Marshal(image)
        // w.Write(jsonBytes)
	}
}

