# Telegram Mini App Template
## Prerequisites
## Local environment
## Production deployment
This repository provides a [workflow](https://docs.github.com/actions) to automatically deploy the code to [Google Cloud Platform](https://cloud.google.com/). Deploy job is triggered on each push to the [main](https://github.com/st-matskevich/tg-mini-app-template/tree/main) branch.

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
0. [Create a secret](https://cloud.google.com/secret-manager/docs/creating-and-accessing-secrets#create) with [Telegram Bot token](#prerequisites) in Secret Manager
0. Copy the secret name and save it to `GCP_TG_TOKEN_SECRET` GitHub variable
0. Define the following GitHub variables:
   - `GCP_UI_SERVICE_NAME` with the desired name of UI Cloud Run instance 
   - `GCP_API_SERVICE_NAME` with the desired name of API Cloud Run instance 

After successful deployment, obtain API service URL from either `deploy-api` job results or from [GCP Project Console](https://console.cloud.google.com) and proceed to [switching bot environment](#switching-bot-environment).

## Switching bot environment
After the bot is either [launched locally](#local-environment) or [deployed in GCP](#production-deployment), Telegram needs to be configured with a proper webhook URL. To set it, use:
```sh
curl https://api.telegram.org/bot${TELEGRAM_BOT_TOKEN}/setWebhook?url=${BOT_API_URL}/bot
```

## Built with
- [Docker](https://www.docker.com/)
- [Go](https://go.dev/)
- [React](https://react.dev/)
- [gotgbot](https://github.com/PaulSonOfLars/gotgbot)
- [nginx](https://www.nginx.com/)
- [ngrok](https://ngrok.com/)

## License
Distributed under the MIT License. See [LICENSE](LICENSE) for more information.

## Contributing
Want a new feature added? Found a bug?
Go ahead an open [a new issue](https://github.com/st-matskevich/tg-mini-app-template/issues/new) or feel free to submit a pull request.