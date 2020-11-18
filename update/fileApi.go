package update

import (
	"api-update/dao"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, request *http.Request) {

	nomeUser := request.FormValue("nome")

	idUser := request.FormValue("idUser")

	nameDirectory := createDirectory(nomeUser)

	file := downloadFile(request)

	saveFileInDirectory(nameDirectory, idUser, file)

	fmt.Println(w, "Successfully Uploading")
}

func saveFileInDirectory(nameDirectory string, idUser string, file multipart.File) {

	tempFile, err := ioutil.TempFile(nameDirectory, "*.png") //salvando imagem no diretorio

	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	dao.InsertTest(idUser, tempFile.Name())

}

func createDirectory(nameUser string) string {
	nomeDiretorio := "images/" + nameUser

	_, err := os.Stat("images/" + nameUser)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("images/"+nameUser, 0777)
		if errDir != nil {
			fmt.Println(errDir)
		}

	}

	return nomeDiretorio
}

func downloadFile(r *http.Request) multipart.File {
	r.ParseMultipartForm(10 << 20) //tamanho imagem

	file, handler, err := r.FormFile("myFile")

	if err != nil {
		fmt.Println(err)
	}

	if file == nil {
		fmt.Println("Arquivo vazio")
		file.Close()
	}

	defer file.Close()

	fmt.Println("Nome do Arquivo: ", string(handler.Filename))
	fmt.Println("Tamanho do Arquivo: ", handler.Size)
	fmt.Println("Tipo do Arquivo: ", string(handler.Header.Get("Content-Type")))
	fmt.Println("****-------------------------****")

	return file
}

// SetupRoutes !
func SetupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}
