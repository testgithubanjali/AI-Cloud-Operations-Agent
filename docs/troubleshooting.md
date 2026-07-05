# Troubleshooting

## CrashLoopBackOff

Usually caused by:

- Application crash
- Wrong configuration
- Missing environment variables

## OOMKilled

The container exceeded its memory limit.

Possible fixes:

- Increase memory limits.
- Investigate memory leaks.

## ImagePullBackOff

Usually caused by:

- Wrong image name
- Missing registry credentials

## Pending Pods

Usually caused by:

- Insufficient resources
- Missing Persistent Volumes