# CodeDeploy
## Instance startup

```
aws ec2 run-instances \
  --image-id ami-0c5204531f799e0c6  \
  --key-name test-codedeploy-gettingstarted-keypair \
  --user-data file://instance-setup.sh \
  --count 1 \
  --instance-type t2.micro \
  --iam-instance-profile Name=CodeDeployDemo-EC2-Instance-Profile
```

## CodeDeploy setup

```
aws deploy create-application --application-name HelloWorld_App
aws deploy create-deployment-group \
  --application-name HelloWorld_App \
  --deployment-group-name HelloWorld_DepGroup \
  --deployment-config-name CodeDeployDefault.OneAtATime \
  --ec2-tag-filters Key=Name,Value=CodeDeployDemo,Type=KEY_AND_VALUE \
  --service-role-arn arn:aws:iam::037361087574:role/CodeDeployServiceRole
```

## Compiling for CodeDeploy instance

```
cd src/hello/
GOOS=linux GOARCH=amd64 go build
```

## Upload

```
aws deploy push \
  --application-name HelloWorld_App \
  --s3-location s3://marvr-codedeploydemobucket/HelloWorldApp.zip \
  --ignore-hidden-files
```

## Deploy

```
aws deploy create-deployment \
  --application-name HelloWorld_App \
  --deployment-config-name CodeDeployDefault.OneAtATime \
  --deployment-group-name HelloWorld_DepGroup \
  --s3-location bucket=marvr-codedeploydemobucket,bundleType=zip,key=HelloWorldApp.zip
```

# Lambda

## Create function

```
aws lambda create-function --function-name my-test-golang-function \
  --zip-file fileb://lambdahello.zip --handler lambdahello --runtime go1.x \
  --role arn:aws:iam::037361087574:role/lambda-role
```

## Delete function

```
aws lambda delete-function --function-name my-test-golang-function
```

## Update function

```
aws lambda update-function-code --function-name my-test-golang-function \
  --zip-file fileb://lambdahello.zip
```

## Invoke function

```
aws lambda invoke --function-name my-test-golang-function /tmp/output.json; cat /tmp/output.json
```
