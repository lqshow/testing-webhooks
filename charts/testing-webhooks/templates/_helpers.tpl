{{/*
Expand the name of the chart.
*/}}
{{- define "testing-webhooks.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "testing-webhooks.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{- define "testing-webhooks.clusterrole" -}}
{{- printf "%s.%s" .Release.Namespace .Values.fullnameOverride | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "testing-webhooks.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "testing-webhooks.labels" -}}
helm.sh/chart: {{ include "testing-webhooks.chart" . }}
{{ include "testing-webhooks.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "testing-webhooks.selectorLabels" -}}
app.kubernetes.io/name: {{ include "testing-webhooks.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "testing-webhooks.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "testing-webhooks.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}


{{/*
admissionWebhooks secret
*/}}
{{- define "testing-webhooks.admissionWebhookSecret" -}}
{{- if .Values.admissionWebhooks.existingSecret }}
{{- default "default" .Values.admissionWebhooks.existingSecret }}
{{- else }}
{{- printf "%s-server-cert" (include "testing-webhooks.fullname" .) }}
{{- end }}
{{- end }}