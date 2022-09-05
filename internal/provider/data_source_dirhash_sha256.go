package provider

import (
	"context"

	"github.com/Think-iT-Labs/dirhash/lib"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDirhashSha256() *schema.Resource {
	return &schema.Resource{
		Description: "Compute SHA256 checksum of a given directory",

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

func dataSourceDirhashSha256Read(ctx context.Context, rd *schema.ResourceData, meta any) diag.Diagnostics {
	dir := rd.Get("directory").(string)
	rd.SetId(dir)
	var ignore []string
	if v, ok := rd.GetOk("ignore"); ok {
		ignore = stringInterfaceListToStringSlice(v.([]interface{}))
	}
	checksum := lib.DirHash(dir, ignore)
	rd.Set("checksum", checksum)
	return nil
}

func stringInterfaceListToStringSlice(stringList []interface{}) []string {
	var ret []string
	for _, v := range stringList {
		if v != nil {
			ret = append(ret, v.(string))
		}
	}
	return ret
}
