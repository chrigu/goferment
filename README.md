# Goferment - Fermentation controller in go

## Setup

1. Follow [](https://github.com/yryz/ds18b20) to setup the 1-wire sensor
2. Setup your AWS account [](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-dynamodb-with-go-sdk.html)
3. Setup a DynamoDB with a table
3. Add the credentials to `.aws/credentials`
4. Setup a `.env` file. Use `env.example` as template