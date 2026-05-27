package util

import "crypto/rsa"

var (
	RsaPrivateKey                       *rsa.PrivateKey
	RsaPublicKey                        *rsa.PublicKey
	AppLabelKey                         = ""
	NamespaceLabelKey                   = ""
	PodAnnotationControllerKindLabelKey = ""
)
