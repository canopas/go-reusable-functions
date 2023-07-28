module github.com/cp-sumi-k/go-scaffolds/email

go 1.20

replace file => ../file

require (
	file v0.0.0-00010101000000-000000000000
	github.com/aws/aws-sdk-go v1.44.311
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.7.0
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
