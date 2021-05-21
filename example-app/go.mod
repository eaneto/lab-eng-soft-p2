module gitihub.com/lab-end-soft-p2/example-app

go 1.16

require (
	github.com/aws/aws-sdk-go-v2/config v1.2.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v1.4.0 // indirect
	github.com/eaneto/notify v0.17.1 // indirect
	github.com/nikoksr/notify v0.16.1 // indirect
)

replace github.com/nikiksr/notify v0.16.1 => github.com/eaneto/notify v0.17.1

replace github.com/nikiksr/notify/amazonsqs v0.16.1 => github.com/eaneto/notify/amazonsqs v0.17.1
