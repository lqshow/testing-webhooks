/*
Copyright 2023 LQ.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var guestbooklog = logf.Log.WithName("guestbook-resource")

// SetupWebhookWithManager will setup the manager to manage the webhooks
func (r *Guestbook) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

var _ webhook.Defaulter = &Guestbook{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Guestbook) Default() {
	guestbooklog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-webapp-foobar-ai-v1-guestbook,mutating=false,failurePolicy=fail,sideEffects=None,groups=webapp.foobar.ai,resources=guestbooks,verbs=create;update,versions=v1,name=vguestbook.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Guestbook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Guestbook) ValidateCreate() (admission.Warnings, error) {
	guestbooklog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.

	return nil, r.validateGuestbook()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Guestbook) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	guestbooklog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Guestbook) ValidateDelete() (admission.Warnings, error) {
	guestbooklog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}

func (r *Guestbook) validateGuestbook() error {
	var allErrs field.ErrorList
	if err := r.validateGuestbookName(); err != nil {
		allErrs = append(allErrs, err)
	}
	if len(allErrs) == 0 {
		return nil
	}

	return errors.NewInvalid(
		schema.GroupKind{Group: "webapp.foobar.ai", Kind: "Guestbook"},
		r.Name, allErrs)
}

func (r *Guestbook) validateGuestbookName() *field.Error {
	if len(r.ObjectMeta.Name) > 5 {
		return field.Invalid(field.NewPath("metadata").Child("name"), r.Name, "Guestbook name must be no more than 5 characters for test purposes")
	}
	return nil
}
