package s3

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Upload(bucket, region, fileName string, upload io.Reader) error {

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region),
		//Credentials: credentials.NewSharedCredentials("", "devaccount"),
	})

	uploader := s3manager.NewUploader(sess)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fileName),
		Body:        upload,
		ContentType: aws.String("application/json"),
	})
	if err != nil {
		exitErrorf("Error uploading %q to S3, %v", err)
	}
	log.Println("Uploaded", fileName, "to", bucket, "bucket.")
	return nil
}

func Download(bucket, region, object string) error {

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region),
		//Credentials: credentials.NewSharedCredentials("", "devaccount"),
	})

	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create("/tmp/" + object)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", object, err)
	}
	defer file.Close()

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(object),
		})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Object %q not found in s3 bucket - %v", object, err)
		return err
	}
	log.Println("Downloaded", object, "from", bucket, "bucket.")
	return nil
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
