# PostgreSQL

PostgreSQL runs as a StatefulSet.

Persistent Volumes store database files.

Daily backups run at 2 AM.

Database credentials should be stored in Kubernetes Secrets.