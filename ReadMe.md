A rather thrown together serverless Go API to calculate the packs required to for fill an order.

First attempt at writing Go 🤔 Not happy with the directory structure, not got the time to fiddle with it. Unit tests feel like they could be composed in a better way, but I don't have time to look into how to properly compose unit tests in Go. TLDR, 🤷‍♂️ time is short.

## Publicly Accessible Deployment
API URI: [packs-example.projects.robertcurran.uk/packs/requestedCount=345](packs-example.projects.robertcurran.uk/packs/requestedCount=345)

GET /packs?requestedCount={requestedCount}

Where `requestedCount` is the number of items to get a packing list for. All other routes and methods will be 403.

## Installing
In order to create the executable used by the lambda run time please run:
`GOOS=linux go build -o ./adapters/rest/packs/main ./adapters/rest/packs/main.go`


## Running 

To run this lambda locally using aws sam ensure that you have sam available locally 
and have followed the installation instructions above then
please run:

`sam local start-api --template infrastructure_template.yaml`

Otherwise, see aws cloudformation docs to deploy to your environment.

```shell
aws cloudformation package --template-file ./infrastructure_template.yaml --s3-bucket {{BUCKETNAME}} --s3-prefix packingListService  --region eu-west-1 --output-template-file ./packaged.yaml --profile robertCurranAccount

aws cloudformation deploy --template-file ./packaged.yaml --stack-name packingListService --region eu-west-1 --s3-bucket {{BUCKETNAME}} --s3-prefix packingListService --capabilities CAPABILITY_NAMED_IAM CAPABILITY_IAM CAPABILITY_AUTO_EXPAND --profile robertCurranAccount

```

## Structure

The service cde is all contained within the service directory. web_app contains the React web application to accompny this.

Main rest handler for /`packs` route: `adapters/rest/packs`
Domain logic in  `packingList` along with matching unit test

Cloudformation template at `./infrastructure_template.yaml` and contains the SAM definitions for the api and lambda.
This is deployed as a GoLang lambda.