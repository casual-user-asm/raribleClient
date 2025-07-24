
# Rarible Client

Service to interact with the Rarible API.




## 🚀 Features

 - Retrieve a list of rarible nft ownerships
 - Retrieve nft trait rarities
 - Be able to run in a docker container
 - Have an automated test suite
 - Have a helm chart to deploy it to a kubernetes cluster




## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`RARIBLE_API_KEY=your api key`



## 🏁 Running with Docker

Clone the repository

```bash
  git clone https://github.com/casual-user-asm/raribleClient.git
  cd raribleClient
```

Build the Docker image

```bash
  docker build -t rarible-client .
```

Run the container

```bash
  docker run -p 8080:8080 rarible-client
```
## API Reference

#### Returns Ownership by Id

```http
  GET /ownership/:id
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |
| `id` | `string` | **Required**. Ownership Id in format: 'ETHEREUM:${token}:${tokenId}:${owner}' |

#### Returns the rarity of the trait

```http
  POST /traits
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `api_key` | `string` | **Required**. Your API key |
| `collectionid`      | `string` | **Required**. Collection id in format ETHEREUM:${collectionid} |



