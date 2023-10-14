# Telegram Mini App Template
Template for building Telegram Bots with Mini Apps in React and Go. Deployed instance available at [@tg_mini_app_template_bot](https://t.me/tg_mini_app_template_bot).

## Features
- Example of creating [Inline Keyboard](https://core.telegram.org/bots/features#inline-keyboards) to interact with the user. Keyboard creation can be found in [main.go](api/main.go)
- Example of adjusting UI styles to the user's current color theme using [ThemeParams](https://core.telegram.org/bots/webapps#themeparams). CSS variable usage can be found in [index.css](ui/src/index.css)
- Example of interaction with [telegram-web-app.js](https://core.telegram.org/bots/webapps#initializing-mini-apps) to show a native alert. Script usage can be found in [App.jsx](ui/src/App.jsx) 
- [Local environment](#local-environment) with [docker-compose](https://docs.docker.com/compose/), that serves traffic over HTTPS proxy and automatically applies UI changes without rebuilding the container
- [Continious deliviery](#production-deployment) to [Google Cloud Platform](https://cloud.google.com/)

## Setup prerequisites
[Telegram Bot](https://core.telegram.org/bots) token is required to interact with [Telegram Bot API](https://core.telegram.org/bots/api). To get one, сreate a bot using [@BotFather](https://t.me/botfather) or follow [Telegram instructions](https://core.telegram.org/bots#how-do-i-create-a-bot).

## Local environment
This repository provides an easy-to-use local development environment. Using it you can start writing your bot business logic without spending time on the environment.

Local environment includes:
- [ngrok](https://ngrok.com/) reverse proxy to server local mini-app and bot deployment over HTTPS
- [nginx](https://www.nginx.com/) reverse proxy to host both API and UI on one ngrok domain and thus fit into the [free plan](https://ngrok.com/pricing)
- React fast refresh to avoid rebuilding docker container on each change of the UI code

Local environment setup:
1. Create an account on [ngrok](https://ngrok.com/)
0. Get a [ngrok auth token](https://ngrok.com/docs/secure-tunnels/ngrok-agent/tunnel-authtokens/) and save it to `NGROK_AUTHTOKEN` variable in `.env` file in the project root directory
0. Claim a [free ngrok domain](https://ngrok.com/blog-post/free-static-domains-ngrok-users) and save it to `NGROK_DOMAIN` variable in `.env` file in the project root directory
0. Copy [Telegram Bot token](#setup-prerequisites) and save it to `TELEGRAM_BOT_TOKEN` variable in `.env` file in the project root directory
0. Install [Docker](https://docs.docker.com/get-docker/)

To start or update the environment with the latest code changes, use:
```sh
docker compose up --build -d
```

After successful deployment, your local bot API will be available at https://ngrok-domain/api. Use this URL to set the bot webhook as described [switching bot environment](#switching-bot-environment).

## Production deployment
This repository provides a [workflow](https://docs.github.com/actions) to automatically deploy the code to [Google Cloud Platform](https://cloud.google.com/). The deploy job is triggered on each push to the [main](https://github.com/st-matskevich/tg-mini-app-template/tree/main) branch.

GCP services used for deployment:
- [Cloud Run](https://cloud.google.com/run) to host dockerized API and UI code
- [Artifact Registry](https://cloud.google.com/artifact-registry) to store docker images
- [Secret Manager](https://cloud.google.com/secret-manager) to store sensitive data

Deployment setup:
1. [Create a project](https://cloud.google.com/resource-manager/docs/creating-managing-projects#creating_a_project) in GCP
0. Copy project ID to `GCP_PROJECT_ID` GitHub variable
0. [Pick a region](https://cloud.withgoogle.com/region-picker/) for your app and save it to `GCP_PROJECT_REGION` GitHub variable
0. [Create a service account](https://cloud.google.com/iam/docs/service-accounts-create#creating) with the following rights:
   - Service Account User (to create resources by the name of this account)
   - Cloud Run Admin (to create Cloud Run instances)
   - Artifact Registry Administrator (to manage images in the registry)
   - Secret Manager Secret Accessor (to access GCP secrets)
0. Copy the service account email and save it to `GCP_SA_EMAIL` GitHub variable
0. [Export the service account key](https://cloud.google.com/iam/docs/keys-create-delete#creating) and save it to `GCP_SA_KEY` GitHub secret
0. Enable the following GCP APIs:
   - Cloud Run Admin API (to create Cloud Run instances)
   - Secret Manager API (to securely store secrets)
0. Create [Artifact Registry for Docker images](https://cloud.google.com/artifact-registry/docs/docker/store-docker-container-images#create) in `GCP_PROJECT_REGION` region
0. Copy Artifact Registry name and save it to `GCP_ARTIFACT_REGISTRY` GitHub variable
0. [Create a secret](https://cloud.google.com/secret-manager/docs/creating-and-accessing-secrets#create) with [Telegram Bot token](#setup-prerequisites) in Secret Manager
0. Copy the secret name and save it to `GCP_SECRET_TG_TOKEN` GitHub variable
0. Define the following GitHub variables:
   - `GCP_SERVICE_UI_NAME` with the desired name of UI Cloud Run instance 
   - `GCP_SERVICE_UI_MAX_INSTANCES` with the desired maximum number of UI service instances
   - `GCP_SERVICE_API_NAME` with the desired name of API Cloud Run instance 
   - `GCP_SERVICE_API_MAX_INSTANCES` with the desired maximum number of API service instances   

After successful deployment, obtain the bot API URL from either `deploy-api` job results or from [GCP Project Console](https://console.cloud.google.com) and proceed to [switching bot environment](#switching-bot-environment).

## Switching bot environment
After the bot is either [launched locally](#local-environment) or [deployed in GCP](#production-deployment), Telegram needs to be configured with a proper webhook URL. To set it, use:
```sh
curl https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/setWebhook?url=${BOT_API_URL}/bot
```

## Template structure
- Project root directory - contains files for [local environment](#local-environment)
  - [docker-compose.yml](docker-compose.yml) - docker-compose file to setup  [local environment](#local-environment)
  - [proxy.template](proxy.template) - nginx config template to route ngrok domain to API and UI containers
- [api](/api/) - contains files for the bot backend written in Go
  - [Dockerfile](/api/Dockerfile) - docker file for the API service image
  - Other files - minimal Go package layout with HTTP endpoint to handle [Telegram Bot API](https://core.telegram.org/bots/api) calls
- [ui](/ui/) - contains files for the mini app UI written in JS with React
  - [dev.dockerfile](/ui/dev.dockerfile) - docker file for the [local environment](#local-environment) to track code changes and updated served UI bundle
  - [prod.dockerfile](/ui/prod.dockerfile) - docker file for the [production deployment](#production-deployment) to bundle UI and serve via nginx
  - [nginx.template](/ui/nginx.template) - nginx config template for production UI service image
  - Other files - React template with minimal changes for the integration with Telegram

## Built with
- [Docker](https://www.docker.com/)
- [Go](https://go.dev/)
- [React](https://react.dev/)
- [Vite](https://vitejs.dev)
- [gotgbot](https://github.com/PaulSonOfLars/gotgbot)
- [nginx](https://www.nginx.com/)
- [ngrok](https://ngrok.com/)

## License
Distributed under the MIT License. See [LICENSE](LICENSE) for more information.

## Contributing
Want a new feature added? Found a bug?
Go ahead and open [a new issue](https://github.com/st-matskevich/tg-mini-app-template/issues/new) or feel free to submit a pull request.