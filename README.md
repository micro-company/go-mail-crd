# go-mail-crd
Service(CRD) for work with mail by SMTP in kubernete

### Feature

- Integration with k8s
- mail service
- gRPC API

### Getting start

```
go get -u github.com/golang/protobuf/proto

protoc -I grpc/mail/ grpc/mail/mail.proto --go_out=plugins=grpc:grpc/mail
```

### Stack

- Go
- gRPC
- CustomResourceDefinition (Kubernetes)
