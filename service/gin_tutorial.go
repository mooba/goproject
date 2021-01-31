package service

import (
	"archive/zip"
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	SR_File_Max_Bytes = 1024 * 1024 * 2
	CSV_COL_NUM       = 2
)

func ParameterInPath1(c *gin.Context) {

	log.Info("calling ParameterInPath1")
	log.Debug("calling ParameterInPath1")
	log.Warn("calling ParameterInPath1")
	log.Error("calling ParameterInPath1")
	name := c.Param("name")
	msg := fmt.Sprintf("Hello %s", name)
	c.String(http.StatusOK, msg)
}

func ParameterInPath2(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func QueryStringParam(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func UploadSingleFile(c *gin.Context) {
	formFile, err := c.FormFile("file")
	if err != nil {
		log.Error(err.Error())
		c.String(http.StatusBadRequest, "formFile format error")
		return
	}
	log.Infof("receiving %s", formFile.Filename)

	if formFile.Size > SR_File_Max_Bytes {
		c.String(http.StatusBadRequest, "formFile sieze exceeds 2M")
		return
	}

	file, err := formFile.Open()
	if err != nil {
		wrapErr := fmt.Errorf("couldn't open the csv file: %w", err)
		log.Error(wrapErr.Error())
		c.String(http.StatusBadRequest, wrapErr.Error())
	}
	defer file.Close()

	//processFileLineByLine(file)
	processFile(file)
}

func processFileLineByLine(file multipart.File) {
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Error("line error:", err.Error())
			continue
		}
		if len(line) != CSV_COL_NUM {
			log.Warn("invalid line ", line)
			continue
		}
		log.Info("line: ", line[0], line[1])
	}
}

func processFile(file multipart.File) error {
	reader := csv.NewReader(bufio.NewReader(file))
	lines, err := reader.ReadAll()
	if err != nil {
		we := fmt.Errorf("error reading all lines: %w", err)
		log.Error(we.Error())
		return we
	}
	var header []string
	for i, line := range lines {
		// skip header
		if i == 0 {
			header = line
			log.Info("header:", header)
			continue
		}

		log.Infof("line: %s, %s", line[0], line[1])
	}
	return nil
}

type Login struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required,len=6"`
}

func BindWithJson(c *gin.Context) {
	region := c.GetHeader("country")
	log.Info("region in header ", region)
	var jsonReq Login
	if err := c.ShouldBindJSON(&jsonReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if jsonReq.User != "manu" || jsonReq.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func Export(c *gin.Context) {
	c.Writer.Header().Set("Content-type", "application/octet-stream")
	c.Writer.Header().Set("Content-Disposition", "attachment; filename='filename.zip'")
	ar :=  zip.NewWriter(c.Writer)
	file1, _ := os.Open("filename1")
	file2, _ := os.Open("filename2")
	f1, _ := ar.Create("filename1")
	io.Copy(f1, file1)
	f2, _ := ar.Create("filename1")
	io.Copy(f2, file2)
	ar.Close()
}
