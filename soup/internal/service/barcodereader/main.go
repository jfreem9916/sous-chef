package barcodereader

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type BarcodeCodeResponse struct {
	Code string `json:"code"`
}

func ReadBarcodeFromFile(filepath string) (barcode string, err error) {
	reader, err := os.Open(filepath)

	if err != nil {
		return "", err
	}

	//does not work when qr code doesn't occupy whole image
	img, _, _ := image.Decode(reader)

	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

	barcodeReader := oned.NewUPCAReader()

	result, err := barcodeReader.Decode(bmp, nil)

	if err != nil {
		return "", err
	}

	return result.GetText(), nil
}

func CreateServer() {

	router := gin.Default()

	router.POST("/api/v0/actions/read-barcode", func(ctx *gin.Context) {
		// single file
		file, _ := ctx.FormFile("file")

		timenano := time.Now().UnixNano()
		timestr := strconv.FormatInt(timenano, 10)

		// Upload the file to specific dst.
		final_filename := "./tmp/barcode" + timestr
		ctx.SaveUploadedFile(file, final_filename)

		code, err := ReadBarcodeFromFile(final_filename)

		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to read barcode"})
			return
		}

		ctx.IndentedJSON(http.StatusAccepted, BarcodeCodeResponse{Code: code})
	})

	router.Run("localhost:8080")
}
