package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	S3_CLIENT *s3.S3
	S3_BUCKET string
	wg        sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("FAKE_ID", "FAKE_SECRET", ""),
	})
	if err != nil {
		panic(err)
	}

	S3_CLIENT = s3.New(sess)
	S3_BUCKET = "goexpert-bucket-example-waldrey"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	uploadControl := make(chan struct{}, 100)
	errorFileUpload := make(chan string, 10)

	go func() {
		for {
			select {
			case file := <-errorFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(file, uploadControl, errorFileUpload)
			}
		}
	}()

	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %v\n", err)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	fmt.Printf("Uploading file %s to bucket %s\n", completeFileName, S3_BUCKET)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", completeFileName, err)
		<-uploadControl
		errorFileUpload <- filename
		return
	}
	defer f.Close()
	_, err = S3_CLIENT.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s: %v\n", filename, err)
		<-uploadControl
		errorFileUpload <- filename
		return
	}

	fmt.Printf("File %s uploaded successfully\n", filename)
	<-uploadControl
}
