package main

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

func main() {
    svc := s3.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

    params := &s3.ListObjectsInput{
        Bucket: aws.String("bucket"),
    }

    resp, _ := svc.ListObjects(params)
    for _, key := range resp.Contents {
        fmt.Println(*key.Key)
    }
}


uploader := s3manager.NewUploader(sess)

_, err = uploader.Upload(&s3manager.UploadInput{
    Bucket: aws.String(AWS_S3_BUCKET), // Bucket to be used
    Key:    aws.String(filename),      // Name of the file to be saved
    Body:   file,                      // File
})
if err != nil {
    // Do your error handling here
    return
}
