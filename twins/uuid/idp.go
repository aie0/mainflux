// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

// Package uuid provides a UUID identity provider.
package uuid

import (
	"github.com/gofrs/uuid"
	"github.com/mainflux/mainflux/twins"
)

var _ twins.IdentityProvider = (*uuidIdentityProvider)(nil)

type uuidIdentityProvider struct{}

// New instantiates a UUID identity provider.
func New() twins.IdentityProvider {
	return &uuidIdentityProvider{}
}

func (idp *uuidIdentityProvider) ID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (idp *uuidIdentityProvider) IsValid(u4 string) error {
	if _, err := uuid.FromString(u4); err != nil {
		return twins.ErrMalformedEntity
	}

	return nil
}
