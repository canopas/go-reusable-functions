module examples

go 1.20

replace email => ../email

replace file => ../file

require (
	github.com/canopas/go-scaffolds/email v0.0.0-20230731052909-8a5dc87ed95c
	github.com/canopas/go-scaffolds/file v0.0.0-20230728124028-c1cf54ceb9ba
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/aws/aws-sdk-go v1.44.313 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
)
