# Telegram Mini App Template
## Deployment
1. Create project in GCP
2. Copy project ID to `GCP_PROJECT_ID` GitHub variable
3. Pick a region for your app and save it to `GCP_PROJECT_REGION`
4. Create service account with following rights:
   - Service Account User (to create resource by the name of this account)
   - Cloud Run Admin (to create Cloud Run instances)
   - Artifact Registry Administrator (to manage images in registry)
   - Secret Manager Secret Accessor (to access GCP secrets)
5. Copy service account email and save it to `GCP_SA_EMAIL` GitHub variable
6. Export service account key and save it to `GCP_SA_KEY` GitHub secret
7. Enable following GCP APIs:
   - Cloud Run Admin API (to create Cloud Run instances)
   - Secret Manager API (to securely store secrets)
8. Create [Artifact Registry for Docker images](https://cloud.google.com/artifact-registry/docs/docker/store-docker-container-images#create) in `GCP_PROJECT_REGION` region
9. Copy Artifact Registry name and save it to `GCP_ARTIFACT_REGISTRY` GitHub variable
10. Create secret with Telegram Bot token in Secret Manager
11. Copy secret name and save it to `GCP_TG_TOKEN_SECRET` GitHub variable
12. Define following GitHub variables:
   - `GCP_UI_SERVICE_NAME` with desired name of UI Cloud Run instance 
   - `GCP_API_SERVICE_NAME` with desired name of API Cloud Run instance 