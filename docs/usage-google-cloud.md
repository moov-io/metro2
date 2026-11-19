---
layout: page
title: Google Cloud Run
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Google Cloud Run

To get started in a hosted environment you can deploy this project to the Google Cloud Platform.

From your [Google Cloud dashboard](https://console.cloud.google.com/home/dashboard) create a new project and call it:
```
moov-metro2-demo
```

Enable the [Artifact Registry](https://cloud.google.com/artifact-registry/docs/docker/store-docker-container-images) and Cloud Run APIs for your project and associate a [billing account](https://cloud.google.com/billing/docs/how-to/manage-billing-account) if needed. Then, open the Cloud Shell terminal and run the following commands, substituting your unique project ID and a region close to you (for example `us-central1`):

```
gcloud services enable artifactregistry.googleapis.com run.googleapis.com

gcloud artifacts repositories create moov \
  --repository-format=docker \
  --location=us-central1 \
  --description="Docker images"

gcloud auth configure-docker us-central1-docker.pkg.dev

docker pull moov/metro2
docker tag moov/metro2 us-central1-docker.pkg.dev/<PROJECT-ID>/moov/metro2
docker push us-central1-docker.pkg.dev/<PROJECT-ID>/moov/metro2
```

Deploy the container to Cloud Run:
```
gcloud run deploy metro2 --image us-central1-docker.pkg.dev/<PROJECT-ID>/moov/metro2 --port 8080 --region us-central1
```

You can also deploy the public image directly, without pushing a copy to Artifact Registry:

```
gcloud run deploy metro2 --image moov/metro2:latest --port 8080 --region us-central1
```

Upon a successful deploy you will be given a URL where the API has been deployed:

```
https://YOUR-METRO2-APP-URL.a.run.app
```

Now you can complete a health check:
```
curl https://YOUR-METRO2-APP-URL.a.run.app/health
```
You should get this response:
```
{"status":"alive"}
```
