package storage

import (
    "bytes"
    "fmt"
    "mime/multipart"
    "net/http"
    "os"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

func UploadToS3(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
    bucket := os.Getenv("AWS_BUCKET_NAME")
    region := os.Getenv("AWS_REGION")

    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(region),
    })
    if err != nil {
        return "", err
    }

    uploader := s3.New(sess)

    buf := bytes.NewBuffer(nil)
    if _, err := buf.ReadFrom(file); err != nil {
        return "", err
    }

    key := fileHeader.Filename

    _, err = uploader.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
        Body:   bytes.NewReader(buf.Bytes()),
        ACL:    aws.String("public-read"),
        ContentType: aws.String(http.DetectContentType(buf.Bytes())),
    })

    if err != nil {
        return "", err
    }

    return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, key), nil
}
