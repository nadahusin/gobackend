package middleware

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/nadahusin/rental/src/libs"
)

func FileUpload(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseMultipartForm(32 << 20); err != nil {
			libs.Respone(err.Error(), 400, true).Send(w)
			return
		}

		file, fileHeader, err := r.FormFile("image")

		defer file.Close()

		if err != nil {
			libs.Respone(err.Error(), 400, true).Send(w)
			return
		}

		//validasi extension
		Extension := fileHeader.Header.Get("Content-Type") ==
			"image/jpeg" || fileHeader.Header.Get("Content-Type") == "image/jpg" ||
			fileHeader.Header.Get("Content-Type") == "image/png"

		if !Extension {
			libs.Respone(err.Error(), 400, true).Send(w)
			return
		}

		fmt.Println(fileHeader.Header.Get("Content-Type"))

		img := strings.ReplaceAll(strings.ReplaceAll(time.Now().Format(time.ANSIC),
			":", "-")+"-"+fileHeader.Filename, " ", "_")
		result, err := os.Create("image/" + img)

		if err != nil {
			libs.Respone(err.Error(), 400, true).Send(w)
			return
		}

		if _, error := io.Copy(result, file); error != nil {
			libs.Respone(err.Error(), 400, true).Send(w)
			return
		}

		libs.Respone(nil, 200, false)

		ctx := context.WithValue(r.Context(), "imgName", img)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
