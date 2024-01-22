package minio

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func resourceMinioStsKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: minioCreateSTSKey,
		ReadContext:   minioReadSTSKey,
		DeleteContext: minioDeleteSTSKey,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"oidc_access_token": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"oidc_id_token": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"access_key_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"secret_access_key": {
				Type:      schema.TypeString,
				Computed:  true,
				Optional:  true,
				Sensitive: true,
			},
			"session_token": {
				Type:      schema.TypeString,
				Computed:  true,
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func minioCreateSTSKey(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	stsConfig := STSKeyConfig(d, meta)

	endpoint := stsConfig.MinioClient.EndpointURL()

	var getWebToken = func() (*credentials.WebIdentityToken, error) {

		return &credentials.WebIdentityToken{
			Token:       stsConfig.MinioOIDCIdToken,
			AccessToken: stsConfig.MinioOIDCAccessToken,
		}, nil
	}

	sts, err := credentials.NewSTSWebIdentity(endpoint.String(), getWebToken)
	if err != nil {
		return NewResourceError("error creating STS credentials", d.Id(), err)
	}

	cred, err := sts.Get()
	if err != nil {
		return NewResourceError("error retrieving STS credentials", d.Id(), err)
	}

	d.SetId(string("key-" + cred.AccessKeyID))
	_ = d.Set("access_key_id", string(cred.AccessKeyID))
	_ = d.Set("secret_access_key", string(cred.SecretAccessKey))
	_ = d.Set("session_token", string(cred.SessionToken))

	log.Printf("[WARN] create (%v)", d)

	return nil
}

func minioReadSTSKey(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// noop
	log.Printf("[WARN] read (%v)", d)
	return nil
}

func minioDeleteSTSKey(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// noop
	log.Printf("[WARN] delete (%v)", d)
	return nil
}
