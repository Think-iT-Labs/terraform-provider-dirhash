package provider

import (
	"context"

	"github.com/Think-iT-Labs/dirhash/lib"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDirhashSha256() *schema.Resource {
	return &schema.Resource{
		Description: "Calculate SHA256 checksum of a given directory",

		ReadContext: dataSourceDirhashSha256Read,

		Schema: map[string]*schema.Schema{
			"directory": {
				Type:        schema.TypeString,
				Description: "Directory to hash.",
				Required:    true,
			},
			"ignore": {
				Type:        schema.TypeList,
				Description: "List of ignored glob patterns.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
			},
			"checksum": {
				Type:        schema.TypeString,
				Description: "Directory SHA 256 Checksum",
				Computed:    true,
			},
		},
	}
}

func dataSourceDirhashSha256Read(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var dir = d.Get("directory").(string)
	d.SetId(dir)
	var ignore = []string{}
	if v, ok := d.GetOk("ignore"); ok {
		ignore = stringListToStringSlice(v.([]interface{}))
	}
	var checksum = lib.DirHash(dir, ignore)
	d.Set("checksum", checksum)
	return nil
}

func stringListToStringSlice(stringList []interface{}) []string {
	ret := []string{}
	for _, v := range stringList {
		if v == nil {
			ret = append(ret, "")
			continue
		}
		ret = append(ret, v.(string))
	}
	return ret
}
