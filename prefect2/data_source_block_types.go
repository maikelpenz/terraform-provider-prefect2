package prefect2

import (
	"context"
	"strconv"
	"time"

	hc "terraform-provider-prefect2/prefect2_api"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBlockTypes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBlockTypesRead,
		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"block_type_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"block_types": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceBlockTypesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*hc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	workspaceId := d.Get("workspace_id").(string)
	blockTypeId := d.Get("block_type_id").(string)
	slug := d.Get("slug").(string)

	var blockTypesOutput []interface{}

	if blockTypeId != "" {
		blockType, err := c.GetBlockTypeById(blockTypeId, workspaceId)
		if err != nil {
			return diag.FromErr(err)
		}

		blockTypes := make([]hc.BlockType, 1, 1)
		blockTypes[0] = *blockType
		blockTypesOutput = tfBlockTypesSchemaOutput(blockTypes)
	} else if slug != "" {
		blockType, err := c.GetBlockTypeBySlug(slug, workspaceId)
		if err != nil {
			return diag.FromErr(err)
		}

		blockTypes := make([]hc.BlockType, 1, 1)
		blockTypes[0] = *blockType
		blockTypesOutput = tfBlockTypesSchemaOutput(blockTypes)
	} else {
		blockTypes, err := c.GetAllBlockTypes(workspaceId)
		if err != nil {
			return diag.FromErr(err)
		}
		blockTypesOutput = tfBlockTypesSchemaOutput(blockTypes)
	}

	if err := d.Set("block_types", blockTypesOutput); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func tfBlockTypesSchemaOutput(block_types []hc.BlockType) []interface{} {
	schemaOutput := make([]interface{}, len(block_types), len(block_types))

	for i, block_type := range block_types {
		schema := make(map[string]interface{})

		schema["id"] = block_type.Id
		schema["name"] = block_type.Name

		schemaOutput[i] = schema
	}
	return schemaOutput
}
