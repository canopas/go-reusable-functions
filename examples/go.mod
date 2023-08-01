module examples

go 1.20

replace email => ../email

replace file => ../file

require (
	github.com/canopas/go-reusables/email v0.0.0-20230801123715-3957df64671f
	github.com/canopas/go-reusables/file v0.0.0-20230801123715-3957df64671f
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/aws/aws-sdk-go v1.44.311 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
)
